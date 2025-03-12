package clip

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

//go:generate go run ./generators/flags/flags.go
//go:generate go run ./generators/cli/cli.go

type BaseSetter func(v *Clip)

func Name(n string) BaseSetter {
	return func(v *Clip) { v.Name = n }
}

func Description(des string) BaseSetter {
	return func(v *Clip) { v.Description = des }
}

func Copyright(cop string) BaseSetter {
	return func(v *Clip) { v.Copyright = cop }
}

func Author(author string) BaseSetter {
	return func(v *Clip) { v.Authors = append(v.Authors, author) }
}

func Version(version string) BaseSetter {
	return func(v *Clip) { v.Version = version }
}

func Build(build string) BaseSetter {
	return func(v *Clip) { v.Build = build }
}

func TakeArguments() BaseSetter {
	return func(v *Clip) { v.takeArgs = true }
}

func Entrypoint(entrypoint func(*CLI)) BaseSetter {
	return func(v *Clip) {
		v.entrypoint = entrypoint
	}
}

func ArgsDescription(des string) BaseSetter {
	return func(v *Clip) { v.ArgsDescription = des }
}

type Clip struct {
	Name            string
	Description     string
	Copyright       string
	Authors         []string
	Version         string
	Build           string
	ArgsDescription string
	flags           []baseFlag
	takeArgs        bool
	entrypoint      func(*CLI)
}

func New(opts ...BaseSetter) *Clip {
	v := &Clip{}
	for _, fn := range opts {
		fn(v)
	}
	return v
}

func (v *Clip) String(name string) StringFlag {
	str := &stringFlag{
		basicFlag: basicFlag[string]{
			name: name,
		},
		isFs:      false,
		isDir:     false,
		mustExist: false,
		value:     "",
	}
	v.flags = append(v.flags, str)
	return str
}

func (v *Clip) StringSlice(name string) StringSliceFlag {
	str := &stringSliceFlag{
		basicFlag: basicFlag[[]string]{
			name: name,
		},
	}
	v.flags = append(v.flags, str)
	return str
}

func (v *Clip) Boolean(name string) BoolFlag {
	b := &boolFlag{
		basicFlag: basicFlag[bool]{
			name: name,
		},
	}
	v.flags = append(v.flags, b)
	return b
}

func (v *Clip) Options(name string, options ...string) OptionsFlag {
	opts := &optionsFlag{
		basicFlag: basicFlag[string]{
			name: name,
		},
		options: options,
	}
	v.flags = append(v.flags, opts)
	return opts
}

func (v *Clip) KVFlag(name string) KVFlag {
	f := &kvFlag{
		basicFlag: basicFlag[map[string]string]{
			name: name,
		},
	}
	v.flags = append(v.flags, f)
	return f
}

func (v *Clip) TakeArguments() {
	v.takeArgs = true
}

func In[V comparable, S ~[]V](s S, v ...V) bool {
	for _, i := range v {
		if slices.Contains(s, i) {
			return true
		}
	}
	return false
}

func (v *Clip) Run() { v.RunArgs(os.Args) }

func (v *Clip) RunArgs(args []string) {
	if v.Name == "" && len(args) > 0 {
		v.Name = args[0]
	}

	if v.entrypoint == nil {
		panic("No entrypoint set. Use clip.Entrypoint to define in the New function")
	}

	if In(args[1:], "--help", "-h") {
		v.printHelpExit(intp(0))
		return
	}

	cli := v.parse(args[1:])

	v.entrypoint(cli)
}

func intp(i int) *int { return &i }

func (v *Clip) printHelpExit(code *int) {
	var output []string

	if v.Description != "" {
		output = append(output, fmt.Sprintf("%s: %s\n", v.Name, v.Description))
	} else {
		output = append(output, fmt.Sprintf("%s\n", v.Name))
	}

	fl := append(v.flags, &boolFlag{
		basicFlag: basicFlag[bool]{
			name:        "help",
			shorthand:   "h",
			required:    false,
			description: "Shows this message and exits",
		},
	})

	var usage []string
	line := ""
	prefix := "usage: " + v.Name + " "
	prefixLen := len(prefix)

	for _, f := range fl {
		uName := f.usageText() + " "
		if prefixLen+len(line)+len(uName) > 80 && line != "" {
			usage = append(usage, line)
			line = ""
		}
		line += uName
	}
	if v.takeArgs {
		if v.ArgsDescription == "" {
			v.ArgsDescription = "ARGS"
		}
		line += " " + v.ArgsDescription
		usage = append(usage, strings.TrimSpace(line))
		line = ""
	}
	if len(line) >= 0 {
		usage = append(usage, line)
	}

	if len(usage) > 0 {
		padding := strings.Repeat(" ", prefixLen)
		for i, o := range usage {
			if i == 0 {
				output = append(output, prefix+o+"\n")
			} else {
				output = append(output, padding+o+"\n")
			}
		}
		output = append(output, "\n")
	}

	if v.Copyright != "" {
		output = append(output, fmt.Sprintf("Copyright %s\n", v.Copyright))
	}

	if l := len(v.Authors); l == 1 {
		output = append(output, fmt.Sprintf("Author: %s\n", v.Authors[0]))
	} else if l > 1 {
		output = append(output, "Authors:\n")
		for _, a := range v.Authors {
			output = append(output, fmt.Sprintf("\t%s\n", a))
		}
	}

	if v.Version != "" {
		if v.Build != "" {
			output = append(output, fmt.Sprintf("Version: %s (%s)\n", v.Version, v.Build))
		} else {
			output = append(output, fmt.Sprintf("Version: %s\n", v.Version))
		}
	}

	output = append(output, "\n")
	output = append(output, "Flags:\n")

	for _, f := range fl {
		prefix := "    " + f.usageText() + " "
		var lines []string
		maxLen := 80 - len(prefix)
		line = ""
		for _, w := range strings.Split(f.descriptionText(), " ") {
			if len(line)+len(w) > maxLen {
				lines = append(lines, line)
				line = ""
			}
			line += w + " "
		}
		if len(line) > 0 {
			lines = append(lines, line)
		}
		if envs := f.readsFromEnv(); len(envs) > 0 {
			lines = append(lines, fmt.Sprintf("Reads from environment: %s", strings.Join(envs, " ")))
		}
		if file := f.readsFromFile(); file != nil {
			lines = append(lines, fmt.Sprintf("Reads from %s, if present.", *file))
		}

		padding := strings.Repeat(" ", len(prefix))
		for i, l := range lines {
			o := ""
			if i == 0 {
				o += prefix
			} else {
				o += padding
			}
			o += l
			output = append(output, o+"\n")
		}
	}

	fmt.Println(strings.Join(output, ""))

	if code != nil {
		os.Exit(*code)
	}
}

func (v *Clip) parse(rawArgs []string) *CLI {
	flags := make([]*flag, 0, len(v.flags))
	for _, f := range v.flags {
		flags = append(flags, f.intoFlag())
	}

	var errs []string
	args := parseArguments(rawArgs)
	var otherArgs []string

	for !args.end() {
		arg := args.next()
		switch arg.kind {
		case argumentKindLong:
			set := false
			for _, f := range flags {
				if f.name == arg.value {
					if err := v.tryTakeValue(f, args); err != nil {
						errs = append(errs, err.Error())
					}
					set = true
					break
				}
			}
			if !set {
				errs = append(errs, fmt.Sprintf("Unknown option --%s", arg.value))
			}
		case argumentKindShort:
			set := false
			for _, f := range flags {
				if f.shorthand == arg.value {
					if err := v.tryTakeValue(f, args); err != nil {
						errs = append(errs, err.Error())
					}
					set = true
					break
				}
			}
			if !set {
				errs = append(errs, fmt.Sprintf("Unknown option -%s", arg.value))
			}
		case argumentKindValue:
			otherArgs = append(otherArgs, arg.value)
		}
	}

loop:
	for _, f := range flags {
		if f.value == nil && len(f.fromEnvs) > 0 {
			for _, k := range f.fromEnvs {
				val, ok := os.LookupEnv(k)
				if ok {
					if err := f.setter(val); err != nil {
						errs = append(errs, fmt.Sprintf("Invalid value %q for %s: %s", val, k, err))
					}
					continue loop
				}
			}
		}
	}

	for _, f := range flags {
		if f.value == nil && len(f.fromFile) > 0 {
			data, err := os.ReadFile(f.fromFile)
			if err != nil {
				continue
			}
			_ = f.setter(string(data))
		}
	}

	for _, f := range flags {
		if f.value == nil && f.defaultValue != nil {
			f.value = f.defaultValue
		}

		if f.value == nil && f.required {
			errs = append(errs, fmt.Sprintf("Required flag --%s is missing", f.name))
			continue
		} else if f.value == nil && !f.required {
			continue
		}

		if f.isFS {
			path := f.value.(string)
			path, err := filepath.Abs(path)
			if err != nil {
				errs = append(errs, fmt.Sprintf("Invalid value %q for %s: %s", path, f.name, err.Error()))
				continue
			}

			stat, err := os.Stat(path)
			if os.IsNotExist(err) {
				if f.mustExist {
					errs = append(errs, fmt.Sprintf("Path %q, provided to %s does not exist", path, f.name))
				}
				continue
			} else if err != nil {
				// FIXME: Not sure what to do in this case.
				continue
			}

			if stat.IsDir() && !f.isDir {
				errs = append(errs, fmt.Sprintf("Path %q, provided to %s is a directory. Expected a file.", path, f.name))
			} else if !stat.IsDir() && f.isDir {
				errs = append(errs, fmt.Sprintf("Path %q, provided to %s is a file. Expected a directory.", path, f.name))
			}

			f.value = path
		}
	}

	if len(errs) > 0 {
		fmt.Println(strings.Join(errs, "\n"))
		v.printHelpExit(intp(2))
		return nil
	}

	cli := cliFromFlags(flags, otherArgs)
	cli.helpFn = v.printHelpExit
	return cli
}

func cliFromFlags(flags []*flag, otherArgs []string) *CLI {
	return &CLI{
		flags:     flags,
		otherArgs: otherArgs,
	}
}

func (v *Clip) tryTakeValue(into *flag, args *argSet) error {
	if into.takesValue {
		v, ok := args.takeValue()
		if !ok {
			return errors.New(fmt.Sprintf("Missing value for argument %s", into.name))
		}
		return into.setter(v.value)
	}

	return into.setter("")
}

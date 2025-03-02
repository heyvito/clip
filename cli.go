package clip

import (
	"fmt"
	"slices"
)

type CLI struct {
	flags     []*flag
	otherArgs []string
	helpFn    func(code *int)
}

func (c *CLI) PrintHelp() {
	c.helpFn(nil)
}

func (c *CLI) PrintHelpExit(code int) {
	c.helpFn(intp(code))
}

func (c *CLI) IsSet(name string) bool {
	return slices.ContainsFunc(c.flags, func(flag *flag) bool {
		return flag.name == name && flag.value != nil
	})
}

func (c *CLI) find(name string) *flag {
	fi := slices.IndexFunc(c.flags, func(flag *flag) bool {
		return flag.name == name
	})
	if fi < 0 {
		return nil
	}
	return c.flags[fi]
}

func (c *CLI) String(name string) string {
	f := c.find(name)
	if f == nil {
		return ""
	}
	switch fv := f.value.(type) {
	case nil:
		return ""
	case string:
		return fv
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", fv)
	case map[string]string:
		panic("Cannot get value of " + name + " as string: It's a KV.")
	case []string:
		panic("Cannot get value of " + name + " as string: It's a slice.")
	case bool:
		return fmt.Sprintf("%t", fv)
	default:
		panic(fmt.Sprintf("Cannot get value of %s as string: It's a %T", name, fv))
	}
}

func (c *CLI) StringSlice(name string) []string {
	f := c.find(name)
	if f == nil {
		return nil
	}
	switch fv := f.value.(type) {
	case nil:
		return nil
	case []string:
		return fv
	case map[string]string:
		panic("Cannot get value of " + name + " as string: It's a KV.")
	default:
		panic(fmt.Sprintf("Cannot get value of %s as string slice: It's a %T", name, fv))
	}
}
func (c *CLI) KV(name string) map[string]string {
	f := c.find(name)
	if f == nil {
		return nil
	}
	switch fv := f.value.(type) {
	case nil:
		return nil
	case map[string]string:
		return fv
	default:
		panic(fmt.Sprintf("Cannot get value of %s as string: It's a %T", name, fv))
	}
}

func (c *CLI) Option(name string) string { return c.String(name) }

func (c *CLI) Boolean(name string) bool {
	f := c.find(name)
	if f == nil {
		return false
	}
	switch fv := f.value.(type) {
	case nil:
		return false
	case bool:
		return fv
	case map[string]string:
		panic("Cannot get value of " + name + " as string: It's a KV.")
	default:
		panic(fmt.Sprintf("Cannot get value of %s as bool: It's a %T", name, fv))
	}
}

func (c *CLI) FetchString(name string) (string, bool) {
	f := c.find(name)
	if f == nil {
		return "", false
	}
	switch fv := f.value.(type) {
	case nil:
		return "", false
	case string:
		return fv, true
	case map[string]string:
		panic("Cannot get value of " + name + " as string: It's a KV.")
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", fv), true
	case []string:
		panic("Cannot get value of " + name + " as string: It's a slice.")
	case bool:
		return fmt.Sprintf("%t", fv), true
	default:
		panic(fmt.Sprintf("Cannot get value of %s as string: It's a %T", name, fv))
	}
}

func (c *CLI) FetchStringSlice(name string) ([]string, bool) {
	f := c.find(name)
	if f == nil {
		return nil, false
	}
	switch fv := f.value.(type) {
	case nil:
		return nil, false
	case []string:
		return fv, true
	case map[string]string:
		panic("Cannot get value of " + name + " as string: It's a slice.")
	default:
		panic(fmt.Sprintf("Cannot get value of %s as string slice: It's a %T", name, fv))
	}
}

func (c *CLI) FetchKV(name string) (map[string]string, bool) {
	f := c.find(name)
	if f == nil {
		return nil, false
	}
	switch fv := f.value.(type) {
	case nil:
		return nil, false
	case map[string]string:
		return fv, true
	default:
		panic(fmt.Sprintf("Cannot get value of %s as string: It's a %T", name, fv))
	}
}

func (c *CLI) FetchOption(name string) (string, bool) { return c.FetchString(name) }

func (c *CLI) FetchBoolean(name string) (bool, bool) {
	f := c.find(name)
	if f == nil {
		return false, false
	}
	switch fv := f.value.(type) {
	case nil:
		return false, false
	case bool:
		return fv, true
	case map[string]string:
		panic("Cannot get value of " + name + " as bool: It's a KV.")
	case []string:
		panic("Cannot get value of " + name + " as bool: It's a slice.")
	default:
		panic(fmt.Sprintf("Cannot get value of %s as bool: It's a %T", name, fv))
	}
}

func (c *CLI) NArgs() int { return len(c.otherArgs) }

func (c *CLI) Args() []string { return c.otherArgs }

func (c *CLI) Arg(idx int) string {
	if idx >= c.NArgs() {
		return ""
	}
	return c.otherArgs[idx]
}

package clip

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type flag struct {
	name         string
	shorthand    string
	required     bool
	defaultValue any
	value        any
	setter       func(string) error
	fromEnvs     []string
	fromFile     string
	takesValue   bool
	validator    func(any) error

	// only set for strings
	isFS      bool
	isDir     bool
	mustExist bool
}

func makeUsageText(name, shorthand, usageName string, required, takesValue bool) string {
	if usageName == "" && takesValue {
		usageName = "value"
	}

	v := []string{
		strings.TrimSpace("--" + name + " " + usageName),
	}
	if shorthand != "" {
		v = append(v, strings.TrimSpace("-"+shorthand+" "+usageName))
	}

	r := strings.Join(v, " | ")
	if !required {
		r = "[" + r + "]"
	}
	return r
}

type basicFlag[T any] struct {
	name        string
	shorthand   string
	required    bool
	description string
	defaultVal  T
	fromEnvs    []string
	fromFile    string
	validatorFn func(T) error
	usageName   string
}

func (f *basicFlag[T]) readsFromEnv() []string {
	return f.fromEnvs
}

func (f *basicFlag[T]) readsFromFile() *string {
	if f.fromFile != "" {
		return &f.fromFile
	}
	return nil
}

func (f *basicFlag[T]) makeFlag() *flag {
	a := &flag{
		name:         f.name,
		shorthand:    f.shorthand,
		required:     f.required,
		defaultValue: f.defaultVal,
		value:        nil,
		setter:       nil,
		fromEnvs:     f.fromEnvs,
		fromFile:     f.fromFile,
	}
	if f.validatorFn != nil {
		a.validator = func(a any) error {
			if v, ok := a.(T); ok {
				return f.validatorFn(v)
			}
			return nil
		}
	}
	return a
}

type stringFlag struct {
	basicFlag[string]
	isFs      bool
	isDir     bool
	mustExist bool
	value     string
}

func (s *stringFlag) UsageName(n string) StringFlag {
	s.usageName = n
	return s
}

func (s *stringFlag) Shorthand(v string) StringFlag {
	s.shorthand = v
	return s
}

func (s *stringFlag) Required() StringFlag {
	s.required = true
	return s
}

func (s *stringFlag) Description(v string) StringFlag {
	s.description = v
	return s
}

func (s *stringFlag) Default(v string) StringFlag {
	s.defaultVal = v
	return s
}

func (s *stringFlag) FromEnv(ss ...string) StringFlag {
	s.fromEnvs = append(s.fromEnvs, ss...)
	return s
}

func (s *stringFlag) FromFile(v string) StringFlag {
	s.fromFile = v
	return s
}

func (s *stringFlag) Validate(valFn func(string) error) StringFlag {
	s.validatorFn = valFn
	return s
}

func (s *stringFlag) File() FSItemFlag {
	s.isFs = true
	return s
}

func (s *stringFlag) Directory() FSItemFlag {
	s.isFs = true
	s.isDir = true
	return s
}

func (s *stringFlag) MustExist() FSItemFlag {
	s.mustExist = true
	return s
}

func (s *stringFlag) intoFlag() *flag {
	f := s.makeFlag()
	f.takesValue = true
	f.setter = func(v string) error {
		f.value = v
		return nil
	}
	f.isFS = s.isFs
	f.isDir = s.isDir
	f.mustExist = s.mustExist
	return f
}

func (s *stringFlag) takesValue() bool { return true }

func (s *stringFlag) usageText() string {
	return makeUsageText(s.name, s.shorthand, s.usageName, s.required, s.takesValue())
}

func (s *stringFlag) descriptionText() string { return s.description }

type numberFlag[V number] struct {
	basicFlag[V]
}

func (n *numberFlag[V]) UsageName(s string) NumberFlag[V] {
	n.usageName = s
	return n
}

func (n *numberFlag[V]) Shorthand(s string) NumberFlag[V] {
	n.shorthand = s
	return n
}

func (n *numberFlag[V]) Required() NumberFlag[V] {
	n.required = true
	return n
}

func (n *numberFlag[V]) Description(s string) NumberFlag[V] {
	n.description = s
	return n
}

func (n *numberFlag[V]) Default(v V) NumberFlag[V] {
	n.defaultVal = v
	return n
}

func (n *numberFlag[V]) FromEnv(s ...string) NumberFlag[V] {
	n.fromEnvs = append(n.fromEnvs, s...)
	return n
}

func (n *numberFlag[V]) FromFile(s string) NumberFlag[V] {
	n.fromFile = s
	return n
}

func (n *numberFlag[V]) Validate(valFn func(V) error) NumberFlag[V] {
	n.validatorFn = valFn
	return n
}

func (n *numberFlag[V]) takesValue() bool { return true }

func (n *numberFlag[V]) usageText() string {
	return makeUsageText(n.name, n.shorthand, n.usageName, n.required, n.takesValue())
}

func (n *numberFlag[V]) descriptionText() string { return n.description }

func (n *numberFlag[V]) intoFlag() *flag {
	f := n.makeFlag()
	f.takesValue = true

	numberKind := reflect.TypeFor[V]().Kind()
	f.setter = func(val string) error {
		switch numberKind {
		case reflect.Int:
			v, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			f.value = v
		case reflect.Int8:
			v, err := strconv.ParseInt(val, 10, 8)
			if err != nil {
				return err
			}
			f.value = int8(v)
		case reflect.Int16:
			v, err := strconv.ParseInt(val, 10, 16)
			if err != nil {
				return err
			}
			f.value = int16(v)
		case reflect.Int32:
			v, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return err
			}
			f.value = int32(v)
		case reflect.Int64:
			v, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			f.value = int64(v)
		case reflect.Uint:
			v, err := strconv.ParseUint(val, 10, strconv.IntSize)
			if err != nil {
				return err
			}
			f.value = uint(v)
		case reflect.Uint8:
			v, err := strconv.ParseUint(val, 10, 8)
			if err != nil {
				return err
			}
			f.value = uint8(v)
		case reflect.Uint16:
			v, err := strconv.ParseUint(val, 10, 16)
			if err != nil {
				return err
			}
			f.value = uint16(v)
		case reflect.Uint32:
			v, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				return err
			}
			f.value = uint32(v)
		case reflect.Uint64:
			v, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return err
			}
			f.value = uint64(v)
		}
		return nil
	}
	return f
}

type boolFlag struct {
	basicFlag[bool]
}

func (b *boolFlag) Shorthand(s string) BoolFlag {
	b.shorthand = s
	return b
}

func (b *boolFlag) Required() BoolFlag {
	b.required = true
	return b
}

func (b *boolFlag) Description(s string) BoolFlag {
	b.description = s
	return b
}

func (b *boolFlag) Default(v bool) BoolFlag {
	b.defaultVal = v
	return b
}

func (b *boolFlag) FromEnv(s ...string) BoolFlag {
	b.fromEnvs = append(b.fromEnvs, s...)
	return b
}

func (b *boolFlag) FromFile(s string) BoolFlag {
	b.fromFile = s
	return b
}

func (b *boolFlag) Validate(valFn func(bool) error) BoolFlag {
	b.validatorFn = valFn
	return b
}

func (b *boolFlag) takesValue() bool { return false }

func (b *boolFlag) usageText() string {
	return makeUsageText(b.name, b.shorthand, b.usageName, b.required, b.takesValue())
}

func (b *boolFlag) descriptionText() string { return b.description }

func (b *boolFlag) intoFlag() *flag {
	f := b.makeFlag()
	f.takesValue = false
	f.setter = func(val string) error { f.value = val != "false"; return nil }
	return f
}

type stringSliceFlag struct {
	basicFlag[[]string]
	separator string
}

func (s *stringSliceFlag) Separator(v string) StringSliceFlag {
	s.separator = v
	return s
}

func (s *stringSliceFlag) UsageName(n string) StringSliceFlag {
	s.usageName = n
	return s
}

func (s *stringSliceFlag) Shorthand(v string) StringSliceFlag {
	s.shorthand = v
	return s
}

func (s *stringSliceFlag) Required() StringSliceFlag {
	s.required = true
	return s
}

func (s *stringSliceFlag) Description(desc string) StringSliceFlag {
	s.description = desc
	return s
}

func (s *stringSliceFlag) Default(v []string) StringSliceFlag {
	s.defaultVal = v
	return s
}

func (s *stringSliceFlag) FromEnv(v ...string) StringSliceFlag {
	s.fromEnvs = append(s.fromEnvs, v...)
	return s
}

func (s *stringSliceFlag) FromFile(v string) StringSliceFlag {
	s.fromFile = v
	return s
}

func (s *stringSliceFlag) Validate(valFn func([]string) error) StringSliceFlag {
	s.validatorFn = valFn
	return s
}

func (s *stringSliceFlag) takesValue() bool { return true }

func (s *stringSliceFlag) usageText() string {
	v := makeUsageText(s.name, s.shorthand, s.usageName, s.required, s.takesValue())
	if s.required {
		v += "..."
	} else {
		v = "(" + v + ")..."
	}
	return v
}

func (s *stringSliceFlag) descriptionText() string { return s.description }

func (s *stringSliceFlag) intoFlag() *flag {
	f := s.makeFlag()
	f.takesValue = true
	f.setter = func(val string) error {
		v, ok := f.value.([]string)
		if !ok {
			v = []string{}
		}
		v = append(v, val)
		f.value = v
		return nil
	}
	return f
}

type optionsFlag struct {
	basicFlag[string]
	options []string
}

func (o *optionsFlag) UsageName(n string) OptionsFlag {
	o.usageName = n
	return o
}

func (o *optionsFlag) Shorthand(s string) OptionsFlag {
	o.shorthand = s
	return o
}

func (o *optionsFlag) Required() OptionsFlag {
	o.required = true
	return o
}

func (o *optionsFlag) Description(s string) OptionsFlag {
	o.description = s
	return o
}

func (o *optionsFlag) Default(v string) OptionsFlag {
	o.defaultVal = v
	return o
}

func (o *optionsFlag) FromEnv(s ...string) OptionsFlag {
	o.fromEnvs = append(o.fromEnvs, s...)
	return o
}

func (o *optionsFlag) FromFile(s string) OptionsFlag {
	o.fromFile = s
	return o
}

func (o *optionsFlag) Validate(valFn func(string) error) OptionsFlag {
	o.validatorFn = valFn
	return o
}

func (o *optionsFlag) takesValue() bool { return true }

func (o *optionsFlag) usageText() string {
	return makeUsageText(o.name, o.shorthand, o.usageName, o.required, o.takesValue())
}

func (o *optionsFlag) descriptionText() string { return o.description }

func (o *optionsFlag) intoFlag() *flag {
	f := o.makeFlag()
	f.takesValue = true
	f.setter = func(val string) error {
		if In(o.options, val) {
			f.value = val
			return nil
		}
		return fmt.Errorf("invalid value %q for %q; only the following values are accpted: %s", val, o.name, strings.Join(o.options, ", "))
	}
	return f
}

type kvFlag struct {
	basicFlag[map[string]string]
	kName string
	vName string
}

func (k *kvFlag) UsageName(kName, vValue string) KVFlag {
	k.kName = kName
	k.vName = vValue
	return k
}

func (k *kvFlag) Shorthand(s string) KVFlag {
	k.shorthand = s
	return k
}

func (k *kvFlag) Required() KVFlag {
	k.required = true
	return k
}

func (k *kvFlag) Description(s string) KVFlag {
	k.description = s
	return k
}

func (k *kvFlag) Default(v map[string]string) KVFlag {
	k.defaultVal = v
	return k
}

func (k *kvFlag) FromEnv(s ...string) KVFlag {
	k.fromEnvs = append(k.fromEnvs, s...)
	return k
}

func (k *kvFlag) FromFile(s string) KVFlag {
	k.fromFile = s
	return k
}

func (k *kvFlag) Validate(valFn func(map[string]string) error) KVFlag {
	k.validatorFn = valFn
	return k
}

func (k *kvFlag) takesValue() bool { return true }

func (k *kvFlag) usageText() string {
	kName := k.kName
	if kName == "" {
		kName = "key"
	}
	vName := k.vName
	if vName == "" {
		vName = "value"
	}

	return makeUsageText(k.name, k.shorthand, kName+"="+vName, k.required, true)
}

func (k *kvFlag) descriptionText() string { return k.description }

func (k *kvFlag) invalidValueError(value string) error {
	kName := k.kName
	if kName == "" {
		kName = "key"
	}
	vName := k.vName
	if vName == "" {
		vName = "value"
	}
	return fmt.Errorf("invalid value %q for %s. Must be in the format %s=%s", value, k.name, kName, vName)

}

func (k *kvFlag) intoFlag() *flag {
	f := k.makeFlag()
	f.takesValue = true
	f.setter = func(val string) error {
		if !strings.ContainsRune(val, '=') {
			return k.invalidValueError(val)
		}

		m, ok := f.value.(map[string]string)
		if !ok {
			m = map[string]string{}
		}
		comps := strings.SplitN(val, "=", 2)
		if len(comps) != 2 {
			return k.invalidValueError(val)
		}
		m[comps[0]] = comps[1]
		f.value = m
		return nil
	}
	return f
}

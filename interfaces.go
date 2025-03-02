package clip

type number interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

type baseFlag interface {
	takesValue() bool
	usageText() string
	descriptionText() string
	readsFromEnv() []string
	readsFromFile() *string
	intoFlag() *flag
}

type Flag[T any, V any] interface {
	baseFlag
	Shorthand(string) T
	Required() T
	Description(string) T
	Default(V) T
	FromEnv(...string) T
	FromFile(string) T
	Validate(valFn func(V) error) T
}

type ValueBasedFlag[T any, V any] interface {
	Flag[T, V]
	UsageName(string) T
}

type StringSliceFlag interface {
	ValueBasedFlag[StringSliceFlag, []string]
	Separator(string) StringSliceFlag
}

type StringFlag interface {
	ValueBasedFlag[StringFlag, string]
	File() FSItemFlag
	Directory() FSItemFlag
}

type FSItemFlag interface {
	StringFlag
	MustExist() FSItemFlag
}

type NumberFlag[V number] interface {
	ValueBasedFlag[NumberFlag[V], V]
}

type BoolFlag interface {
	Flag[BoolFlag, bool]
}

type OptionsFlag interface {
	ValueBasedFlag[OptionsFlag, string]
}

type KVFlag interface {
	Flag[KVFlag, map[string]string]
	UsageName(kName, vName string) KVFlag
}

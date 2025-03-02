package clip

type argumentKind int

const (
	argumentKindLong argumentKind = iota
	argumentKindShort
	argumentKindValue
)

type argument struct {
	kind  argumentKind
	value string
}

func parseArguments(args []string) *argSet {
	var a []argument
	for _, v := range args {
		l := len(v)
		if l == 0 || l == 1 {
			a = append(a, argument{
				kind:  argumentKindValue,
				value: v,
			})
			continue
		}

		if v[0] == '-' && v[1] == '-' && l > 2 {
			a = append(a, argument{
				kind:  argumentKindLong,
				value: v[2:],
			})
		} else if v[0] == '-' && v[1] != '-' && l >= 2 {
			a = append(a, argument{
				kind:  argumentKindShort,
				value: v[1:],
			})
		} else {
			a = append(a, argument{
				kind:  argumentKindValue,
				value: v,
			})
		}
	}
	return &argSet{args: a, cur: 0, len: len(a)}
}

type argSet struct {
	args []argument
	cur  int
	len  int
}

func (a *argSet) peek() *argument {
	return &a.args[a.cur]
}

func (a *argSet) next() *argument {
	arg := a.peek()
	a.cur++
	return arg
}

func (a *argSet) end() bool { return a.cur >= a.len }

func (a *argSet) takeValue() (*argument, bool) {
	if a.end() {
		return nil, false
	}
	p := a.peek()
	if p.kind == argumentKindValue {
		return a.next(), true
	}
	return nil, false
}

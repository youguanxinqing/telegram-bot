package box

type Set map[any]bool

func (s Set) Add(o any) {
	s[o] = true
}

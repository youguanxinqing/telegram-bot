package app

type DescMsg struct {
	Name   string // service name
	Detail string // service detail
}

type Server interface {
	Desc() DescMsg
}

type ServerFunc func() DescMsg

func (f ServerFunc) Desc() DescMsg {
	return f()
}

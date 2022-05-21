package help

import (
	"fmt"
	"strings"
	"telegram-bot/internal/app"
)

type Service struct {
	servers []app.Server
}

func New() *Service {
	s := &Service{servers: nil}
	s.AddService(s)
	return s
}

func (srv *Service) Desc() app.DescMsg {
	return app.DescMsg{
		Name:   "help",
		Detail: "帮助信息",
	}
}

func (srv *Service) AddService(others ...app.Server) {
	if others == nil {
		return
	}

	for _, o := range others {
		srv.servers = append(srv.servers, o)
	}
}

// String 返回服务的帮助信息.
func (srv *Service) String() string {
	var buf strings.Builder
	for idx, s := range srv.servers {
		desc := s.Desc()
		buf.WriteString(fmt.Sprintf("%s %s", desc.Name, desc.Detail))

		if idx+1 < len(srv.servers) {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

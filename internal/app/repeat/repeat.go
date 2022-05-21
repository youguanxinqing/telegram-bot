package repeat

import "telegram-bot/internal/app"

type Service struct{}

func (*Service) Desc() app.DescMsg {
	return app.DescMsg{
		Name:   "repeat",
		Detail: "鹦鹉学舌",
	}
}

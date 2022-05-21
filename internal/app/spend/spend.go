package spend

import "telegram-bot/internal/app"

type Service struct{}

func (*Service) Desc() app.DescMsg {
	return app.DescMsg{
		Name:   "spend",
		Detail: "记账",
	}
}

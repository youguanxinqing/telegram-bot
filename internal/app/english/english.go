package english

import "telegram-bot/internal/app"

type Service struct{}

func (*Service) Desc() app.DescMsg {
	return app.DescMsg{
		Name:   "english",
		Detail: "英语单词",
	}
}

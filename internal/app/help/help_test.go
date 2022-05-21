package help

import (
	"github.com/stretchr/testify/assert"
	"telegram-bot/internal/app"
	"testing"
)

func TestService_String(t *testing.T) {
	srv := NewService()
	srv.AddService(app.ServerFunc(func() app.DescMsg {
		return app.DescMsg{Name: "spend", Detail: "记账"}
	}), app.ServerFunc(func() app.DescMsg {
		return app.DescMsg{Name: "english", Detail: "背英语"}
	}))

	assert.Equal(t, "help 帮助信息\nspend 记账\nenglish 背英语", srv.String())
}

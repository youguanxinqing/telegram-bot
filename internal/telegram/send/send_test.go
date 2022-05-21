package send

import (
	"context"
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSend(t *testing.T) {
	target := Target{
		Token:  "fake token string",
		ChatId: 0000000,
	}
	text := "hello world"

	patch1 := gomonkey.ApplyFunc(tgbotapi.NewBotAPI, func(token string) (*tgbotapi.BotAPI, error) {
		return &tgbotapi.BotAPI{
			Token: target.Token,
		}, nil
	})
	defer patch1.Reset()

	patch2 := gomonkey.ApplyMethod(
		reflect.TypeOf(&tgbotapi.BotAPI{}),
		"Send", func(bot *tgbotapi.BotAPI, c tgbotapi.Chattable) (tgbotapi.Message, error) {
			reqMsg, ok := c.(tgbotapi.MessageConfig)
			assert.Equal(t, true, ok, "消息不是 tgbotapi.MessageConfig 类型")

			assert.Equal(t, bot.Token, target.Token, fmt.Sprintf("token!=%s, but %s", target.Token, bot.Token))
			assert.Equal(t, target.ChatId, reqMsg.ChatID, fmt.Sprintf("chatid!=%d, but %d", target.ChatId, reqMsg.ChatID))
			assert.Equal(t, text, reqMsg.Text, fmt.Sprintf("text!=%s, but %s", text, reqMsg.Text))

			return tgbotapi.Message{}, nil
		})
	defer patch2.Reset()

	if err := Send(context.TODO(), target, text); err != nil {
		t.Error(err)
	}
}

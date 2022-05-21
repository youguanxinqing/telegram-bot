package send

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Target 聊天信息.
type Target struct {
	Token  string
	ChatId int64
}

// Send telegram bot 发消息给聊天室.
func Send(ctx context.Context, target Target, text string) error {
	var (
		bot *tgbotapi.BotAPI
		err error
	)
	if bot, err = tgbotapi.NewBotAPI(target.Token); err != nil {
		return fmt.Errorf("new bot err: %v", err)
	}

	msg := tgbotapi.NewMessage(target.ChatId, text)
	if _, err = bot.Send(msg); err != nil {
		return fmt.Errorf("send msg err: %v", err)
	}
	return nil
}

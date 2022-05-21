package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"telegram-bot/internal/app/english"
	"telegram-bot/internal/config"
	"telegram-bot/internal/dispatch"
	"telegram-bot/internal/telegram/message"
	"telegram-bot/internal/telegram/send"
	"time"
)

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "hello world",
	})
}

// BotHandle telegram bot 入口函数.
func BotHandle(c *gin.Context) {
	now := time.Now()

	target := send.Target{
		Token:  config.Basic.BotToken,
		ChatId: config.Basic.ChatId,
	}
	var m message.HookMsg
	if err := c.BindJSON(&m); err != nil {
		var tmpM map[string]interface{}
		_ = c.BindJSON(&tmpM)
		log.Errorf("parse hook msg err: %v, detail: %#v", err, tmpM)
	} else {

		var text string
		if srv, err := dispatch.New(m.Message.Text); err != nil {
			text = err.Error()
		} else {
			text = srv.Do()
		}

		if err = send.Send(context.TODO(), target, text); err != nil {
			log.Errorf("failed to send msg: %+v, err: %v\n", m, err)
		} else {
			log.Infof("finished! cost: %v", time.Since(now))
		}
	}

	c.JSON(200, gin.H{})
}

// EcdictHandle 单词查询服务器.
func EcdictHandle(c *gin.Context) {
	if c.Query("sw") == "" {
		c.JSON(200, []string{})
		return
	}

	srv := english.New()
	ws := strings.Split(c.Query("sw"), ",")
	log.Println(ws)
	stars, err := srv.Search(context.TODO(), ws...)
	log.Println(stars)
	if err != nil {
		c.JSON(http.StatusInternalServerError, []string{err.Error()})
		return
	}

	c.JSON(http.StatusOK, stars)
}

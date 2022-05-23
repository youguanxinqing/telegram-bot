package dispatch

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"regexp"
	"strings"
	"sync"
	"telegram-bot/internal/app/help"
	"telegram-bot/internal/app/word"
	"telegram-bot/internal/config"
	"telegram-bot/internal/ent"
)

var cmdPattern *regexp.Regexp

var (
	once sync.Once
	db   *ent.Client
)

func init() {
	cmdPattern, _ = regexp.Compile(`(\S+)\s+(\S+)\s+(.*)`)
}

func initDB() {
	once.Do(func() {
		var err error
		if db, err = ent.Open("sqlite3", config.Basic.DbDSN); err != nil {
			log.Fatalf("open db err: %v", err)
		}
		if err = db.Schema.Create(context.TODO()); err != nil {
			log.Fatalf("create schema err: %v", err)
		}
	})
}

type cmd struct {
	cmd    string
	subCmd string
	args   string
}

func newCmd(s string) (cmd, error) {
	var c cmd
	subs := cmdPattern.FindStringSubmatch(s)
	if len(subs) != 4 {
		return c, fmt.Errorf("not correct cmd fmt, detail: '%s'\n", s)
	}

	c.cmd = subs[1]
	c.subCmd = subs[2]
	c.args = subs[3]
	return c, nil
}

type Service struct {
	rawCmd string
	cmd    cmd
}

func New(cmd string) (*Service, error) {
	c, err := newCmd(cmd)
	if err != nil {
		return nil, err
	}

	return &Service{
		rawCmd: cmd,
		cmd:    c,
	}, nil
}

func (srv *Service) Do() string {
	switch srv.cmd.cmd {
	case "help":
		return callHelp()
	case "eng":
		return callWord(srv.rawCmd, srv.cmd.subCmd, srv.cmd.args)
	}
	return fmt.Sprintf("not support cmd '%s'", srv.rawCmd)
}

func callHelp() string {
	srv := help.New()
	srv.AddService((*word.Service)(nil))

	return srv.String()
}

func callWord(rawCmd string, action string, args string) string {
	initDB()

	srv := word.New(db)
	switch action {
	case "s", "save":
		if w, err := srv.SaveWord(context.TODO(), args); err != nil {
			return fmt.Sprintf("fail cmd '%s', err: %v", rawCmd, err)
		} else {
			var buf strings.Builder
			buf.WriteString("SAVE OK!")
			buf.WriteString("\n")
			buf.WriteString(w.String())
			return buf.String()
		}
	case "r", "rand":
		if ws, err := srv.RandomWords(context.TODO(), cast.ToInt(args)); err != nil {
			return fmt.Sprintf("fail cmd '%s', err: %v", rawCmd, err)
		} else {
			var buf strings.Builder
			for _, w := range ws {
				buf.WriteString(w.String())
				buf.WriteString("\n")
			}
			return buf.String()
		}
	case "g", "get":
		if w, err := srv.QueryWord(context.TODO(), args); err != nil {
			return fmt.Sprintf("fail cmd '%s', err: %v", rawCmd, err)
		} else {
			var buf strings.Builder
			buf.WriteString(w.String())
			return buf.String()
		}
	case "h", "hide":
		if err := srv.HideWord(context.TODO(), args); err != nil {
			return fmt.Sprintf("fail cmd '%s', err: %v", rawCmd, err)
		} else {
			return fmt.Sprintf("HIDE '%s' OK ", args)
		}
	case "uh", "unhide":
		if err := srv.UnHideWord(context.TODO(), args); err != nil {
			return fmt.Sprintf("fail cmd '%s', err: %v", rawCmd, err)
		} else {
			return fmt.Sprintf("UNHIDE '%s' OK ", args)
		}
	}
	return fmt.Sprintf("not support %s", action)
}

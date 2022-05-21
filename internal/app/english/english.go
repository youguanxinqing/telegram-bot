package english

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sync"
	"telegram-bot/internal/app"
	"telegram-bot/internal/config"
	"telegram-bot/internal/ent"
	"telegram-bot/internal/ent/stardict"
)

var (
	once sync.Once
	db   *ent.Client
)

func initDB() {
	once.Do(func() {
		var err error
		if db, err = ent.Open("sqlite3", config.Basic.StarDictDSN); err != nil {
			log.Fatalf("open db err: %v", err)
		}
		//if err = db.Schema.Create(context.TODO()); err != nil {
		//	log.Fatalf("create schema err: %v", err)
		//}
	})
}

type Service struct {
	db *ent.Client
}

func New() *Service {
	initDB()
	return &Service{db: db}
}

func (*Service) Desc() app.DescMsg {
	return app.DescMsg{
		Name:   "english",
		Detail: "英语单词",
	}
}

func (srv *Service) Search(ctx context.Context, ws ...string) ([]*ent.StarDict, error) {
	x, _ := srv.db.Debug().StarDict.Query().First(ctx)
	log.Println(x)

	return srv.db.Debug().StarDict.Query().Where(stardict.WordEQ(ws[0])).All(ctx)
}

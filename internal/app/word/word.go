package word

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"telegram-bot/internal/app"
	"telegram-bot/internal/ecdict"
	"telegram-bot/internal/ent"
	"telegram-bot/internal/ent/words"
	"telegram-bot/pkg/rand"
)

// DisplayWord 用于对前端展示.
type DisplayWord struct {
	ID        int    `json:"id"`
	Name      string `json:"name" copier:"Word"`
	Explain   string `json:"explain"`
	Phonetic  string `json:"phonetic"`
	Frequency int    `json:"frequency"`
}

func (w DisplayWord) String() string {
	return fmt.Sprintf(`
%d. %s [%s]
%s
`, w.ID, w.Name, w.Phonetic, w.Explain)
}

type Service struct {
	client *ent.Client
}

func New(db *ent.Client) *Service {
	return &Service{client: db}
}

func (*Service) Desc() app.DescMsg {
	return app.DescMsg{
		Name:   "word",
		Detail: "单词本",
	}
}

// SaveWord 创建单词
func (srv *Service) SaveWord(ctx context.Context, word string) (DisplayWord, error) {
	var showWord DisplayWord

	// 1. 查询单词
	queryResult, err := ecdict.QueryWord(ctx, word)
	if err != nil {
		return showWord, fmt.Errorf("query word err: %v", err)
	}

	// 2. 创建词条
	w, err := srv.client.Words.
		Create().
		SetWord(word).
		SetPhonetic(string(queryResult.Phonetic)).
		SetExplain(queryResult.Translation).
		Save(ctx)
	if err != nil {
		return showWord, err
	}

	// 3. 返回词条
	err = copier.Copy(&showWord, &w)
	return showWord, err
}

// HideWord 隐藏单词.
func (srv *Service) HideWord(ctx context.Context, word string) error {
	_, err := srv.client.Words.
		Update().
		Where(words.WordEQ(word)).
		SetIsHide(true).
		Save(ctx)
	return err
}

// HideWords 隐藏单词.
func (srv *Service) HideWords(ctx context.Context, ws ...string) error {
	_, err := srv.client.Words.
		Update().
		Where(words.WordIn(ws...)).
		SetIsHide(true).
		Save(ctx)
	return err
}

// RandomWords 随机展示单词.
func (srv *Service) RandomWords(ctx context.Context, num int) ([]DisplayWord, error) {
	var sws []DisplayWord

	// 1. 找到所有有效 id
	ws, err := srv.client.Words.
		Query().
		Select(words.FieldID).
		Where(words.IsHideEQ(false), words.IsDeleteEQ(false)).
		All(ctx)
	if err != nil {
		return sws, err
	}

	// 2. 需要查询详情的 ids
	var ids []int
	if num > len(ws) {
		// 需要展示的数量 > 可以展示的数量, 全查出来
		ids = lo.Map(ws, func(w *ent.Words, _ int) int {
			return w.ID
		})
	} else {
		// 生成随机 ids
		idxNeed, _ := rand.GenIntsAmong(num, 0, len(ws))
		for i := range idxNeed {
			ids = append(ids, ws[i].ID)
		}
	}

	// 3. 查询数据库
	ws, err = srv.client.Words.Query().Where(words.IDIn(ids...)).All(ctx)
	if err != nil {
		return sws, err
	}
	err = copier.Copy(&sws, &ws)
	return sws, err
}

func (srv *Service) QueryWord(ctx context.Context, word string) (DisplayWord, error) {
	var sw DisplayWord

	w, err := srv.client.Words.
		Query().
		Where(words.WordEQ(word)).
		First(ctx)
	if err != nil {
		return sw, err
	}

	if err = copier.Copy(&sw, w); err != nil {
		return sw, err
	}

	return sw, nil
}

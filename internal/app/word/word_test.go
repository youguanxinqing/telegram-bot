package word

import (
	"context"
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"telegram-bot/internal/ecdict"
	"telegram-bot/internal/ent"
	"telegram-bot/pkg/op"
	"testing"
)

var dbClient *ent.Client

func TestMain(m *testing.M) {
	var err error
	dbClient, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("open db err: %v", err)
	}
	defer dbClient.Close()
	_ = dbClient.Schema.Create(context.TODO())

	m.Run()
}

func testQueryNotExistedWord(t *testing.T) {
	srv := New(dbClient)
	_, err := srv.QueryWord(context.TODO(), "table")
	assert.NotEmpty(t, err, "单词'table'不应该存在")
}

func testQueryExistedWord(t *testing.T) {
	// patch
	patch := gomonkey.ApplyFunc(ecdict.QueryWord, func(_ context.Context, word string) (ecdict.Word, error) {
		return ecdict.Word{
			SW:          "table",
			Translation: "n.桌子",
		}, nil
	})
	defer patch.Reset()

	srv := New(dbClient)
	displayWord, err := srv.SaveWord(context.TODO(), "table")
	assert.Equal(t, nil, err, "创建词条失败")
	assert.Equal(t, DisplayWord{
		ID:      1,
		Name:    "table",
		Explain: "n.桌子",
	}, displayWord, "创建词条异常")

	displayWord2, _ := srv.QueryWord(context.TODO(), "table")
	assert.Equal(t, DisplayWord{
		ID:      1,
		Name:    "table",
		Explain: "n.桌子",
	}, displayWord2, "查询词条异常")
}

func TestService_QueryAndSave(t *testing.T) {
	t.Run("testQueryNotExistedWord", testQueryNotExistedWord)
	t.Run("testQueryExistedWord", testQueryExistedWord)
}

var WordList = []string{
	"table",
	"love",
	"they",
	"sum",
	"tomorrow",
}

func TestService_RandomWords(t *testing.T) {
	patch := gomonkey.ApplyFunc(ecdict.QueryWord, func(_ context.Context, word string) (ecdict.Word, error) {
		return ecdict.Word{
			SW:          word,
			Translation: "假装翻译:" + word,
		}, nil
	})
	defer patch.Reset()

	srv := New(dbClient)

	for _, w := range WordList {
		_, _ = srv.SaveWord(context.TODO(), w)
	}

	t.Run("testRandomWords_Get3", testRandomWords_Get3)
	t.Run("testRandomWords_Get20", testRandomWords_Get20)
}

func testRandomWords_Get3(t *testing.T) {
	srv := New(dbClient)
	ws, _ := srv.RandomWords(context.TODO(), 3)
	assert.Len(t, ws, 3)

	for _, w := range ws {
		assert.True(t, op.In(w.Name, WordList), fmt.Sprintf("%s not in %+v", w.Name, ws))
	}
}

func testRandomWords_Get20(t *testing.T) {
	srv := New(dbClient)
	ws, _ := srv.RandomWords(context.TODO(), 20)
	assert.Len(t, ws, len(WordList))

	for _, w := range ws {
		assert.True(t, op.In(w.Name, WordList), fmt.Sprintf("%s not in %+v", w.Name, ws))
	}
}

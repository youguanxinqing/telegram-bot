package ecdict

import (
	"context"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/zhshch2002/goreq"
	"telegram-bot/pkg/conv"
)

type CollinsLevel string

func (level CollinsLevel) String() string {
	return string(level) + "星"
}

type PhoneticType string

func (tp PhoneticType) String() string {
	return "[" + string(tp) + "]"
}

type Word struct {
	SW          string       `conv:"sw"` // 单词
	Phonetic    PhoneticType // 音标
	Translation string       // 中文翻译
	Collins     CollinsLevel
}

const url = "http://127.0.0.1:5121?sw=%s"

// QueryWord 查询单词.
func QueryWord(ctx context.Context, word string) (Word, error) {
	var w Word

	resp := goreq.Get(fmt.Sprintf(url, word)).Do()
	if resp.Err != nil {
		return w, resp.Err
	}

	var m []map[string]any
	if err := json.Unmarshal(resp.Body, &m); err != nil || len(m) == 0 {
		return w, fmt.Errorf("not found '%s' or exist err: %v", word, err)
	}

	var sw Word
	bs, _ := json.Marshal(m[0])
	if err := conv.BS2Struct(bs, &sw); err != nil {
		return w, err
	}

	return sw, nil
}

package ecdict

import (
	"context"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/zhshch2002/goreq"
	"reflect"
	"testing"
)

func TestQueryWord(t *testing.T) {
	text := `[{"id":"2989309","word":"table","sw":"table","phonetic":"'teibl","definition":"n. a set of data arranged in rows and columns\nn. a piece of furniture having a smooth flat top that is usually supported by one or more vertical legs\nn. a piece of furniture with tableware for a meal laid out on it\nn. a company of people assembled at a table for a meal or game","translation":"n. 桌子, 餐桌, 工作台, 铭文, 表格, 表, 高原, 平地层\nvt. 搁置, 嵌合, 制表, 把...列入议事日程\n[计] 表格, 模拟运算表","pos":"v:1\/n:99","collins":"4","oxford":"1","tag":"zk gk ielts gre","bnc":"410","frq":"539","exchange":"s:tables\/d:tabled\/p:tabled\/i:tabling\/3:tables","detail":null,"audio":""}]`

	var request *goreq.Request
	patch := gomonkey.ApplyMethod(reflect.TypeOf(request), "Do", func(req *goreq.Request) *goreq.Response {
		resp := new(goreq.Response)
		resp.Text = text
		resp.Body = []byte(resp.Text)
		return resp
	})
	defer patch.Reset()

	w, _ := QueryWord(context.TODO(), "table")
	t.Logf("%v", w)
}

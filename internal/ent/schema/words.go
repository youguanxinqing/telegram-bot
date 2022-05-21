package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type Words struct {
	ent.Schema
}

// Fields of the Words.
func (Words) Fields() []ent.Field {
	return []ent.Field{
		field.String("word").NotEmpty().Unique().Comment("词汇"),
		field.Text("explain").Default("").Comment("解释词汇"),
		field.String("phonetic").Default("").Comment("音标"),
		field.Int("frequency").Default(0).NonNegative().Comment("出现频率"),
		field.Time("last_show_time").Default(time.Now).Comment("最近一次出现时间"),
		field.Bool("is_hide").Default(false).Comment("隐藏单词"),
		field.Time("create_time").Default(time.Now).Comment("创建时间"),
		field.Time("delete_time").Default(time.Now).Comment("删除时间"),
		field.Bool("is_delete").Default(false).Comment("软删除"),
	}
}

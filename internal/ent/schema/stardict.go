package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type StarDict struct {
	ent.Schema
}

func (StarDict) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "stardict"},
	}
}

// Fields of the StarDict.
func (StarDict) Fields() []ent.Field {
	return []ent.Field{
		field.String("word").NotEmpty().Unique().Comment("词汇"),
		field.String("sw").NotEmpty().Comment("词汇"),
		field.String("phonetic").Default("").Comment("音标"),
		field.Text("definition").Default("").Comment("定义"),
		field.Text("translation").Default("").Comment("翻译"),
		field.String("pos").Default("").Comment(""),
		field.Int("collins").Default(0).Comment("柯林斯词汇等级"),
		field.Int("oxford").Default(0).Comment("牛津词汇等级"),
		field.String("tag").Default("").Comment(""),
		field.Int("bnc").Comment(""),
		field.Int("frq").Comment(""),
		field.Text("exchange").Default("").Comment(""),
		field.Text("detail").Default("").Comment(""),
		field.Text("audio").Default("").Comment("音频"),
	}
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"telegram-bot/internal/ent/schema"
	"telegram-bot/internal/ent/stardict"
	"telegram-bot/internal/ent/words"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	stardictFields := schema.StarDict{}.Fields()
	_ = stardictFields
	// stardictDescWord is the schema descriptor for word field.
	stardictDescWord := stardictFields[0].Descriptor()
	// stardict.WordValidator is a validator for the "word" field. It is called by the builders before save.
	stardict.WordValidator = stardictDescWord.Validators[0].(func(string) error)
	// stardictDescSw is the schema descriptor for sw field.
	stardictDescSw := stardictFields[1].Descriptor()
	// stardict.SwValidator is a validator for the "sw" field. It is called by the builders before save.
	stardict.SwValidator = stardictDescSw.Validators[0].(func(string) error)
	// stardictDescPhonetic is the schema descriptor for phonetic field.
	stardictDescPhonetic := stardictFields[2].Descriptor()
	// stardict.DefaultPhonetic holds the default value on creation for the phonetic field.
	stardict.DefaultPhonetic = stardictDescPhonetic.Default.(string)
	// stardictDescDefinition is the schema descriptor for definition field.
	stardictDescDefinition := stardictFields[3].Descriptor()
	// stardict.DefaultDefinition holds the default value on creation for the definition field.
	stardict.DefaultDefinition = stardictDescDefinition.Default.(string)
	// stardictDescTranslation is the schema descriptor for translation field.
	stardictDescTranslation := stardictFields[4].Descriptor()
	// stardict.DefaultTranslation holds the default value on creation for the translation field.
	stardict.DefaultTranslation = stardictDescTranslation.Default.(string)
	// stardictDescPos is the schema descriptor for pos field.
	stardictDescPos := stardictFields[5].Descriptor()
	// stardict.DefaultPos holds the default value on creation for the pos field.
	stardict.DefaultPos = stardictDescPos.Default.(string)
	// stardictDescCollins is the schema descriptor for collins field.
	stardictDescCollins := stardictFields[6].Descriptor()
	// stardict.DefaultCollins holds the default value on creation for the collins field.
	stardict.DefaultCollins = stardictDescCollins.Default.(int)
	// stardictDescOxford is the schema descriptor for oxford field.
	stardictDescOxford := stardictFields[7].Descriptor()
	// stardict.DefaultOxford holds the default value on creation for the oxford field.
	stardict.DefaultOxford = stardictDescOxford.Default.(int)
	// stardictDescTag is the schema descriptor for tag field.
	stardictDescTag := stardictFields[8].Descriptor()
	// stardict.DefaultTag holds the default value on creation for the tag field.
	stardict.DefaultTag = stardictDescTag.Default.(string)
	// stardictDescExchange is the schema descriptor for exchange field.
	stardictDescExchange := stardictFields[11].Descriptor()
	// stardict.DefaultExchange holds the default value on creation for the exchange field.
	stardict.DefaultExchange = stardictDescExchange.Default.(string)
	// stardictDescDetail is the schema descriptor for detail field.
	stardictDescDetail := stardictFields[12].Descriptor()
	// stardict.DefaultDetail holds the default value on creation for the detail field.
	stardict.DefaultDetail = stardictDescDetail.Default.(string)
	// stardictDescAudio is the schema descriptor for audio field.
	stardictDescAudio := stardictFields[13].Descriptor()
	// stardict.DefaultAudio holds the default value on creation for the audio field.
	stardict.DefaultAudio = stardictDescAudio.Default.(string)
	wordsFields := schema.Words{}.Fields()
	_ = wordsFields
	// wordsDescWord is the schema descriptor for word field.
	wordsDescWord := wordsFields[0].Descriptor()
	// words.WordValidator is a validator for the "word" field. It is called by the builders before save.
	words.WordValidator = wordsDescWord.Validators[0].(func(string) error)
	// wordsDescExplain is the schema descriptor for explain field.
	wordsDescExplain := wordsFields[1].Descriptor()
	// words.DefaultExplain holds the default value on creation for the explain field.
	words.DefaultExplain = wordsDescExplain.Default.(string)
	// wordsDescPhonetic is the schema descriptor for phonetic field.
	wordsDescPhonetic := wordsFields[2].Descriptor()
	// words.DefaultPhonetic holds the default value on creation for the phonetic field.
	words.DefaultPhonetic = wordsDescPhonetic.Default.(string)
	// wordsDescFrequency is the schema descriptor for frequency field.
	wordsDescFrequency := wordsFields[3].Descriptor()
	// words.DefaultFrequency holds the default value on creation for the frequency field.
	words.DefaultFrequency = wordsDescFrequency.Default.(int)
	// words.FrequencyValidator is a validator for the "frequency" field. It is called by the builders before save.
	words.FrequencyValidator = wordsDescFrequency.Validators[0].(func(int) error)
	// wordsDescLastShowTime is the schema descriptor for last_show_time field.
	wordsDescLastShowTime := wordsFields[4].Descriptor()
	// words.DefaultLastShowTime holds the default value on creation for the last_show_time field.
	words.DefaultLastShowTime = wordsDescLastShowTime.Default.(func() time.Time)
	// wordsDescIsHide is the schema descriptor for is_hide field.
	wordsDescIsHide := wordsFields[5].Descriptor()
	// words.DefaultIsHide holds the default value on creation for the is_hide field.
	words.DefaultIsHide = wordsDescIsHide.Default.(bool)
	// wordsDescCreateTime is the schema descriptor for create_time field.
	wordsDescCreateTime := wordsFields[6].Descriptor()
	// words.DefaultCreateTime holds the default value on creation for the create_time field.
	words.DefaultCreateTime = wordsDescCreateTime.Default.(func() time.Time)
	// wordsDescDeleteTime is the schema descriptor for delete_time field.
	wordsDescDeleteTime := wordsFields[7].Descriptor()
	// words.DefaultDeleteTime holds the default value on creation for the delete_time field.
	words.DefaultDeleteTime = wordsDescDeleteTime.Default.(func() time.Time)
	// wordsDescIsDelete is the schema descriptor for is_delete field.
	wordsDescIsDelete := wordsFields[8].Descriptor()
	// words.DefaultIsDelete holds the default value on creation for the is_delete field.
	words.DefaultIsDelete = wordsDescIsDelete.Default.(bool)
}

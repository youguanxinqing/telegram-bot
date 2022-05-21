package conv

import (
	json "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

const tag = "conv"

func BS2Struct(bs []byte, o any) error {
	var m map[string]any
	if err := json.Unmarshal(bs, &m); err != nil {
		return err
	}

	ovVal := reflect.ValueOf(o).Elem()
	ovTyp := reflect.TypeOf(o).Elem()

	for i := 0; i < ovVal.NumField(); i++ {
		field := ovVal.Field(i)

		fieldName := ovTyp.Field(i).Name
		fieldTag := ovTyp.Field(i).Tag.Get(tag)
		// 如果存在 tag
		if fieldTag != "" {
			if v, ok := m[fieldTag]; ok {
				field.Set(reflect.ValueOf(castTo(field.Kind(), v)).Convert(field.Type()))
			}
		} else { // 如果没有 tag, 使用 field name, 支持自动转换大小写
			if v, ok := m[fieldName]; ok {
				field.Set(reflect.ValueOf(castTo(field.Kind(), v)).Convert(field.Type()))
			} else if v, ok = m[strings.ToLower(fieldName)]; ok {
				field.Set(reflect.ValueOf(castTo(field.Kind(), v)).Convert(field.Type()))
			}
		}
	}
	return nil
}

func castTo(kind reflect.Kind, v any) any {
	switch kind {
	case reflect.String:
		return cast.ToString(v)
	case reflect.Int:
		return cast.ToInt(v)
	case reflect.Int8:
		return cast.ToInt8(v)
	case reflect.Int16:
		return cast.ToInt16(v)
	case reflect.Int32:
		return cast.ToInt32(v)
	case reflect.Int64:
		return cast.ToInt64(v)
	case reflect.Float32:
		return cast.ToFloat32(v)
	case reflect.Float64:
		return cast.ToFloat64(v)
	default:
		return v
	}
}

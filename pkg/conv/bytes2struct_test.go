package conv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBS2Struct(t *testing.T) {
	jsonStr := `
{"Name": "zhong", "age": 16, "gender": "girl", "boy_friend": "guan"}
`

	var s struct {
		Name      string
		Age       int
		Alias     string
		BoyFriend string `conv:"boy_friend"`
	}

	err := BS2Struct([]byte(jsonStr), &s)
	assert.Equal(t, nil, err)

	t.Logf("%+v", s)
}

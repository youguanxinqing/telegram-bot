package rand

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"telegram-bot/pkg/box"
	"testing"
)

func testGenIntsAmong_Normal(t *testing.T) {
	set := make(box.Set)
	min, max := 126, 200
	rands, _ := GenIntsAmong(10, min, max)
	for _, r := range rands {

		assert.Conditionf(t, func() bool {
			return r >= min && r < max && !set[r]
		}, fmt.Sprintf("r(%d) not in [%d, %d)", r, min, max), rands)

		set.Add(r)
	}
	t.Logf("[%d, %d): %v", min, max, rands)
}

func testGenIntsAmong_Num(t *testing.T) {
	set := make(box.Set)
	min, max := 126, 136
	rands, _ := GenIntsAmong(10, min, max)
	for _, r := range rands {
		assert.Conditionf(t, func() bool {
			return r >= min && r < max && !set[r]
		}, fmt.Sprintf("r(%d) not in [%d, %d)", r, min, max), rands)

		set.Add(r)
	}
	t.Logf("[%d, %d): %v", min, max, rands)
}

func TestGenIntsAmong(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run("testGenIntsAmong_Normal", testGenIntsAmong_Normal)
		t.Run("testGenIntsAmong_Num", testGenIntsAmong_Num)
	}
}

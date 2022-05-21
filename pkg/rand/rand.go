package rand

import (
	"fmt"
	"math/rand"
	"telegram-bot/pkg/op"
	"time"
)

//// GenInts 生成 num 个随机 int.
//func GenInts(num int) []int {
//	ints := make([]int, num)
//	for idx, _ := range ints {
//		rand.Seed(time.Now().UnixNano())
//		ints[idx] = rand.Int()
//	}
//	return ints
//}

// GenIntsAmong 生成 num 个 && 在 [min, max) 范围内的随机 int.
// 生成不相同的随机数, 如果存在相等, 就累加1, 直到不重复为止.
// 要求 [min, max) 总数应该 >= num.
func GenIntsAmong(num int, min, max int) ([]int, error) {
	if num <= 0 || min >= max || max-min < num {
		return nil, fmt.Errorf("require num>0 && max(%d)-min(%d)>=num(%d)", max, min, num)
	}

	// 初始化切片
	ints := make([]int, num)
	for idx, _ := range ints {
		ints[idx] = min - 1
	}

	for idx, _ := range ints {
		rand.Seed(time.Now().UnixNano())
		target := rand.Intn(max-min) + min
		if op.In(target, ints) {
			for {
				target++
				target = target%(max-min) + min
				if !op.In(target, ints) {
					break
				}
			}
		}
		ints[idx] = target
	}
	return ints, nil
}

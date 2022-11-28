package common

import (
	"fmt"
	"strconv"
)

// Division 两个 uint64 做除法返回五位小数点的结果
// up为分数线上，down为分数线下
func Division(up, down uint64) (result float64, err error) {
	result, err = strconv.ParseFloat(fmt.Sprintf("%.5f", float64(up)/float64(down)), 64)
	if err != nil {
		return
	}
	return
}

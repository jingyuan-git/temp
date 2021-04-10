package main

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func fractionToDecimal(numerator int, denominator int) string {
	minus := false

	// 先定义数组 然后再转换
	var res []string

	// 1. 判断正负
	if numerator < 0 {
		minus = !minus
		numerator = -numerator
	}

	if denominator < 0 {
		minus = !minus
		denominator = -denominator
	}

	if minus {
		res = append(res, "-")
	}

	// 整数，小数 integer decimal
	// 但是用余数更贴切
	integer := numerator / denominator
	remainder := numerator % denominator

	if remainder == 0 {
		return strings.Join(res, "") + strconv.Itoa(integer)
	}

	res = append(res, ".")
	// 计算小数是否有循环，通过buf记录每次余数出现时，res的当前的位置
	loopStart := -1
	buf := map[int]int{}

	curPosition := len(res)
	for remainder != 0 {
		if v, ok := buf[remainder]; ok {
			loopStart = v
			break
		} else {
			buf[remainder] = curPosition
			remainder *= 10
			curPosition++
		}


		numerator = remainder / denominator
		remainder = remainder % denominator
		res = append(res, strconv.Itoa(numerator))

	}

	if loopStart != -1 {
		// 塞进去 做括号
		res = append(res[:loopStart], append([]string{"("}, res[loopStart:]...)...)
		res = append(res, ")")
	}
	return strings.Join(res, "")
}

func main() {
	fractionToDecimal()
}
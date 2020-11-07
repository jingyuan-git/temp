package main

import (
	"fmt"
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func divide(dividend int, divisor int) int {
	// 处理溢出: 只有一种情况
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	positive := true
	if dividend < 0 {
		positive = !positive
		dividend = -dividend
	}
	if divisor < 0 {
		positive = !positive
		divisor = -divisor
	}
	if positive {
		return div(dividend, divisor)
	} else {
		return -div(dividend, divisor)
	}
}

func div(dividend, divisor int) int {
	if dividend < divisor {
		return 0
	}
	count := 1
	divisor_temp := divisor
	for dividend > divisor_temp+divisor_temp {
		count += count
		divisor_temp += divisor_temp
	}
	return count + div(dividend-divisor_temp, divisor)
}

func main() {
	fmt.Println(divide(9, 4))
}

//Given two integers dividend and divisor, divide two integers without using mul
//tiplication, division, and mod operator. 
//
// Return the quotient after dividing dividend by divisor. 
//
// The integer division should truncate toward zero, which means losing its frac
//tional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2. 
//
// Note: 
//
// 
// Assume we are dealing with an environment that could only store integers with
//in the 32-bit signed integer range: [−231, 231 − 1]. For this problem, assume th
//at your function returns 231 − 1 when the division result overflows. 
// 
//
// 
// Example 1: 
//
// 
//Input: dividend = 10, divisor = 3
//Output: 3
//Explanation: 10/3 = truncate(3.33333..) = 3.
// 
//
// Example 2: 
//
// 
//Input: dividend = 7, divisor = -3
//Output: -2
//Explanation: 7/-3 = truncate(-2.33333..) = -2.
// 
//
// Example 3: 
//
// 
//Input: dividend = 0, divisor = 1
//Output: 0
// 
//
// Example 4: 
//
// 
//Input: dividend = 1, divisor = 1
//Output: 1
// 
//
// 
// Constraints: 
//
// 
// -231 <= dividend, divisor <= 231 - 1 
// divisor != 0 
// 
// Related Topics 数学 二分查找 
// 👍 453 👎 0

package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
// 递归时注意返回条件
// for循环时结束循环的条件，不能跳不出循环

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
	if positive{
		return div(dividend, divisor)
	} else {
		return -div(dividend, divisor)
	}
}

func div(dividend, divisor int) int {
	if dividend < divisor{
		return 0
	}
	count := 1
	divisor_temp := divisor
	for dividend > divisor_temp + divisor_temp{
		count += count
		divisor_temp += divisor_temp
	}
	return count + div(dividend-divisor_temp, divisor)
}
//leetcode submit region end(Prohibit modification and deletion)

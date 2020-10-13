//Implement atoi which converts a string to an integer. 
//
// The function first discards as many whitespace characters as necessary until 
//the first non-whitespace character is found. Then, starting from this character 
//takes an optional initial plus or minus sign followed by as many numerical digit
//s as possible, and interprets them as a numerical value. 
//
// The string can contain additional characters after those that form the integr
//al number, which are ignored and have no effect on the behavior of this function
//. 
//
// If the first sequence of non-whitespace characters in str is not a valid inte
//gral number, or if no such sequence exists because either str is empty or it con
//tains only whitespace characters, no conversion is performed. 
//
// If no valid conversion could be performed, a zero value is returned. 
//
// Note: 
//
// 
// Only the space character ' ' is considered a whitespace character. 
// Assume we are dealing with an environment that could only store integers with
//in the 32-bit signed integer range: [−231, 231 − 1]. If the numerical value is o
//ut of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is 
//returned. 
// 
//
// 
// Example 1: 
//
// 
//Input: str = "42"
//Output: 42
// 
//
// Example 2: 
//
// 
//Input: str = "   -42"
//Output: -42
//Explanation: The first non-whitespace character is '-', which is the minus sig
//n. Then take as many numerical digits as possible, which gets 42.
// 
//
// Example 3: 
//
// 
//Input: str = "4193 with words"
//Output: 4193
//Explanation: Conversion stops at digit '3' as the next character is not a nume
//rical digit.
// 
//
// Example 4: 
//
// 
//Input: str = "words and 987"
//Output: 0
//Explanation: The first non-whitespace character is 'w', which is not a numeric
//al digit or a +/- sign. Therefore no valid conversion could be performed.
// 
//
// Example 5: 
//
// 
//Input: str = "-91283472332"
//Output: -2147483648
//Explanation: The number "-91283472332" is out of the range of a 32-bit signed 
//integer. Thefore INT_MIN (−231) is returned.
// 
//
// 
// Constraints: 
//
// 
// 0 <= s.length <= 200 
// s consists of English letters (lower-case and upper-case), digits, ' ', '+', 
//'-' and '.'. 
// 
// Related Topics 数学 字符串 
// 👍 868 👎 0


//在这里我罗列几个要点：
//
//根据示例 1，需要去掉前导空格；
//根据示例 2，需要判断第 1 个字符为 + 和 - 的情况，因此，可以设计一个变量 sign，初始化的时候为 1，如果遇到 - ，将 sign 修正为 -1；
//判断是否是数字，可以使用字符的 ASCII 码数值进行比较，即 0 <= c <= '9'；
//根据示例 3 和示例 4 ，在遇到第 1 个不是数字的字符的情况下，转换停止，退出循环；
//根据示例 5，如果转换以后的数字超过了 int 类型的范围，需要截取。这里不能将结果 res 变量设计为 long 类型，注意：由于输入的字符串转换以后也有可能超过 long 类型，因此需要在循环内部就判断是否越界，只要越界就退出循环，这样也可以减少不必要的计算；
//由于涉及下标访问，因此全程需要考虑数组下标是否越界的情况。

package main

import (
	"math"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func myAtoi(s string) int {
	return convertAvoidOverflow(clean(s))

}

func clean(s string) (sign int ,abs string) {
	// 去掉左边的空格
	s = strings.TrimSpace(s)
	//fmt.Println(s)
	if s == ""{
		return
	}
	switch s[0] {
	case '0', '1','2','3','4','5','6','7','8','9':
		sign = 1
		abs = s
	case '+':
		sign = 1
		abs = s[1:]
	case '-':
		sign = -1
		abs = s[1:]
	default:
		return
	}
	//fmt.Println(abs,"abs=")
	for i, val := range abs {
		if val < '0' || val >'9'{
			abs = abs[:i]
			break
		}
	}

	return

}

func convertAvoidOverflow(sign int ,abs string) int {
	//fmt.Println(sign, abs)
	res := 0
	for _, val := range abs{
		//fmt.Println(val - '0')
		res = res * 10 + int(val) - '0'
	}
	if sign * res >math.MaxInt32{
		return math.MaxInt32
	}
	if sign * res < math.MinInt32{
		return math.MaxInt32
	}
	return sign * res
}

//leetcode submit region end(Prohibit modification and deletion)

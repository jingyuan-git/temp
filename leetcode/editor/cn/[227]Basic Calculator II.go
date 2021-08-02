//Given a string s which represents an expression, evaluate this expression and 
//return its value. 
//
// The integer division should truncate toward zero. 
//
// Note: You are not allowed to use any built-in function which evaluates string
//s as mathematical expressions, such as eval(). 
//
// 
// Example 1: 
// Input: s = "3+2*2"
//Output: 7
// Example 2: 
// Input: s = " 3/2 "
//Output: 1
// Example 3: 
// Input: s = " 3+5 / 2 "
//Output: 5
// 
// 
// Constraints: 
//
// 
// 1 <= s.length <= 3 * 105 
// s consists of integers and operators ('+', '-', '*', '/') separated by some n
//umber of spaces. 
// s represents a valid expression. 
// All the integers in the expression are non-negative integers in the range [0,
// 231 - 1]. 
// The answer is guaranteed to fit in a 32-bit integer. 
// 
// Related Topics 栈 数学 字符串 
// 👍 431 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) (res int) {
	// 因此不妨考虑先进行所有乘除运算，
	// 并将这些乘除运算后的整数值放回原表达式的相应位置，
	// 则随后整个表达式的值，就等于一系列整数加减后的值。
	stash := []int{}
	presign := '+'
	num := 0
	for i, v := range s {
		isDigit := v >= '0' && v <= '9'
		if isDigit {
			num = num * 10 + int(v - '0')
		}
		// 每当遇到运算符号的时候进入逻辑，处理栈
		if !isDigit && v != ' ' || i == len(s) - 1 {
			switch presign {
			case '+':
				stash = append(stash, num)
			case '-':
				stash = append(stash, -num)
			case '*':
				stash[len(stash)-1] *= num
			case '/':
				stash[len(stash)-1] /= num
			}
			presign = v
			num = 0
		}
	}
	for _, v := range stash {
		res += v
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

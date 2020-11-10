//Given a string containing just the characters '(' and ')', find the length of 
//the longest valid (well-formed) parentheses substring. 
//
// 
// Example 1: 
//
// 
//Input: s = "(()"
//Output: 2
//Explanation: The longest valid parentheses substring is "()".
// 
//
// Example 2: 
//
// 
//Input: s = ")()())"
//Output: 4
//Explanation: The longest valid parentheses substring is "()()".
// 
//
// Example 3: 
//
// 
//Input: s = ""
//Output: 0
// 
//
// 
// Constraints: 
//
// 
// 0 <= s.length <= 3 * 104 
// s[i] is '(', or ')'. 
// 
// Related Topics 字符串 动态规划 
// 👍 1057 👎 0

package main

// 关键：栈首会记录最长括号对的开始值 -1 的索引
// 由于，从首先从0开始遍历，所以先将 -1 存入
//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	res := 0
	stack := []int{-1}

	for i := 0; i < len(s); i++ {
		if s[i] == '('	{
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i - stack[len(stack)-1])
			}
		}
	}

	return res
}
func max(x, y int) int {
	if x > y {
	 	return x
	} else {
		return y
	}
}
//leetcode submit region end(Prohibit modification and deletion)

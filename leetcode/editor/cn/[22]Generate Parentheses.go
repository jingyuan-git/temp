//Given n pairs of parentheses, write a function to generate all combinations of
// well-formed parentheses. 
//
// 
// Example 1: 
// Input: n = 3
//Output: ["((()))","(()())","(())()","()(())","()()()"]
// Example 2: 
// Input: n = 1
//Output: ["()"]
// 
// 
// Constraints: 
//
// 
// 1 <= n <= 8 
// 
// Related Topics 字符串 回溯算法 
// 👍 1387 👎 0
package main

// 1. （的数量是否大于n
// 2. ）的数量要小于（
//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	var res []string
	var bactrack func(left int, right int, temp string)
	bactrack = func(left int, right int, temp string){
		if len(temp) == n*2{
			res = append(res, temp)
			return
		}
		if left < n {
			bactrack(left+1, right, temp + "(")
		}
		if left > right {
			bactrack(left, right+1, temp + ")")
		}
	}
	bactrack(0,0 , "" )
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

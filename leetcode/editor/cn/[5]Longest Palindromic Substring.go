//Given a string s, return the longest palindromic substring in s. 
//
// 
// Example 1: 
//
// 
//Input: s = "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
// 
//
// Example 2: 
//
// 
//Input: s = "cbbd"
//Output: "bb"
// 
//
// Example 3: 
//
// 
//Input: s = "a"
//Output: "a"
// 
//
// Example 4: 
//
// 
//Input: s = "ac"
//Output: "a"
// 
//
// 
// Constraints: 
//
// 
// 1 <= s.length <= 1000 
// s consist of only digits and English letters (lower-case and/or upper-case), 
//
// 
// Related Topics 字符串 动态规划 
// 👍 2776 👎 0
package main

// TODO: 动态规划 1. 状态转移方程 2. 边界条件
//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	n := len(s)
	if n<2 {
		return s
	}
	var maxLength = 1
	var left int
	dp := make([][]int, n)
	for i := 0; i<n; i++{
		dp[i] = make([]int, n)
	}
	// 对角线的值都为Ture,
	for i := 0; i<n; i++{
		dp[i][i] = 1
	}

	// 填满右上三角的那部分
	for j:=1; j<n; j++{
		for i:=0; i<j; i++{
			if s[i] != s[j]{
				dp[i][j] = 0
			} else {
				if j - i < 3{
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == 1 && (j - i + 1 > maxLength){
				maxLength = j - i + 1
				left = i
			}
		}
	}
	return s[left: left+maxLength]
}
//leetcode submit region end(Prohibit modification and deletion)

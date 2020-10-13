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


//leetcode submit region begin(Prohibit modification and deletion)

// Method1
// TODO: 动态规划 1. 状态转移方程 2. 边界条件
//func longestPalindrome(s string) string {
//	n := len(s)
//	if n<2 {
//		return s
//	}
//	var maxLength = 1
//	var left int
//	dp := make([][]int, n)
//	for i := 0; i<n; i++{
//		dp[i] = make([]int, n)
//	}
//	// 对角线的值都为Ture,
//	for i := 0; i<n; i++{
//		dp[i][i] = 1
//	}
//
//	// 填满右上三角的那部分
//	for j:=1; j<n; j++{
//		for i:=0; i<j; i++{
//			if s[i] != s[j]{
//				dp[i][j] = 0
//			} else {
//				if j - i < 3{
//					dp[i][j] = 1
//				} else {
//					dp[i][j] = dp[i+1][j-1]
//				}
//			}
//			if dp[i][j] == 1 && (j - i + 1 > maxLength){
//				maxLength = j - i + 1
//				left = i
//			}
//		}
//	}
//	return s[left: left+maxLength]
//}


// Method2
// 中心扩展算法

func longestPalindrome(s string) string {
	l, r := 0, 0
	for i:=0; i<len(s); i++ {
		// 奇数扩散
		l1,r1:=expend(s, i, i)
		// 记录最大情况下的左右的值
		if r1-l1 > r-l {
			l, r = l1, r1
		}
		// 偶数扩散
		l2,r2:=expend(s, i, i+1)
		if r2-l2 > r-l {
			l, r = l2, r2
	 	}
	}
	return s[l:r+1]
}


func expend(s string, l, r int) (int, int) {
	for l>=0 && r<len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l+1,r-1
}
//leetcode submit region end(Prohibit modification and deletion)

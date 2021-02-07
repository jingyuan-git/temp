//Given an integer n, return the number of structurally unique BST's (binary sea
//rch trees) which has exactly n nodes of unique values from 1 to n. 
//
// 
// Example 1: 
//
// 
//Input: n = 3
//Output: 5
// 
//
// Example 2: 
//
// 
//Input: n = 1
//Output: 1
// 
//
// 
// Constraints: 
//
// 
// 1 <= n <= 19 
// 
// Related Topics 树 动态规划 
// 👍 989 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	dp := make([]int, n + 1)
	dp[0], dp[1] = 1, 1
	// TODO: 因为dp 0，1已经算好了所以遍历从2开始
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
//leetcode submit region end(Prohibit modification and deletion)

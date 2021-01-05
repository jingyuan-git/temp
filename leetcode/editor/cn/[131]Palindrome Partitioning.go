//Given a string s, partition s such that every substring of the partition is a 
//palindrome. Return all possible palindrome partitioning of s. 
//
// A palindrome string is a string that reads the same backward as forward. 
//
// 
// Example 1: 
// Input: s = "aab"
//Output: [["a","a","b"],["aa","b"]]
// Example 2: 
// Input: s = "a"
//Output: [["a"]]
// 
// 
// Constraints: 
//
// 
// 1 <= s.length <= 16 
// s contains only lowercase English letters. 
// 
// Related Topics 深度优先搜索 动态规划 回溯算法 
// 👍 451 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
var (
	res		[][]string
	path	[]string
)

func isPalindrome(s string) bool {
	for i,j := 0,len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func dfs(start int, s string)  {
	if start == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s[start:i+1]) {
			path = append(path, s[start:i+1])
			dfs(i + 1, s)
			path = path[:len(path) - 1]
		}

	}
}

func partition(s string) [][]string {
	res = [][]string{}
	path = []string{}
	dfs(0, s)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

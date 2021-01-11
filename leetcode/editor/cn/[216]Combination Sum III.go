//Find all valid combinations of k numbers that sum up to n such that the follow
//ing conditions are true: 
//
// 
// Only numbers 1 through 9 are used. 
// Each number is used at most once. 
// 
//
// Return a list of all possible valid combinations. The list must not contain t
//he same combination twice, and the combinations may be returned in any order. 
//
// 
// Example 1: 
//
// 
//Input: k = 3, n = 7
//Output: [[1,2,4]]
//Explanation:
//1 + 2 + 4 = 7
//There are no other valid combinations. 
//
// Example 2: 
//
// 
//Input: k = 3, n = 9
//Output: [[1,2,6],[1,3,5],[2,3,4]]
//Explanation:
//1 + 2 + 6 = 9
//1 + 3 + 5 = 9
//2 + 3 + 4 = 9
//There are no other valid combinations.
// 
//
// Example 3: 
//
// 
//Input: k = 4, n = 1
//Output: []
//Explanation: There are no valid combinations. [1,2,1] is not valid because 1 i
//s used twice.
// 
//
// Example 4: 
//
// 
//Input: k = 3, n = 2
//Output: []
//Explanation: There are no valid combinations.
// 
//
// Example 5: 
//
// 
//Input: k = 9, n = 45
//Output: [[1,2,3,4,5,6,7,8,9]]
//Explanation:
//1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 = 45
//​​​​​​​There are no other valid combinations.
// 
//
// 
// Constraints: 
//
// 
// 2 <= k <= 9 
// 1 <= n <= 60 
// 
// Related Topics 数组 回溯算法 
// 👍 245 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k int, n int) [][]int {
	var (
		temp 	[]int
		res		[][]int
	)
	var dfs func(cur, rest int)
	dfs = func(cur, rest int) {
		if rest == 0 && len(temp) == k {
			res = append(res, append([]int(nil), temp...))
			return
		}

		// TODO: len(temp)+10-cur < k 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
		if rest < 0 || len(temp)+10-cur < k {
			return
		}

		// 不选
		dfs(cur+1, rest)

		// 选
		temp = append(temp, cur)
		dfs(cur+1, rest-cur)
		temp = temp[:len(temp)-1]
	}
	dfs(1, n)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

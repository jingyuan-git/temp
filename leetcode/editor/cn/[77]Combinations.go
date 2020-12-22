//Given two integers n and k, return all possible combinations of k numbers out 
//of 1 ... n. 
//
// You may return the answer in any order. 
//
// 
// Example 1: 
//
// 
//Input: n = 4, k = 2
//Output:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
// 
//
// Example 2: 
//
// 
//Input: n = 1, k = 1
//Output: [[1]]
// 
//
// 
// Constraints: 
//
// 
// 1 <= n <= 20 
// 1 <= k <= n 
// 
// Related Topics 回溯算法 
// 👍 457 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	var res [][]int
	visit := make([]bool, n)

	var dfs func(index int, visit []bool, path []int)
	dfs = func(index int, visit []bool, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			if visit[i] || len(path)>0 && i+1 < path[len(path)-1] {
				continue
			}
			visit[i] = true
			path = append(path, i+1)
			dfs(index+1, visit, path)
			path = path[:len(path)-1]
			visit[i] = false
		}
	}
	dfs(0, visit, []int{})
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

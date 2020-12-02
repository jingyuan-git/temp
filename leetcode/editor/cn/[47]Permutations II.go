//Given a collection of numbers, nums, that might contain duplicates, return all
// possible unique permutations in any order. 
//
// 
// Example 1: 
//
// 
//Input: nums = [1,1,2]
//Output:
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
// 
//
// Example 2: 
//
// 
//Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 8 
// -10 <= nums[i] <= 10 
// 
// Related Topics 回溯算法 
// 👍 534 👎 0

package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) (res [][]int) {
	sort.Ints(nums)
	visited := make([]bool, len(nums))
	var path []int

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}

		for i, v := range nums {
			path = append(path, v)
			visited[i] = true
			dfs(index+1)
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return

	//visited := make(map[int]bool)
	//var path []int
	//var dfs func(index int)
	//dfs = func(index int) {
	//	if index == len(nums) {
	//		res = append(res, append([]int(nil), path...))
	//	}
	//	for i, n := range nums {
	//		if visited[n] || i>0 && nums[i] == nums[i-1] {
	//			continue
	//		}
	//		path = append(path, n)
	//		visited[n] = true
	//		dfs(index+1)
	//		path = path[:len(path)-1]
	//		visited[n] = false
	//	}
	//}
	//dfs(0)
	//return
}
//leetcode submit region end(Prohibit modification and deletion)

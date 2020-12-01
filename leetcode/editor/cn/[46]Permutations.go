//Given an array nums of distinct integers, return all the possible permutations
//. You can return the answer in any order. 
//
// 
// Example 1: 
// Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// Example 2: 
// Input: nums = [0,1]
//Output: [[0,1],[1,0]]
// Example 3: 
// Input: nums = [1]
//Output: [[1]]
// 
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 6 
// -10 <= nums[i] <= 10 
// All the integers of nums are unique. 
// 
// Related Topics 回溯算法 
// 👍 1024 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) (res [][]int) {
	var dfs func(path []int)
	visited := make(map[int]bool)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}
	dfs([]int{})
	return

	//var dfs func(start int, temp []int)
	//dfs = func(start int, temp []int) {
	//	if len(temp) == len(nums) {
	//		res = append(res, append([]int(nil), temp...))
	//	}
	//	for i := start; i < len(nums); i++ {
	//		temp = append(temp, nums[i])
	//		dfs(i, temp)
	//		temp = temp[:len(temp)-1]
	//	}
	//}
	//dfs(0, []int{})
	//return
}
//leetcode submit region end(Prohibit modification and deletion)

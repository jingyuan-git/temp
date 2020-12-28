//Given a collection of integers that might contain duplicates, nums, return all
// possible subsets (the power set). 
//
// Note: The solution set must not contain duplicate subsets. 
//
// Example: 
//
// 
//Input: [1,2,2]
//Output:
//[
//  [2],
//  [1],
//  [1,2,2],
//  [2,2],
//  [1,2],
//  []
//]
// 
// Related Topics 数组 回溯算法 
// 👍 360 👎 0

package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	// TODO: 注意排序！
	sort.Ints(nums)
	res := make([][]int, 0)
	//var path []int
	var dfs func (index int, path []int)
	dfs = func(index int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := index; i < len(nums); i++ {
			// TODO: i > index是精髓
			if i > index && nums[i - 1] == nums[i]{
				continue
			}
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
		return
	}
	dfs(0, []int{})
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

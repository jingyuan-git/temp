//Given an integer array nums, return all possible subsets (the power set). 
//
// The solution set must not contain duplicate subsets. 
//
// 
// Example 1: 
//
// 
//Input: nums = [1,2,3]
//Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 
//
// Example 2: 
//
// 
//Input: nums = [0]
//Output: [[],[0]]
// 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 10 
// -10 <= nums[i] <= 10 
// 
// Related Topics 位运算 数组 回溯算法 
// 👍 920 👎 0

package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	sort.Ints(nums)

}
//leetcode submit region end(Prohibit modification and deletion)

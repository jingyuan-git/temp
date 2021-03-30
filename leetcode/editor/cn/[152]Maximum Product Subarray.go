//Given an integer array nums, find a contiguous non-empty subarray within the a
//rray that has the largest product, and return the product. 
//
// It is guaranteed that the answer will fit in a 32-bit integer. 
//
// A subarray is a contiguous subsequence of the array. 
//
// 
// Example 1: 
//
// 
//Input: nums = [2,3,-2,4]
//Output: 6
//Explanation: [2,3] has the largest product 6.
// 
//
// Example 2: 
//
// 
//Input: nums = [-2,0,-1]
//Output: 0
//Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
// 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 2 * 104 
// -10 <= nums[i] <= 10 
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit 
//integer. 
// 
// Related Topics 数组 动态规划 
// 👍 1016 👎 0
package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	dpMax, dpMin, res := nums[0], nums[0], nums[0]


	for i := 1; i < len(nums); i++ {
		vMin, vMax := dpMin, dpMax
		dpMin = min(min(vMin*nums[i], vMax*nums[i]),nums[i])
		dpMax = max(max(vMin*nums[i], vMax*nums[i]),nums[i])
		res = max(res, dpMax)
	}
	return res
}

func max(x int, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)

//Given an integer array nums, find the contiguous subarray (containing at least
// one number) which has the largest sum and return its sum. 
//
// Follow up: If you have figured out the O(n) solution, try coding another solu
//tion using the divide and conquer approach, which is more subtle. 
//
// 
// Example 1: 
//
// 
//Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
// 
//
// Example 2: 
//
// 
//Input: nums = [1]
//Output: 1
// 
//
// Example 3: 
//
// 
//Input: nums = [0]
//Output: 0
// 
//
// Example 4: 
//
// 
//Input: nums = [-1]
//Output: -1
// 
//
// Example 5: 
//
// 
//Input: nums = [-2147483647]
//Output: -2147483647
// 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 2 * 104 
// -231 <= nums[i] <= 231 - 1 
// 
// Related Topics 数组 分治算法 动态规划 
// 👍 2690 👎 0
package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	dp := math.MinInt32
	res := math.MinInt32
	for _, num := range nums {
		if num > dp + num {
			dp = num
		} else {
			dp = dp + num
		}
		res = max(res, dp)
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)

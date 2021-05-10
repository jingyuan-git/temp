//Given an array, rotate the array to the right by k steps, where k is non-negat
//ive. 
//
// 
// Example 1: 
//
// 
//Input: nums = [1,2,3,4,5,6,7], k = 3
//Output: [5,6,7,1,2,3,4]
//Explanation:
//rotate 1 steps to the right: [7,1,2,3,4,5,6]
//rotate 2 steps to the right: [6,7,1,2,3,4,5]
//rotate 3 steps to the right: [5,6,7,1,2,3,4]
// 
//
// Example 2: 
//
// 
//Input: nums = [-1,-100,3,99], k = 2
//Output: [3,99,-1,-100]
//Explanation: 
//rotate 1 steps to the right: [99,-1,-100,3]
//rotate 2 steps to the right: [3,99,-1,-100]
// 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 105 
// -231 <= nums[i] <= 231 - 1 
// 0 <= k <= 105 
// 
//
// 
// Follow up: 
//
// 
// Try to come up with as many solutions as you can. There are at least three di
//fferent ways to solve this problem. 
// Could you do it in-place with O(1) extra space? 
// 
// Related Topics 数组 
// 👍 957 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int)  {
	le := len(nums)
	k = k % le
	nums1 := append(nums[le - k:], nums[:le - k]...)
	copy(nums, nums1)
}
//leetcode submit region end(Prohibit modification and deletion)

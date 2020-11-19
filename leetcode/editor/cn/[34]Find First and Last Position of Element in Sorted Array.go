//Given an array of integers nums sorted in ascending order, find the starting a
//nd ending position of a given target value. 
//
// If target is not found in the array, return [-1, -1]. 
//
// Follow up: Could you write an algorithm with O(log n) runtime complexity? 
//
// 
// Example 1: 
// Input: nums = [5,7,7,8,8,10], target = 8
//Output: [3,4]
// Example 2: 
// Input: nums = [5,7,7,8,8,10], target = 6
//Output: [-1,-1]
// Example 3: 
// Input: nums = [], target = 0
//Output: [-1,-1]
// 
// 
// Constraints: 
//
// 
// 0 <= nums.length <= 105 
// -109 <= nums[i] <= 109 
// nums is a non-decreasing array. 
// -109 <= target <= 109 
// 
// Related Topics 数组 二分查找 
// 👍 654 👎 0
// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/solution/yi-ge-mo-ban-tong-sha-suo-you-er-fen-cha-zhao-wen-/
package main

//leetcode submit region begin(Prohibit modification and deletion)
func findindex(nums []int, target int, leftbool bool) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > target || (target==nums[mid]&&leftbool){
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func searchRange(nums []int, target int) []int {
	leftIdx := findindex(nums, target, true)
	if leftIdx == len(nums) || nums[leftIdx] != target{
		 return []int{-1, -1}
	}

	return []int{leftIdx, findindex(nums, target, false)-1}
}
//leetcode submit region end(Prohibit modification and deletion)

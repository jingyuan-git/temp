//Given an array of integers numbers that is already sorted in ascending order, 
//find two numbers such that they add up to a specific target number. 
//
// Return the indices of the two numbers (1-indexed) as an integer array answer 
//of size 2, where 1 <= answer[0] < answer[1] <= numbers.length. 
//
// You may assume that each input would have exactly one solution and you may no
//t use the same element twice. 
//
// 
// Example 1: 
//
// 
//Input: numbers = [2,7,11,15], target = 9
//Output: [1,2]
//Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
// 
//
// Example 2: 
//
// 
//Input: numbers = [2,3,4], target = 6
//Output: [1,3]
// 
//
// Example 3: 
//
// 
//Input: numbers = [-1,0], target = -1
//Output: [1,2]
// 
//
// 
// Constraints: 
//
// 
// 2 <= numbers.length <= 3 * 104 
// -1000 <= numbers[i] <= 1000 
// numbers is sorted in increasing order. 
// -1000 <= target <= 1000 
// Only one valid answer exists. 
// 
// Related Topics 数组 双指针 二分查找 
// 👍 501 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
// 双指针
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1

	for left < right {
		if numbers[left] + numbers[right] < target {
			left++
		} else if numbers[left] + numbers[right] > target {
			right--
		} else {
			return []int{left+1, right+1}
		}
	}
	return nil
}
//leetcode submit region end(Prohibit modification and deletion)

/**
https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/solution/liang-shu-zhi-he-ii-shu-ru-you-xu-shu-zu-by-leet-2/
 */
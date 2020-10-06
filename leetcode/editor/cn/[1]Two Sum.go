//Given an array of integers nums and an integer target, return indices of the t
//wo numbers such that they add up to target. 
//
// You may assume that each input would have exactly one solution, and you may n
//ot use the same element twice. 
//
// You can return the answer in any order. 
//
// 
// Example 1: 
//
// 
//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Output: Because nums[0] + nums[1] == 9, we return [0, 1].
// 
//
// Example 2:
//
// 
//Input: nums = [3,2,4], target = 6
//Output: [1,2]
// 
//
// Example 3: 
//
// 
//Input: nums = [3,3], target = 6
//Output: [0,1]
// 
//
// 
// Constraints: 
//
// 
// 2 <= nums.length <= 105 
// -109 <= nums[i] <= 109 
// -109 <= target <= 109 
// Only one valid answer exists. 
// 
// Related Topics 数组 哈希表 
// 👍 9297 👎 0)

package main

// method1
//func twoSum(nums []int, target int) []int {
//	//for i, one := range nums {
//	//	for j := i + 1; j < len(nums); j++ {
//	//		if target == one + nums[j]{
//	//			return []int{i, j}
//	//		}
//	//	}
//	//}
//	//return nil
//
//	for i, x := range nums {
//		for j := i + 1; j < len(nums); j++ {
//			if x+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}

//leetcode submit region begin(Prohibit modification and deletion)

// method1
// 暴力破解
func twoSum(nums []int, target int) []int {
	for i, one := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target == one + nums[j]{
				return []int{i, j}
			}
		}
	}
	return nil
}

// method2
//能够快速寻找数组中是否存在目标元素。如果存在，我们需要找出它的索引。
//func twoSum(nums []int, target int) []int {
//	hashTable := map[int]int{}
//	for i, x := range nums {
//		if p, ok := hashTable[target-x]; ok {
//			return []int{p, i}
//		}
//		hashTable[x] = i
//	}
//	return nil
//}
//leetcode submit region end(Prohibit modification and deletion)



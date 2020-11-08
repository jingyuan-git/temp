//Implement next permutation, which rearranges numbers into the lexicographicall
//y next greater permutation of numbers. 
//
// If such an arrangement is not possible, it must rearrange it as the lowest po
//ssible order (i.e., sorted in ascending order). 
//
// The replacement must be in place and use only constant extra memory. 
//
// 
// Example 1: 
// Input: nums = [1,2,3]
// Output: [1,3,2]
// Example 2: 
// Input: nums = [3,2,1]
// Output: [1,2,3]
// Example 3: 
// Input: nums = [1,1,5]
// Output: [1,5,1]
// Example 4: 
// Input: nums = [1]
// Output: [1]
// 
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 100 
// 0 <= nums[i] <= 100 
// 
// Related Topics 数组 
// 👍 726 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int)  {
	if len(nums) <= 1{
		return
	}
	before, after, littleOverIndex := len(nums) - 2, len(nums) - 1, len(nums) - 1
	// TODO： 注意等号
	for before >= 0 && nums[before] >= nums[after]{
		before--
		after--
	}
	if before >= 0{
		// TODO： 注意等号
		for nums[before] >= nums[littleOverIndex]{
			littleOverIndex--
		}
		nums[before], nums[littleOverIndex] = nums[littleOverIndex], nums[before]
	}
	// reverse A[after:end]
	// 之后部分本来会是升序，现在倒叙调转
	for i, j := after, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
//leetcode submit region end(Prohibit modification and deletion)

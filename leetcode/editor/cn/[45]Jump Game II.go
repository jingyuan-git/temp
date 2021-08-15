//Given an array of non-negative integers nums, you are initially positioned at 
//the first index of the array. 
//
// Each element in the array represents your maximum jump length at that positio
//n. 
//
// Your goal is to reach the last index in the minimum number of jumps. 
//
// You can assume that you can always reach the last index. 
//
// 
// Example 1: 
//
// 
//Input: nums = [2,3,1,1,4]
//Output: 2
//Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 
//step from index 0 to 1, then 3 steps to the last index.
// 
//
// Example 2: 
//
// 
//Input: nums = [2,3,0,1,4]
//Output: 2
// 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 3 * 104 
// 0 <= nums[i] <= 105 
// 
// Related Topics 贪心算法 数组 
// 👍 772 👎 0

// 我们维护当前能够到达的最大下标位置，记为边界。我们从左到右遍历数组，到达边界时，更新边界并将跳跃次数增加 1。
package main
//leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	var maxFar, count int
	var end int // 上次跳跃可达范围右边界（下次的最右起跳点）
	for i := 0; i < len(nums) - 1; i++ {
		maxFar = max(maxFar, i + nums[i])
		// 在本次跳跃的右边界范围内寻找最大的maxFar
		// 从而到达右边界后，重新赋值end
		if end == i {
			count++
			end = maxFar
		}
	}
	return count
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
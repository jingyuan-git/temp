//You are a professional robber planning to rob houses along a street. Each hous
//e has a certain amount of money stashed, the only constraint stopping you from r
//obbing each of them is that adjacent houses have security systems connected and 
//it will automatically contact the police if two adjacent houses were broken into
// on the same night. 
//
// Given an integer array nums representing the amount of money of each house, r
//eturn the maximum amount of money you can rob tonight without alerting the polic
//e. 
//
// 
// Example 1: 
//
// 
//Input: nums = [1,2,3,1]
//Output: 4
//Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//Total amount you can rob = 1 + 3 = 4.
// 
//
// Example 2: 
//
// 
//Input: nums = [2,7,9,3,1]
//Output: 12
//Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 
//(money = 1).
//Total amount you can rob = 2 + 9 + 1 = 12.
// 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 100 
// 0 <= nums[i] <= 400 
// 
// Related Topics 动态规划 
// 👍 1467 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	lth := len(nums)
	if lth == 0 {
		return 0
	}
	sum1, sum2 := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		temp := sum2
		sum2 = max(sum1 + nums[i], sum2)
		sum1 = temp
		//fmt.Println(sum1, sum2)
	}
	return sum2
}

func max(x int, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
//leetcode submit region end(Prohibit modification and deletion)

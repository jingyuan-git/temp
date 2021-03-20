//Given an array nums of n integers and an integer target, find three integers i
//n nums such that the sum is closest to target. Return the sum of the three integ
//ers. You may assume that each input would have exactly one solution.
//
//
// Example 1:
//
//
//Input: nums = [-1,2,1,-4], target = 1
//Output: 2
//Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
//
//
//
// Constraints:
//
//
// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4
//
// Related Topics 数组 双指针
// 👍 608 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var (
		n    = len(nums)
		best = math.MaxInt32
	)

	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			print("right=", right)
			sum := nums[i] + nums[left] + nums[right]
			update(sum)
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}

			if target > sum {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else if target < sum {
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				right--
			}
			fmt.Println(sum, i, left, right)
		}
	}
	return best
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

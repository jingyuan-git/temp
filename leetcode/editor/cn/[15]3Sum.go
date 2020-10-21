//Given an array nums of n integers, are there elements a, b, c in nums such tha
//t a + b + c = 0? Find all unique triplets in the array which gives the sum of ze
//ro. 
//
// Notice that the solution set must not contain duplicate triplets. 
//
// 
// Example 1: 
// Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
// Example 2: 
// Input: nums = []
//Output: []
// Example 3: 
// Input: nums = [0]
//Output: []
// 
// 
// Constraints: 
//
// 
// 0 <= nums.length <= 3000 
// -105 <= nums[i] <= 105 
// 
// Related Topics 数组 双指针 
// 👍 2690 👎 0

package main

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int

	length := len(nums)
	if length <3{
		return res
	}

	for i:=0; i<length; i++{
		// 如果固定位的值已经大于0，因为已经排好序了，后面的两个指针对应的值也肯定大于0，则和不可能为0，所以返回
		if nums[i] > 0 {
			return res
		}
		// 排除值重复的固定位
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i+1
		right := length -1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum == 0:
				// 将结果加入二元组
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 去重，如果l < r且下一个数字一样，则继续挪动
				for left < right && nums[left] == nums[left+1]{
					left ++
				}
				// 同理
				for left < right && nums[right] == nums[right-1]{
					right --
				}
				left ++
				right --
			case sum > 0:
				right --
			case sum < 0:
				left ++
			}

		}

		//for left < right{
		//	if nums[left] + nums[right] + nums[i]==0{
		//		res = append(res, []int{nums[i], nums[left], nums[right]})
		//		break
		//	} else if nums[i] + nums[left] + nums[right] < 0 {
		//		left ++
		//	} else {
		//		right --
		//	}
		//}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
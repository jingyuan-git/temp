//Given a list of non-negative integers nums, arrange them such that they form t
//he largest number. 
//
// Note: The result may be very large, so you need to return a string instead of
// an integer. 
//
// 
// Example 1: 
//
// 
//Input: nums = [10,2]
//Output: "210"
// 
//
// Example 2: 
//
// 
//Input: nums = [3,30,34,5,9]
//Output: "9534330"
// 
//
// Example 3: 
//
// 
//Input: nums = [1]
//Output: "1"
// 
//
// Example 4: 
//
// 
//Input: nums = [10]
//Output: "10"
// 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 100 
// 0 <= nums[i] <= 109 
// 
// Related Topics 排序 
// 👍 691 👎 0
package main

import (
	"sort"
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		lx, ly := 10, 10
		// 注意有等号
		for lx <= y {
			lx *= 10
		}

		for ly <= x {
			ly *= 10
		}

		return lx*x + y > ly*y + x
	})

	/**
	防止全零的情况
	测试用例:[0,0]
	测试结果:"00"
	期望结果:"0"
	 */
	if nums[0] == 0 {
		return "0"
	}

	var ans []byte
	for _, v := range nums {
		//strconv.Itoa(v)... 是为了string转换为byte
		ans = append(ans, strconv.Itoa(v)...)
	}
	return string(ans)
}
//leetcode submit region end(Prohibit modification and deletion)











/**

func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        x, y := nums[i], nums[j]
        sx, sy := 10, 10
        for sx <= x {
            sx *= 10
        }
        for sy <= y {
            sy *= 10
        }
        return sy*x+y > sx*y+x
    })
    if nums[0] == 0 {
        return "0"
    }
    ans := []byte{}
    for _, x := range nums {
        ans = append(ans, strconv.Itoa(x)...)
    }
    return string(ans)
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/largest-number/solution/zui-da-shu-by-leetcode-solution-sid5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */
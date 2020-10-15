//Given n non-negative integers a1, a2, ..., an , where each represents a point
//at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of
// the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x
//-axis forms a container, such that the container contains the most water.
//
// Notice that you may not slant the container.
//
//
// Example 1:
//
//
//Input: height = [1,8,6,2,5,4,8,3,7]
//Output: 49
//Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,
//3,7]. In this case, the max area of water (blue section) the container can conta
//in is 49.
//
//
// Example 2:
//
//
//Input: height = [1,1]
//Output: 1
//
//
// Example 3:
//
//
//Input: height = [4,3,2,1,4]
//Output: 16
//
//
// Example 4:
//
//
//Input: height = [1,2,1]
//Output: 2
//
//
//
// Constraints:
//
//
// 2 <= height.length <= 3 * 104
// 0 <= height[i] <= 3 * 104
//
// Related Topics 数组 双指针
// 👍 1914 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

// Method1 暴力破解
// Method2 双指针
// 两个关键点：
// 1. 左右指针的移动如何选择
// 		移动较短的一边
// 2. 什么情况下，才有必要重新计算面积
//		移动后的边长大于移动前的边长
func maxArea(height []int) int {
	if len(height) < 1 {
		return 0
	}
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		areaNow := calculate(height[left], height[right], right-left)
		area = max(area, areaNow)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func calculate(leftSide int, rightSide int, distance int) int {
	return min(leftSide, rightSide) * distance
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x int, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	array := []int{4, 3, 2, 1, 4}
	fmt.Println(maxArea(array))
}

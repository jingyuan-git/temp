//Given an array of intervals where intervals[i] = [starti, endi], merge all ove
//rlapping intervals, and return an array of the non-overlapping intervals that co
//ver all the intervals in the input. 
//
// 
// Example 1: 
//
// 
//Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
// 
//
// Example 2: 
//
// 
//Input: intervals = [[1,4],[4,5]]
//Output: [[1,5]]
//Explanation: Intervals [1,4] and [4,5] are considered overlapping.
// 
//
// 
// Constraints: 
//
// 
// 1 <= intervals.length <= 104 
// intervals[i].length == 2 
// 0 <= starti <= endi <= 104 
// 
// Related Topics 排序 数组 
// 👍 728 👎 0

package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	left := intervals[0][0]
	right := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		// 当两个元素可以被融合
		if intervals[i][0] <= right {
			right = max(intervals[i][1], right)
		} else {
			// TODO: 添加之前的res （当2个数组已经不能融合）
			res = append(res, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	// 添加当前的res
	res = append(res, []int{left, right})
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)

/*
TODO: 注意测试案例，的边界设置

	nums := [][]int{

		{15,18},
		{2,6},
		{1,3},
		{8,10},

		{1, 4},
		{4, 5},

		{1, 4},
		{2, 3},
	}

*/

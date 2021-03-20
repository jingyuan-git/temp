package main

import (
	"fmt"
	"sort"
)

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

func main() {
	nums := [][]int{
		{15, 18}, /* 第三行索引为 2 */
		{2, 6},   /*  第二行索引为 1 */
		{1, 3},   /*  第一行索引为 0 */
		{8, 10},  /* 第三行索引为 2 */

		//{1, 4},
		//{4, 5},

		//{1, 4},
		//{2, 3},
	}

	fmt.Println(merge(nums))
}

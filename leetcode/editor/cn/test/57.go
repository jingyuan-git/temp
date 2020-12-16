package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	insertBefore := false
	var res [][]int

	for _, interval := range intervals {
		if right < interval[0] {
			if !insertBefore {
				res = append(res, []int{left, right})
				insertBefore = true
			}
			res = append(res, interval)
		} else if left > interval[1] {
			res = append(res, interval)
		} else {
			// 有融合的部分,融合为新的newInterval
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	// 如果之前的迭代都没有完成前插
	// 则一定要将最后的newInterval后插如整个数组
	if !insertBefore {
		res = append(res, []int{left, right})
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums := [][]int{
		//{1, 3},
		//{6, 9},

		//{1, 5},

		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	//newInterval := []int{2, 7}
	newInterval := []int{4, 8}
	fmt.Println(insert(nums, newInterval))
}

package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	dp := math.MinInt32
	res := math.MinInt32
	for _, num := range nums {
		if num > dp+num {
			dp = num
		} else {
			dp = dp + num
		}
		res = max(res, dp)
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

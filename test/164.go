package main

import "fmt"

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	maxValue := max(nums...)

	for exp := 1; exp <= maxValue; exp*=10 {
		cnt := make([]int, 10)

		var digit int
		for _, v := range nums {
			digit = v/exp%10
			cnt[digit]++
		}

		//for i, _ := range cnt[1:] {
		//	cnt[i] += cnt[i-1]
		//}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}

		buf := make([]int, n)
		//for _, v := range nums{
		//	digit = v/exp%10
		//	buf[cnt[digit]-1] = v
		//	cnt[digit]--
		//}
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / exp % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums, buf)
	}

	fmt.Println(nums)
	res := 0
	for i := 1; i < n; i++ {
		res = max(res, nums[i]-nums[i-1])
	}
	return res
}

func max(nums ...int) int {
	res := nums[0]
	for _, v := range nums[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	//input := []int{3,6,9,1, 11, 12, 432, 100}
	input := []int{1,10000000}
	fmt.Println(maximumGap(input))
}

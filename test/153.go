package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (left + right)/2
		if nums[left] <= nums[right] {
			return nums[left]
		} else {
			if nums[left] > nums[mid] {
				right = mid
			} else if nums[mid] > nums[right] {
				left = mid + 1
			}
		}
	}

	return -1
}

func main() {
	num := []int{4}
	fmt.Println(findMin(num))
}
package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// if nums[left]<nums[mid] {
		// TODO："="等号一定要加，第一层比较的是两端，长度不需要动。
		// 第二层判断才需要缩短范围
		if nums[0] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			//} else if nums[right] > nums[mid] {
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{3, 1}
	fmt.Println(search(nums, 1))
}

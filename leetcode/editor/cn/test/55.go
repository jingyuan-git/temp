package main

import "fmt"

func canJump(nums []int) bool {
	maxLength := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if maxLength >= i {
			maxLength = max(maxLength, i+nums[i])
			if maxLength >= length-1 {
				return true
			}
		} else {
			break
		}
	}
	return false
}

func max(x int, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	//fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

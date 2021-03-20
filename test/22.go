package main

import "fmt"

func generateParenthesis(n int) []string {
	var res []string
	var bactrack func(left int, right int, temp string)
	bactrack = func(left int, right int, temp string) {
		if len(temp) == n*2 {
			res = append(res, temp)
			return
		}
		if left < n {
			bactrack(left+1, right, temp+"(")
		}
		if left > right {
			bactrack(left, right+1, temp+")")
		}
	}
	bactrack(0, 0, "")
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}

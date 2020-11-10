package main

import "fmt"

func longestValidParentheses(s string) int {
	res, temp := 0, 0
	stack := []int{-1}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			val := stack[len(stack)-1]
			if val != -1 {
				stack = stack[:len(stack)-1]
				temp += 2
			} else {
				temp = 0
			}
		}
		res = max(res, temp)
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(longestValidParentheses(")()())"))
	//	()(()
	//	会错误
}

func longestValidParentheses(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

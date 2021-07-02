//Evaluate the value of an arithmetic expression in Reverse Polish Notation. 
//
// Valid operators are +, -, *, and /. Each operand may be an integer or another
// expression. 
//
// Note that division between two integers should truncate toward zero. 
//
// It is guaranteed that the given RPN expression is always valid. That means th
//e expression would always evaluate to a result, and there will not be any divisi
//on by zero operation. 
//
// 
// Example 1: 
//
// 
//Input: tokens = ["2","1","+","3","*"]
//Output: 9
//Explanation: ((2 + 1) * 3) = 9
// 
//
// Example 2: 
//
// 
//Input: tokens = ["4","13","5","/","+"]
//Output: 6
//Explanation: (4 + (13 / 5)) = 6
// 
//
// Example 3: 
//
// 
//Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
//Output: 22
//Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//= ((10 * (6 / (12 * -11))) + 17) + 5
//= ((10 * (6 / -132)) + 17) + 5
//= ((10 * 0) + 17) + 5
//= (0 + 17) + 5
//= 17 + 5
//= 22
// 
//
// 
// Constraints: 
//
// 
// 1 <= tokens.length <= 104 
// tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the 
//range [-200, 200]. 
// 
// Related Topics 栈 
// 👍 327 👎 0
package main

import (
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
func evalRPN(tokens []string) int {
	//21.7.2
	stack := []int{}
	for _, v := range tokens {
		int, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, int)
		} else {
			res := 0
			v1, v2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				res = v1 + v2
			case "-":
				res = v1 - v2
			case "*":
				res = v1 * v2
			case "/":
				res = v1 / v2
			}
			stack = append(stack, res)
		}
	}
	return stack[0]


	//var stack []int
	//for _, v := range tokens {
	//	// 1. 转换为数字看是否会报错
	//	i, err := strconv.Atoi(v)
	//	if err == nil {
	//		stack = append(stack, i)
	//	} else {
	//		// val1, val2 顺序不能错
	//		val1, val2 := stack[len(stack)-2], stack[len(stack)-1]
	//		stack = stack[:len(stack)-2]
	//
	//		switch v {
	//		case "+":
	//			stack = append(stack, val1+val2)
	//		case "-":
	//			stack = append(stack, val1-val2)
	//		case "*":
	//			stack = append(stack, val1*val2)
	//		case "/":
	//			stack = append(stack, val1/val2)
	//		}
	//	}
	//}
	//fmt.Println(stack)
	//return stack[0]
}
//leetcode submit region end(Prohibit modification and deletion)

/*

func evalRPN(tokens []string) int {
    stack := []int{}
    for _, token := range tokens {
        val, err := strconv.Atoi(token)
        if err == nil {
            stack = append(stack, val)
        } else {
            num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            switch token {
            case "+":
                stack = append(stack, num1+num2)
            case "-":
                stack = append(stack, num1-num2)
            case "*":
                stack = append(stack, num1*num2)
            default:
                stack = append(stack, num1/num2)
            }
        }
    }
    return stack[0]
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/solution/ni-bo-lan-biao-da-shi-qiu-zhi-by-leetcod-wue9/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */

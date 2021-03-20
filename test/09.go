//Determine whether an integer is a palindrome. An integer is a palindrome when
//it reads the same backward as forward.
//
// Follow up: Could you solve it without converting the integer to a string?
//
//
// Example 1:
//
//
//Input: x = 121
//Output: true
//
//
// Example 2:
//
//
//Input: x = -121
//Output: false
//Explanation: From left to right, it reads -121. From right to left, it becomes
// 121-. Therefore it is not a palindrome.
//
//
// Example 3:
//
//
//Input: x = 10
//Output: false
//Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
//
// Example 4:
//
//
//Input: x = -101
//Output: false
//
//
//
// Constraints:
//
//
// -231 <= x <= 231 - 1
//
// Related Topics 数学
// 👍 1270 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
//func isPalindrome(x int) bool {
//	if x < 0{
//		return false
//	}
//	raw := x
//	reverseVal := 0
//	for x != 0 {
//		reverseVal = reverseVal*10 + x%10
//		x /= 10
//		fmt.Println(x, reverseVal)
//	}
//	return reverseVal==raw
//}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reservedNumber := 0
	for reservedNumber < x {
		reservedNumber = reservedNumber*10 + x%10
		x /= 10
	}
	fmt.Println(reservedNumber, x)
	return reservedNumber == x || reservedNumber/10 == x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	x := 1221
	fmt.Println(isPalindrome(x))
}

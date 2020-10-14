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

//leetcode submit region begin(Prohibit modification and deletion)
//method1
//执行耗时:20 ms,击败了61.85% 的Go用户
//内存消耗:5.1 MB,击败了54.25% 的Go用户
//func isPalindrome(x int) bool {
//	if x < 0{
//		return false
//	}
//	raw := x
//	reverseVal := 0
//	for x != 0 {
//		reverseVal = reverseVal*10 + x%10
//		x /= 10
//	}
//	return reverseVal==raw
//}

// method2
// 反转一半数字
//执行耗时:36 ms,击败了6.71% 的Go用户
//内存消耗:5.2 MB,击败了43.66% 的Go用户

// 链接：https://leetcode-cn.com/problems/palindrome-number/solution/hui-wen-shu-by-leetcode-solution/

func isPalindrome(x int) bool{
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x!=0){
		return false
	}
	reservedNumber := 0
	for reservedNumber < x{
		reservedNumber = reservedNumber*10 + x%10
		x /= 10
	}
	//fmt.Println(reservedNumber,x)

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return reservedNumber == x || reservedNumber/10 == x

}
//leetcode submit region end(Prohibit modification and deletion)


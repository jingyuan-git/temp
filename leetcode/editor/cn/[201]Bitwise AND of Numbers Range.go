//Given two integers left and right that represent the range [left, right], retu
//rn the bitwise AND of all numbers in this range, inclusive. 
//
// 
// Example 1: 
//
// 
//Input: left = 5, right = 7
//Output: 4
// 
//
// Example 2: 
//
// 
//Input: left = 0, right = 0
//Output: 0
// 
//
// Example 3: 
//
// 
//Input: left = 1, right = 2147483647
//Output: 0
// 
//
// 
// Constraints: 
//
// 
// 0 <= left <= right <= 231 - 1 
// 
// Related Topics 位运算 
// 👍 294 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left < right {
		left, right = left >> 1, right >> 1
		shift++
	}
	return left << shift
}
//leetcode submit region end(Prohibit modification and deletion)

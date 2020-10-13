//Given a 32-bit signed integer, reverse digits of an integer. 
//
// Note: 
//Assume we are dealing with an environment that could only store integers withi
//n the 32-bit signed integer range: [−231, 231 − 1]. For the purpose of this prob
//lem, assume that your function returns 0 when the reversed integer overflows. 
//
// 
// Example 1: 
// Input: x = 123
//Output: 321
// Example 2: 
// Input: x = -123
//Output: -321
// Example 3: 
// Input: x = 120
//Output: 21
// Example 4: 
// Input: x = 0
//Output: 0
// 
// 
// Constraints: 
//
// 
// -231 <= x <= 231 - 1 
// 
// Related Topics 数学 
// 👍 2250 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	var res int
	for x != 0{
		if temp := int32(res); temp != (temp*10)/10{
			return 0
		}
		res = res* 10 + x%10
		x /= 10
	}
	return res

	//y := 0
	//for x!=0 {
	//	y = y*10 + x%10
	//	if !( -(1<<31) <= y && y <= (1<<31)-1) {
	//		return 0
	//	}
	//	x /= 10
	//}
	//return y





	// 特别挫的方法
	////var positive bool
	//res := 0
	//
	////if x > 0{
	////	positive = true
	////}else {
	////	positive = false
	////	//x = -x
	////}
	//i := x
	//for i != 0 {
	//	yushu := i % 10
	//	res = res*10 + yushu
	//	i = i/10
	//	if -(1<<31) < res && res > (1<<31)-1 {
	//		return 0
	//	}
	//}
	////if positive{
	////	return res
	////}else {
	////	return -res
	////}
	//return res
}
//leetcode submit region end(Prohibit modification and deletion)

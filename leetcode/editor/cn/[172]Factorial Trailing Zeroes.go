//Given an integer n, return the number of trailing zeroes in n!. 
//
// Follow up: Could you write a solution that works in logarithmic time complexi
//ty? 
//
// 
// Example 1: 
//
// 
//Input: n = 3
//Output: 0
//Explanation: 3! = 6, no trailing zero.
// 
//
// Example 2: 
//
// 
//Input: n = 5
//Output: 1
//Explanation: 5! = 120, one trailing zero.
// 
//
// Example 3: 
//
// 
//Input: n = 0
//Output: 0
// 
//
// 
// Constraints: 
//
// 
// 0 <= n <= 104 
// 
// Related Topics 数学 
// 👍 448 👎 0
package main

// TODO: 如何防止溢出的问题？
//leetcode submit region begin(Prohibit modification and deletion)
func trailingZeroes(n int) int {
	//res := 1
	//for n > 0 {
	//	res *= n
	//	n--
	//}
	//
	//nums := strconv.Itoa(res)
	//fmt.Println(nums)
	//
	//ans := 0
	//for i := len(nums) - 1; i >= 0; i-- {
	//	//fmt.Println(nums[i])
	//	if nums[i] == '0' {
	//		ans++
	//	} else {
	//		break
	//	}
	//}
	//return ans

	//// 依然溢出
	//product := 1
	//ans := 0
	//for n > 0 {
	//	product *= n
	//	fmt.Println(product)
	//	for product%10==0{
	//		product/=10
	//		ans += 1
	//	}
	//	n--
	//}
	//return ans

	// 计算5的个数
	res:=0
	for n>0{
		res+=n/5
		n/=5
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

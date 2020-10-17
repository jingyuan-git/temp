//Roman numerals are represented by seven different symbols: I, V, X, L, C, D an
//d M. 
//
// 
//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000 
//
// For example, 2 is written as II in Roman numeral, just two one's added togeth
//er. 12 is written as XII, which is simply X + II. The number 27 is written as XX
//VII, which is XX + V + II. 
//
// Roman numerals are usually written largest to smallest from left to right. Ho
//wever, the numeral for four is not IIII. Instead, the number four is written as 
//IV. Because the one is before the five we subtract it making four. The same prin
//ciple applies to the number nine, which is written as IX. There are six instance
//s where subtraction is used: 
//
// 
// I can be placed before V (5) and X (10) to make 4 and 9. 
// X can be placed before L (50) and C (100) to make 40 and 90. 
// C can be placed before D (500) and M (1000) to make 400 and 900. 
// 
//
// Given an integer, convert it to a roman numeral. 
//
// 
// Example 1: 
//
// 
//Input: num = 3
//Output: "III"
// 
//
// Example 2: 
//
// 
//Input: num = 4
//Output: "IV"
// 
//
// Example 3: 
//
// 
//Input: num = 9
//Output: "IX"
// 
//
// Example 4: 
//
// 
//Input: num = 58
//Output: "LVIII"
//Explanation: L = 50, V = 5, III = 3.
// 
//
// Example 5: 
//
// 
//Input: num = 1994
//Output: "MCMXCIV"
//Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
// 
//
// 
// Constraints: 
//
// 
// 1 <= num <= 3999 
// 
// Related Topics 数学 字符串 
// 👍 429 👎 0

package main

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)

//// Method 1:
//// 解答成功:
//// 执行耗时:12 ms,击败了59.56% 的Go用户
//// 内存消耗:6.3 MB,击败了5.46% 的Go用户
//func intToRoman(num int) string {
//	res := []string{}
//	var dev int
//	convertMap := map[int]string{
//		1: 	"I",
//		4: 	"IV",
//		5: 	"V",
//		9: 	"IX",
//		10: "X",
//		40: "XL",
//		50: "L",
//		90: "XC",
//		100:"C",
//		400: "CD",
//		500: "D",
//		900: "CM",
//		1000: "M",
//	}
//	convertList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
//	for i := 0; i<len(convertList); i++ {
//		dev = num/convertList[i]
//		num -=  convertList[i] * dev
//		for dev != 0 {
//			res = append(res, convertMap[convertList[i]])
//			dev --
//		}
//
//	}
//	return strings.Join(res,"")
//}

//Method 2:
// 解答成功:
// 执行耗时:12 ms,击败了59.56% 的Go用户
// 内存消耗:6.3 MB,击败了5.46% 的Go用户
func intToRoman(num int) string {
	res := []string{}
	var dev int
	roman := []string{"I", "IV", "V",	"IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	convertList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i<len(convertList); i++ {
		dev = num/convertList[i]
		num -=  convertList[i] * dev
		for dev != 0 {
			res = append(res, roman[convertList[i]])
			dev --
		}

	}
	return strings.Join(res,"")
}
//leetcode submit region end(Prohibit modification and deletion)

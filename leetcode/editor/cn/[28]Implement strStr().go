//Implement strStr(). 
//
// Return the index of the first occurrence of needle in haystack, or -1 if need
//le is not part of haystack. 
//
// Clarification: 
//
// What should we return when needle is an empty string? This is a great questio
//n to ask during an interview. 
//
// For the purpose of this problem, we will return 0 when needle is an empty str
//ing. This is consistent to C's strstr() and Java's indexOf(). 
//
// 
// Example 1: 
// Input: haystack = "hello", needle = "ll"
//Output: 2
// Example 2: 
// Input: haystack = "aaaaa", needle = "bba"
//Output: -1
// Example 3: 
// Input: haystack = "", needle = ""
//Output: 0
// 
// 
// Constraints: 
//
// 
// 0 <= haystack.length, needle.length <= 5 * 104 
// haystack and needle consist of only lower-case English characters. 
// 
// Related Topics 双指针 字符串 
// 👍 615 👎 0

package main

// 注意边界！（for 开始，结束的边界）
//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	hl := len(haystack)
	nl := len(needle)
	if nl <= 0{
		return 0
	}
	if hl <= 0{
		return -1
	}
	//for i := 0; i < hl - nl + 1; i++{
	//	if haystack[i] == needle[0]{
	//		flag := true
	//		for j := 1; j < nl; j++{
	//			if haystack[i+j] != needle[j]{
	//				flag = false
	//				break
	//			}
	//		}
	//		if flag {
	//			return i
	//		}
	//	}
	//}
	for i := 0; i < hl - nl + 1; i++{
		if haystack[i] == needle[0]{
			for j := 0; j < nl; j++{
				if haystack[i+j] != needle[j]{
					break
				}
				if nl == j + 1{
					return i
				}
			}
		}
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)

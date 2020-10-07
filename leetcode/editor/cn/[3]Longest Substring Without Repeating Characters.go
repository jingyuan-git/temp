//Given a string s, find the length of the longest substring without repeating c
//haracters. 
//
// 
// Example 1: 
//
// 
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
// 
//
// Example 2: 
//
// 
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
// 
//
// Example 3: 
//
// 
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a 
//substring.
// 
//
// Example 4: 
//
// 
//Input: s = ""
//Output: 0
// 
//
// 
// Constraints: 
//
// 
// 0 <= s.length <= 5 * 104 
// s consists of English letters, digits, symbols and spaces. 
// 
// Related Topics 哈希表 双指针 字符串 Sliding Window 
// 👍 4408 👎 0
package main

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
// method 1 sliding windows
func lengthOfLongestSubstring(s string) int {
	var s1 string
	var Length int
	left, right := 0, 0
	s1 = s[left:right]

	for ; right < len(s); right++ {
		// 注意在s1中找s[right]，index是滑动窗口中的下标
		// 因此left要在原有基础上增加
		if index := strings.IndexByte(s1, s[right]); index != -1{
				left += index + 1
		}
		s1 = s[left:right+1]
		if len(s1) > Length {
			Length = len(s1)
		}
	}
	return Length
}
//leetcode submit region end(Prohibit modification and deletion)

//Given a string s consists of some words separated by spaces, return the length
// of the last word in the string. If the last word does not exist, return 0. 
//
// A word is a maximal substring consisting of non-space characters only. 
//
// 
// Example 1: 
// Input: s = "Hello World"
//Output: 5
// Example 2: 
// Input: s = " "
//Output: 0
// 
// 
// Constraints: 
//
// 
// 1 <= s.length <= 104 
// s consists of only English letters and spaces ' '. 
// 
// Related Topics 字符串 
// 👍 263 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLastWord(s string) int {
	length := len(s)
	count := 0
	for i := length - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if count != 0 {
			// 如果已经遍历到不为‘ ’的字符，s[i]还依然为空。说明，最后一个单词遍历结束，可以返回
			return count
		}
	}
	// index到了开头，依然没有返回
	return count

	//count := 0
	//for i := len(s) - 1 ;i >= 0 ;i--{
	//	if s[i] != ' '{
	//		count ++
	//	}
	//	if count != 0 && s[i] == ' '{
	//		break
	//	}
	//}
	//
	//return count
}
//leetcode submit region end(Prohibit modification and deletion)

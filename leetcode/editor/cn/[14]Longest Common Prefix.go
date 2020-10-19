//Write a function to find the longest common prefix string amongst an array of 
//strings. 
//
// If there is no common prefix, return an empty string "". 
//
// 
// Example 1: 
//
// 
//Input: strs = ["flower","flow","flight"]
//Output: "fl"
// 
//
// Example 2: 
//
// 
//Input: strs = ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.
// 
//
// 
// Constraints: 
//
// 
// 0 <= strs.length <= 200 
// 0 <= strs[i].length <= 200 
// strs[i] consists of only lower-case English letters. 
// 
// Related Topics 字符串 
// 👍 1313 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {

	if len(strs)==0{
		return ""
	}
	for i := 0; i < len(strs[0]); i++{
		for j := 1; j < len(strs); j++{
			if len(strs[j]) == i || strs[0][i] != strs[j][i]{
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
//leetcode submit region end(Prohibit modification and deletion)

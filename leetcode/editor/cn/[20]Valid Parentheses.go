//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']
//', determine if the input string is valid. 
//
// An input string is valid if: 
//
// 
// Open brackets must be closed by the same type of brackets. 
// Open brackets must be closed in the correct order. 
// 
//
// 
// Example 1: 
//
// 
//Input: s = "()"
//Output: true
// 
//
// Example 2: 
//
// 
//Input: s = "()[]{}"
//Output: true
// 
//
// Example 3: 
//
// 
//Input: s = "(]"
//Output: false
// 
//
// Example 4: 
//
// 
//Input: s = "([)]"
//Output: false
// 
//
// Example 5: 
//
// 
//Input: s = "{[]}"
//Output: true
// 
//
// 
// Constraints: 
//
// 
// 1 <= s.length <= 104 
// s consists of parentheses only '()[]{}'. 
// 
// Related Topics 栈 字符串 
// 👍 1941 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	dic := map[byte]byte{'}': '{', ')': '(', ']': '['}
	stack := []byte{}
	for _, v := range s {
		if vv, ok := dic[byte(v)]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == vv {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, byte(v))
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
//leetcode submit region end(Prohibit modification and deletion)

/**
length := len(s)
if length % 2 == 1{
	return false
}
pairs := map[byte]byte{
	'}':'{',
	')':'(',
	']':'[',
}
var stack []byte
for i:=0; i<length; i++{
	if pairs[s[i]] > 0{
		if len(stack) >0 && pairs[s[i]] == stack[len(stack)-1]{
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	} else {
		stack = append(stack, s[i])
	}
}
return len(stack) == 0
 */
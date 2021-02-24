//Given a string s, determine if it is a palindrome, considering only alphanumer
//ic characters and ignoring cases. 
//
// 
// Example 1: 
//
// 
//Input: s = "A man, a plan, a canal: Panama"
//Output: true
//Explanation: "amanaplanacanalpanama" is a palindrome.
// 
//
// Example 2: 
//
// 
//Input: s = "race a car"
//Output: false
//Explanation: "raceacar" is not a palindrome.
// 
//
// 
// Constraints: 
//
// 
// 1 <= s.length <= 2 * 105 
// s consists only of printable ASCII characters. 
// 
// Related Topics 双指针 字符串 
// 👍 334 👎 0

package main

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	//TODO: 双指针的问题！
	left, right := 0, len(s) - 1
	for left < right {
		for left < right && !isValid(s[left]) {
			left++
		}

		for left < right && !isValid(s[right]) {
			right--
		}

		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func isValid(ch byte) bool {
	return  (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'Z')
}
//leetcode submit region end(Prohibit modification and deletion)

/**

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    left, right := 0, len(s) - 1
    for left < right {
        for left < right && !isalnum(s[left]) {
            left++
        }
        for left < right && !isalnum(s[right]) {
            right--
        }
        if left < right {
            if s[left] != s[right] {
                return false
            }
            left++
            right--
        }
    }
    return true
}

func isalnum(ch byte) bool {
    return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/valid-palindrome/solution/yan-zheng-hui-wen-chuan-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
//Given a string containing digits from 2-9 inclusive, return all possible lette
//r combinations that the number could represent. Return the answer in any order. 
//
//
// A mapping of digit to letters (just like on the telephone buttons) is given b
//elow. Note that 1 does not map to any letters. 
//
// 
//
// 
// Example 1: 
//
// 
//Input: digits = "23"
//Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// 
//
// Example 2: 
//
// 
//Input: digits = ""
//Output: []
// 
//
// Example 3: 
//
// 
//Input: digits = "2"
//Output: ["a","b","c"]
// 
//
// 
// Constraints: 
//
// 
// 0 <= digits.length <= 4 
// digits[i] is a digit in the range ['2', '9']. 
// 
// Related Topics 字符串 回溯算法 
// 👍 965 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	combinations = []string{}
	// 全局变量用时要初始化
	backtrack(digits, 0, "")
	return combinations
}
func backtrack(digits string, index int, combination string){
	if len(combination) == len(digits){
		combinations = append(combinations, combination)
		return
	}
	letters := phoneMap[string(digits[index])]
	for _, letter := range letters{
		backtrack(digits, index+1, combination + string(letter))
	}
}

//leetcode submit region end(Prohibit modification and deletion)

//Given an integer columnNumber, return its corresponding column title as it app
//ears in an Excel sheet. 
//
// For example: 
//
// 
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28 
//...
// 
//
// 
// Example 1: 
//
// 
//Input: columnNumber = 1
//Output: "A"
// 
//
// Example 2: 
//
// 
//Input: columnNumber = 28
//Output: "AB"
// 
//
// Example 3: 
//
// 
//Input: columnNumber = 701
//Output: "ZY"
// 
//
// Example 4: 
//
// 
//Input: columnNumber = 2147483647
//Output: "FXSHRXW"
// 
//
// 
// Constraints: 
//
// 
// 1 <= columnNumber <= 231 - 1 
// 
// Related Topics 数学 
// 👍 331 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func convertToTitle(columnNumber int) string {
	result := []byte{}
	for columnNumber != 0 {
		// columnNumber - 1 这个卡了好久。。。
		remainder := (columnNumber - 1) % 26

		result = append([]byte{'A' + byte(remainder)}, result...)
		columnNumber = (columnNumber - 1) / 26
		//fmt.Println(columnNumber)
	}
	return string(result)
}
//leetcode submit region end(Prohibit modification and deletion)

//Given an integer rowIndex, return the rowIndexth row of the Pascal's triangle.
// 
//
// Notice that the row index starts from 0. 
//
// 
//In Pascal's triangle, each number is the sum of the two numbers directly above
// it. 
//
// Follow up: 
//
// Could you optimize your algorithm to use only O(k) extra space? 
//
// 
// Example 1: 
// Input: rowIndex = 3
//Output: [1,3,3,1]
// Example 2: 
// Input: rowIndex = 0
//Output: [1]
// Example 3: 
// Input: rowIndex = 1
//Output: [1,1]
// 
// 
// Constraints: 
//
// 
// 0 <= rowIndex <= 33 
// 
// Related Topics 数组 
// 👍 260 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		// TODO: 从后往前
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}
//leetcode submit region end(Prohibit modification and deletion)

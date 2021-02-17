//Given a non-negative integer numRows, generate the first numRows of Pascal's t
//riangle. 
//
// 
//In Pascal's triangle, each number is the sum of the two numbers directly above
// it. 
//
// Example: 
//
// 
//Input: 5
//Output:
//[
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
//]
// 
// Related Topics 数组 
// 👍 453 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		res[i] = make([]int, i + 1)
		res[i][0] = 1
		// 头尾赋值
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j]+res[i-1][j-1]
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

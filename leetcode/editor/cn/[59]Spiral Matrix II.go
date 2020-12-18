//Given a positive integer n, generate an n x n matrix filled with elements from
// 1 to n2 in spiral order. 
//
// 
// Example 1: 
//
// 
//Input: n = 3
//Output: [[1,2,3],[8,9,4],[7,6,5]]
// 
//
// Example 2: 
//
// 
//Input: n = 1
//Output: [[1]]
// 
//
// 
// Constraints: 
//
// 
// 1 <= n <= 20 
// 
// Related Topics 数组 
// 👍 273 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	if n <= 0 {
		// 注意return可以是nil
		return nil
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// TODO: 在设定的框框的移动。想要将4个方向的代码合并，没写对
	top, bottom := 0, n-1
	left, right := 0, n-1

	start := 1
	for top <= bottom {
		for i := left; i <= right; i++ {
			res[top][i] = start
			start++
		}
		top++

		for i := top; i <= bottom; i++ {
			res[i][right] = start
			start++
		}
		right--

		for i := right; i >= left; i-- {
			res[bottom][i] = start
			start++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			res[i][left] = start
			start++
		}
		left++
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

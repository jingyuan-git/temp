//Given an m x n matrix, return all elements of the matrix in spiral order. 
//
// 
// Example 1: 
//
// 
//Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
//Output: [1,2,3,6,9,8,7,4,5]
// 
//
// Example 2: 
//
// 
//Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//Output: [1,2,3,4,8,12,11,10,9,5,6,7]
// 
//
// 
// Constraints: 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 10 
// -100 <= matrix[i][j] <= 100 
// 
// Related Topics 数组 
// 👍 562 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		rows, columns = len(matrix), len(matrix[0])
		order = make([]int, rows*columns)
		index = 0
		top, bottom, left, right = 0, rows-1, 0, columns-1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}

		for row := top+1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}

		// TODO：注意if与for的区别！！！
		if left < right && top < bottom {
			for column := right-1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}

			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}

		left++
		top++
		right--
		bottom--
	}
	return order
}
//leetcode submit region end(Prohibit modification and deletion)

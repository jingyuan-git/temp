package main

import "fmt"

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

	//if len(matrix) == 0 || len(matrix[0]) == 0 {
	//	return []int{}
	//}
	//var (
	//	rows, columns = len(matrix), len(matrix[0])
	//	order = make([]int, rows * columns)
	//	index = 0
	//	left, right, top, bottom = 0, columns - 1, 0, rows - 1
	//)
	//
	//for left <= right && top <= bottom {
	//	for column := left; column <= right; column++ {
	//		order[index] = matrix[top][column]
	//		index++
	//	}
	//	for row := top + 1; row <= bottom; row++ {
	//		order[index] = matrix[row][right]
	//		index++
	//	}
	//	if left < right && top < bottom {
	//		for column := right - 1; column > left; column-- {
	//			order[index] = matrix[bottom][column]
	//			index++
	//		}
	//		for row := bottom; row > top; row-- {
	//			order[index] = matrix[row][left]
	//			index++
	//		}
	//	}
	//	left++
	//	right--
	//	top++
	//	bottom--
	//}
	//return order

}

func main() {
	matrix := [][]int{
		{0, 1, 2},  /*  第一行索引为 0 */
		{4, 5, 6},  /*  第二行索引为 1 */
		{8, 9, 10}, /* 第三行索引为 2 */
	}

	fmt.Println(spiralOrder(matrix))
}
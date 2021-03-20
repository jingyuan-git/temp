package main

import "fmt"

func rotate(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		for j := 0; j < (length+1)/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[length-j-1][i]
			matrix[length-j-1][i] = matrix[length-i-1][length-j-1]
			matrix[length-i-1][length-j-1] = matrix[j][length-i-1]
			matrix[j][length-i-1] = temp
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 1, 2},  /*  第一行索引为 0 */
		{4, 5, 6},  /*  第二行索引为 1 */
		{8, 9, 10}, /* 第三行索引为 2 */
	}
	rotate(matrix)
	fmt.Println(matrix)
}

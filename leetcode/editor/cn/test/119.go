package main

import "fmt"

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

func main() {
	fmt.Println(getRow(3))
}

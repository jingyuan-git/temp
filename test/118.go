package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		// 头尾赋值
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res
}

func main() {
	fmt.Println(generate(3))
}

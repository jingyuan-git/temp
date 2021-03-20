package main

import "fmt"

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	row, column := len(board), len(board[0])
	visit := make([][]bool, row)
	for i := 0; i < row; i++ {
		visit[i] = make([]bool, column)
	}

	var dfs func(x int, y int, index int) bool
	dfs = func(x int, y int, index int) bool {
		// TODO: 这里先剪枝！
		if board[x][y] != word[index] {
			return false
		}

		if index == len(word)-1 {
			return true
		}

		//if index == len(word) - 1 {
		//	if board[x][y] == word[index] {
		//		return true
		//	} else {
		//		return false
		//	}
		//}
		visit[x][y] = true
		defer func() { visit[x][y] = false }() // TODO: 回溯时还原已访问的单元格

		for _, v := range directions {
			if x+v.x >= 0 && x+v.x < row && y+v.y >= 0 && y+v.y < column && !visit[x+v.x][y+v.y] {
				if dfs(x+v.x, y+v.y, index+1) {
					return true
				}
			}
		}
		//visit[x][y]  = false
		return false
	}

	// TODO: 寻找出是否匹配已有路径。和回溯例题中列出所有可能算法还是很不一样的
	for i, row := range board {
		for j := range row {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	nums := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	str := "ABCE"

	//nums := [][]byte{
	//	{'a', 'a'},
	//}
	//str := "aaa"
	fmt.Println(exist(nums, str))
}

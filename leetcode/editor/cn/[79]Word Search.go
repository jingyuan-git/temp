//Given an m x n board and a word, find if the word exists in the grid. 
//
// The word can be constructed from letters of sequentially adjacent cells, wher
//e "adjacent" cells are horizontally or vertically neighboring. The same letter c
//ell may not be used more than once. 
//
// 
// Example 1: 
//
// 
//Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
// "ABCCED"
//Output: true
// 
//
// Example 2: 
//
// 
//Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
// "SEE"
//Output: true
// 
//
// Example 3: 
//
// 
//Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
// "ABCB"
//Output: false
// 
//
// 
// Constraints: 
//
// 
// m == board.length 
// n = board[i].length 
// 1 <= m, n <= 200 
// 1 <= word.length <= 103 
// board and word consists only of lowercase and uppercase English letters. 
// 
// Related Topics 数组 回溯算法 
// 👍 723 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
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

		if index == len(word) - 1 {
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
		defer func() { visit[x][y]  = false }() // TODO: 回溯时还原已访问的单元格

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
//leetcode submit region end(Prohibit modification and deletion)

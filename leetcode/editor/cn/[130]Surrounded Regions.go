//Given an m x n matrix board containing 'X' and 'O', capture all regions surrou
//nded by 'X'. 
//
// A region is captured by flipping all 'O's into 'X's in that surrounded region
//. 
//
// 
// Example 1: 
//
// 
//Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O"
//,"X","X"]]
//Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X
//"]]
//Explanation: Surrounded regions should not be on the border, which means that 
//any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not o
//n the border and it is not connected to an 'O' on the border will be flipped to 
//'X'. Two cells are connected if they are adjacent cells connected horizontally o
//r vertically.
// 
//
// Example 2: 
//
// 
//Input: board = [["X"]]
//Output: [["X"]]
// 
//
// 
// Constraints: 
//
// 
// m == board.length 
// n == board[i].length 
// 1 <= m, n <= 200 
// board[i][j] is 'X' or 'O'. 
// 
// Related Topics 深度优先搜索 广度优先搜索 并查集 
// 👍 478 👎 0
package main


//leetcode submit region begin(Prohibit modification and deletion)
var n, m int

func solve(board [][]byte)  {
	// TODO: 使用深度优先搜索实现标记操作。先把边界上和 O 连通点找到, 把这些变成 B,然后遍历整个 board 把 O 变成 X, 把 B 变成 O
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n = len(board), len(board[0])
	// 左右两列的遍历
	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}

	// 上下两行的遍历
	for j := 1; j < n - 1; j++ {
		dfs(board, 0, j)
		dfs(board, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, x int, y int)  {
	// 注意条件 != 'O' 而不是== 'X'
	if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs(board, x-1, y)
	dfs(board, x+1, y)
	dfs(board, x, y-1)
	dfs(board, x, y+1)
}
//leetcode submit region end(Prohibit modification and deletion)

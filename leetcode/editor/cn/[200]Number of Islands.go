//Given an m x n 2D binary grid grid which represents a map of '1's (land) and '
//0's (water), return the number of islands. 
//
// An island is surrounded by water and is formed by connecting adjacent lands h
//orizontally or vertically. You may assume all four edges of the grid are all sur
//rounded by water. 
//
// 
// Example 1: 
//
// 
//Input: grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//Output: 1
// 
//
// Example 2: 
//
// 
//Input: grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//Output: 3
// 
//
// 
// Constraints: 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 300 
// grid[i][j] is '0' or '1'. 
// 
// Related Topics 深度优先搜索 广度优先搜索 并查集 
// 👍 1165 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	//h := len(grid)
	//w := len(grid[0])
	//walkJudge := make([][]bool, h)
	//for i, _ := range walkJudge {
	//	walkJudge[i] = make([]bool, w)
	//}
	count := 0

	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count += 1
				//fullPath(&walkJudge, grid,i, j)
			}
		}
	}
	return count
}


func dfs(grid [][]byte, i, j int)  {
	// 判断 base case
	if !inArea(grid, i, j) {
		return
	}

	if grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func inArea(grid [][]byte, i, j int) bool {
	// 注意 ”=“ 的范围
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}
//leetcode submit region end(Prohibit modification and deletion)

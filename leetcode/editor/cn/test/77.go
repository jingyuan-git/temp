package main

import "fmt"

func combine(n int, k int) [][]int {
	var res [][]int
	visit := make([]bool, n)

	var dfs func(index int, visit []bool, path []int)
	dfs = func(index int, visit []bool, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			if visit[i] || len(path) > 0 && i+1 < path[len(path)-1] {
				continue
			}
			visit[i] = true
			path = append(path, i+1)
			dfs(index+1, visit, path)
			path = path[:len(path)-1]
			visit[i] = false
		}
	}
	dfs(0, visit, []int{})
	return res
}

func main() {
	//fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}

package main

import "fmt"

type Edge struct {
	val  int
	prob int
}

func minSum(n int, edges [][]int, succProb []int) int {
	nums := make([][]Edge, n)

	for k, _ := range edges {
		n0, n1 := edges[k][0], edges[k][1]
		nums[n0] = append(nums[n0], Edge{n1, succProb[k]})
	}
	fmt.Println(nums)
	return 0
}

func main() {
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 3},
		{3, 5},
		{4, 3},
		{4, 5},
		{4, 6},
		{5, 6},
	}
	succProb := []int{1, 12, 3, 9, 5, 4, 13, 15, 4}
	res := minSum(6, edges, succProb)
	fmt.Println(res)
}

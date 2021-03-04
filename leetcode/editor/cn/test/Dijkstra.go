package main

import (
	"fmt"
	"math"
)

type Edge struct {
	kval  map[int]int
	Visit bool
}

var dp []int

func minSum(n int, edges [][]int, succProb []int) int {
	nums := make(map[int]Edge)

	dp = make([]int, n)
	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}
	for i := 1; i < n; i++ {
		nums[i] = Edge{make(map[int]int), false}
	}

	for k, _ := range edges {
		start, end := edges[k][0], edges[k][1]
		nums[start].kval[end] = succProb[k]
	}
	fmt.Println(nums)

	for k := 1; k < n; k++ {
		fmt.Println("k = ", k)
		minNode := findMinNode(nums[k])
		aaa := nums[minNode]
		aaa.Visit = true
		nums[minNode] = aaa
		fmt.Println(nums)

	}



	return dp[len(dp)-1]
}

func findMinNode(nodes Edge) int {
	minValue := math.MaxInt32
	minIndex := 0
	for k, v := range nodes.kval {
		if v < minValue{
			minIndex = k
			minValue = v
		}
	}
	return minIndex
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

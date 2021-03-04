package main

import (
	"fmt"
	"math"
)

var n int

func main() {
	n = 5
	graph := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			graph[i][j] = math.MaxInt32
		}
	}
	graph[1][2] = 10
	graph[1][5] = 100
	graph[1][4] = 30
	graph[2][3] = 50
	graph[3][5] = 10
	graph[4][3] = 20
	graph[4][5] = 60

	dist := make([]int, n+1)
	prev := make([]int, n+1)
	fmt.Println(graph)
	res := Dijkstra1(1, graph, dist, prev)
	fmt.Println("=========")
	fmt.Println(res)

}

func Dijkstra1(v int, graph [][]int, dist []int, pre []int) []int {
	visit := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		dist[i] = graph[v][i]
		//// 源点到各个节点的距离
		//if graph[v][i] < math.MaxInt32 {
		//	pre[i] = 0
		//} else {
		//	pre[i] = v
		//}
	}
	fmt.Println(dist, visit)

	visit[0] = true
	dist[v] = 0

	for i := 2; i <= n; i++ {
		temp := math.MaxInt32
		index := 0

		for j := 2; j <= n; j++ {
			if !visit[j] && temp > dist[j] {
				temp = dist[j]
				index = j
			}
		}
		fmt.Println("index = ", index)
		visit[index] = true
		for j := 2; j <= n; j++ {
			// 寻找当前节点的下一个节点
			if !visit[j] && graph[index][j] < math.MaxInt32 {
				newdist := temp + graph[index][j]
				fmt.Println(newdist)
				if newdist < dist[j] {
					dist[j] = newdist
				}
			}
		}
		fmt.Println(dist, visit)
	}

	return dist
}

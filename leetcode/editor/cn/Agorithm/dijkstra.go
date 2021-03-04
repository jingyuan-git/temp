package main

import (
	"fmt"
	"math"
)

func dijkstra(ans map[string]map[string]int) map[string]int {
	visited := make(map[string]bool)
	res := make(map[string]int)

	for _, v := range ans {
		for k1, _ := range v {
			res[k1] = math.MaxInt32
		}
	}

	cur := "A"
	res[cur] = 0
	visited["A"] = true
	fmt.Println(res, visited)
	for i := 0; i < len(ans); i++ {
		tmp := res[cur]
		minIndex := math.MaxInt32

		var minValue string
		for k, v := range ans[cur] {
			if v < minIndex {
				minValue = k
			}
		}

		for k, v := range ans[cur] {
			if !visited[k] && v+tmp < res[k] {
				res[k] = v + tmp
			}
		}
		visited[minValue] = true
		cur = minValue

	}
	return res
}

func main() {
	ans := make(map[string]map[string]int)
	ans["A"] = map[string]int{"B": 1, "C": 12}
	ans["B"] = map[string]int{"C": 9, "D": 3}
	ans["C"] = map[string]int{"E": 5}
	ans["D"] = map[string]int{"C": 4, "E": 13, "F": 15}
	ans["E"] = map[string]int{"F": 4}

	fmt.Println(ans)
	res := dijkstra(ans)
	fmt.Println(res)

}

/**

start := "A"

s := make(map[string]bool)
dis := make(map[string]int)

cur := start

if _, ok := s[cur]; !ok {
	s[cur] = true
	dis[cur] = 0
}

for {
	fmt.Println(cur)

	minValues := 9999
	lastCur := ""

	for k, v := range ans[cur] {
		if _, ok := s[k]; !ok {
			s[k] = true
		}

		values, _ := dis[cur]
		values += v

		if values < minValues {
			minValues = values
			lastCur = k
		}

		if i, ok := dis[k]; !ok {
			dis[k] = values
		} else {
			if values < i {
				dis[k] = values
			}
		}
	}

	cur = lastCur

	if len(ans) == len(dis)-1 {
		break
	}
}

fmt.Println(dis)
*/

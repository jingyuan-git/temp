package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) (res [][]int) {
	//sort.Ints(candidates)
	//var dfs func(start int, temp []int, sum int)
	//dfs = func(start int, temp []int, sum int) {
	//	if sum >= target {
	//		if sum == target {
	//			res = append(res, append([]int(nil), temp...))
	//		}
	//		return
	//	}
	//	for i:=start; i<len(candidates); i++ {
	//		if i-1 >= start && candidates[i-1] == candidates[i] {
	//			continue
	//		}
	//		temp = append(temp, candidates[i])
	//		dfs(start+1, temp, sum+candidates[i])
	//		temp = temp[:len(temp)-1]
	//	}
	//}
	//dfs(0, []int{}, 0)
	//return
	sort.Ints(candidates)

	var dfs func(start int, temp []int, sum int)
	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				t := make([]int, len(temp))
				copy(t, temp)
				res = append(res, t)
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			if i-1 >= start && candidates[i-1] == candidates[i] {
				continue
			}
			temp = append(temp, candidates[i])
			dfs(i+1, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

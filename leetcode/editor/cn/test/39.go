package main

import "fmt"

func combinationSum(candidates []int, target int) (res [][]int) {
	var dfs func(start int, temp []int, sum int)
	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				res = append(res, append([]int(nil), temp...))
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			dfs(i, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return

}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

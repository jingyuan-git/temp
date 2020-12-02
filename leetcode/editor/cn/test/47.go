package main

import (
	"fmt"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) (res [][]int) {
	sort.Ints(nums)
	visited := make([]bool, len(nums))
	var path []int

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}

		for i, v := range nums {
			if visited[i] || i > 0 && !visited[i-1] && v == nums[i-1] {
				continue
			}
			path = append(path, v)
			visited[i] = true
			dfs(index + 1)
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
	//sort.Ints(nums)
	//visited := make([]bool, len(nums))
	//var path []int
	//var dfs func(index int)
	//dfs = func(index int) {
	//	if index == len(nums) {
	//		res = append(res, append([]int(nil), path...))
	//		return
	//	}
	//	for i, n := range nums {
	//		if visited[i] || i>0  && nums[i] == nums[i-1]  {
	//		//if visited[i] || i>0 && !visited[i-1] && nums[i] == nums[i-1]  {
	//			continue
	//		}
	//		path = append(path, n)
	//		visited[i] = true
	//		dfs(index+1)
	//		path = path[:len(path)-1]
	//		visited[i] = false
	//	}
	//}
	//dfs(0)
	//return
}

//
//func permuteUnique(nums []int) (ans [][]int) {
//	sort.Ints(nums)
//	n := len(nums)
//	perm := []int{}
//	vis := make([]bool, n)
//	var backtrack func(int)
//	backtrack = func(idx int) {
//		if idx == n {
//			ans = append(ans, append([]int(nil), perm...))
//			return
//		}
//		for i, v := range nums {
//			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
//				continue
//			}
//			perm = append(perm, v)
//			vis[i] = true
//			backtrack(idx + 1)
//			vis[i] = false
//			perm = perm[:len(perm)-1]
//		}
//	}
//	backtrack(0)
//	return
//}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 3}))
}

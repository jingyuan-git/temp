package main

import "fmt"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	//var path []int
	var dfs func(index int, path []int)
	dfs = func(index int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := index; i < len(nums); i++ {
			// TODO: i > index是精髓
			if i > index && nums[i-1] == nums[i] {
				continue
			}
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
		return
	}
	dfs(0, []int{})
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

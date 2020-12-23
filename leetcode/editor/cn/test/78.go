package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var path []int
	var dfs func(start int, n int)
	dfs = func(start int, n int) {
		if len(path) == n {
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
		}
		dfs(0, 0)

	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

//func subsets(nums []int) (ans [][]int) {
//	set := []int{}
//	var dfs func(int)
//	dfs = func(cur int) {
//		if cur == len(nums) {
//			ans = append(ans, append([]int(nil), set...))
//			return
//		}
//		// 选择nums[cur]这个元素
//		set = append(set, nums[cur])
//		dfs(cur + 1)
//		// 不选择nums[cur]这个元素
//		set = set[:len(set)-1]
//		dfs(cur + 1)
//	}
//	dfs(0)
//	return
//}
//
//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/subsets/solution/zi-ji-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

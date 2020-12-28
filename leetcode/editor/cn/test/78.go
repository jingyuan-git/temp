package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	//sort.Ints(nums)
	var res [][]int

	var dfs func(start int, list []int)
	dfs = func(start int, list []int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			list = append(list, nums[i])
			// TODO: 注意dfs开始的下标 之前写成启示为i+1了，那dfs的开始就没有变化
			dfs(i+1, list)
			list = list[:len(list)-1]
		}
	}
	dfs(0, []int{})
	return res
}
//
//func subsets(nums []int) [][]int {
//	res := [][]int{}
//
//	var dfs func(i int, list []int)
//	dfs = func(i int, list []int) {
//		tmp := make([]int, len(list))
//		copy(tmp, list)
//		res = append(res, tmp)
//
//		for j := i; j < len(nums); j++ {
//			list = append(list, nums[j])
//			dfs(j+1, list)
//			list = list[:len(list)-1]
//		}
//	}
//
//	dfs(0, []int{})
//	return res
//}
//
//作者：xiao_ben_zhu
//链接：https://leetcode-cn.com/problems/subsets/solution/shou-hua-tu-jie-zi-ji-hui-su-fa-xiang-jie-wei-yun-/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

//
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

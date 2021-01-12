package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var (
		temp []int
		res  [][]int
	)
	var dfs func(cur, rest int)
	dfs = func(cur, rest int) {
		if rest == 0 && len(temp) == k {
			res = append(res, append([]int(nil), temp...))
			return
		}

		// TODO: len(temp)+10-cur < k
		if rest < 0 || len(temp)+10-cur < k {
			return
		}

		// 不选
		dfs(cur+1, rest)

		// 选
		temp = append(temp, cur)
		dfs(cur+1, rest-cur)
		temp = temp[:len(temp)-1]
	}
	dfs(1, n)
	return res
}

//func combinationSum3(k int, n int) (ans [][]int) {
//	var temp []int
//	var dfs func(cur, rest int)
//	dfs = func(cur, rest int) {
//		// 找到一个答案
//		if len(temp) == k && rest == 0 {
//			ans = append(ans, append([]int(nil), temp...))
//			return
//		}
//		// 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
//		if len(temp)+10-cur < k || rest < 0 {
//			return
//		}
//		// 跳过当前数字
//		dfs(cur+1, rest)
//		// 选当前数字
//		temp = append(temp, cur)
//		dfs(cur+1, rest-cur)
//		temp = temp[:len(temp)-1]
//	}
//	dfs(1, n)
//	return
//}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/combination-sum-iii/solution/zu-he-zong-he-iii-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	fmt.Println(combinationSum3(4, 10))
}

package main

import (
	"fmt"
	//"strconv"
)

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	factorial[1] = 1
	for i := 2; i < len(factorial); i++ {
		factorial[i] = factorial[i-1] * i
	}

	var path, res []byte
	used := make([]bool, n+1)

	var dfs func(n int, k int, index int, path []byte)
	dfs = func(n int, k int, index int, path []byte) {
		if index == n {
			// TODO: 重新开辟一个内存，定义一个变量返回。不然递归的结果会被覆盖
			res = path
			return
		}
		cnt := factorial[n-1-index]

		for i := 1; i < n+1; i++ {
			if used[i] {
				continue
			}
			if cnt < k {
				k -= cnt
				continue
			}
			path = append(path, byte(i+'0'))
			used[i] = true
			dfs(n, k, index+1, path)
		}
	}

	if n == 0 {
		return ""
	}

	dfs(n, k, 0, path)
	return string(res)
}

//
//// 树形查找，并及时剪枝，提高效率
//func getPermutation(n int, k int) string {
//	var (
//		dfs  func(process int)
//		used = make([]bool, n+1)
//		ans  string
//	)
//
//	factorial := 1
//	for i := 1; i <= n; i++ {
//		factorial *= i
//	}
//
//	// 当前已经获取了多少个结果了
//	dfs = func(process int) {
//		// 当前结果集已经有几个了
//		if process == n {
//			ans = path
//			return
//		}
//		// 每一组的数量
//		groupNum := factorial / (n - process)
//		// 下一轮的阶乘也就是当前每组的数量
//		factorial = groupNum
//		for i := 1; i <= n; i++ {
//			if used[i] {
//				continue
//			}
//			// 如果当前数已经比这一组的个数都多，那就直接放弃
//			if k > groupNum {
//				k -= groupNum
//				continue
//			}
//			used[i] = true
//			ans = ans + strconv.Itoa(i)
//			dfs(process+1)
//		}
//	}
//
//	dfs(0)
//	return ans
//}

//作者：fbbin
//链接：https://leetcode-cn.com/problems/permutation-sequence/solution/shen-du-bian-li-jian-zhi-by-fbbin/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	fmt.Println(getPermutation(3, 3))
}

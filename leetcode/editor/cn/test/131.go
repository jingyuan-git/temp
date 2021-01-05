package main

import "fmt"

var (
	res  [][]string
	path []string
)

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func dfs(start int, s string) {
	if start == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s[start : i+1]) {
			path = append(path, s[start:i+1])
			dfs(i+1, s)
			path = path[:len(path)-1]
		}

	}
}

func partition(s string) [][]string {
	res = [][]string{}
	path = []string{}
	dfs(0, s)
	return res
}

func main() {
	fmt.Println(partition("aabac"))
}

////func checkP(s string) bool {
////	for i,j := 0,len(s)-1; i < j; i,j = i+1,j-1 {
////		if s[i] != s[j] {
////			return false
////		}
////	}
////	return true
////}
//
//func recurseP(cur []string, ret *[][]string, index int, s string) {
//	if index == len(s) {
//		tmp := make([]string,len(cur))
//		copy(tmp,cur)
//		*ret = append(*ret,tmp)
//		return
//	}
//
//	for i:=index; i<len(s); i++ {
//		if checkP(s[index:i+1]) {
//			cur = append(cur, s[index:i+1])
//			recurseP(cur,ret,i+1,s)
//			cur = cur[:len(cur)-1]
//		}
//	}
//}
//
//func partition(s string) [][]string {
//	ret := make([][]string,0)
//	cur := make([]string,0)
//	recurseP(cur,&ret,0,s)
//	return ret
//}
//
//作者：sailor-8
//链接：https://leetcode-cn.com/problems/palindrome-partitioning/solution/golang-di-gui-hui-su-by-sailor-8/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func partition(s string) [][]string {
	ans := make([][]string, 0)
	size := len(s)
	if size == 0 {
		return ans
	}
	// 预处理
	// 使用 dp[i][j] 表示 s[i:j] 是否回文
	dp := make([][]bool, size)
	for i := range dp {
		dp[i] = make([]bool, size)
	}
	for r := 0; r < size; r++ {
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
			}
		}
	}

	path := make([]string, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		if start >= size {
			tmp := make([]string, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := start; i < size; i++ {
			if !dp[start][i] {
				continue
			}
			path = append(path, s[start:i+1])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}

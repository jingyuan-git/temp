//The set [1, 2, 3, ..., n] contains a total of n! unique permutations. 
//
// By listing and labeling all of the permutations in order, we get the followin
//g sequence for n = 3: 
//
// 
// "123" 
// "132" 
// "213" 
// "231" 
// "312" 
// "321" 
// 
//
// Given n and k, return the kth permutation sequence. 
//
// 
// Example 1: 
// Input: n = 3, k = 3
//Output: "213"
// Example 2: 
// Input: n = 4, k = 9
//Output: "2314"
// Example 3: 
// Input: n = 3, k = 1
//Output: "123"
// 
// 
// Constraints: 
//
// 
// 1 <= n <= 9 
// 1 <= k <= n! 
// 
// Related Topics 数学 回溯算法 
// 👍 450 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
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
//leetcode submit region end(Prohibit modification and deletion)

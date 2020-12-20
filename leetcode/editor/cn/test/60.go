package main

import "fmt"

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	factorial[1] = 1

	for i := 2; i < len(factorial); i++ {
		factorial[i] = factorial[i-1] * i
	}
	fmt.Println(factorial)

	used := make([]bool, n+1)
	fmt.Println(used)

	var dfs func(n int, k int, index int, path []byte)
	dfs = func(n int, k int, index int, path []byte) {
		if index == n {
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
			return

		}
	}

	if n == 0 {
		return ""
	}

	var path []byte
	dfs(n, k, 0, path)
	return string(path)
}

func main() {
	fmt.Println(getPermutation(3, 3))
}

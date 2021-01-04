package main

import "fmt"

var (
	res		[][]string
	path	[]string
)

func isPalindrome(s string) bool {
	return false
}

func dfs(start int, s string)  {
	if start == len(s) {
		temp := make([]string, len(s))
		res = append(res, temp)
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s[start:i+1]) {
			path = append(path, s[start:i+1])
		}
		dfs(i + 1, s)
	}
}

func partition(s string) [][]string {
	res = [][]string{}
	path = []string{}
	dfs(0, s)
	return res
}

func main() {
	fmt.Println(partition("aab"))
}
package main

import "fmt"

func strStr(haystack string, needle string) int {
	hl := len(haystack)
	nl := len(needle)
	if nl <= 0 {
		return 0
	}
	if hl <= 0 {
		return -1
	}
	for i := 0; i < hl-nl+1; i++ {
		if haystack[i] == needle[0] {
			flag := true
			for j := 1; j < nl; j++ {
				if haystack[i+j] != needle[j] {
					flag = false
					break
				}
			}
			if flag {
				return i
			}
		}
	}
	return -1
}
func main() {
	fmt.Println(strStr("a", "a"))
}

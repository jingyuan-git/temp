package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := len(s)
	count := 0
	for i := length - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if i == 0 || count != 0 {
			return count
		}
	}
	return count
}

func main() {
	//fmt.Println(lengthOfLastWord("sdf asdf  "))
	fmt.Println(lengthOfLastWord(" "))
}

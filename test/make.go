package main

import "fmt"

func main() {
	slice := make([]string, 0)
	sss := []string{"1", "1", "1"}
	slice = sss
	fmt.Println(slice)

}

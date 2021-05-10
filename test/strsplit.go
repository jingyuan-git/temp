package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "node1LrMplsLabel-lrvrf-subnet123"
	strs :=strings.Split(str, "MplsLabel-")
	fmt.Println(strs)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	//demo := "I&love&Go,&and&I&also&love&Python."
	demo := "edgeportid_host"
	string_slice := strings.Split(demo, "_")
	edgeportid, host := string_slice[0], string_slice[1]

	fmt.Println(edgeportid)
	fmt.Println(host)

	fmt.Println("result:", string_slice)
	fmt.Println("len:", len(string_slice))
	fmt.Println("cap:", cap(string_slice))
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	mac := "fa:16:3d:d9:49:bb"
	mac = strings.ReplaceAll(mac, ":", "_")
	fmt.Println(mac)
}
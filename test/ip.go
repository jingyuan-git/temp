package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.ParseIP("sdf"))
	fmt.Println(net.ParseIP("0.0.0.0"))
	fmt.Println(net.ParseIP(""))
	fmt.Println(net.ParseIP("sdf"))


}

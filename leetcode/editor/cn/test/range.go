package main

import "fmt"

type MyType struct {
	field int
}

func main() {
	var array [10]MyType

	for i, _ := range array {
		array[i].field = i
	}
	for _, value := range array {
		value.field = 100
	}

	for _, e := range array {
		fmt.Println(e.field)
		//fmt.Println("--")
	}
}

//}func main() {
//	var array [10]MyType
//
//	for _, e := range array {
//		e.field = "foo"
//	}
//
//	for _, e := range array {
//		fmt.Println(e.field)
//		//fmt.Println("--")
//	}
//}

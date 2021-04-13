package main

import "fmt"

func main() {

	existedIndexes := map[string]bool{}

	existedIndexes["xxx"] = true
	existedIndexes["222"] = false

	k, v := existedIndexes["xxx"]
	//fmt.Println(k, v)

	if k, v = existedIndexes["222"]; v {
		fmt.Println(k, v)

	}
	//edgePortConfigMap := make(map[string][]string)
	//edgePortConfigMap["sdf"] = []string{"sdf", "asdf"}
	//fmt.Println(edgePortConfigMap["lu"])
	//lbls := map[string]string{}
	//lbls["nexthop_mode_left"] = "port"
	//lbls["nexthop_mode_right"] = "network"
	//
	//policy := make(map[string]map[string]string)
	//policy["go"] = lbls
	////fmt.Println(policy)
	//lbls1 := lbls
	//lbls1["nexthop_mode"] = "port_1"
	//policy["goes"] = lbls1
	////fmt.Println(policy)
	//
	//for k, v := range policy["go"] {
	//	fmt.Println(k)
	//
	//	fmt.Println("//")
	//	fmt.Println(v)
	//	fmt.Println("--")
	//}
	////fmt.Println(policy)
	//if _, ok := policy["wer"]; ok {
	//	fmt.Println(ok)
	//} else {
	//	fmt.Println(ok)
	//	fmt.Println("sdfa")
	//}
	//
	//fmt.Println(policy["wer"])
}

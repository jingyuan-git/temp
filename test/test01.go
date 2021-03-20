package main

import (
	"fmt"
	"strings"
)

func getPortPrefix(deviceValue string) (portSg string) {
	if strings.Index(deviceValue, "tap") == 0 {
		portSg = deviceValue[3:]
	}

	if strings.Index(deviceValue, "hpeer") == 0 {
		portSg = deviceValue[5:]
	}
	return
}

func main() {
	temp := getPortPrefix("sdf")
	if temp == "" {
		fmt.Println("is null")
	}
}

//func convertAvoidOverflow(sign int ,abs string) int {
//	//fmt.Println(sign, abs)
//	res := 0
//	for _, val := range abs{
//		//fmt.Println(val - '0')
//		res = res * 10 + int(val) - '0'
//	}
//	//溢出 切记
//	if sign * res >math.MaxInt32{
//		return math.MaxInt32
//	}
//	if sign * res < math.MinInt32{
//		return math.MinInt32
//	}
//	return sign * res
//}
//
//func main() {
//	//i := -11
//	//fmt.Println(i/10)
//	fmt.Println(convertAvoidOverflow(1,"922337203685477580"))
//}

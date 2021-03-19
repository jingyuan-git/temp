package main

import (
	"fmt"
	"strings"
)

func isHump(s string) (string, bool) {
	stringSlice := strings.Split(s, "_")
	if len(stringSlice) >= 2 {
		var humpString string

		for _, v := range stringSlice {
			var capitalizeString string
			vv := []rune(v)
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					capitalizeString += string(vv[i])
				} else {
					if vv[i] >= 65 && vv[i] <= 90 { // 后文有介绍
						vv[i] += 32 // string的码表相差32位
						capitalizeString += string(vv[i])
					}
				}
			}
			humpString += capitalizeString
		}
		return humpString, true
	} else {
		return s, false
	}
}

func convertLine(line string) string {
	var result string
	stringLine := strings.Split(line, " ")
	for _, v := range stringLine {
		convertVarious, isConvertVarious := isHump(v)
		if isConvertVarious {
			result += convertVarious
		} else {
			result += v
		}
		result += " "
	}

	return result
}

func main() {
	line := "VIF_TYPE_HOST         VIFType = 0"
	fmt.Println(convertLine(line))

	////demo := "I&love&Go,&and&I&also&love&Python."
	//demo := "VIF_TYPE_HOST"
	//string_slice := strings.Split(demo, "_")
	////edgeportid, host := string_slice[0], string_slice[1]
	////
	////fmt.Println(edgeportid)
	////fmt.Println(host)
	//
	//if len(string_slice) >= 1 {
	//	var humpString string
	//
	//	for _, v := range string_slice {
	//		var capitalizeString string
	//		vv := []rune(v)
	//		for i := 0; i < len(vv); i++ {
	//			if i == 0 {
	//				capitalizeString += string(vv[i])
	//			} else {
	//				if vv[i] >= 65 && vv[i] <= 90 {  // 后文有介绍
	//					vv[i] += 32 // string的码表相差32位
	//					capitalizeString += string(vv[i])
	//				}
	//			}
	//		}
	//		humpString += capitalizeString
	//	}
	//	fmt.Println(humpString)
	//}
	//fmt.Println("result:", string_slice)
	//fmt.Println("len:", len(string_slice))
	//fmt.Println("cap:", cap(string_slice))
}

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text :="Abc a7c MFC 8ca. 你好！ Golang  VIF_TYPE_HOST     VIF_TYPE_sd    /"

	////reg, _ := regexp.Compile(`[_]{1,}`)           // 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
	//reg, _ := regexp.Compile(`(\w+)_+(\w+)`)           // 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
	//fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello" "Go"]
	//fmt.Println(reg.ReplaceAllString(text, `\U$1\E`))
	////fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["Hello" "Go"]


	re := regexp.MustCompile(`(\w+)_+(\w+)`)
	fmt.Println(re.ReplaceAllStringFunc(text, toHump))
}

func toHump (s string) string {
	stringSlice := strings.Split(s, "_")
	var humpString string

	for _, v := range stringSlice {
		var capitalizeString string
		vv := []rune(v)
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				capitalizeString += string(vv[i])
			} else {
				if vv[i] >= 65 && vv[i] <= 90 {
					vv[i] += 32
					capitalizeString += string(vv[i])
				}
			}
		}
		humpString += capitalizeString
	}
	return humpString
}

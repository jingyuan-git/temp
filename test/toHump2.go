package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	filepath := "./input.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		writeFile(convertLine(line))
		//fmt.Println(line)

		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
}

func writeFile(line string) {
	var filename = "./output.txt"
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()

	//fmt.Printf("写入 %d 个字节n", n)
	_, err1 = f.WriteString(line) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	//fmt.Printf("写入 %d 个字节n", n)
	f.Sync()
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

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
					if vv[i] >= 65 && vv[i] <= 90 {
						vv[i] += 32
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

func convertLine(line string) string {
	re := regexp.MustCompile(`(\w+)_+(\w+)`)
	return re.ReplaceAllStringFunc(line, toHump) + "\n"

	//fmt.Println(re.ReplaceAllStringFunc(line, toHump))

	//stringLine := strings.Split(line, " ")
	//for _, v := range stringLine {
	//	convertVarious, isConvertVarious := isHump(v)
	//	if isConvertVarious {
	//		result += convertVarious
	//	} else {
	//		result += v
	//	}
	//	result += " "
	//}
	//result += "\n"
	//return result
}

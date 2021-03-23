package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var inputfile = flag.String("inputfile", "", "input file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	start()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}

func start()  {
	filepath := *inputfile
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
	fmt.Println("file size = ", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		s := convertLine(line)
		writeFile(s)
		//fmt.Println(line)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Success!")
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
		f, err1 = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend) //打开文件
		//fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		if err1 != nil {
			fmt.Printf("ERROR: %+v\n", err1)
		}
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
	re := regexp.MustCompile(`([A-Z]+)(_+([A-Z]+))+`)
	return re.ReplaceAllStringFunc(line, toHump) + "\n"
}

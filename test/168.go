package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func convertToTitle(columnNumber int) string {
	result := []byte{}
	for columnNumber != 0 {
		remainder := (columnNumber - 1) % 26

		result = append([]byte{'A' + byte(remainder)}, result...)
		columnNumber = (columnNumber - 1) / 26
		fmt.Println(columnNumber)
	}
	return string(result)
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(convertToTitle(701))
}
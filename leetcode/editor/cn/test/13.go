package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func romanToInt(s string) int {
	sByte := []byte(s)
	var _map = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	preValue := 0
	res := 0
	for _, v := range sByte {
		if _map[v] <= preValue {
			res += _map[v]
			preValue = _map[v]
		} else {
			res = res - 2*preValue + _map[v]
			preValue = _map[v]
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	input := "LVIII"
	fmt.Println(romanToInt(input))
}

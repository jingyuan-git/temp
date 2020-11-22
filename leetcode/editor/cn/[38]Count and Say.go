//The count-and-say sequence is a sequence of digit strings defined by the recur
//sive formula: 
//
// 
// countAndSay(1) = "1" 
// countAndSay(n) is the way you would "say" the digit string from countAndSay(n
//-1), which is then converted into a different digit string. 
// 
//
// To determine how you "say" a digit string, split it into the minimal number o
//f groups so that each group is a contiguous section all of the same character. T
//hen for each group, say the number of characters, then say the character. To con
//vert the saying into a digit string, replace the counts with a number and concat
//enate every saying. 
//
// For example, the saying and conversion for digit string "3322251": 
//
// Given a positive integer n, return the nth term of the count-and-say sequence
//. 
//
// 
// Example 1: 
//
// 
//Input: n = 1
//Output: "1"
//Explanation: This is the base case.
// 
//
// Example 2: 
//
// 
//Input: n = 4
//Output: "1211"
//Explanation:
//countAndSay(1) = "1"
//countAndSay(2) = say "1" = one 1 = "11"
//countAndSay(3) = say "11" = two 1's = "21"
//countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"
// 
//
// 
// Constraints: 
//
// 
// 1 <= n <= 30 
// 
// Related Topics 字符串 
// 👍 601 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
//func countAndSay(n int) string {
//	str :="1"
//	for i:=2;i<=n;i++{
//		var tmp strings.Builder
//		preByte := str[0]
//		count := 1
//		for j:=1;j<len(str);j++{
//			if str[j]==preByte{
//				count++
//			}else {
//				tmp.WriteString(strconv.Itoa(count))
//				tmp.WriteByte(preByte)
//				preByte =str[j]
//				count=1
//			}
//		}
//		tmp.WriteString(strconv.Itoa(count))
//		tmp.WriteByte(preByte)
//		str = tmp.String()
//	}
//	return str
//}

func countAndSay(n int) string {
	buf := []byte{'1'}

	for n > 1 {
		buf = say(buf)
		n--
	}
	return string(buf)
}

func say(buf []byte) []byte {
	// res 长度不会超过 buf 的两倍，所以，可以事先指定容量，加快append的速度
	res := make([]byte, 0, len(buf)*2)
	i, j := 0, 1
	for i < len(buf) {
		// 利用 j ，找到下一个不同的元素
		for j < len(buf) && buf[i] == buf[j] {
			j++
		}
		// res 中 res[i] 表示 res[i+1] 的个数，i 为0,2,4,6,...
		res = append(res, byte(j-i+'0'), buf[i])
		// 移动 i 到 j
		i = j
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

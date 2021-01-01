//Given a string s containing only digits, return all possible valid IP addresse
//s that can be obtained from s. You can return them in any order. 
//
// A valid IP address consists of exactly four integers, each integer is between
// 0 and 255, separated by single dots and cannot have leading zeros. For example,
// "0.1.2.201" and "192.168.1.1" are valid IP addresses and "0.011.255.245", "192.
//168.1.312" and "192.168@1.1" are invalid IP addresses. 
//
// 
// Example 1: 
// Input: s = "25525511135"
//Output: ["255.255.11.135","255.255.111.35"]
// Example 2: 
// Input: s = "0000"
//Output: ["0.0.0.0"]
// Example 3: 
// Input: s = "1111"
//Output: ["1.1.1.1"]
// Example 4: 
// Input: s = "010010"
//Output: ["0.10.0.10","0.100.1.0"]
// Example 5: 
// Input: s = "101023"
//Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
// 
// 
// Constraints: 
//
// 
// 0 <= s.length <= 3000 
// s consists of digits only. 
// 
// Related Topics 字符串 回溯算法 
// 👍 475 👎 0

package main

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
const SEG_COUNT = 4

var (
	ans      []string
	segments []int
)

func restoreIpAddresses(s string) []string {
	segments = make([]int, SEG_COUNT)
	ans = []string{}
	dfs(0, 0, s)
	return ans
}



func dfs(segStart int, segId int, s string) {
	// segId == SEG_COUNT 此处的写法不一样，所以之后会需要多写一个else
	if SEG_COUNT == segId && segStart == len(s) {
		ipAddr := ""
		for i := 0; i < SEG_COUNT; i++ {
			ipAddr += strconv.Itoa(segments[i])
			if i != SEG_COUNT-1 {
				ipAddr += "."
			}
		}
		ans = append(ans, ipAddr)
		return
	}

	if SEG_COUNT == segId {
		return
	}

	if segStart == len(s) {
		return
	}

	if s[segStart] == '0' {
		segments[segId] = 0
		dfs(segStart + 1, segId + 1, s)
	}

	addr := 0
	for segEnd := segStart; segEnd < len(s); segEnd++ {
		addr = addr * 10 + int(s[segEnd] - '0')
		if addr > 0 && addr <= 0xFF {
			segments[segId] = addr
			dfs(segEnd+1, segId + 1, s)
		} else {
			break
		}
	}
}

//func dfs(s string, segId, segStart int) {
//	// TODO: 如果找到了 4 段 IP 地址并且遍历完了字符串，那么就是一种答案
//	if segId == SEG_COUNT {
//		if segStart == len(s) {
//			ipAddr := ""
//			for i := 0; i < SEG_COUNT; i++ {
//				ipAddr += strconv.Itoa(segments[i])
//				if i != SEG_COUNT - 1 {
//					ipAddr += "."
//				}
//			}
//			ans = append(ans, ipAddr)
//		}
//		return
//	}
//
//	// 如果还没有找到 4 段 IP 地址就已经遍历完了字符串，那么提前回溯
//	if segStart == len(s) {
//		return
//	}
//	// TODO: 由于不能有前导零，如果当前数字为 0，那么这一段 IP 地址只能为 0
//	if s[segStart] == '0' {
//		segments[segId] = 0
//		dfs(s, segId + 1, segStart + 1)
//	}
//	// 一般情况，TODO: 枚举每一种可能性并递归
//	addr := 0
//	for segEnd := segStart; segEnd < len(s); segEnd++ {
//		addr = addr * 10 + int(s[segEnd] - '0')
//		if addr > 0 && addr <= 0xFF {
//			segments[segId] = addr
//			dfs(s, segId + 1, segEnd + 1)
//		} else {
//			break
//		}
//	}
//}
//leetcode submit region end(Prohibit modification and deletion)

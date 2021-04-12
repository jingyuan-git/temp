//Given two version numbers, version1 and version2, compare them. 
//
// 
// 
//
// Version numbers consist of one or more revisions joined by a dot '.'. Each re
//vision consists of digits and may contain leading zeros. Every revision contains
// at least one character. Revisions are 0-indexed from left to right, with the le
//ftmost revision being revision 0, the next revision being revision 1, and so on.
// For example 2.5.33 and 0.1 are valid version numbers. 
//
// To compare version numbers, compare their revisions in left-to-right order. R
//evisions are compared using their integer value ignoring any leading zeros. This
// means that revisions 1 and 001 are considered equal. If a version number does n
//ot specify a revision at an index, then treat the revision as 0. For example, ve
//rsion 1.0 is less than version 1.1 because their revision 0s are the same, but t
//heir revision 1s are 0 and 1 respectively, and 0 < 1. 
//
// Return the following: 
//
// 
// If version1 < version2, return -1. 
// If version1 > version2, return 1. 
// Otherwise, return 0. 
// 
//
// 
// Example 1: 
//
// 
//Input: version1 = "1.01", version2 = "1.001"
//Output: 0
//Explanation: Ignoring leading zeroes, both "01" and "001" represent the same i
//nteger "1".
// 
//
// Example 2: 
//
// 
//Input: version1 = "1.0", version2 = "1.0.0"
//Output: 0
//Explanation: version1 does not specify revision 2, which means it is treated a
//s "0".
// 
//
// Example 3: 
//
// 
//Input: version1 = "0.1", version2 = "1.1"
//Output: -1
//Explanation: version1's revision 0 is "0", while version2's revision 0 is "1".
// 0 < 1, so version1 < version2.
// 
//
// Example 4: 
//
// 
//Input: version1 = "1.0.1", version2 = "1"
//Output: 1
// 
//
// Example 5: 
//
// 
//Input: version1 = "7.5.2.4", version2 = "7.5.3"
//Output: -1
// 
//
// 
// Constraints: 
//
// 
// 1 <= version1.length, version2.length <= 500 
// version1 and version2 only contain digits and '.'. 
// version1 and version2 are valid version numbers. 
// All the given revisions in version1 and version2 can be stored in a 32-bit in
//teger. 
// 
// Related Topics 字符串 
// 👍 152 👎 0
package main

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")

	for i := 0; i < len(s1) || i < len(s2) ;i++ {
		var n1, n2 int
		if i < len(s1) {
			n1, _ = strconv.Atoi(s1[i])
		} else {
			// 超过本来输入长度，则补0
			n1 = 0
		}

		if i < len(s2) {
			n2, _ = strconv.Atoi(s2[i])
		} else {
			n2 = 0
		}

		if n1 > n2 {
			return 1
		}

		if n1 < n2 {
			return -1
		}

	}

	return 0
}
//leetcode submit region end(Prohibit modification and deletion)
/**
法一：分割+解析
https://leetcode-cn.com/problems/compare-version-numbers/solution/bi-jiao-ban-ben-hao-by-leetcode/

TODO: 法二：双指针
func compareVersion(version1 string, version2 string) int {
	i :=0
	j := 0

	for i<len(version1) || j < len(version2){
        var tm1 = ""
        var tm2 = ""
		for i<len(version1) && version1[i]!='.'{
			tm1+=string(version1[i])
			i++
		}
		for j < len(version2) && version2[j]!='.'{
			tm2+=string(version2[j])
			j++
		}
		in1,_ := strconv.Atoi(tm1)
		in2,_ := strconv.Atoi(tm2)
		if in1>in2{
			return 1
		}else if in1<in2{
			return -1
		}
		i++
		j++

	}
	return 0
}

作者：HaA4qccPPa
链接：https://leetcode-cn.com/problems/compare-version-numbers/solution/golang-zui-kuai-jie-fa-by-haa4qccppa/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
//The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C
//', 'G', and 'T'. 
//
// 
// For example, "ACGAATTCCG" is a DNA sequence. 
// 
//
// When studying DNA, it is useful to identify repeated sequences within the DNA
//. 
//
// Given a string s that represents a DNA sequence, return all the 10-letter-lon
//g sequences (substrings) that occur more than once in a DNA molecule. You may re
//turn the answer in any order. 
//
// 
// Example 1: 
// Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
//Output: ["AAAAACCCCC","CCCCCAAAAA"]
// Example 2: 
// Input: s = "AAAAAAAAAAAAA"
//Output: ["AAAAAAAAAA"]
// 
// 
// Constraints: 
//
// 
// 1 <= s.length <= 105 
// s[i] is either 'A', 'C', 'G', or 'T'. 
// 
// Related Topics 位运算 哈希表 
// 👍 159 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func findRepeatedDnaSequences(s string) []string {
	// 一定要先初始化
	res := []string{}
	dict := map[string]bool{}

	for i, n := 0, len(s)-9; i < n; i++ {
		if _, ok := dict[s[i:i+10]]; ok {
			dict[s[i:i+10]] = true
		} else {
			dict[s[i:i+10]] = false
		}
	}

	for v, ok := range dict {
		if ok {
			res = append(res, v)
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)





/**
func findRepeatedDnaSequences(s string) []string {
    ans, d := []string{}, map[string]bool{}
	// 根据题目：滑动窗口的长度=10
	// len(s) - 9 ：滑动窗口会移动多少次
    for i, n := 0, len(s) - 9; i < n; i++ {
        if _, ok := d[s[i: i + 10]]; ok {
			// 含有重复元素
            d[s[i: i + 10]] = true
        } else {
			// 第一次搜寻到此元素
            d[s[i: i + 10]] = false
        }
    }
    for si, ok := range d {
        if ok {
            ans = append(ans, si)
        }
    }
    return ans
}

作者：tuotuoli
链接：https://leetcode-cn.com/problems/repeated-dna-sequences/solution/liang-ge-ji-he-by-tuotuoli/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
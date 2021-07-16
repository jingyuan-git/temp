//Given a string path, which is an absolute path (starting with a slash '/') to 
//a file or directory in a Unix-style file system, convert it to the simplified ca
//nonical path. 
//
// In a Unix-style file system, a period '.' refers to the current directory, a 
//double period '..' refers to the directory up a level, and any multiple consecut
//ive slashes (i.e. '//') are treated as a single slash '/'. For this problem, any
// other format of periods such as '...' are treated as file/directory names. 
//
// The canonical path should have the following format: 
//
// 
// The path starts with a single slash '/'. 
// Any two directories are separated by a single slash '/'. 
// The path does not end with a trailing '/'. 
// The path only contains the directories on the path from the root directory to
// the target file or directory (i.e., no period '.' or double period '..') 
// 
//
// Return the simplified canonical path. 
//
// 
// Example 1: 
//
// 
//Input: path = "/home/"
//Output: "/home"
//Explanation: Note that there is no trailing slash after the last directory nam
//e.
// 
//
// Example 2: 
//
// 
//Input: path = "/../"
//Output: "/"
//Explanation: Going one level up from the root directory is a no-op, as the roo
//t level is the highest level you can go.
// 
//
// Example 3: 
//
// 
//Input: path = "/home//foo/"
//Output: "/home/foo"
//Explanation: In the canonical path, multiple consecutive slashes are replaced 
//by a single one.
// 
//
// Example 4: 
//
// 
//Input: path = "/a/./b/../../c/"
//Output: "/c"
// 
//
// 
// Constraints: 
//
// 
// 1 <= path.length <= 3000 
// path consists of English letters, digits, period '.', slash '/' or '_'. 
// path is a valid absolute Unix path. 
// 
// Related Topics 栈 字符串 
// 👍 295 👎 0
package main

import (
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func simplifyPath(path string) string {
	dic := []string{}
	// 0. 先切分
	// 1. 若是.不管
	// 2. 若是.. 删除上一个元素
	// 3. 连续多个斜杠转为一个斜杠！！！
	paths := strings.Split(path, "/")
	for _, v := range paths {
		switch v {
		case ".":
			continue
		case "..":
			if len(dic) > 0 {
				dic = dic[:len(dic)-1]
			}
		case "":
			continue
		default:
			dic = append(dic, v)
		}
	}
	res := strings.Join(dic, "/")
	// 4. 第一个斜杠号还需自己添加
	return "/" + res
}
//leetcode submit region end(Prohibit modification and deletion)

/**
class Solution:
    def simplifyPath(self, path: str) -> str:
        # 1. 连续多个斜杠转为一个斜杠
        # 2. 删去.
        # 3. 删去..和其前一个目录
        paths = path.split("/")
        res = []
        for p in paths:
            if p=='.':
                continue
            if p=='..':
                if len(res)>0:
                    del res[-1]
                continue
            if p=='':
                continue
            res.append(p)
        res="/".join(res)
        return '/'+res

作者：jiang-liu-wan-zhuan
链接：https://leetcode-cn.com/problems/simplify-path/solution/zhi-jie-yi-ge-yi-ge-pan-duan-by-jiang-li-uvvb/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
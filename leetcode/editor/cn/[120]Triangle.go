//Given a triangle array, return the minimum path sum from top to bottom. 
//
// For each step, you may move to an adjacent number of the row below. More form
//ally, if you are on index i on the current row, you may move to either index i o
//r index i + 1 on the next row. 
//
// 
// Example 1: 
//
// 
//Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
//Output: 11
//Explanation: The triangle looks like:
//   2
//  3 4
// 6 5 7
//4 1 8 3
//The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined abov
//e).
// 
//
// Example 2: 
//
// 
//Input: triangle = [[-10]]
//Output: -10
// 
//
// 
// Constraints: 
//
// 
// 1 <= triangle.length <= 200 
// triangle[0].length == 1 
// triangle[i].length == triangle[i - 1].length + 1 
// -104 <= triangle[i][j] <= 104 
// 
//
// 
//Follow up: Could you do this using only O(n) extra space, where n is the total
// number of rows in the triangle? Related Topics 数组 动态规划 
// 👍 691 👎 0
package main

import (
	"fmt"
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]
	for i := 1 ; i < len(triangle); i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	//fmt.Println(dp)
	ans := math.MaxInt32
	for i := 0; i < len(dp); i++ {
		ans = min(ans, dp[i])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)

/**
func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    f := make([][]int, n)
    for i := 0; i < n; i++ {
        f[i] = make([]int, n)
    }
    f[0][0] = triangle[0][0]
    for i := 1; i < n; i++ {
        f[i][0] = f[i - 1][0] + triangle[i][0]
        for j := 1; j < i; j++ {
            f[i][j] = min(f[i - 1][j - 1], f[i - 1][j]) + triangle[i][j]
        }
        f[i][i] = f[i - 1][i - 1] + triangle[i][i]
    }
    ans := math.MaxInt32
    for i := 0; i < n; i++ {
        ans = min(ans, f[n-1][i])
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/triangle/solution/san-jiao-xing-zui-xiao-lu-jing-he-by-leetcode-solu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
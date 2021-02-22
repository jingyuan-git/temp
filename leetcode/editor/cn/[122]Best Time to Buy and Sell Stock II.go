//You are given an array prices for which the ith element is the price of a give
//n stock on day i. 
//
// Find the maximum profit you can achieve. You may complete as many transaction
//s as you like (i.e., buy one and sell one share of the stock multiple times). 
//
// Note: You may not engage in multiple transactions simultaneously (i.e., you m
//ust sell the stock before you buy again). 
//
// 
// Example 1: 
//
// 
//Input: prices = [7,1,5,3,6,4]
//Output: 7
//Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 
//5-1 = 4.
//Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
//
// 
//
// Example 2: 
//
// 
//Input: prices = [1,2,3,4,5]
//Output: 4
//Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 
//5-1 = 4.
//Note that you cannot buy on day 1, buy on day 2 and sell them later, as you ar
//e engaging multiple transactions at the same time. You must sell before buying a
//gain.
// 
//
// Example 3: 
//
// 
//Input: prices = [7,6,4,3,1]
//Output: 0
//Explanation: In this case, no transaction is done, i.e., max profit = 0.
// 
//
// 
// Constraints: 
//
// 
// 1 <= prices.length <= 3 * 104 
// 0 <= prices[i] <= 104 
// 
// Related Topics 贪心算法 数组 
// 👍 1094 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {

}
//leetcode submit region end(Prohibit modification and deletion)

/**
动态规划
func maxProfit(prices []int) int {
    n := len(prices)
    dp0, dp1 := 0, -prices[0]
    for i := 1; i < n; i++ {
        dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
    }
    return dp0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode-s/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

/**
贪心

func maxProfit(prices []int) (ans int) {
    for i := 1; i < len(prices); i++ {
        ans += max(0, prices[i]-prices[i-1])
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode-s/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
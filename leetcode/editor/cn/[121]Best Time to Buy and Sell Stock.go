//You are given an array prices where prices[i] is the price of a given stock on
// the ith day. 
//
// You want to maximize your profit by choosing a single day to buy one stock an
//d choosing a different day in the future to sell that stock. 
//
// Return the maximum profit you can achieve from this transaction. If you canno
//t achieve any profit, return 0. 
//
// 
// Example 1: 
//
// 
//Input: prices = [7,1,5,3,6,4]
//Output: 5
//Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 
//6-1 = 5.
//Note that buying on day 2 and selling on day 1 is not allowed because you must
// buy before you sell.
// 
//
// Example 2: 
//
// 
//Input: prices = [7,6,4,3,1]
//Output: 0
//Explanation: In this case, no transactions are done and the max profit = 0.
// 
//
// 
// Constraints: 
//
// 
// 1 <= prices.length <= 105 
// 0 <= prices[i] <= 104 
// 
// Related Topics 数组 动态规划 
// 👍 1445 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	dp := 0
	minV := prices[0]
	for i := 1; i < len(prices); i++ {
		minV = min(prices[i], minV)
		dp = max(dp, prices[i] - minV)
	}
	return dp
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
//leetcode submit region end(Prohibit modification and deletion)

/**

// 运行时间可能更快，思路差不多。没有像我一样独立封装max，min
func maxProfit(prices []int) int {
    if (prices == nil || len(prices) == 0){
        return 0
    }
    minprice:= prices[0]
    maxProfit:=0
    for i:=1;i<len(prices);i++{
        if prices[i]<minprice{
            minprice = prices[i]
        }
        if (prices[i]- minprice>maxProfit){
            maxProfit= prices[i]- minprice
        }
    }
    return maxProfit
}
*/
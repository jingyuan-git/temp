//Given a collection of candidate numbers (candidates) and a target number (targ
//et), find all unique combinations in candidates where the candidate numbers sum 
//to target. 
//
// Each number in candidates may only be used once in the combination. 
//
// Note: The solution set must not contain duplicate combinations. 
//
// 
// Example 1: 
//
// 
//Input: candidates = [10,1,2,7,6,1,5], target = 8
//Output: 
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
// 
//
// Example 2: 
//
// 
//Input: candidates = [2,5,2,1,2], target = 5
//Output: 
//[
//[1,2,2],
//[5]
//]
// 
//
// 
// Constraints: 
//
// 
// 1 <= candidates.length <= 100 
// 1 <= candidates[i] <= 50 
// 1 <= target <= 30 
// 
// Related Topics 数组 回溯算法 
// 👍 450 👎 0
package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	var dfs func(start int, temp []int, sum int)
	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				res = append(res, append([]int(nil), temp...))
			}
			return
		}
		for i:=start; i<len(candidates); i++ {
			if i-1 >= start && candidates[i-1] == candidates[i] {
				continue
			}
			temp = append(temp, candidates[i])
			dfs(i+1, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return
}
//leetcode submit region end(Prohibit modification and deletion)

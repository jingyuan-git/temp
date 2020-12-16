//Given a set of non-overlapping intervals, insert a new interval into the inter
//vals (merge if necessary). 
//
// You may assume that the intervals were initially sorted according to their st
//art times. 
//
// 
// Example 1: 
//
// 
//Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
//Output: [[1,5],[6,9]]
// 
//
// Example 2: 
//
// 
//Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//Output: [[1,2],[3,10],[12,16]]
//Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10]. 
//
//
// Example 3: 
//
// 
//Input: intervals = [], newInterval = [5,7]
//Output: [[5,7]]
// 
//
// Example 4: 
//
// 
//Input: intervals = [[1,5]], newInterval = [2,3]
//Output: [[1,5]]
// 
//
// Example 5: 
//
// 
//Input: intervals = [[1,5]], newInterval = [2,7]
//Output: [[1,7]]
// 
//
// 
// Constraints: 
//
// 
// 0 <= intervals.length <= 104 
// intervals[i].length == 2 
// 0 <= intervals[i][0] <= intervals[i][1] <= 105 
// intervals is sorted by intervals[i][0] in ascending order. 
// newInterval.length == 2 
// 0 <= newInterval[0] <= newInterval[1] <= 105 
// 
// Related Topics 排序 数组 
// 👍 338 👎 0
package main

// 此题中的newInterval 相当于上一题 所遍历的 [intervals[i][0], intervals[i][1]].
// 区别在于插入的灵活性更多一些：
//		前插可以立马写入res
//		后插（在所有元素都比newInterval小的情况下）。已知没有前插，那一定需要后插。
//leetcode submit region begin(Prohibit modification and deletion)
func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	insertBefore := false
	var res [][]int

	for _, interval := range intervals {
		if right < interval[0] {
			if !insertBefore {
				res = append(res, []int{left, right})
				insertBefore = true
			}
			res = append(res, interval)
		} else if left > interval[1] {
			res = append(res, interval)
		} else {
			// 有融合的部分,融合为新的newInterval
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	// 如果之前的迭代都没有完成前插
	// 则一定要将最后的newInterval后插如整个数组
	if !insertBefore {
		res = append(res, []int{left, right})
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)

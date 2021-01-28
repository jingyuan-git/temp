//Given an array of integers nums, write a method that returns the "pivot" index
// of this array. 
//
// We define the pivot index as the index where the sum of all the numbers to th
//e left of the index is equal to the sum of all the numbers to the right of the i
//ndex. 
//
// If no such index exists, we should return -1. If there are multiple pivot ind
//exes, you should return the left-most pivot index. 
//
// 
// Example 1: 
//
// 
//Input: nums = [1,7,3,6,5,6]
//Output: 3
//Explanation:
//The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the su
//m of numbers to the right of index 3.
//Also, 3 is the first index where this occurs.
// 
//
// Example 2: 
//
// 
//Input: nums = [1,2,3]
//Output: -1
//Explanation:
//There is no index that satisfies the conditions in the problem statement.
// 
//
// 
// Constraints: 
//
// 
// The length of nums will be in the range [0, 10000]. 
// Each element nums[i] will be an integer in the range [-1000, 1000]. 
// 
// Related Topics 数组 
// 👍 265 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	sum := 0
	for i, v := range nums {
		if sum*2 == total-v {
			return i
		}
		sum += v
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)







/*
func pivotIndex(nums []int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    sum := 0
    for i, v := range nums {
        if 2*sum+v == total {
            return i
        }
        sum += v
    }
    return -1
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/find-pivot-index/solution/xun-zhao-shu-zu-de-zhong-xin-suo-yin-by-gzjle/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
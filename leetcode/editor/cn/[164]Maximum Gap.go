//Given an integer array nums, return the maximum difference between two success
//ive elements in its sorted form. If the array contains less than two elements, r
//eturn 0. 
//
// 
// Example 1: 
//
// 
//Input: nums = [3,6,9,1]
//Output: 3
//Explanation: The sorted form of the array is [1,3,6,9], either (3,6) or (6,9) 
//has the maximum difference 3.
// 
//
// Example 2: 
//
// 
//Input: nums = [10]
//Output: 0
//Explanation: The array contains less than 2 elements, therefore return 0.
// 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 104 
// 0 <= nums[i] <= 109 
// 
//
// 
//Follow up: Could you solve it in linear time/space? Related Topics 排序 
// 👍 364 👎 0
package main
//TODO: 基排序 & 桶排序 可以在 O(N)O(N) 的时间内完成整数之间的排序。
//leetcode submit region begin(Prohibit modification and deletion)
func maximumGap(nums []int) int {

}
//leetcode submit region end(Prohibit modification and deletion)

/**
基数排序

func maximumGap(nums []int) (ans int) {
    n := len(nums)
    if n < 2 {
        return
    }

    buf := make([]int, n)
    maxVal := max(nums...)
    for exp := 1; exp <= maxVal; exp *= 10 {
        cnt := [10]int{}
        for _, v := range nums {
            digit := v / exp % 10
            cnt[digit]++
        }
        for i := 1; i < 10; i++ {
            cnt[i] += cnt[i-1]
        }
        for i := n - 1; i >= 0; i-- {
            digit := nums[i] / exp % 10
            buf[cnt[digit]-1] = nums[i]
            cnt[digit]--
        }
        copy(nums, buf)
    }

    for i := 1; i < n; i++ {
        ans = max(ans, nums[i]-nums[i-1])
    }
    return
}

func max(a ...int) int {
    res := a[0]
    for _, v := range a[1:] {
        if v > res {
            res = v
        }
    }
    return res
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/maximum-gap/solution/zui-da-jian-ju-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

/**
桶的算法

type pair struct{ min, max int }

func maximumGap(nums []int) (ans int) {
    n := len(nums)
    if n < 2 {
        return
    }

    minVal := min(nums...)
    maxVal := max(nums...)
    d := max(1, (maxVal-minVal)/(n-1))
    bucketSize := (maxVal-minVal)/d + 1

    // 存储 (桶内最小值，桶内最大值) 对，(-1, -1) 表示该桶是空的
    buckets := make([]pair, bucketSize)
    for i := range buckets {
        buckets[i] = pair{-1, -1}
    }
    for _, v := range nums {
        bid := (v - minVal) / d
        if buckets[bid].min == -1 {
            buckets[bid].min = v
            buckets[bid].max = v
        } else {
            buckets[bid].min = min(buckets[bid].min, v)
            buckets[bid].max = max(buckets[bid].max, v)
        }
    }

    prev := -1
    for i, b := range buckets {
        if b.min == -1 {
            continue
        }
        if prev != -1 {
            ans = max(ans, b.min-buckets[prev].max)
        }
        prev = i
    }
    return
}

func min(a ...int) int {
    res := a[0]
    for _, v := range a[1:] {
        if v < res {
            res = v
        }
    }
    return res
}

func max(a ...int) int {
    res := a[0]
    for _, v := range a[1:] {
        if v > res {
            res = v
        }
    }
    return res
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/maximum-gap/solution/zui-da-jian-ju-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
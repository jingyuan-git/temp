//A peak element is an element that is strictly greater than its neighbors. 
//
// Given an integer array nums, find a peak element, and return its index. If th
//e array contains multiple peaks, return the index to any of the peaks. 
//
// You may imagine that nums[-1] = nums[n] = -∞. 
//
// 
// Example 1: 
//
// 
//Input: nums = [1,2,3,1]
//Output: 2
//Explanation: 3 is a peak element and your function should return the index num
//ber 2. 
//
// Example 2: 
//
// 
//Input: nums = [1,2,1,3,5,6,4]
//Output: 5
//Explanation: Your function can return either index number 1 where the peak ele
//ment is 2, or index number 5 where the peak element is 6. 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 1000 
// -231 <= nums[i] <= 231 - 1 
// nums[i] != nums[i + 1] for all valid i. 
// 
//
// 
//Follow up: Could you implement a solution with logarithmic complexity? Related
// Topics 数组 二分查找 
// 👍 400 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func findPeakElement(nums []int) int {
	// TODO: 学习如何当作一个数学问题来简化
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}
//leetcode submit region end(Prohibit modification and deletion)

/*
public class Solution {
    public int findPeakElement(int[] nums) {
        int l = 0, r = nums.length - 1;
        while (l < r) {
            int mid = (l + r) / 2;
// 代表右边是降
            if (nums[mid] > nums[mid + 1])
                r = mid;
// 代码左边是升
            else
                l = mid + 1;
        }
        return l;
    }
}

作者：LeetCode
链接：https://leetcode-cn.com/problems/find-peak-element/solution/xun-zhao-feng-zhi-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
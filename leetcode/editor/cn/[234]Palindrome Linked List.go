//Given the head of a singly linked list, return true if it is a palindrome. 
//
// 
// Example 1: 
//
// 
//Input: head = [1,2,2,1]
//Output: true
// 
//
// Example 2: 
//
// 
//Input: head = [1,2]
//Output: false
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the list is in the range [1, 105]. 
// 0 <= Node.val <= 9 
// 
//
// 
//Follow up: Could you do it in O(n) time and O(1) space? Related Topics 栈 递归 链表
// 双指针 
// 👍 1039 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// 21.7.7
	dic := []int{}
	for head != nil {
		dic = append(dic, head.Val)
		head = head.Next
	}

	for i := 0; i < len(dic)/2; i++ {
		if dic[i] != dic[len(dic)-1-i] {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)

//Given the head of a linked list, remove the nth node from the end of the list 
//and return its head. 
//
// Follow up: Could you do this in one pass? 
//
// 
// Example 1: 
//
// 
//Input: head = [1,2,3,4,5], n = 2
//Output: [1,2,3,5]
// 
//
// Example 2: 
//
// 
//Input: head = [1], n = 1
//Output: []
// 
//
// Example 3: 
//
// 
//Input: head = [1,2], n = 1
//Output: [1]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the list is sz. 
// 1 <= sz <= 30 
// 0 <= Node.val <= 100 
// 1 <= n <= sz 
// 
// Related Topics 链表 双指针 
// 👍 1084 👎 0
package main

type ListNode struct {
	Val 	int
	Next 	*ListNode
}
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := head, dummy
	for i := 0; i < n; i++{
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next{
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)

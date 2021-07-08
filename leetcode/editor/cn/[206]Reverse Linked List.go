//Given the head of a singly linked list, reverse the list, and return the rever
//sed list. 
//
// 
// Example 1: 
//
// 
//Input: head = [1,2,3,4,5]
//Output: [5,4,3,2,1]
// 
//
// Example 2: 
//
// 
//Input: head = [1,2]
//Output: [2,1]
// 
//
// Example 3: 
//
// 
//Input: head = []
//Output: []
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the list is the range [0, 5000]. 
// -5000 <= Node.val <= 5000 
// 
//
// 
// Follow up: A linked list can be reversed either iteratively or recursively. C
//ould you implement both? 
// Related Topics 链表 
// 👍 1627 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 21.7.8
	// TODO: pre这两种定义差别很大！
	var pre *ListNode
	//pre := &ListNode{}
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre

	// 不必大费周章了
	//if head == nil {
	//	return nil
	//}
	//
	//root := &ListNode{0, nil}
	////root.Next = head
	//pre := root
	//root2 := head

	//var pre *ListNode
	//for head != nil {
	//	temp := head.Next
	//	head.Next = pre
	//	pre = head
	//	head = temp
	//}
	////root2.Next = nil
	//return pre
}
//leetcode submit region end(Prohibit modification and deletion)

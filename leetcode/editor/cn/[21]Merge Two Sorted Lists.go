//Merge two sorted linked lists and return it as a new sorted list. The new list
// should be made by splicing together the nodes of the first two lists. 
//
// 
// Example 1: 
//
// 
//Input: l1 = [1,2,4], l2 = [1,3,4]
//Output: [1,1,2,3,4,4]
// 
//
// Example 2: 
//
// 
//Input: l1 = [], l2 = []
//Output: []
// 
//
// Example 3: 
//
// 
//Input: l1 = [], l2 = [0]
//Output: [0]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in both lists is in the range [0, 50]. 
// -100 <= Node.val <= 100 
// Both l1 and l2 are sorted in non-decreasing order. 
// 
// Related Topics 链表 
// 👍 1345 👎 0

package main

//基本上马上可以想到需要设置一个哨兵节点，这可以在最后让我们比较容易地返回合并后的链表。
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	prehead := &ListNode{}
//	result := prehead
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			prehead.Next = l1
//			prehead = prehead.Next
//			l1 = l1.Next
//		} else {
//			prehead.Next = l2
//			prehead = prehead.Next
//			l2 = l2.Next
//		}
//	}
//
//	if l1 != nil{
//		prehead.Next = l1
//	}
//	if l2 != nil{
//		prehead.Next = l2
//	}
//
//	return result.Next
//}

// 21.6.10
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	}
	if l2 == nil {
		cur.Next = l1
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

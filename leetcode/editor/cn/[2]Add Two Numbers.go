//You are given two non-empty linked lists representing two non-negative integer
//s. The digits are stored in reverse order, and each of their nodes contains a si
//ngle digit. Add the two numbers and return the sum as a linked list. 
//
// You may assume the two numbers do not contain any leading zero, except the nu
//mber 0 itself. 
//
// 
// Example 1: 
//
// 
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.
// 
//
// Example 2: 
//
// 
//Input: l1 = [0], l2 = [0]
//Output: [0]
// 
//
// Example 3: 
//
// 
//Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//Output: [8,9,9,9,0,0,0,1]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in each linked list is in the range [1, 100]. 
// 0 <= Node.val <= 9 
// It is guaranteed that the list represents a number that does not have leading
// zeros. 
// 
// Related Topics 链表 数学 
// 👍 5046 👎 0

package main

import "aaa/test/link"

type ListNode struct {
	Val int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var tail, head *ListNode
//	carry := 0
//
//	for l1 != nil || l2 != nil {
//		v1, v2 := 0, 0
//		if l1 != nil{
//			v1 = l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil{
//			v2 = l2.Val
//			l2 = l2.Next
//		}
//		sum := v1 + v2 +carry
//		carry = sum/10
//		sum = sum%10
//		if head == nil{
//			head = &ListNode{sum, nil}
//			tail = head
//		}else {
//			tail.Next = &ListNode{sum, nil}
//			tail = tail.Next
//		}
//	}
//	if carry > 0{
//		tail.Next = &ListNode{carry, nil}
//	}
//	return head
//}

//  21.6.5
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	mod := 0
	res := &ListNode{}
	dummy := res

	var v1, v2 int
	for l1 != nil || l2 != nil {
		v1, v2 = 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := (v1 + v2 + mod) % 10
		mod = (v1 + v2 + mod) / 10
		res.Next = &ListNode{Val: sum}
		res = res.Next
	}
	if mod != 0 {
		res.Next = &ListNode{mod, nil}
	}
	return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}

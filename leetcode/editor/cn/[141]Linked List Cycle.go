//Given head, the head of a linked list, determine if the linked list has a cycl
//e in it. 
//
// There is a cycle in a linked list if there is some node in the list that can 
//be reached again by continuously following the next pointer. Internally, pos is 
//used to denote the index of the node that tail's next pointer is connected to. N
//ote that pos is not passed as a parameter. 
//
// Return true if there is a cycle in the linked list. Otherwise, return false. 
//
//
// 
// Example 1: 
//
// 
//Input: head = [3,2,0,-4], pos = 1
//Output: true
//Explanation: There is a cycle in the linked list, where the tail connects to t
//he 1st node (0-indexed).
// 
//
// Example 2: 
//
// 
//Input: head = [1,2], pos = 0
//Output: true
//Explanation: There is a cycle in the linked list, where the tail connects to t
//he 0th node.
// 
//
// Example 3: 
//
// 
//Input: head = [1], pos = -1
//Output: false
//Explanation: There is no cycle in the linked list.
// 
//
// 
// Constraints: 
//
// 
// The number of the nodes in the list is in the range [0, 104]. 
// -105 <= Node.val <= 105 
// pos is -1 or a valid index in the linked-list. 
// 
//
// 
// Follow up: Can you solve it using O(1) (i.e. constant) memory? 
// Related Topics 链表 双指针 
// 👍 916 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	//21.6.24
	// 1. 特殊情况：如果只有1个节点，或者链表为空。则返回
	if head == nil || head.Next == nil {
		return false
	}

	// 2. 快慢链表
	// 刚开始low，fast就要错位，来避免链表太短 fast根本没有移动就 一致的情况
	low, fast := head, head.Next
	for fast != nil && fast.Next != nil && fast != low {
		low = low.Next
		fast = fast.Next.Next
	}

	if fast == low {
		return true
	} else {
		return false
	}

	//if head == nil || head.Next == nil {
	//	return false
	//}
	//
    //root := &ListNode{
    //	Next: head,
	//}
	//slow := root
	//fast := root.Next
	//
	//for fast != nil && fast.Next != nil && fast != slow {
	//	slow = slow.Next
	//	fast = fast.Next.Next
	//}
	//
	//if fast == slow {
	//	return true
	//} else {
	//	return false
	//}
}
//leetcode submit region end(Prohibit modification and deletion)

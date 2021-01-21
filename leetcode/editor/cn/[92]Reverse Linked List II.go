//Reverse a linked list from position m to n. Do it in one-pass. 
//
// Note: 1 ≤ m ≤ n ≤ length of list. 
//
// Example: 
//
// 
//Input: 1->2->3->4->5->NULL, m = 2, n = 4
//Output: 1->4->3->2->5->NULL
// 
// Related Topics 链表 
// 👍 646 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	root := &ListNode{}

	root.Next = head
	front := root
	cur := head

	// 越过不用反转的部分
	for i := 1; i < m; i++ {
		front = cur
		cur = cur.Next
	}
	// 标记反转链内的头头，等反转之后就是尾部了
	end := cur

	// 链内反转
	pre := &ListNode{}
	for i := m; i <= n; i++ {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur =temp
	}

	// 反转的链和外部拼接。pre是反转链内的末尾，反转后为头部
	front.Next = pre
	end.Next = cur
	return root.Next
}
//leetcode submit region end(Prohibit modification and deletion)

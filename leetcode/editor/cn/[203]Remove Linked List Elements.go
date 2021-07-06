//Given the head of a linked list and an integer val, remove all the nodes of th
//e linked list that has Node.val == val, and return the new head. 
//
// 
// Example 1: 
//
// 
//Input: head = [1,2,6,3,4,5,6], val = 6
//Output: [1,2,3,4,5]
// 
//
// Example 2: 
//
// 
//Input: head = [], val = 1
//Output: []
// 
//
// Example 3: 
//
// 
//Input: head = [7,7,7,7], val = 7
//Output: []
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the list is in the range [0, 104]. 
// 1 <= Node.val <= 50 
// 0 <= val <= 50 
// 
// Related Topics 递归 链表 
// 👍 653 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// 21.7.6
	root := &ListNode{Next: head}
	pre := root
	for pre.Next != nil {
		// 若等于则pre位置不变，pre.Next向后移动
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			// pre向后移一位 继续试探
			pre = pre.Next
		}
	}
	return root.Next
}
//leetcode submit region end(Prohibit modification and deletion)

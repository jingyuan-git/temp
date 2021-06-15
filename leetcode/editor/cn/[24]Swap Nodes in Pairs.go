//Given a linked list, swap every two adjacent nodes and return its head. 
//
// You may not modify the values in the list's nodes. Only nodes itself may be c
//hanged. 
//
// 
// Example 1: 
//
// 
//Input: head = [1,2,3,4]
//Output: [2,1,4,3]
// 
//
// Example 2: 
//
// 
//Input: head = []
//Output: []
// 
//
// Example 3: 
//
// 
//Input: head = [1]
//Output: [1]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the list is in the range [0, 100]. 
// 0 <= Node.val <= 100 
// 
// Related Topics 链表 
// 👍 728 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	//pre := &ListNode{Val: 0, Next: head}
	//res := pre
	//cur := head
	//if head.Next == nil{
	//	return head
	//}
	//post := head.Next
	//for cur.Next != nil{
	//	temp := pre.Next
	//	pre.Next = post
	//	post.Next = cur
	//	cur.Next = temp
	//
	//	pre = pre.Next.Next
	//	cur = cur.Next.Next
	//	post = post.Next.Next
	//}
	//return res.Next

	//dummyHead := &ListNode{Val: 0, Next: head}
	//pre := dummyHead
	//for pre.Next != nil && pre.Next.Next != nil{
	//	// 赋值
	//	cur := pre.Next
	//	post := pre.Next.Next
	//	// swap every two adjacent nodes
	//	pre.Next = post
	//	cur.Next = post.Next
	//	post.Next = cur
	//	pre = cur	// 此时的cur已经是调转之后（前post的位置）, 相当于pre往后移动两个链
	//}
	//return dummyHead.Next

	// 21.6.15
	dummy := &ListNode{Next: head}
	pre := dummy
	// 此行非常重要，不然没有进入for循环的必要了
	for pre.Next != nil && pre.Next.Next != nil {
		cur := pre.Next
		post := pre.Next.Next

		pre.Next = cur.Next
		cur.Next = post.Next
		post.Next = cur

		pre = cur
	}
	return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)

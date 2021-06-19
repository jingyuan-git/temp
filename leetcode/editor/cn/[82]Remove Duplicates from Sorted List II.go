//Given the head of a sorted linked list, delete all nodes that have duplicate n
//umbers, leaving only distinct numbers from the original list. Return the linked 
//list sorted as well. 
//
// 
// Example 1: 
//
// 
//Input: head = [1,2,3,3,4,4,5]
//Output: [1,2,5]
// 
//
// Example 2: 
//
// 
//Input: head = [1,1,1,2,3]
//Output: [2,3]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the list is in the range [0, 300]. 
// -100 <= Node.val <= 100 
// The list is guaranteed to be sorted in ascending order. 
// 
// Related Topics 链表 
// 👍 429 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// TODO: 我之前考虑的是pre，current，next。比较好的写法：把head指针提前一位，然后判断head.next与head.next.next
func deleteDuplicates(head *ListNode) *ListNode {
	// 21.6.18
	// 21.6.17
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur != nil && cur.Next != nil {
		post := cur.Next.Next
		count := 0
		// current保持不动，post继续向后试探
		for post != nil && post.Val == cur.Next.Val {
			count++
			post = post.Next

		}
		// 根据计数判断，若是重复节点则把cur.Next的节点覆盖掉。
		//否则cur继续向后移动：当前cur.Next一定一定不是重复元素，cur往后移动后，再判断cur.Next
		if count > 0 {
			cur.Next = post
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)
/*
func deleteDuplicates(head *ListNode) *ListNode {
    node := &ListNode{0, head}
	head, first := node, node
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			first.Val = head.Next.Val
			for head.Next != nil && first.Val == head.Next.Val {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return node.Next
}

作者：user1775a
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/solution/golang-yi-ci-xun-huan-by-user1775a/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
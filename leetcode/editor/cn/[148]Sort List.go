//Given the head of a linked list, return the list after sorting it in ascending
// order. 
//
// Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.
//e. constant space)? 
//
// 
// Example 1: 
//
// 
//Input: head = [4,2,1,3]
//Output: [1,2,3,4]
// 
//
// Example 2: 
//
// 
//Input: head = [-1,5,3,4,0]
//Output: [-1,0,3,4,5]
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
// The number of nodes in the list is in the range [0, 5 * 104]. 
// -105 <= Node.val <= 105 
// 
// Related Topics 排序 链表 
// 👍 1064 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// TODO: 归并排序的方法
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyNode := &ListNode{Next: head}
	lastedSorted, cur := head, head.Next

	for cur != nil {
		if cur.Val >= lastedSorted.Val {
			lastedSorted = lastedSorted.Next
			//lastedSorted = cur
			//cur = cur.Next
		} else {
			pre := dummyNode
			for pre.Next.Val <= cur.Val {
				pre = pre.Next
			}
			lastedSorted.Next = cur.Next

			cur.Next = pre.Next
			pre.Next = cur
			//NeedSorted := cur
			//cur = cur.Next
			//fmt.Printf("NeedSorted.Next=%+v, \"cur=%+v\n", NeedSorted.Next,cur)
			//lastedSorted.Next = cur
			//fmt.Printf("cur=%+v\n", cur, NeedSorted, "\n")
			//
			//pre := dummyNode
			//for pre.Next.Val <= NeedSorted.Val {
			//	pre = pre.Next
			//}
			//
			//temp := pre.Next
			//pre.Next = NeedSorted
			//NeedSorted.Next = temp
			//
			//// 如果在这里重新赋值，cur已经被修改了！！！
			////fmt.Printf("cur=%+v\n", cur, NeedSorted)
			////cur = cur.Next
			////lastedSorted.Next = cur
			//
			////insertNode := dummyNode.Next
			////for insertNode.Next.Val <= cur.Val {
			////	insertNode = insertNode.Next
			////}
			////
			////temp := insertNode.Next
			////insertNode.Next = cur
			////lastedSorted.Next = cur.Next
			////cur.Next = temp
			////
			////cur = lastedSorted.Next
		}
		cur = lastedSorted.Next
	}
	return dummyNode.Next
}
//leetcode submit region end(Prohibit modification and deletion)

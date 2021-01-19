//Given the head of a linked list and a value x, partition it such that all node
//s less than x come before nodes greater than or equal to x. 
//
// You should preserve the original relative order of the nodes in each of the t
//wo partitions. 
//
// 
// Example 1: 
//
// 
//Input: head = [1,4,3,2,5,2], x = 3
//Output: [1,2,2,4,3,5]
// 
//
// Example 2: 
//
// 
//Input: head = [2,1], x = 2
//Output: [1,2]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the tree is in the range [0, 200]. 
// -100 <= Node.val <= 100 
// -200 <= x <= 200 
// 
// Related Topics 链表 双指针 
// 👍 354 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	// TODO：由于要将链表分为两部分。所以定义2个前置定位链表头，和2个移动的节点
	smail := &ListNode{}
	large := &ListNode{}
	smailHead := smail
	largeHead := large

	for head != nil {
		if head.Val < x {
			smail.Next = head
			smail = smail.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	// TODO: 注意要置链尾为空
	large.Next = nil
	smail.Next = largeHead.Next
	return smailHead.Next
}
//leetcode submit region end(Prohibit modification and deletion)

/*
func partition(head *ListNode, x int) *ListNode {
    small := &ListNode{}
    smallHead := small
    large := &ListNode{}
    largeHead := large
    for head != nil {
        if head.Val < x {
            small.Next = head
            small = small.Next
        } else {
            large.Next = head
            large = large.Next
        }
        head = head.Next
    }
    large.Next = nil
    small.Next = largeHead.Next
    return smallHead.Next
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/partition-list/solution/fen-ge-lian-biao-by-leetcode-solution-7ade/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
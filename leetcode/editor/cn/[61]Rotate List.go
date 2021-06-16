//Given the head of a linked list, rotate the list to the right by k places. 
//
// 
// Example 1: 
//
// 
//Input: head = [1,2,3,4,5], k = 2
//Output: [4,5,1,2,3]
// 
//
// Example 2: 
//
// 
//Input: head = [0,1,2], k = 4
//Output: [2,0,1]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the list is in the range [0, 500]. 
// -100 <= Node.val <= 100 
// 0 <= k <= 2 * 109 
// 
// Related Topics 链表 双指针 
// 👍 391 👎 0
package main

import "fmt"

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
func rotateRight(head *ListNode, k int) *ListNode {
	//if head == nil || k == 0 {
	//	return head
	//}
	//
	//n := 1
	//temp := head
	//for temp.Next != nil {
	//	n++
	//	temp = temp.Next
	//}
	//
	//m := k % n
	////fmt.Println(m, n)
	//temp.Next = head
	//
	//// 此时的temp还在原本链的末尾
	//for i := 1; i <= n - m; i++ {
	//	temp = temp.Next
	//}
	//
	//head = temp.Next
	//temp.Next = nil
	//return head

	//21.6.16
	//0. 处理特殊值
	if head == nil || k == 0 {
		return head
	}

	//1. 计算长度
	dummy := &ListNode{Next: head}
	length := 0
	for head != nil {
		length++
		head = head.Next
	}

	//2. 设置快慢指针
	head = dummy.Next
	mod := k % length
	low, fast := head, head
	for mod > 0 {
		fast = fast.Next
		mod--
	}

	//3. 快慢指针同时进行，直到快指针结束
	for fast.Next != nil {
		fast = fast.Next
		low = low.Next
	}

	//4. 旋转指针
	fast.Next = dummy.Next
	dummy.Next = low.Next
	// 结尾置空
	low.Next = nil
	return dummy.Next
}
//func rotateRight(head *ListNode, k int) *ListNode {
//	if head == nil || k == 0 {
//		return head
//	}
//
//	// 获得整个链的长度
//	n, p := 1, head
//	for p.Next != nil {
//		p = p.Next
//		n++
//	}
//
//	// 将原先end的空next填补
//	p.Next = head
//
//	// 将旋转节点的前半部分末尾的next置空
//	k %= n
//	for i := 1; i <= n-k; i++ {
//		p = p.Next
//	}
//	head, p.Next = p.Next, nil
//	return head
//}
//leetcode submit region end(Prohibit modification and deletion)

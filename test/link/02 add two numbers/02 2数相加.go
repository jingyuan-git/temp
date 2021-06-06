package main

import (
	"aaa/test/link"
	"fmt"
)

//// 缺点：有太多冗余的代码，其实就是计算 sum和mod，可以用一个for大循环，里面加一些条件判断
//func addTwoNumbers(l1 *link.ListNode, l2 *link.ListNode) *link.ListNode {
//	if l1 == nil && l2 == nil {
//		return &link.ListNode{}
//	}
//	res := &link.ListNode{}
//	dummy := res
//	mod := 0
//	for l1 != nil && l2 != nil {
//		res.Next = &link.ListNode{(l1.Val + l2.Val + mod) % 10, nil}
//		mod = (l1.Val + l2.Val + mod) / 10
//		res = res.Next
//		l1 = l1.Next
//		l2 = l2.Next
//	}
//
//	for l1 == nil && l2 == nil {
//		if mod == 1 {
//			res.Next = &link.ListNode{1, nil}
//		}
//		return dummy.Next
//	}
//
//	for l1 != nil {
//		res.Next = &link.ListNode{(l1.Val + mod) % 10, nil}
//		mod = (l1.Val + mod) / 10
//		res = res.Next
//		l1 = l1.Next
//	}
//
//	for l2 != nil {
//		res.Next = &link.ListNode{(l2.Val + mod) % 10, nil}
//		mod = (l2.Val + mod) / 10
//		res = res.Next
//		l2 = l2.Next
//	}
//	if mod == 1 {
//		res.Next = &link.ListNode{1, nil}
//	}
//	return dummy.Next
//}

func addTwoNumbers(l1 *link.ListNode, l2 *link.ListNode) *link.ListNode {
	mod := 0
	res := &link.ListNode{}
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
		res.Next = &link.ListNode{Val: sum}
		res = res.Next
	}
	if mod != 0 {
		res.Next = &link.ListNode{mod, nil}
	}
	return dummy.Next
}

func main() {
	//nums1 := []int{2,4,3}
	//nums2 := []int{5,6,4}
	nums1 := []int{9,9,9,9,9,9,9}
	nums2 := []int{9,9,9,9}
	l1 := link.CreatLink(nums1)
	l2 := link.CreatLink(nums2)
	res := link.PrintLink(addTwoNumbers(l1, l2))
	fmt.Println(res)
}
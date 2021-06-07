package main

import (
	"aaa/test/link"
	"fmt"
)

func removeNthFromEnd(head *link.ListNode, n int) *link.ListNode {
	dummy := &link.ListNode{0, head}
	pre, cur := dummy, dummy
	ppre :=  &link.ListNode{}
	for n > 0 {
		cur = cur.Next
		n--
	}
	for cur != nil {
		cur = cur.Next
		ppre = pre
		pre = pre.Next
	}

	ppre.Next = pre.Next

	return dummy.Next
}

func main() {
	//nums := []int{1,2,3,4,5}
	nums := []int{1,2}
	l1 := link.CreatLink(nums)
	res := link.PrintLink(removeNthFromEnd(l1, 1))
	fmt.Println(res)
}
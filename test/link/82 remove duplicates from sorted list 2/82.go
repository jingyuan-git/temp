package main

import (
	"aaa/test/link"
	"fmt"
)

func deleteDuplicates(head *link.ListNode) *link.ListNode {
	// 21.6.17
	dummy := &link.ListNode{Next: head}
	cur := dummy
	//for cur != nil && cur.Next != nil && cur.Next.Next != nil {
	for cur != nil && cur.Next != nil {
		post := cur.Next.Next
		count := 0
		// current保持不动，post继续向后试探
		for post != nil && post.Val == cur.Next.Val {
			count++
			post = post.Next

		}
		// 根据计数判断，若是重复节点则把cur.Next的节点覆盖掉。否则cur继续向后移动
		if count > 0 {
			cur.Next = post
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}


func main() {
	nums := []int{1,1,1,2,3}
	//nums := []int{1,2,3,3,4,4,5}
	l1 := link.CreatLink(nums)
	res := link.PrintLink(deleteDuplicates(l1))
	fmt.Println(res)
}

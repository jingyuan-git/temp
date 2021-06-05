package link

type ListNode struct {
	Val int
	Next *ListNode
}

func CreatLink(nums []int) *ListNode {
	l := &ListNode{}
	dummy := l

	for _, v := range nums {
		l.Next = &ListNode{v, nil}
		l = l.Next
	}
	l = dummy.Next
	return l
}

//func main() {
//	nums := []int{1, 2, 3, 4}
//	l := CreatLink(nums)
//
//	for l != nil {
//		fmt.Println(l.Val)
//		l = l.Next
//	}
//}
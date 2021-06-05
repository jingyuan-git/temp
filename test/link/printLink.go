package link

func PrintLink(l *ListNode) []int {
	nums := []int{}
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return nums
}

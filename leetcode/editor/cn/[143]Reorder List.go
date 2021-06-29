//You are given the head of a singly linked-list. The list can be represented as
//: 
//
// 
//L0 → L1 → … → Ln - 1 → Ln
// 
//
// Reorder the list to be on the following form: 
//
// 
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 
//
// You may not modify the values in the list's nodes. Only nodes themselves may 
//be changed. 
//
// 
// Example 1: 
//
// 
//Input: head = [1,2,3,4]
//Output: [1,4,2,3]
// 
//
// Example 2: 
//
// 
//Input: head = [1,2,3,4,5]
//Output: [1,5,2,4,3]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the list is in the range [1, 5 * 104]. 
// 1 <= Node.val <= 1000 
// 
// Related Topics 栈 递归 链表 双指针 
// 👍 604 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	// 21.6.29
	// 1. 特殊情况
	if head == nil || head.Next == nil {
		return
	}

	// 2. 遍历，并存储
	dic := []*ListNode{}
	for head != nil {
	dic = append(dic, head)
	head = head.Next
	}

	// 3. 首尾拼接
	for i := 0; i < len(dic)/2; i++ {
		dic[i].Next = dic[len(dic) - 1 - i]
		dic[len(dic) - 1 - i].Next = dic[i+1]
	}
	// 4. 中介链表置空，防止回环
	dic[len(dic)/2].Next = nil
	return
}
//leetcode submit region end(Prohibit modification and deletion)

/*
   //leetcode submit region begin(Prohibit modification and deletion)
   /**
    * Definition for singly-linked list.
    * type ListNode struct {
    *     Val int
    *     Next *ListNode
    * }
*/


//leetcode submit region end(Prohibit modification and deletion)


//TODO: 方法二
//方法二：寻找链表中点 + 链表逆序 + 合并链表
//注意到目标链表即为将原链表的左半端和反转后的右半端合并后的结果。
//
//这样我们的任务即可划分为三步：
//
//找到原链表的中点（参考「876. 链表的中间结点」）。
//
//我们可以使用快慢指针来 O(N)O(N) 地找到链表的中间节点。
//将原链表的右半端反转（参考「206. 反转链表」）。
//
//我们可以使用迭代法实现链表的反转。
//将原链表的两端合并。
//
//因为两链表长度相差不超过 11，因此直接合并即可。
//
//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/reorder-list/solution/zhong-pai-lian-biao-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

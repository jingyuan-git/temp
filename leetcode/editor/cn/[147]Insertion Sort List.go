//Given the head of a singly linked list, sort the list using insertion sort, an
//d return the sorted list's head. 
//
// The steps of the insertion sort algorithm: 
//
// 
// Insertion sort iterates, consuming one input element each repetition and grow
//ing a sorted output list. 
// At each iteration, insertion sort removes one element from the input data, fi
//nds the location it belongs within the sorted list and inserts it there. 
// It repeats until no input elements remain. 
// 
//
// The following is a graphical example of the insertion sort algorithm. The par
//tially sorted list (black) initially contains only the first element in the list
//. One element (red) is removed from the input data and inserted in-place into th
//e sorted list with each iteration. 
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
// 
// Constraints: 
//
// 
// The number of nodes in the list is in the range [1, 5000]. 
// -5000 <= Node.val <= 5000 
// 
// Related Topics 排序 链表 
// 👍 368 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 精辟，只需赋值next，总而增加链头
	root := &ListNode{Next: head}

	lastSorted, cur := head, head.Next

	for cur != nil {
		// trick1: 根据边界，判断是否需要插入
		if cur.Val >= lastSorted.Val {
			lastSorted = cur
			cur = cur.Next
		} else {
			needInsert := cur
			cur = cur.Next
			lastSorted.Next = cur

			pre := root
			// trick2: 快速遍历无需变化的前半部分。找到pre
			for pre.Next.Val <= needInsert.Val {
				pre = pre.Next
			}
			temp := pre.Next
			pre.Next = needInsert
			needInsert.Next = temp
		}
	}

	// 思路混乱，毫无逻辑。就很容易写错了
	//for head.Next != nil {
	//	cur := root
	//	for cur.Next != nil {
	//		if cur.Next.Val > head.Val {
	//			temp := cur.Next
	//			cur.Next = head
	//			pre.Next = head.Next
	//			head.Next = temp
	//			break
	//		} else {
	//			pre = cur
	//			cur = cur.Next
	//		}
	//	}
	//	pre = pre.Next
	//}
	return root.Next
}
//leetcode submit region end(Prohibit modification and deletion)

/**
func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummyHead := &ListNode{Next: head}
    lastSorted, curr := head, head.Next
    for curr != nil {
        if lastSorted.Val <= curr.Val {
            lastSorted = lastSorted.Next
        } else {
            prev := dummyHead
            for prev.Next.Val <= curr.Val {
                prev = prev.Next
            }
            lastSorted.Next = curr.Next
            curr.Next = prev.Next
            prev.Next = curr
        }
        curr = lastSorted.Next
    }
    return dummyHead.Next
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/insertion-sort-list/solution/dui-lian-biao-jin-xing-cha-ru-pai-xu-by-leetcode-s/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
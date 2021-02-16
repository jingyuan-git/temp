//Given a binary tree 
//
// 
//struct Node {
//  int val;
//  Node *left;
//  Node *right;
//  Node *next;
//}
// 
//
// Populate each next pointer to point to its next right node. If there is no ne
//xt right node, the next pointer should be set to NULL. 
//
// Initially, all next pointers are set to NULL. 
//
// 
//
// Follow up: 
//
// 
// You may only use constant extra space. 
// Recursive approach is fine, you may assume implicit stack space does not coun
//t as extra space for this problem. 
// 
//
// 
// Example 1: 
//
// 
//
// 
//Input: root = [1,2,3,4,5,null,7]
//Output: [1,#,2,3,#,4,5,7,#]
//Explanation: Given the above binary tree (Figure A), your function should popu
//late each next pointer to point to its next right node, just like in Figure B. T
//he serialized output is in level order as connected by the next pointers, with '
//#' signifying the end of each level.
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the given tree is less than 6000. 
// -100 <= node.val <= 100 
// 
// Related Topics 树 深度优先搜索 
// 👍 361 👎 0

package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	//if root == nil {
	//	return nil
	//}
	//
	//queue := []*Node{root}
	//
	//for len(queue) > 0 {
	//	temp := queue
	//	queue = []*Node{}
	//	for i := 0; i < len(temp); i++ {
	//		if i > 0{
	//			temp[i-1].Next = temp[i]
	//		}
	//
	//		if temp[i].Left != nil {
	//			queue = append(queue, temp[i].Left)
	//		}
	//
	//		if temp[i].Right != nil {
	//			queue = append(queue, temp[i].Right)
	//		}
	//	}
	//}
	//return root

	// method 2:
	if root == nil {
		return root
	}

	start := root
	for start != nil {
		var first, last *Node
		// 可以在上一层为下一层建立 \text{next}next 指针。该方法最重要的一点是：位于第 xx 层时为第 x + 1x+1 层建立 \text{next}next 指针。
		//
		//作者：LeetCode-Solution
		//链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/solution/tian-chong-mei-ge-jie-dian-de-xia-yi-ge-you-ce-15/
		//来源：力扣（LeetCode）
		//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
		handle := func(cur *Node) {
			if cur == nil {
				return
			}

			// first为下一行第一个元素
			if first == nil {
				first = cur
			}

			// last为cur的同一行的前一个元素
			if last != nil {
				last.Next = cur
			}

			last = cur
		}

		for cur := start; cur != nil; cur = cur.Next {
			handle(cur.Left)
			handle(cur.Right)
		}
		start = first
	}
	return root
}
//leetcode submit region end(Prohibit modification and deletion)

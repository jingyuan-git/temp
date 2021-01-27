//Given a binary tree, determine if it is height-balanced. 
//
// For this problem, a height-balanced binary tree is defined as: 
//
// 
// a binary tree in which the left and right subtrees of every node differ in he
//ight by no more than 1. 
// 
//
// 
// Example 1: 
//
// 
//Input: root = [3,9,20,null,null,15,7]
//Output: true
// 
//
// Example 2: 
//
// 
//Input: root = [1,2,2,3,3,null,null,4,4]
//Output: false
// 
//
// Example 3: 
//
// 
//Input: root = []
//Output: true
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the tree is in the range [0, 5000]. 
// -104 <= Node.val <= 104 
// Related Topics 树 深度优先搜索 递归 
// 👍 576 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(x int, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func abs(x int) int {
	if x < 0{
		return -x
	}
	return x
}

// TODO：差点忘了怎么算二叉树的最大高度了
func depth(head *TreeNode) int {
	if head ==nil {
		return 0
	}
	return max(depth(head.Left), depth(head.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(depth(root.Left) - depth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}
//leetcode submit region end(Prohibit modification and deletion)

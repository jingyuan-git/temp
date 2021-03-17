//Given two integer arrays inorder and postorder where inorder is the inorder tr
//aversal of a binary tree and postorder is the postorder traversal of the same tr
//ee, construct and return the binary tree. 
//
// 
// Example 1: 
//
// 
//Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
//Output: [3,9,20,null,null,15,7]
// 
//
// Example 2: 
//
// 
//Input: inorder = [-1], postorder = [-1]
//Output: [-1]
// 
//
// 
// Constraints: 
//
// 
// 1 <= inorder.length <= 3000 
// postorder.length == inorder.length 
// -3000 <= inorder[i], postorder[i] <= 3000 
// inorder and postorder consist of unique values. 
// Each value of postorder also appears in inorder. 
// inorder is guaranteed to be the inorder traversal of the tree. 
// postorder is guaranteed to be the postorder traversal of the tree. 
// 
// Related Topics 树 深度优先搜索 数组 
// 👍 465 👎 0
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]
	root := &TreeNode{Val: val}
	index := -1

	for i, k := range inorder {
		if k == val {
			index = i
		}
	}
	root.Left = buildTree(inorder[:index], postorder[index+1:])
	root.Right = buildTree(inorder[:index], postorder[index:len(postorder)-2])
	return root
}
//leetcode submit region end(Prohibit modification and deletion)

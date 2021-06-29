//Given the root of a binary tree, return the preorder traversal of its nodes' v
//alues. 
//
// 
// Example 1: 
//
// 
//Input: root = [1,null,2,3]
//Output: [1,2,3]
// 
//
// Example 2: 
//
// 
//Input: root = []
//Output: []
// 
//
// Example 3: 
//
// 
//Input: root = [1]
//Output: [1]
// 
//
// Example 4: 
//
// 
//Input: root = [1,2]
//Output: [1,2]
// 
//
// Example 5: 
//
// 
//Input: root = [1,null,2]
//Output: [1,2]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the tree is in the range [0, 100]. 
// -100 <= Node.val <= 100 
// 
//
// 
// Follow up: Recursive solution is trivial, could you do it iteratively? 
// Related Topics 栈 树 
// 👍 540 👎 0
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

var result []int
func preorderTraversal(root *TreeNode) (result []int) {
	// 21.6.29
	if root == nil {
		return
	}
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result

	//// Method 1: recursion
	////var preorder func(*TreeNode)
	////preorder = func(node *TreeNode) {
	////	if node == nil {
	////		return
	////	}
	////
	////	result = append(result, node.Val)
	////	preorder(node.Left)
	////	preorder(node.Right)
	////}
	////preorder(root)
	////return
	//////if root == nil {
	//////	return nil
	//////}
	//////
	//////result = append(result, root.Val)
	//////preorderTraversal(root.Left)
	//////preorderTraversal(root.Right)
	//////return
	//
	//// Method 2: iteration
	//stack := []*TreeNode{}
	//node := root
	//
	//for node != nil || len(stack) > 0 {
	//
	//	for node != nil {
	//		result = append(result, node.Val)
	//		stack = append(stack, node)
	//		node = node.Left
	//	}
	//
	//	node = stack[len(stack)-1].Right
	//	stack = stack[:len(stack)-1]
	//
	//	//stack = append(stack, node)
	//	//node = stack[0]
	//	//
	//	//if node != nil && node.Left != nil {
	//	//	stack = append(stack, node.Left)
	//	//} else {
	//	//	if node != nil && node.Right != nil {
	//	//		stack = append(stack, node.Right)
	//	//	}
	//	//}
	//	//
	//	//result = append(result, node.Val)
	//	//stack = stack[1:]
	//}
	//return result
}
//leetcode submit region end(Prohibit modification and deletion)

//Given the root of a binary tree, determine if it is a valid binary search tree
// (BST). 
//
// A valid BST is defined as follows: 
//
// 
// The left subtree of a node contains only nodes with keys less than the node's
// key. 
// The right subtree of a node contains only nodes with keys greater than the no
//de's key. 
// Both the left and right subtrees must also be binary search trees. 
// 
//
// 
// Example 1: 
//
// 
//Input: root = [2,1,3]
//Output: true
// 
//
// Example 2: 
//
// 
//Input: root = [5,1,4,null,null,3,6]
//Output: false
//Explanation: The root node's value is 5 but its right child's value is 4.
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the tree is in the range [1, 104]. 
// -231 <= Node.val <= 231 - 1 
// 
// Related Topics 树 深度优先搜索 递归 
// 👍 917 👎 0

package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func helper(root *TreeNode, low int, high int) bool {
	if root == nil {
		return true
	}

	if root.Val <= low || root.Val >= high {
		return false
	}

	return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, high)
}

func isValidBST(root *TreeNode) bool {
	// TODO: 注意最大值 最小值的范围
	return helper(root, math.MinInt64, math.MaxInt64)


	//// TODO: 不能单纯的比较左节点小于中间节点，右节点大于中间节点就完事了。
	//// **我们要比较的是 左子树所有节点小于中间节点，右子树所有节点大于中间节点。**所以以上代码的判断逻辑是错误的。
	//if root == nil{
	//	return true
	//}
	//
	//if root.Right != nil && root.Right.Val <= root.Val {
	//	return false
	//}
	//
	//if root.Left != nil && root.Left.Val >= root.Val {
	//	return false
	//}
	//
	////if root.Left.Val < root.Val && root.Right.Val > root.Val {
	////	return true
	////}
	////
	////return isValidBST(root.Left) && isValidBST(root.Right)
	//if !isValidBST(root.Left) || !isValidBST(root.Right) {
	//	return false
	//}
	//return true
}
//leetcode submit region end(Prohibit modification and deletion)

/**

func isValidBST(root *TreeNode) bool {
    return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
    if root == nil {
        return true
    }
    if root.Val <= lower || root.Val >= upper {
        return false
    }
    return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/validate-binary-search-tree/solution/yan-zheng-er-cha-sou-suo-shu-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

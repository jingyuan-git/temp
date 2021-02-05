//Given the root of a binary tree and an integer targetSum, return all root-to-l
//eaf paths where each path's sum equals targetSum. 
//
// A leaf is a node with no children. 
//
// 
// Example 1: 
//
// 
//Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//Output: [[5,4,11,2],[5,8,4,5]]
// 
//
// Example 2: 
//
// 
//Input: root = [1,2,3], targetSum = 5
//Output: []
// 
//
// Example 3: 
//
// 
//Input: root = [1,2], targetSum = 0
//Output: []
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the tree is in the range [0, 5000]. 
// -1000 <= Node.val <= 1000 
// -1000 <= targetSum <= 1000 
// 
// Related Topics 树 深度优先搜索 
// 👍 424 👎 0
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


func pathSum(root *TreeNode, targetSum int) (res [][]int) {
	var dfs func(path []int, root *TreeNode, targetSum int)
	dfs = func(path []int, root *TreeNode, targetSum int) {
		if root == nil {
			return
		}

		defer func() { path = path[:len(path)-1] }()
		if root.Left == nil && root.Right == nil && targetSum == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		dfs(append(path, root.Val), root.Left, targetSum - root.Val)
		dfs(append(path, root.Val), root.Right, targetSum - root.Val)
	}
	dfs([]int{}, root, targetSum)

	return
}
//leetcode submit region end(Prohibit modification and deletion)

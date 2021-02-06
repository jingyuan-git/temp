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
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}

		// TODO: 注意defer函数的顺序
		path = append(path, root.Val)
		defer func() { path = path[:len(path)-1] }()

		if root.Left == nil && root.Right == nil && targetSum - root.Val == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}

		dfs(root.Left, targetSum - root.Val)
		dfs(root.Right, targetSum - root.Val)
	}
	dfs(root, targetSum)

	return
}
//leetcode submit region end(Prohibit modification and deletion)

/**
func pathSum(root *TreeNode, sum int) (ans [][]int) {
    path := []int{}
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, left int) {
        if node == nil {
            return
        }
        left -= node.Val
        path = append(path, node.Val)
        defer func() { path = path[:len(path)-1] }()
        if node.Left == nil && node.Right == nil && left == 0 {
            ans = append(ans, append([]int(nil), path...))
            return
        }
        dfs(node.Left, left)
        dfs(node.Right, left)
    }
    dfs(root, sum)
    return
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/path-sum-ii/solution/lu-jing-zong-he-ii-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
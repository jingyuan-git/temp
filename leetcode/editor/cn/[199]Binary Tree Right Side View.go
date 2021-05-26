//Given the root of a binary tree, imagine yourself standing on the right side o
//f it,o return the values of the nodes you can see ordered from top to bottm.
//
// 
// Example 1: 
//
// 
//Input: root = [1,2,3,null,5,null,4]
//Output: [1,3,4]
// 
//
// Example 2: 
//
// 
//Input: root = [1,null,3]
//Output: [1,3]
// 
//
// Example 3: 
//
// 
//Input: root = []
//Output: []
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the tree is in the range [0, 100]. 
// -100 <= Node.val <= 100 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 递归 队列 
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	temp := []*TreeNode{root}
	res := []int{root.Val}
	for len(temp) != 0 {
		var child []*TreeNode
		for _, v := range temp {
			if v.Left != nil {
				child = append(child, v.Left)
			}
			if v.Right != nil {
				child = append(child, v.Right)
			}
		}
		if len(child) != 0 {
			res = append(res, child[len(child)-1].Val)
			//fmt.Println(res, "child = " ,child)
		}
		temp = child
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)

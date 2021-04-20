//Implement the BSTIterator class that represents an iterator over the in-order 
//traversal of a binary search tree (BST): 
//
// 
// BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. Th
//e root of the BST is given as part of the constructor. The pointer should be ini
//tialized to a non-existent number smaller than any element in the BST. 
// boolean hasNext() Returns true if there exists a number in the traversal to t
//he right of the pointer, otherwise returns false. 
// int next() Moves the pointer to the right, then returns the number at the poi
//nter. 
// 
//
// Notice that by initializing the pointer to a non-existent smallest number, th
//e first call to next() will return the smallest element in the BST. 
//
// You may assume that next() calls will always be valid. That is, there will be
// at least a next number in the in-order traversal when next() is called. 
//
// 
// Example 1: 
//
// 
//Input
//["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext
//", "next", "hasNext"]
//[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
//Output
//[null, 3, 7, true, 9, true, 15, true, 20, false]
//
//Explanation
//BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
//bSTIterator.next();    // return 3
//bSTIterator.next();    // return 7
//bSTIterator.hasNext(); // return True
//bSTIterator.next();    // return 9
//bSTIterator.hasNext(); // return True
//bSTIterator.next();    // return 15
//bSTIterator.hasNext(); // return True
//bSTIterator.next();    // return 20
//bSTIterator.hasNext(); // return False
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the tree is in the range [1, 105]. 
// 0 <= Node.val <= 106 
// At most 105 calls will be made to hasNext, and next. 
// 
//
// 
// Follow up: 
//
// 
// Could you implement next() and hasNext() to run in average O(1) time and use 
//O(h) memory, where h is the height of the tree? 
// 
// Related Topics 栈 树 设计 
// 👍 438 👎 0

package main

// next的次序是前序遍历
// Method 1: 递归遍历，结果存在数组中
// Method 2：栈
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
	cur *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		cur: root,
	}
}


func (this *BSTIterator) Next() int {
	for node := this.cur ; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}

	this.cur, this.stack = this.stack[len(this.stack) - 1], this.stack[:len(this.stack) - 1]
	res := this.cur.Val
	// 试探右侧的子树是否存在
	this.cur = this.cur.Right

	return res
}


func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.stack) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
//leetcode submit region end(Prohibit modification and deletion)

/**

type BSTIterator struct {
    arr []int
}

func Constructor(root *TreeNode) (it BSTIterator) {
    it.inorder(root)
    return
}

func (it *BSTIterator) inorder(node *TreeNode) {
    if node == nil {
        return
    }
    it.inorder(node.Left)
    it.arr = append(it.arr, node.Val)
    it.inorder(node.Right)
}

func (it *BSTIterator) Next() int {
    val := it.arr[0]
    it.arr = it.arr[1:]
    return val
}

func (it *BSTIterator) HasNext() bool {
    return len(it.arr) > 0
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/binary-search-tree-iterator/solution/er-cha-sou-suo-shu-die-dai-qi-by-leetcod-4y0y/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */

/**

type BSTIterator struct {
    stack []*TreeNode
    cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
    return BSTIterator{cur: root}
}

func (it *BSTIterator) Next() int {
    for node := it.cur; node != nil; node = node.Left {
        it.stack = append(it.stack, node)
    }
    it.cur, it.stack = it.stack[len(it.stack)-1], it.stack[:len(it.stack)-1]
    val := it.cur.Val
    it.cur = it.cur.Right
    return val
}

func (it *BSTIterator) HasNext() bool {
    return it.cur != nil || len(it.stack) > 0
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/binary-search-tree-iterator/solution/er-cha-sou-suo-shu-die-dai-qi-by-leetcod-4y0y/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */
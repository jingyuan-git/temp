//A linked list is given such that each node contains an additional random point
//er which could point to any node in the list or null. 
//
// Return a deep copy of the list. 
//
// The Linked List is represented in the input/output as a list of n nodes. Each
// node is represented as a pair of [val, random_index] where: 
//
// 
// val: an integer representing Node.val 
// random_index: the index of the node (range from 0 to n-1) where random pointe
//r points to, or null if it does not point to any node. 
// 
//
// 
// Example 1: 
//
// 
//Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
//Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 
//
// Example 2: 
//
// 
//Input: head = [[1,1],[2,1]]
//Output: [[1,1],[2,1]]
// 
//
// Example 3: 
//
// 
//
// 
//Input: head = [[3,null],[3,0],[3,null]]
//Output: [[3,null],[3,0],[3,null]]
// 
//
// Example 4: 
//
// 
//Input: head = []
//Output: []
//Explanation: Given linked list is empty (null pointer), so return null.
// 
//
// 
// Constraints: 
//
// 
// -10000 <= Node.val <= 10000 
// Node.random is null or pointing to a node in the linked list. 
// The number of nodes will not exceed 1000. 
// 
// Related Topics 哈希表 链表 
// 👍 471 👎 0
package main

type Node struct {
	Val int
	Next *Node
	Random *Node
}
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// TODO: 复制链表中的指针都不应指向原链表中的节点
func copyRandomList(head *Node) *Node {
	newHead := head
	tempHead := newHead
	for head != nil {
		newHead = head
		newHead.Random = head.Random
		head = head.Next
		newHead.Next = head.Next
	}
	return tempHead
}
//leetcode submit region end(Prohibit modification and deletion)

/*
class Solution(object):
    """
    :type head: Node
    :rtype: Node
    """
    def __init__(self):
        # Dictionary which holds old nodes as keys and new nodes as its values.
        self.visitedHash = {}

    def copyRandomList(self, head):

        if head == None:
            return None

        # If we have already processed the current node, then we simply return the cloned version of it.
        if head in self.visitedHash:
            return self.visitedHash[head]

        # create a new node
        # with the value same as old node.
        node = Node(head.val, None, None)

        # Save this value in the hash map. This is needed since there might be
        # loops during traversal due to randomness of random pointers and this would help us avoid them.
        self.visitedHash[head] = node

        # Recursively copy the remaining linked list starting once from the next pointer and then from the random pointer.
        # Thus we have two independent recursive calls.
        # Finally we update the next and random pointers for the new node created.
        node.next = self.copyRandomList(head.next)
        node.random = self.copyRandomList(head.random)

        return node

作者：LeetCode
链接：https://leetcode-cn.com/problems/copy-list-with-random-pointer/solution/fu-zhi-dai-sui-ji-zhi-zhen-de-lian-biao-by-leetcod/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func copyRandomList(head *Node) *Node {
    res:=new(Node)

    p:=res
    cur:=head
    m:=map[*Node]*Node{}
    for cur!=nil{
        p.Next=&Node{Val:cur.Val}
		// TODO: m存的key是当前所遍历节点的指针，value是复制的节点的指针
        m[cur]=p.Next
        p=p.Next
        cur=cur.Next
    }

    cur =head
    p = res.Next

    for cur!=nil{
        p.Random=m[cur.Random]
        p=p.Next
        cur=cur.Next
    }
    return res.Next
}

作者：ba-fei-niu
链接：https://leetcode-cn.com/problems/copy-list-with-random-pointer/solution/go-shuang-bai-bian-li-liang-ci-by-ba-fei-niu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
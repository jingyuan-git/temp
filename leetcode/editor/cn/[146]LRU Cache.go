//Design a data structure that follows the constraints of a Least Recently Used 
//(LRU) cache. 
//
// Implement the LRUCache class: 
//
// 
// LRUCache(int capacity) Initialize the LRU cache with positive size capacity. 
//
// int get(int key) Return the value of the key if the key exists, otherwise ret
//urn -1. 
// void put(int key, int value) Update the value of the key if the key exists. O
//therwise, add the key-value pair to the cache. If the number of keys exceeds the
// capacity from this operation, evict the least recently used key. 
// 
//
// Follow up: 
//Could you do get and put in O(1) time complexity? 
//
// 
// Example 1: 
//
// 
//Input
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//Output
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//Explanation
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // cache is {1=1}
//lRUCache.put(2, 2); // cache is {1=1, 2=2}
//lRUCache.get(1);    // return 1
//lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
//lRUCache.get(2);    // returns -1 (not found)
//lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
//lRUCache.get(1);    // return -1 (not found)
//lRUCache.get(3);    // return 3
//lRUCache.get(4);    // return 4
// 
//
// 
// Constraints: 
//
// 
// 1 <= capacity <= 3000 
// 0 <= key <= 3000 
// 0 <= value <= 104 
// At most 3 * 104 calls will be made to get and put. 
// 
// Related Topics 设计 
// 👍 1267 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	size int
	capacity int
	cache map[int]*DlinkedNode
	head, tail *DlinkedNode
}

type DlinkedNode struct {
	key, value int
	prev, next *DlinkedNode
}

func initDlinkedNode(key, value int) *DlinkedNode {
	//d := DlinkedNode{
	//	key: key,
	//	value: value,
	//}
	//return &d
	return &DlinkedNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache: make(map[int]*DlinkedNode),
		head: initDlinkedNode(0, 0),
		tail: initDlinkedNode(0, 0),
	}
	// TODO: 收尾连接
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}

	//node := this.cache[key]
	//this.moveToTail(node)
	//return node.value

	this.moveToTail(this.cache[key])
	return this.cache[key].value


	//node := this.head.next
	//for node != nil {
	//	if node.key == key {
	//		this.moveToTail(node)
	//		return node.value
	//	} else {
	//		node = node.next
	//	}
	//}
	//return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok {
		node := initDlinkedNode(key, value)
		this.addToTail(node)
		this.cache[key] = node
		this.size++
		if this.size > this.capacity {
			deleted := this.deleteHead()
			delete(this.cache, deleted.key)
			this.size--
		}
	} else {
		// 重新赋值value，转入method
		node := this.cache[key]
		node.value = value
		this.moveToTail(node)

	}


	// 这边逻辑和想的有些不一样：
	//1. 若判断node的数据已经有了，则不必再存入。直接制新状态。
	//
	//if this.size >= this.capacity {
	//	this.deleteHead()
	//	this.size--
	//}
	//
	//this.moveToTail(initDlinkedNode(key, value))
	//this.size++
}

func (this *LRUCache) deleteHead () *DlinkedNode {
	node := this.head.next
	node.prev.next = node.next
	node.next.prev = node.prev
	return node
}

func (this *LRUCache) addToTail(moveNode *DlinkedNode) {
	// 移到最末端
	moveNode.next = this.tail
	moveNode.prev = this.tail.prev
	this.tail.prev.next = moveNode
	this.tail.prev = moveNode
}

func (this *LRUCache) moveToTail(moveNode *DlinkedNode)  {
	moveNode.prev.next = moveNode.next
	moveNode.next.prev = moveNode.prev

	this.addToTail(moveNode)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

/**

type LRUCache struct {
    size int
    capacity int
    cache map[int]*DLinkedNode
    head, tail *DLinkedNode
}

type DLinkedNode struct {
    key, value int
    prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
    return &DLinkedNode{
        key: key,
        value: value,
    }
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        cache: map[int]*DLinkedNode{},
        head: initDLinkedNode(0, 0),
        tail: initDLinkedNode(0, 0),
        capacity: capacity,
    }
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.cache[key]; !ok {
        return -1
    }
    node := this.cache[key]
    this.moveToHead(node)
    return node.value
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.cache[key]; !ok {
        node := initDLinkedNode(key, value)
        this.cache[key] = node
        this.addToHead(node)
        this.size++
        if this.size > this.capacity {
            removed := this.removeTail()
            delete(this.cache, removed.key)
            this.size--
        }
    } else {
        node := this.cache[key]
        node.value = value
        this.moveToHead(node)
    }
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
    node := this.tail.prev
    this.removeNode(node)
    return node
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/lru-cache/solution/lruhuan-cun-ji-zhi-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */
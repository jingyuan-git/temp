//Implement a first in first out (FIFO) queue using only two stacks. The impleme
//nted queue should support all the functions of a normal queue (push, peek, pop, 
//and empty). 
//
// Implement the MyQueue class: 
//
// 
// void push(int x) Pushes element x to the back of the queue. 
// int pop() Removes the element from the front of the queue and returns it. 
// int peek() Returns the element at the front of the queue. 
// boolean empty() Returns true if the queue is empty, false otherwise. 
// 
//
// Notes: 
//
// 
// You must use only standard operations of a stack, which means only push to to
//p, peek/pop from top, size, and is empty operations are valid. 
// Depending on your language, the stack may not be supported natively. You may 
//simulate a stack using a list or deque (double-ended queue) as long as you use o
//nly a stack's standard operations. 
// 
//
// Follow-up: Can you implement the queue such that each operation is amortized 
//O(1) time complexity? In other words, performing n operations will take overall 
//O(n) time even if one of those operations may take longer. 
//
// 
// Example 1: 
//
// 
//Input
//["MyQueue", "push", "push", "peek", "pop", "empty"]
//[[], [1], [2], [], [], []]
//Output
//[null, null, null, 1, 1, false]
//
//Explanation
//MyQueue myQueue = new MyQueue();
//myQueue.push(1); // queue is: [1]
//myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
//myQueue.peek(); // return 1
//myQueue.pop(); // return 1, queue is [2]
//myQueue.empty(); // return false
// 
//
// 
// Constraints: 
//
// 
// 1 <= x <= 9 
// At most 100 calls will be made to push, pop, peek, and empty. 
// All the calls to pop and peek are valid. 
// 
// Related Topics 栈 设计 
// 👍 353 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
type MyQueue struct {
	inStack, outStack	[]int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.inStack = append(this.inStack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.inStack) == 0 && len(this.outStack) == 0 {
		return true
	}
	return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)

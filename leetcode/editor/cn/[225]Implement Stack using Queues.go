//Implement a last-in-first-out (LIFO) stack using only two queues. The implemen
//ted stack should support all the functions of a normal stack (push, top, pop, an
//d empty). 
//
// Implement the MyStack class: 
//
// 
// void push(int x) Pushes element x to the top of the stack. 
// int pop() Removes the element on the top of the stack and returns it. 
// int top() Returns the element on the top of the stack. 
// boolean empty() Returns true if the stack is empty, false otherwise. 
// 
//
// Notes: 
//
// 
// You must use only standard operations of a queue, which means that only push 
//to back, peek/pop from front, size and is empty operations are valid. 
// Depending on your language, the queue may not be supported natively. You may 
//simulate a queue using a list or deque (double-ended queue) as long as you use o
//nly a queue's standard operations. 
// 
//
// 
// Example 1: 
//
// 
//Input
//["MyStack", "push", "push", "top", "pop", "empty"]
//[[], [1], [2], [], [], []]
//Output
//[null, null, null, 2, 2, false]
//
//Explanation
//MyStack myStack = new MyStack();
//myStack.push(1);
//myStack.push(2);
//myStack.top(); // return 2
//myStack.pop(); // return 2
//myStack.empty(); // return False
// 
//
// 
// Constraints: 
//
// 
// 1 <= x <= 9 
// At most 100 calls will be made to push, pop, top, and empty. 
// All the calls to pop and top are valid. 
// 
//
// 
// Follow-up: Can you implement the stack using only one queue? 
// Related Topics 栈 设计 队列 
// 👍 349 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	stack []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{stack: []int{}}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.stack = append(this.stack, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := this.stack[len(this.stack) - 1]
	this.stack = this.stack[:len(this.stack) - 1]
	return res
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.stack) == 0 {
		return true
	} else {
		return false
	}
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)

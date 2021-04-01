//Design a stack that supports push, pop, top, and retrieving the minimum elemen
//t in constant time. 
//
// Implement the MinStack class: 
//
// 
// MinStack() initializes the stack object. 
// void push(val) pushes the element val onto the stack. 
// void pop() removes the element on the top of the stack. 
// int top() gets the top element of the stack. 
// int getMin() retrieves the minimum element in the stack. 
// 
//
// 
// Example 1: 
//
// 
//Input
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//Output
//[null,null,null,null,-3,null,0,-2]
//
//Explanation
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin(); // return -3
//minStack.pop();
//minStack.top();    // return 0
//minStack.getMin(); // return -2
// 
//
// 
// Constraints: 
//
// 
// -231 <= val <= 231 - 1 
// Methods pop, top and getMin operations will always be called on non-empty sta
//cks. 
// At most 3 * 104 calls will be made to push, pop, top, and getMin. 
// 
// Related Topics 栈 设计 
// 👍 857 👎 0
package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: []int{}}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	return
}


func (this *MinStack) Top() int {
	val := this.stack[len(this.stack)-1]
	//this.stack = this.stack[:len(this.stack)-1]
	return val
}


func (this *MinStack) GetMin() int {
	minVal := math.MaxInt32
	for _, v := range this.stack {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)

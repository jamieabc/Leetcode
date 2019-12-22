package main

//Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
//    push(x) -- Push element x onto stack.
//    pop() -- Removes the element on top of the stack.
//    top() -- Get the top element.
//    getMin() -- Retrieve the minimum element in the stack.
//
//
//
//Example:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> Returns -3.
//minStack.pop();
//minStack.top();      --> Returns 0.
//minStack.getMin();   --> Returns -2.

type MinStack struct {
	data   []int
	min    int
	length int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:   make([]int, 0),
		min:    0,
		length: 0,
	}
}

func (this *MinStack) Push(x int) {
	if this.length == 0 {
		this.min = x
	} else if x < this.min {
		this.min = x
	}

	this.data = append(this.data, x)
	this.length++
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}

	// find next min number
	if this.data[this.length-1] == this.min {
		this.min = this.data[0]
		for i := 1; i < this.length-1; i++ {
			if this.data[i] < this.min {
				this.min = this.data[i]
			}
		}
	}

	this.data = this.data[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.data[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

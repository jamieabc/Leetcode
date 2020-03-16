package main

import "math"

//Design a max stack that supports push, pop, top, peekMax and popMax.
//
//push(x) -- Push element x onto stack.
//pop() -- Remove the element on top of the stack and return it.
//top() -- Get the element on the top.
//peekMax() -- Retrieve the maximum element in the stack.
//popMax() -- Retrieve the maximum element in the stack, and remove it. If you find more than one maximum elements, only remove the top-most one.
//Example 1:
//
//MaxStack stack = new MaxStack();
//stack.push(5);
//stack.push(1);
//stack.push(5);
//stack.top(); -> 5
//stack.popMax(); -> 5
//stack.top(); -> 1
//stack.peekMax(); -> 5
//stack.pop(); -> 1
//stack.top(); -> 5
//Note:
//
//-1e7 <= x <= 1e7
//Number of operations won't exceed 10000.
//The last four operations won't be called when stack is empty.

type MaxStack struct {
	data []int
	max  int
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	return MaxStack{
		data: make([]int, 0),
		max:  math.MinInt32,
	}
}

// O(1)
func (this *MaxStack) Push(x int) {
	this.data = append(this.data, x)

	if this.max < x {
		this.max = x
	}
}

// O(n)
func (this *MaxStack) Pop() int {
	length := len(this.data)
	popped := this.data[length-1]
	this.data = this.data[:length-1]
	this.updateMax()

	return popped
}

// O(1)
func (this *MaxStack) Top() int {
	return this.data[len(this.data)-1]
}

// O(1)
func (this *MaxStack) PeekMax() int {
	return this.max
}

// O(n)
func (this *MaxStack) PopMax() int {
	length := len(this.data)
	popped := this.max
	for i := length - 1; i >= 0; i-- {
		if this.data[i] == this.max {
			// remove
			if i == length-1 {
				this.data = this.data[:length-1]
			} else {
				this.data = append(this.data[:i], this.data[i+1:]...)
			}
			break
		}
	}
	this.updateMax()

	return popped
}

// O(n)
func (this *MaxStack) updateMax() {
	this.max = math.MinInt32

	for _, num := range this.data {
		if num > this.max {
			this.max = num
		}
	}
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */

// problems
//	1.	when popping max, forget to update maxIndex
//	2.	when maxIndex is updated, criteria to update also changes
//	3.	fix logic of updating maxIndex
//	4.	use math.MinInt32
//	5.	optimize, no need to store index, just search when popping max
//	6.	remove too much...forget to break loop when value is found

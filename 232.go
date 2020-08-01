package main

//Implement the following operations of a queue using stacks.
//
//    push(x) -- Push element x to the back of queue.
//    pop() -- Removes the element from in front of queue.
//    peek() -- Get the front element.
//    empty() -- Return whether the queue is empty.
//
//Example:
//
//MyQueue queue = new MyQueue();
//
//queue.push(1);
//queue.push(2);
//queue.peek();  // returns 1
//queue.pop();   // returns 1
//queue.empty(); // returns false
//
//Notes:
//
//    You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
//    Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
//    You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).

// lazy operation
type MyQueue struct {
	data, reversed []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		data:     make([]int, 0),
		reversed: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	this.reverse()
	popped := this.reversed[len(this.reversed)-1]
	this.reversed = this.reversed[:len(this.reversed)-1]

	return popped
}

func (this *MyQueue) reverse() {
	if len(this.reversed) == 0 {
		// change stack order (data) to queue order (reversed)
		for len(this.data) > 0 {
			this.reversed = append(this.reversed, this.data[len(this.data)-1])
			this.data = this.data[:len(this.data)-1]
		}
	}
}

func (this *MyQueue) Peek() int {
	this.reverse()
	return this.reversed[len(this.reversed)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.data)+len(this.reversed) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

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

type MyStack struct {
	data []int
}

func (s *MyStack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *MyStack) Pop() int {
	length := len(s.data)
	popped := s.data[length-1]
	s.data = s.data[:length-1]
	return popped
}

func (s *MyStack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.data) == 0
}

type MyQueue struct {
	data     *MyStack
	reversed *MyStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		data: &MyStack{
			data: make([]int, 0),
		},
		reversed: &MyStack{
			data: make([]int, 0),
		},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.data.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var popped int
	for {
		popped = this.data.Pop()
		if this.data.Empty() {
			break
		}
		this.reversed.Push(popped)
	}

	for !this.reversed.Empty() {
		this.data.Push(this.reversed.Pop())
	}

	return popped
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	var popped int
	for !this.data.Empty() {
		popped = this.data.Pop()
		this.reversed.Push(popped)
	}

	for !this.reversed.Empty() {
		this.data.Push(this.reversed.Pop())
	}

	return popped
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.data.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

package main

//Implement the following operations of a stack using queues.
//
//    push(x) -- Push element x onto stack.
//    pop() -- Removes the element on top of the stack.
//    top() -- Get the top element.
//    empty() -- Return whether the stack is empty.
//
//Example:
//
//MyStack stack = new MyStack();
//
//stack.push(1);
//stack.push(2);
//stack.top();   // returns 2
//stack.pop();   // returns 2
//stack.empty(); // returns false
//
//Notes:
//
//    You must use only standard operations of a queue -- which means only push to back, peek/pop from front, size, and is empty operations are valid.
//    Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.
//    You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).

type Queue interface {
	Push(int)
	Pop() int
	Peek() int
	Size() int
	Empty() bool
}

type queue struct {
	data []int
	size int
}

func (q *queue) Push(i int) {
	q.data = append(q.data, i)
	q.size++
}

func (q *queue) Pop() int {
	popped := q.data[0]
	q.data = q.data[1:]
	q.size--
	return popped
}

func (q *queue) Peek() int {
	return q.data[0]
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) Empty() bool {
	return q.size == 0
}

func newQueue() Queue {
	return &queue{
		data: make([]int, 0),
	}
}

type MyStack struct {
	data   Queue
	backup Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		data:   newQueue(),
		backup: newQueue(),
	}
}

//5 4 3 2 1 0
//
//only difference is pop from front

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.data.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.data.Size() == 1 {
		return this.data.Pop()
	}

	for this.data.Size() != 1 {
		this.backup.Push(this.data.Pop())
	}
	popped := this.data.Pop()
	this.data = this.backup
	this.backup = newQueue()
	return popped
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.data.Size() == 1 {
		return this.data.Peek()
	}

	for this.data.Size() != 1 {
		this.backup.Push(this.data.Pop())
	}
	top := this.data.Pop()
	this.backup.Push(top)
	this.data = this.backup
	this.backup = newQueue()
	return top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.data.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

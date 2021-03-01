package main

import "fmt"

// Implement FreqStack, a class which simulates the operation of a stack-like data structure.
//
// FreqStack has two functions:
//
// push(int x), which pushes an integer x onto the stack.
// pop(), which removes and returns the most frequent element in the stack.
// If there is a tie for most frequent element, the element closest to the top of the stack is removed and returned.
//
//
//
// Example 1:
//
// Input:
// ["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
// [[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
// Output: [null,null,null,null,null,null,null,5,7,5,4]
// Explanation:
// After making six .push operations, the stack is [5,7,5,7,4,5] from bottom to top.  Then:
//
// pop() -> returns 5, as 5 is the most frequent.
// The stack becomes [5,7,5,7,4].
//
// pop() -> returns 7, as 5 and 7 is the most frequent, but 7 is closest to the top.
// The stack becomes [5,7,5,4].
//
// pop() -> returns 5.
// The stack becomes [5,7,4].
//
// pop() -> returns 4.
// The stack becomes [5,7].
//
//
//
// Note:
//
// Calls to FreqStack.push(int x) will be such that 0 <= x <= 10^9.
// It is guaranteed that FreqStack.pop() won't be called if the stack has zero elements.
// The total number of FreqStack.push calls will not exceed 10000 in a single test case.
// The total number of FreqStack.pop calls will not exceed 10000 in a single test case.
// The total number of FreqStack.push and FreqStack.pop calls will not exceed 150000 across all test cases.

type FreqStack struct {
	MaxFreq    int
	FreqGroup  [][]int
	NumberFreq map[int]int
}

func Constructor() FreqStack {
	return FreqStack{
		FreqGroup:  make([][]int, 0),
		NumberFreq: make(map[int]int),
	}
}

func (this *FreqStack) Push(x int) {
	this.NumberFreq[x]++
	freq := this.NumberFreq[x]

	this.MaxFreq = max(this.MaxFreq, freq)

	// since its array, need to allocate memory first
	if len(this.FreqGroup) < freq {
		this.FreqGroup = append(this.FreqGroup, []int{})
	}
	this.FreqGroup[freq-1] = append(this.FreqGroup[freq-1], x)
}

func (this *FreqStack) Pop() int {
	arr := this.FreqGroup[this.MaxFreq-1]
	n := arr[len(arr)-1]
	this.FreqGroup[this.MaxFreq-1] = arr[:len(arr)-1]
	this.NumberFreq[n]--

	for ; this.MaxFreq > 0 && len(this.FreqGroup[this.MaxFreq-1]) == 0; this.MaxFreq-- {
	}

	return n
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

type LinkedList1 struct {
	Prev, Next *LinkedList1
	Freq, Val  int
	Idx        []int
}

type FreqStack1 struct {
	Freq         map[int]*LinkedList1
	Lookup       map[int]*LinkedList1
	MaxFreq, Idx int
}

func Constructor1() FreqStack1 {
	return FreqStack1{
		Freq:   make(map[int]*LinkedList1),
		Lookup: make(map[int]*LinkedList1),
	}
}

func (f *FreqStack1) Print() {
	for i := 1; i <= f.MaxFreq; i++ {
		fmt.Println("freq", i)
		for n := f.Freq[i]; n != nil; n = n.Next {
			fmt.Println(n)
		}
	}
}

func (f *FreqStack1) Check() {
	for i := 1; i <= f.MaxFreq; i++ {
		fmt.Println("freq", i)
		for n := f.Freq[i]; n != nil; n = n.Next {
			if i != n.Freq {
				fmt.Println("wrong", n)
			}
		}
	}
}

func (this *FreqStack1) Push(x int) {
	var node *LinkedList1

	if n, ok := this.Lookup[x]; ok {
		// process current frequency list
		if this.Freq[n.Freq] == n {
			this.Freq[n.Freq] = n.Next
			if n.Next == nil {
				delete(this.Freq, n.Freq)
			}
		}

		if n.Next != nil {
			n.Next.Prev = n.Prev
		}

		if n.Prev != nil {
			n.Prev.Next = n.Next
		}

		n.Next, n.Prev = nil, nil
		n.Idx = append(n.Idx, this.Idx)
		n.Freq++
		node = n
	} else {
		node = &LinkedList1{
			Freq: 1,
			Idx:  []int{this.Idx},
			Val:  x,
		}
		this.Lookup[x] = node
	}

	if orig, ok := this.Freq[node.Freq]; ok {
		orig.Prev = node
		node.Next = orig
	}
	this.Freq[node.Freq] = node
	this.MaxFreq = max(this.MaxFreq, node.Freq)
	this.Idx++
}

// O(n)
func (this *FreqStack1) Pop() int {
	node := this.Freq[this.MaxFreq]
	node.Freq--
	node.Idx = node.Idx[:len(node.Idx)-1]

	// no more nodes with same frequency
	this.Freq[this.MaxFreq] = node.Next
	if node.Next == nil {
		delete(this.Freq, this.MaxFreq)
		this.MaxFreq--
	} else {
		node.Next.Prev = node.Prev

		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
	}

	if node.Freq > 0 {
		// put into right place
		if n := this.Freq[node.Freq]; n == nil {
			this.Freq[node.Freq] = node
			node.Next, node.Prev = nil, nil
		} else {
			var prev *LinkedList1
			for ; n != nil && n.Idx[len(n.Idx)-1] > node.Idx[len(node.Idx)-1]; n = n.Next {
				prev = n
			}

			if prev != nil {
				node.Next = prev.Next

				if prev.Next != nil {
					prev.Next.Prev = node
				}
				prev.Next = node
				node.Prev = prev
			} else {
				n := this.Freq[node.Freq]
				node.Next = n
				n.Prev = node
				node.Prev = nil
				this.Freq[node.Freq] = node
			}
		}
	} else {
		delete(this.Lookup, node.Val)
	}

	return node.Val
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */

//	Notes
//	1.	at first I thought use heap to solve this problem, when pushing, need to
//		scan all data in heap which violates its original design, so I change to
//		use hashmap to implement, push with O(1), pop with O(k) k: # number nodes
//		for same frequency group

//	2.	takes a lot of time debugging infinite loop problem, linked-list easily
//		causes this because of prev/next not properly handled

//	3.	inspired from solution & https://leetcode.com/problems/maximum-frequency-stack/discuss/163410/C%2B%2BJavaPython-O(1)

//		stack already contains information of latest, the point here is to store
//		duplicate numbers into different frequency group, and always pop from
//		maximum frequency, which result in problem want...

//		when pushing, need a faster way to find its frequency
//		when popping, need a faster way to know top frequent number

//		how brilliant...

//	4.	inspired from solution, can use [][]int for frequency group of numbers
//		it should be faster, but [][]int need to allocate memory first

//	5.	inspired from https://leetcode.com/problems/maximum-frequency-stack/discuss/163435/Python-Simple-PriorityQueue

//		priority queue also works, not implement

//		key point similar to stack based solution, put multiple numbers into
//		heap

package main

import (
	"container/heap"
	"sort"
)

// You have some number of sticks with positive integer lengths. These lengths are given as an array sticks, where sticks[i] is the length of the ith stick.
//
// You can connect any two sticks of lengths x and y into one stick by paying a cost of x + y. You must connect all the sticks until there is only one stick remaining.
//
// Return the minimum cost of connecting all the given sticks into one stick in this way.
//
//
//
// Example 1:
//
// Input: sticks = [2,4,3]
// Output: 14
// Explanation: You start with sticks = [2,4,3].
// 1. Combine sticks 2 and 3 for a cost of 2 + 3 = 5. Now you have sticks = [5,4].
// 2. Combine sticks 5 and 4 for a cost of 5 + 4 = 9. Now you have sticks = [9].
// There is only one stick left, so you are done. The total cost is 5 + 9 = 14.
//
// Example 2:
//
// Input: sticks = [1,8,3,5]
// Output: 30
// Explanation: You start with sticks = [1,8,3,5].
// 1. Combine sticks 1 and 3 for a cost of 1 + 3 = 4. Now you have sticks = [4,8,5].
// 2. Combine sticks 4 and 5 for a cost of 4 + 5 = 9. Now you have sticks = [9,8].
// 3. Combine sticks 9 and 8 for a cost of 9 + 8 = 17. Now you have sticks = [17].
// There is only one stick left, so you are done. The total cost is 4 + 9 + 17 = 30.
//
// Example 3:
//
// Input: sticks = [5]
// Output: 0
// Explanation: There is only one stick, so you don't need to do anything. The total cost is 0.
//
//
//
// Constraints:
//
// 1 <= sticks.length <= 104
// 1 <= sticks[i] <= 104

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() int          { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func connectSticks(sticks []int) int {
	minHeap := &MinHeap{}
	for _, i := range sticks {
		heap.Push(minHeap, i)
	}

	var cost int

	for minHeap.Len() > 1 {
		p1 := heap.Pop(minHeap).(int)
		p2 := heap.Pop(minHeap).(int)

		cost += p1 + p2
		heap.Push(minHeap, p1+p2)
	}

	return cost
}

//	Notes
//	1.	initially i thought about always merging smallest two numbers, but then
//		i found it's wrong because if array is long enough, always pick from smallest
//		cannot work

//		e.g.
//		[1, 1, 1, 1, 1, 1, ..., 1, 1, 500], length = 1000

//		start from left most 1, it will add for 999 times, because each time add left
//		means original 1 is added again and again, 1 add 999 times > 500, so start from
//		smallest most number and greedily merge right number is wrong

//	2.	then i think of, each time select smallest 2 numbers, add sum of theses 2 numbers
//		back, which is greedily select smallest 2

//		because if [a, b, c, d] sorted ascending
//		a+b < c+d
//		a+c > a+b
//		a+d > a+b

//		smallest 2 of a & b is the best choice

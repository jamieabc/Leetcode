package main

import (
	"container/heap"
)

// Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.
// For example,
//
// [2,3,4], the median is 3
//
// [2,3], the median is (2 + 3) / 2 = 2.5
//
// Design a data structure that supports the following two operations:
//
//     void addNum(int num) - Add a integer number from the data stream to the data structure.
//     double findMedian() - Return the median of all elements so far.
//
//
//
// Example:
//
// addNum(1)
// addNum(2)
// findMedian() -> 1.5
// addNum(3)
// findMedian() -> 2
//
//
//
// Follow up:
//
//     If all integer numbers from the stream are between 0 and 100, how would you optimize it?
//     If 99% of all integer numbers from the stream are between 0 and 100, how would you optimize it?

type InitHeap []int

func (h InitHeap) Len() int            { return len(h) }
func (h InitHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h InitHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h InitHeap) Peak() interface{}   { return h[0] }
func (h *InitHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *InitHeap) Pop() interface{} {
	length := len(*h)
	popped := (*h)[length-1]
	*h = (*h)[:length-1]
	return popped
}

type MedianFinder struct {
	minHeap *InitHeap
	maxHeap *InitHeap
	size    int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	minHeap, maxHeap := &InitHeap{}, &InitHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)

	return MedianFinder{
		minHeap: &InitHeap{},
		maxHeap: &InitHeap{},
		size:    0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxHeap, -num)
	heap.Push(this.minHeap, -heap.Pop(this.maxHeap).(int))
	if this.minHeap.Len() > this.maxHeap.Len() {
		heap.Push(this.maxHeap, -heap.Pop(this.minHeap).(int))
	}
	this.size++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.size&1 == 1 {
		return float64(-this.maxHeap.Peak().(int))
	}

	return float64(this.minHeap.Peak().(int)-this.maxHeap.Peak().(int)) / float64(2)
}

// type MedianFinder struct {
// 	data []int
// }
//
// /** initialize your data structure here. */
// func Constructor() MedianFinder {
// 	return MedianFinder{
// 		data: make([]int, 0),
// 	}
// }
//
// func (this *MedianFinder) AddNum(num int) {
// 	if len(this.data) == 0 {
// 		this.data = append(this.data, num)
// 	} else {
// 		var i, j int
// 		for i, j = 0, len(this.data)-1; i < j; {
// 			mid := i + (j-i)/2
//
// 			if this.data[mid] > num {
// 				j = mid
// 			} else {
// 				i = mid + 1
// 			}
// 		}
//
// 		if this.data[len(this.data)-1] >= num {
// 			tmp := append([]int{}, this.data[:i]...)
// 			tmp = append(tmp, num)
// 			tmp = append(tmp, this.data[i:]...)
// 			this.data = tmp
// 		} else {
// 			this.data = append(this.data, num)
// 		}
// 	}
// }
//
// func (this *MedianFinder) FindMedian() float64 {
// 	if len(this.data) == 0 {
// 		return float64(0)
// 	}
//
// 	if len(this.data)&1 == 1 {
// 		return float64(this.data[len(this.data)/2])
// 	}
//
// 	return float64(this.data[len(this.data)/2]+this.data[len(this.data)/2-1]) / float64(2)
// }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

//	problems
//	1.	too slow, use binary search, tc: O(n) + O(log n), it takes log n to
//		find right position, take another n to move data

//	2.	inspired from solution, uses 2 heaps to find median value. the point
//		is not to have a sorted array, target is to find median value.

//		w/ min & max heap to find median value. min heap stores larger
//		numbers, and max heap stores lower numbers. since median number has
//		2 conditions: data length even or odd

//		a way can be selected to always put number into lower half heap
//		(max heap), after sorted, poll max heap and put into higher half
//		heap (min heap). after min heap is sorted, if min heap with more
//		number, poll from min heap and put into max heap

//		data can be view as 2n or 2n+1:
//		- 2n: both max & min heap stores median value
//		- 2n+1: max heap stores 1 more value from min heap

//		tc: O(5 log n), at worst case, heap with 3 insertions and 2 deletions

//	3.	from sample code, author uses quite clever way to reuse Less: for
//		max heap, put number multiply -1, this reduces extra interface and
//		struct type

//		I don't even know go provides heap package...

//	4.	add reference https://leetcode.com/problems/find-median-from-data-stream/discuss/74047/JavaPython-two-heap-solution-O(log-n)-add-O(1)-find

//	5.	add reference https://leetcode.com/problems/find-median-from-data-stream/discuss/275207/Solutions-to-follow-ups

//		for follow-ups, if all number are within 0-100, use array to store
//		count of those numbers, tc: O(1)

//		for 99% integers in 0-100, use array to store numbers in 0-100,
//		and 2 arrays to store number < 0 and number > 100, because median
//		is not likely to fall in those out of bound numbers

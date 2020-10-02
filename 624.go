package main

import (
	"math"
)

// Given m arrays, and each array is sorted in ascending order. Now you can pick up two integers from two different arrays (each array picks one) and calculate the distance. We define the distance between two integers a and b to be their absolute difference |a-b|. Your task is to find the maximum distance.
//
//Example 1:
//
//Input:
//[[1,2,3],
// [4,5],
// [1,2,3]]
//Output: 4
//Explanation:
//One way to reach the maximum distance 4 is to pick 1 in the first or third array and pick 5 in the second array.
//
//Note:
//
//    Each given array will have at least 1 number. There will be at least two non-empty arrays.
//    The total number of the integers in all the m arrays will be in the range of [2, 10000].
//    The integers in the m arrays will be in the range of [-10000, 10000].

func maxDistance(arrays [][]int) int {
	smallest, largest := arrays[0][0], arrays[0][len(arrays[0])-1]
	var maxDist int

	// two numbers from different rows, so max distance need to be calculated,
	// and previous smallest/largest can update
	for i := 1; i < len(arrays); i++ {
		arr := arrays[i]
		maxDist = max(maxDist, max(abs(arr[0]-largest), abs(arr[len(arr)-1]-smallest)))
		largest = max(largest, arr[len(arr)-1])
		smallest = min(smallest, arr[0])
	}

	return maxDist
}

type priorityQueue struct {
	data        []interface{}
	size        int
	compareFunc func(interface{}, interface{}) bool
}

func (q *priorityQueue) Poll() interface{} {
	if len(q.data) == 0 {
		return nil
	}

	polled := q.data[0]
	q.swap(0, len(q.data)-1)
	q.data = q.data[:len(q.data)-1]

	q.bubbleDown(0)
	return polled
}

func (q *priorityQueue) Offer(i interface{}) {
	if len(q.data) < q.size {
		q.data = append(q.data, i)
		idx := q.bubbleUp(len(q.data) - 1)
		q.bubbleDown(idx)
	} else {
		if q.compareFunc(i, q.data[q.size-1]) {
			q.data[q.size-1] = i
			idx := q.bubbleUp(q.size - 1)
			q.bubbleDown(idx)
		}
	}
}

func (q *priorityQueue) Size() int {
	return len(q.data)
}

func (q *priorityQueue) Peek() interface{} {
	if len(q.data) == 0 {
		return nil
	}
	return q.data[0]
}

func (q *priorityQueue) Comparator(f func(interface{}, interface{}) bool) {
	q.compareFunc = f
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func (q *priorityQueue) swap(i, j int) {
	q.data[i], q.data[j] = q.data[j], q.data[i]
}

func (q *priorityQueue) bubbleUp(idx int) int {
	for current := idx; current != 0; {
		p := parent(current)
		if q.compareFunc(q.data[p], q.data[current]) {
			return current
		} else {
			q.swap(p, current)
			current = p
		}
	}
	return 0
}

func leftChild(idx int) int {
	return idx*2 + 1
}

func rightChild(idx int) int {
	return idx*2 + 2
}

func (q *priorityQueue) bubbleDown(idx int) {
	for current := idx; current < len(q.data)-1; {
		l := leftChild(idx)
		r := rightChild(idx)

		ls, rs := true, true
		if l < len(q.data) {
			ls = q.compareFunc(q.data[current], q.data[l])
		}

		if r < len(q.data) {
			rs = q.compareFunc(q.data[current], q.data[r])
		}

		if ls && rs {
			return
		}

		// in case right child not exist or l is end of data
		if l == len(q.data)-1 || r >= len(q.data) {
			q.swap(l, current)
			current = l
			continue
		}

		if q.compareFunc(q.data[l], q.data[r]) {
			q.swap(l, current)
			current = l
		} else {
			q.swap(r, current)
			current = r
		}
	}
}

func maxDistance1(arrays [][]int) int {
	maxPQ := priorityQueue{
		data: make([]interface{}, 0),
		size: 3,
		compareFunc: func(i, j interface{}) bool {
			return i.(int) >= j.(int)
		},
	}

	minPQ := priorityQueue{
		data: make([]interface{}, 0),
		size: 3,
		compareFunc: func(i, j interface{}) bool {
			return i.(int) <= j.(int)
		},
	}

	var maxDuplicate, minDuplicate, sameline bool

	for i, l := range arrays {
		length := len(l)

		// first item just pushing
		if i == 0 {
			maxPQ.Offer(l[length-1])
			minPQ.Offer(l[0])
		} else {
			prevMax := maxPQ.Peek().(int)
			maxPQ.Offer(l[length-1])

			maxDuplicate = prevMax == l[length-1]

			prevMin := minPQ.Peek().(int)
			minPQ.Offer(l[0])

			minDuplicate = prevMin == l[0]
		}

		if maxPQ.Peek().(int) == l[length-1] && minPQ.Peek().(int) == l[0] {
			sameline = true
		} else if maxPQ.Peek().(int) == l[length-1] || minPQ.Peek().(int) == l[0] {
			sameline = false
		}
	}

	max1 := maxPQ.Poll().(int)
	max2 := math.MinInt32
	if i := maxPQ.Poll(); i != nil {
		max2 = i.(int)
	}

	min1 := minPQ.Poll().(int)
	min2 := math.MaxInt32
	if i := minPQ.Poll(); i != nil {
		min2 = i.(int)
	}

	if !maxDuplicate && !minDuplicate && sameline {
		return max(dist(max1, min2), dist(max2, min1))
	}

	return max1 - min1
}

func dist(i, j int) int {
	if i <= j {
		return j - i
	}

	return i - j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

// 	Notes
//	1.	compare wrong, what I want is value of min index, not index value
//	2.	optimize, no need for second pass, comparison can be done in single
//		run
//	3.	optimize, no need to store indexes, just use value of 0th and last to
//		check if this line should be used or not
//	4.	optimize, max distance comes from max - min, but the problem is if
//		these values stay at the same line, additional checking are needed
//
//		special case comes from if both max & min are unique, and they stays
//		at same line, then result either comes from (max, min2) or (max2, min)
//	5.	fix problem, double poll cause number wrong
//	6.	wrong logic to check if sub-array contains both current max & min,
//		existing variable of sameMax & sameMin is to find duplicates of
//		max & min, but same line needs to check if current max == sub array
//		max, and current min == sub array min
//	7.	wrong logic to check if max/min occurs multiple time
//	8.	a bug when bubble down, cannot always assume 2 children both exist
//	9.	fix priority queue bug, size should not be changed
//	10.	wrong criteria when offering new number
//	11.	optimize, use one pass to calculate. As long as calculated result
//		doesn't come from same sub-array, then result can be calculated.
//	12.	initial result cannot be calculated directly from globalMax - globalMin
//		because those 2 values might come from same sub-array

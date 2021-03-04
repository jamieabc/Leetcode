package main

import "container/heap"

// We have a list of points on the plane.  Find the K closest points to the origin (0, 0).
//
// (Here, the distance between two points on a plane is the Euclidean distance.)
//
// You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.)
//
//
//
// Example 1:
//
// Input: points = [[1,3],[-2,2]], K = 1
// Output: [[-2,2]]
// Explanation:
// The distance between (1, 3) and the origin is sqrt(10).
// The distance between (-2, 2) and the origin is sqrt(8).
// Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
// We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].
// Example 2:
//
// Input: points = [[3,3],[5,-1],[-2,4]], K = 2
// Output: [[3,3],[-2,4]]
// (The answer [[-2,4],[3,3]] would also be accepted.)
//
//
// Note:
//
// 1 <= K <= points.length <= 10000
// -10000 < points[i][0] < 10000
// -10000 < points[i][1] < 10000

// tc: average O(n)
func kClosest(points [][]int, k int) [][]int {
	n := len(points)
	arr := make([][]int, n)
	for i, p := range points {
		arr[i] = []int{p[0]*p[0] + p[1]*p[1], p[0], p[1]}
	}

	quickSelect(arr, 0, n-1, k)

	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		ans[i] = []int{arr[i][1], arr[i][2]}
	}

	return ans
}

func quickSelect(arr [][]int, start, end, k int) {
	if start >= end {
		return
	}

	store := start
	idx := start + rand.Intn(end-start)
	pivot := arr[idx][0]
	arr[end], arr[idx] = arr[idx], arr[end]

	for i := start; i < end; i++ {
		if arr[i][0] < pivot {
			arr[i], arr[store] = arr[store], arr[i]
			store++
		}
	}

	arr[store], arr[end] = arr[end], arr[store]

	if store > k {
		quickSelect(arr, start, store-1, k)
	} else {
		quickSelect(arr, store+1, end, k)
	}
}

type MaxHeap [][]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][2] > h[j][2] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() []int        { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	p := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return p
}

// tc: O(n log(k))
func kClosest1(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, p := range points {
		tmp := p[0]*p[0] + p[1]*p[1]

		if h.Len() == k && h.Peek()[2] > tmp {
			heap.Pop(h)
		}

		if h.Len() < k {
			heap.Push(h, []int{p[0], p[1], tmp})
		}
	}

	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		popped := heap.Pop(h).([]int)
		ans[i] = []int{popped[0], popped[1]}
	}

	return ans
}

//	Notes
//	1.	there could exists equal distance points, since it's to find closest,
//		so put equal or larger to right

//	2.	too slow, because I got wrong about item search range... just do search

//	3.	add reference https://leetcode.com/problems/k-closest-points-to-origin/discuss/220235/Java-Three-solutions-to-this-classical-K-th-problem.

//		author adds some conclusion

//	4.	inspired from sample code

//		every time get a random pivot node

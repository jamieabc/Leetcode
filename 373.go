package main

import (
	"container/heap"
	"fmt"
)

//You are given two integer arrays nums1 and nums2 sorted in ascending order and an integer k.
//
//Define a pair (u,v) which consists of one element from the first array and one element from the second array.
//
//Find the k pairs (u1,v1),(u2,v2) ...(uk,vk) with the smallest sums.
//
//Example 1:
//
//Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//Output: [[1,2],[1,4],[1,6]]
//Explanation: The first 3 pairs are returned from the sequence:
//             [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//Example 2:
//
//Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
//Output: [1,1],[1,1]
//Explanation: The first 2 pairs are returned from the sequence:
//             [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//Example 3:
//
//Input: nums1 = [1,2], nums2 = [3], k = 3
//Output: [1,3],[2,3]
//Explanation: All possible pairs are returned from the sequence: [1,3],[2,3]
type Sum struct {
	val1, val2 int
}

type MaxHeap []Sum

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i].val1+h[i].val2 > h[j].val1+h[j].val2
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() Sum     { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Sum))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)

	for i := 0; i < k && i < len(nums1); i++ {
		for j := 0; j < k && j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]

			if h.Len() < k {
				heap.Push(h, Sum{
					nums1[i], nums2[j],
				})
			} else if sum < h.Peek().val1+h.Peek().val2 {
				heap.Pop(h)
				heap.Push(h, Sum{
					nums1[i], nums2[j],
				})
			}
		}
	}

	result := make([][]int, 0)
	for h.Len() > 0 {
		popped := heap.Pop(h).(Sum)
		result = append(result, []int{popped.val1, popped.val2})
	}

	return result
}

const (
	indexNotFound = -1
)

type node []int

func (n node) val() int {
	return n[0] + n[1]
}

type priorityQueue struct {
	queue []node
	limit int
}

func (q *priorityQueue) insert(target node) {
	q.queue = append(q.queue, target)
	bubbleUp(q.queue, len(q.queue)-1)
}

func (q *priorityQueue) data() [][]int {
	result := make([][]int, q.limit)
	for i := 0; i < q.limit; i++ {
		popped, err := q.pop()
		if nil != err {
			return result
		}
		result[i] = popped
	}
	return result
}

func (q *priorityQueue) pop() (node, error) {
	if len(q.queue) == 0 {
		return nil, fmt.Errorf("empty")
	}
	popped := q.queue[0]
	swap(q.queue, 0, len(q.queue)-1)
	q.queue = q.queue[:len(q.queue)-1]
	bubbleDown(q.queue, 0)
	return popped, nil
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(queue []node, index int) int {
	child := index*2 + 1
	if len(queue) <= child {
		return indexNotFound
	}
	return child
}

func rightChild(queue []node, index int) int {
	child := index*2 + 2
	if len(queue) <= child {
		return indexNotFound
	}
	return child
}

// the problem is this operation cannot guarantee maximum at the last element
// maybe I should use standard pop operation for queue to make sure every pop is smallest.
func bubbleUp(q []node, index int) {
	for index != 0 {
		p := parent(index)
		if smaller(q[index], q[p]) {
			swap(q, p, index)
			index = p
		} else {
			return
		}
	}
}

func bubbleDown(q []node, index int) {
	for index != indexNotFound {
		l := leftChild(q, index)
		r := rightChild(q, index)

		// both index not found
		if l == indexNotFound {
			return
		}

		// one index not found
		if r == indexNotFound && l != indexNotFound {
			if smaller(q[index], q[l]) {
				return
			}
			swap(q, l, index)
			return
		}

		if smaller(q[index], q[l]) && smaller(q[index], q[r]) {
			return
		}

		if smaller(q[l], q[r]) {
			swap(q, l, index)
			index = l
		} else {
			swap(q, r, index)
			index = r
		}
	}
}

func smaller(n1, n2 node) bool {
	return n1.val() < n2.val()
}

func swap(q []node, src, dst int) {
	q[src], q[dst] = q[dst], q[src]
}

func newPriorityQueue(limit int) *priorityQueue {
	return &priorityQueue{queue: make([]node, 0), limit: limit}
}

func kSmallestPairs1(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}
	maxCount := len(nums1) * len(nums2)
	if k > maxCount {
		k = maxCount
	}
	pq := newPriorityQueue(k)
	for _, i := range nums1 {
		for _, j := range nums2 {
			pq.insert(node{i, j})
		}
	}
	return pq.data()
}

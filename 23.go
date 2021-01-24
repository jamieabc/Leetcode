package main

import "container/heap"

// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
//
// Example:
//
// Input:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// Output: 1->1->2->3->4->4->5->6

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// tc: O(n log(k))
func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)

	for i := range lists {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	var head *ListNode
	ptr := head

	for h.Len() > 0 {
		popped := heap.Pop(h).(*ListNode)

		if ptr == nil {
			head = popped
			ptr = head
		} else {
			ptr.Next = popped
			ptr = ptr.Next
		}

		if popped.Next != nil {
			heap.Push(h, popped.Next)
		}
	}

	return head
}

// tc: O(kn)
func mergeKLists1(lists []*ListNode) *ListNode {
	var head, cur *ListNode
	var idx int

	for true {
		var next *ListNode

		for i := range lists {
			if lists[i] != nil {
				if next == nil || next.Val > lists[i].Val {
					next = lists[i]
					idx = i
				}
			}
		}

		if next == nil {
			break
		}

		if head == nil {
			head = next
			cur = next
		} else {
			cur.Next = next
			cur = next
		}
		lists[idx] = next.Next
	}

	return head
}

//	Notes
//	1.	pointer might be nil

//	2.	use original linked list

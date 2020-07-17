package main

// Given a linked list, rotate the list to the right by k places, where k is non-negative.
//
// Example 1:
//
// Input: 1->2->3->4->5->NULL, k = 2
// Output: 4->5->1->2->3->NULL
// Explanation:
// rotate 1 steps to the right: 5->1->2->3->4->NULL
// rotate 2 steps to the right: 4->5->1->2->3->NULL
// Example 2:
//
// Input: 0->1->2->NULL, k = 4
// Output: 2->0->1->NULL
// Explanation:
// rotate 1 steps to the right: 2->0->1->NULL
// rotate 2 steps to the right: 1->2->0->NULL
// rotate 3 steps to the right: 0->1->2->NULL
// rotate 4 steps to the right: 2->0->1->NULL

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 || head.Next == nil {
		return head
	}

	// find size of linked list
	var size int
	var ptr *ListNode
	for ptr = head; ptr.Next != nil; ptr = ptr.Next {
		size++
	}
	size++

	if k%size == 0 {
		return head
	}

	ptr.Next = head

	// go to size - (k % size)
	steps := size - (k % size)
	ptr = head
	var prev *ListNode
	for i := 0; i < steps; i++ {
		prev, ptr = ptr, ptr.Next
	}

	// start from that point, break linked list
	head = ptr
	prev.Next = nil

	return head
}

//	problems
//	1.	be careful about corner cases

//	2.	inspired form solution, I don't need to create link again, just connect
//		tail to head to form a closed loop, and break at some meeting node

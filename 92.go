package main

// Reverse a linked list from position m to n. Do it in one-pass.
//
// Note: 1 ≤ m ≤ n ≤ length of list.
//
// Example:
//
// Input: 1->2->3->4->5->NULL, m = 2, n = 4
// Output: 1->4->3->2->5->NULL
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// move pointer to m
	ptr := head
	var prev *ListNode
	for i := 0; i < m-1; i++ {
		prev, ptr = ptr, ptr.Next
	}

	start := prev
	prev, ptr = ptr, ptr.Next

	for i := 0; i < n-m; i++ {
		next := ptr.Next
		ptr.Next = prev
		prev, ptr = ptr, next
	}

	if m > 1 {
		start.Next.Next = ptr
		start.Next = prev
	} else {
		head.Next = ptr
		head = prev
	}

	return head
}

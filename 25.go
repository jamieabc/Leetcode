package main

// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
// Example:
//
// Given this linked list: 1->2->3->4->5
//
// For k = 2, you should return: 2->1->4->3->5
//
// For k = 3, you should return: 3->2->1->4->5
//
// Note:
//
// Only constant extra memory is allowed.
// You may not alter the values in the list's nodes, only nodes itself may be changed.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}

	for n := reverse(head, dummy, k); n != nil; {
		n = reverse(n.Next, n, k)
	}

	return dummy.Next
}

func reverse(start, prev *ListNode, k int) *ListNode {
	// check remain nodes count >= k
	var i int
	var ptr *ListNode
	for i, ptr = 0, start; i < k-1 && ptr != nil; i++ {
		ptr = ptr.Next
	}

	// reach end of linked list
	if ptr == nil {
		return nil
	}

	end := ptr.Next

	// reverse
	var next *ListNode
	ptr, next = start, start.Next
	for i = 0; i < k-1; i++ {
		n := next.Next
		next.Next = ptr
		next, ptr = n, next
	}

	prev.Next = ptr
	start.Next = end

	return start
}

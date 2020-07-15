package main

// Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
//
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//
// Example 1:
//
// Given 1->2->3->4, reorder it to 1->4->2->3.
// Example 2:
//
// Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// find half of linked list
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 1 2 3 4
	// 1 2 3 4 5
	mid := slow
	var odd bool
	if fast.Next == nil {
		// odd
		slow = slow.Next
		odd = true
	}

	// reverse second half of linked list
	prev := slow
	slow = slow.Next
	for slow != nil {
		tmp := slow.Next
		slow.Next = prev
		prev, slow = slow, tmp
	}
	slow = prev

	// mix first half with second half
	for ptr := head; ptr != mid; {
		forward, backward := ptr.Next, slow.Next
		ptr.Next = slow
		slow.Next = forward
		ptr, slow = forward, backward
	}

	if odd {
		mid.Next = nil
	} else {
		mid.Next.Next = nil
	}
}

//	problems
//	1.	be careful about mid node processing, odd list & even list are different

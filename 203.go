package main

//Remove all elements from a linked list of integers that have value val.
//
//Example:
//
//Input:  1->2->6->3->4->5->6, val = 6
//Output: 1->2->3->4->5

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	// 6 - 6
	// 1 - 2 - 3 - 6
	start := head
	cur := start

	// check for start of 6
	for cur != nil && cur.Val == val {
		cur = cur.Next
	}

	if cur != start {
		start = cur
	}

	if start == nil {
		return nil
	}

	cur = start.Next
	prev := start
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}

	return start
}

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var smallerStart, smallerEnd, largerStart, largerEnd *ListNode

	for node := head; node != nil; {
		next := node.Next
		node.Next = nil

		if node.Val < x {
			if smallerStart == nil {
				smallerStart = node
				smallerEnd = node
			} else {
				smallerEnd.Next = node
				smallerEnd = node
			}
		} else {
			if largerStart == nil {
				largerStart = node
				largerEnd = node
			} else {
				largerEnd.Next = node
				largerEnd = node
			}
		}

		node = next
	}

	if smallerStart != nil {
		smallerEnd.Next = largerStart
		return smallerStart
	}

	return largerStart
}

//	problems
//	1.	separate number may not exist in node, which means smaller/larger could
//		be empty

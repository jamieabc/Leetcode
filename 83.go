package main

//Given a sorted linked list, delete all duplicates such that each element appear only once.
//
//Example 1:
//
//Input: 1->1->2
//Output: 1->2
//
//Example 2:
//
//Input: 1->1->2->3->3
//Output: 1->2->3

func deleteDuplicates(head *ListNode) *ListNode {
	node := head

	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	node, prev := head.Next, head

	for node != nil {
		if node.Val == prev.Val {
			prev.Next = node.Next
			node = node.Next
		} else {
			prev, node = node, node.Next
		}
	}

	return head
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, cur *ListNode
	prev = head
	cur = head.Next
	for cur != nil {
		if cur.Val == prev.Val {
			// 1 1 1
			// 1 1 2
			prev.Next = cur.Next
			cur = cur.Next
		} else {
			prev = cur
			cur = cur.Next
		}
	}

	return head
}

// problems
// 1. forget to check for consecutive same vallue nodes

//	2.	inspired from solutino, only one variable is needed

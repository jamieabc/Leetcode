package main

// Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.
//
//
//
// Example 1:
//
// Input: head = [1,2,3,3,4,4,5]
// Output: [1,2,5]
//
// Example 2:
//
// Input: head = [1,1,1,2,3]
// Output: [2,3]
//
//
//
// Constraints:
//
//     The number of nodes in the list is in the range [0, 300].
//     -100 <= Node.val <= 100
//     The list is guaranteed to be sorted in ascending order.

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}
	pp := dummy

	var cur, prev *ListNode

	for prev, cur = head, head.Next; cur != nil; {
		// find first node with different value than previous node
		for ; cur != nil && cur.Val == prev.Val; cur = cur.Next {
		}

		// check if duplicate exists
		if prev.Next == cur {
			pp.Next = prev
			pp = pp.Next
		}

		if cur != nil {
			prev, cur = cur, cur.Next
		}
	}

	// check if last node is unique
	if prev.Next == nil {
		pp.Next = prev
	} else {
		pp.Next = nil
	}

	return dummy.Next
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next

	for ; next != nil && head.Val == next.Val; next = next.Next {
	}

	if head.Next == next {
		head.Next = deleteDuplicates1(next)
		return head
	}
	return deleteDuplicates1(next)
}

//	Notes
//	1.	duplicates means prev & next values are different, since single linked
//		list can only go forward, needs another variable to store previous node

//	2.	to make logic more consistent, a technique can be used: add dummy node

//	3.	inspired form https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/discuss/28339/My-Recursive-Java-Solution

//		author uses recursion, although it's less efficient than iteration, still
//		an interesting idea, I didn't even think of that

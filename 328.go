package main

// Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.
//
// You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.
//
// Example 1:
//
// Input: 1->2->3->4->5->NULL
// Output: 1->3->5->2->4->NULL
//
// Example 2:
//
// Input: 2->1->3->5->6->4->7->NULL
// Output: 2->3->6->7->1->5->4->NULL
//
//
//
// Constraints:
//
//     The relative order inside both the even and odd groups should remain as it was in the input.
//     The first node is considered odd, the second node even and so on ...
//     The length of the linked list is between [0, 10^4].

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	even, odd, oddHead := head, head.Next, head.Next

	for even.Next != nil && odd.Next != nil {
		even.Next = odd.Next
		even = even.Next
		odd.Next = even.Next
		odd = odd.Next
	}
	even.Next = oddHead
	return head
}

func oddEvenList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	even, odd := head, head.Next
	oddHead := odd

	var tmp1, tmp2 *ListNode
	for odd != nil && even != nil {
		tmp2 = odd.Next
		if tmp2 != nil {
			tmp1 = tmp2.Next
		} else {
			tmp1 = nil
		}

		if tmp1 == nil && tmp2 == nil {
			even.Next = oddHead
			odd.Next = nil
			return head
		}

		if tmp1 == nil {
			even.Next = tmp2
			even.Next.Next = oddHead
			odd.Next = nil
			return head
		}

		even.Next, odd.Next = tmp2, tmp1
		even, odd = tmp2, tmp1
	}

	return head
}

//	problems
//	1.	from solution, there's more elegant solution

package main

// You are given the head of a linked list, and an integer k.
//
// Return the head of the linked list after swapping the values of the kth node from the beginning and the kth node from the end (the list is 1-indexed).
//
//
//
// Example 1:
//
// Input: head = [1,2,3,4,5], k = 2
// Output: [1,4,3,2,5]
//
// Example 2:
//
// Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
// Output: [7,9,6,6,8,7,3,0,9,5]
//
// Example 3:
//
// Input: head = [1], k = 1
// Output: [1]
//
// Example 4:
//
// Input: head = [1,2], k = 1
// Output: [2,1]
//
// Example 5:
//
// Input: head = [1,2,3], k = 2
// Output: [1,2,3]
//
//
//
// Constraints:
//
//     The number of nodes in the list is n.
//     1 <= k <= n <= 105
//     0 <= Node.val <= 100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapNodes(head *ListNode, k int) *ListNode {
	var first, second *ListNode

	for node := head; node != nil; node = node.Next {
		if second != nil {
			second = second.Next
		}

		k--
		if k == 0 {
			first = node
			second = head
		}
	}

	first.Val, second.Val = second.Val, first.Val

	return head
}

func swapNodes1(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	var forward, backward *ListNode
	var end int

	traverse(dummy, k, 0, &end, &forward, &backward)

	if forward == backward {
		return dummy.Next
	}

	if forward.Next == backward {
		forward.Next = backward.Next
		forward.Next.Next, backward.Next = backward, backward.Next.Next
	} else if backward.Next == forward {
		backward.Next = forward.Next
		backward.Next.Next, forward.Next = forward, forward.Next.Next
	} else {
		tmp1, tmp2 := forward.Next.Next, backward.Next.Next
		forward.Next, backward.Next = backward.Next, forward.Next
		forward.Next.Next = tmp1
		backward.Next.Next = tmp2
	}

	return dummy.Next
}

func traverse(head *ListNode, k, cur int, end *int, forward, backward **ListNode) {
	if head == nil {
		*end = cur
		return
	}

	traverse(head.Next, k, cur+1, end, forward, backward)

	if cur == k-1 {
		*forward = head
	}

	if cur == *end-k-1 {
		*backward = head
	}
}

//	Notes
//	1.	it's really messy, something not clearly thought

//		the goal is to find nodes that node.Next is kth node, recursion is possible
//		but not efficient

//	2.	inspired from https://leetcode.com/problems/swapping-nodes-in-a-linked-list/discuss/1009800/C%2B%2B-One-Pass

//		voturbac uses way similar to sliding window, find first kth node then
//		keep moving by this determined range until end reached

//		a b c d e f
//		|     |				kth on node d
//		  |     |			keep moving
//			|     |			end reached, store node c

//	3.	it's swapping *value*, not node

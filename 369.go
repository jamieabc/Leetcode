package main

// Given a non-negative integer represented as non-empty a singly linked list of digits, plus one to the integer.
//
// You may assume the integer do not contain any leading zero, except the number 0 itself.
//
// The digits are stored such that the most significant digit is at the head of the list.
//
// Example :
//
// Input: [1,2,3]
// Output: [1,2,4]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
	if rec(head) {
		return &ListNode{
			Val:  1,
			Next: head,
		}
	}

	return head
}

// return value means needs to continue bubble up
func rec(node *ListNode) bool {
	if node == nil {
		return true
	}

	added := rec(node.Next)

	// last list, or carry shoud be processed
	if added {
		node.Val++
	}

	if node.Val <= 9 {
		return false
	}

	// overflow
	node.Val = 0

	return true
}

//	problems
//	1.	when checking not 9, it should be next node

//	2.	refactor, use a dummy node for simpler code. For nodes not head, check
//		if next node's value is not 9, but for head, it should check self, so
//		it's clever to add a dummy node which can makes all checking easier.

//		inspired from https://leetcode.com/problems/plus-one-linked-list/discuss/84125/Iterative-Two-Pointers-with-dummy-node-Java-O(n)-time-O(1)-space

//	3.	reference from https://leetcode.com/problems/plus-one-linked-list/discuss/84118/9-lines-recursive-*without*-helper

//		this is also a smart solution, it uses newly generated node as an info
//		of carry, if newly generated returned, then current node's value
//		needs to be incremented.

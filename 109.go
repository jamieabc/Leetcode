package main

//Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
//
//For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
//Example:
//
//Given the sorted linked list: [-10,-3,0,5,9],
//
//One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
//
//0
/// \
//-3   9
///   /
//-10  5

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	return mid(head)
}

func mid(node *ListNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Next == nil {
		return &TreeNode{Val: node.Val}
	}

	var slow, fast, prev *ListNode
	slow = node
	fast = node

	// use fast & slow pointer to find mid
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	return &TreeNode{
		Val:   slow.Val,
		Left:  mid(node),
		Right: mid(slow.Next),
	}
}

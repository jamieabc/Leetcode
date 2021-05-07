package main

// Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
//
// For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
// Example:
//
// Given the sorted linked list: [-10,-3,0,5,9],
//
// One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
//
// 0
// / \
// -3   9
// /   /
// -10  5

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// tc: O(n), sc: O(log(n))
func sortedListToBST(head *ListNode) *TreeNode {
	size := findSize(head)
	node := head

	return inOrder(0, size-1, &node)
}

func inOrder(start, end int, node **ListNode) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	left := inOrder(start, mid-1, node)

	current := &TreeNode{
		Val:  (*node).Val,
		Left: left,
	}

	*node = (*node).Next

	current.Right = inOrder(mid+1, end, node)

	return current
}

func findSize(node *ListNode) int {
	var size int
	for ; node != nil; node = node.Next {
		size++
	}

	return size
}

// tc: O(n log(n)), sc: O(log(n))
func sortedListToBST1(head *ListNode) *TreeNode {
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

//	Notes
//	1.	inspired from solution, at first i thought finding middle take O(n) to
//		traverse, but then realize that it's actually n/2 operations

//		tc looks like O(n^2), but take close look:
//		n/2 + 2 * (n/4) + 4 * (n/8) + ... = n/2 + n/2 + n/2 + ...
//		since it takes log(n) to cut, overall tc becomes (n/2) * log(n) which
//		is O(n log(n))

//		sc: O(log(n)), because each recursion takes memory

//	2.	inspired from solution, stores linked list into array, which saves
//		time for finding middle node, tc: O(n)

//	3.	inspired from https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/discuss/596899/JavaPython-Clean-code-O(N)-time-O(logN)-space

//		very beautiful solution, w/o knowing actual mid value, just keep cutting
//		linked list to half until left-most node is touched

//		w/ this way, not need to find middle and convert to BST, it converts
//		when boundary meets...

//		by the merge-sort like way, use middle point to determine the left
//		and right side; but why does middle is not needed?

//		as long as in-order traversal is guaranteed, the order just like
//		linked-list. So, how to make sure sequence in pre-order traversal?

//		in-order: LNR, with proper notation (index), in-order can be guaranteed
//		this is a very brilliant solution

//	4.	inspired from solution, can also put all values into array, then
//		convert to BST, tc: O(n), sc: O(n)

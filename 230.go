package main

// Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
//
// Note:
// You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
//
// Example 1:
//
// Input: root = [3,1,4,null,2], k = 1
//    3
//   / \
//  1   4
//   \
//    2
// Output: 1
//
// Example 2:
//
// Input: root = [5,3,6,2,4,null,null,1], k = 3
//        5
//       / \
//      3   6
//     / \
//    2   4
//   /
//  1
// Output: 3
//
// Follow up:
// What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	node := root
	stack := make([]*TreeNode, 0)

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return node.Val
		}

		node = node.Right
	}

	return 0
}

func kthSmallest1(root *TreeNode, k int) int {
	_, val := inOrder(root, k)

	return val
}

// return sub-tree count, val
func inOrder(node *TreeNode, k int) (int, int) {
	if node == nil {
		return 0, 0
	}

	left, val := inOrder(node.Left, k)

	if left >= k {
		return k, val
	}

	if left == k-1 {
		return k, node.Val
	}

	right, val := inOrder(node.Right, k-1-left)

	if right >= k-1-left {
		return k, val
	}

	return left + right + 1, 0
}

//	Notes
//	1.	not the fasted, use iteration instead

//	2.	not update node

//	3.	not add correct node into stack

//	4.	this is really hard for me to figure out how to write iterative code to
//		traverse tree. I need to practice more about it.

//		I think problem comes from how is stack used. Should I put every node
//		into it and pop it out later? Or should I only put nodes that left child
//		is parsed. If a node has right child, does it need to put into stack?

//	5.	reference from https://leetcode.com/problems/kth-smallest-element-in-a-bst/discuss/63660/3-ways-implemented-in-JAVA-(Python)%3A-Binary-Search-in-order-iterative-and-recursive

//		author provides a cleaner way of iterative traversing, compare to my solution, I think most
//		differences come from how is tree traversed.

//		in-order traversal means if a node has left child, go to left child. If no left child, go
//		self, then go right. for right child, check if left child exist, keep looping this.

//		I should follow this kind of traversing process, initially put all nodes with left child
//		into stack. Then for processing, focus on self and right child. If right child has left child,
//		looping until no left child.

//	6.	from post, there's a more beautiful way to write iterative code. the point here is to use
//		both stack & pointer. stack saves the node that already process left child. node is the one
//		that needs to process left child. I need to clarify whole process: always loop left child,
//		and process self, then right child

//	7.	for followups, I have no idea. reference from https://leetcode.com/problems/kth-smallest-element-in-a-bst/discuss/63659/What-if-you-could-modify-the-BST-node's-structure

//		it stores count of all children below (including self), with this info, every time insertion
//		or deletion happens, the path than includes this node will be updated, which is O(h), h is
//		height of BST

//		and search for kth element is also in O(h) time

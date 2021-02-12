package main

// Given a node in a binary search tree, find the in-order successor of that node in the BST.
//
// If that node has no in-order successor, return null.
//
// The successor of a node is the node with the smallest key greater than node.val.
//
// You will have direct access to the node but not to the root of the tree. Each node will have a reference to its parent node. Below is the definition for Node:
//
// class Node {
//     public int val;
//     public Node left;
//     public Node right;
//     public Node parent;
// }
//
//
// Follow up:
//
// Could you solve it without looking up any of the node's values?
//
//
//
// Example 1:
//
//
// Input: tree = [2,1,3], node = 1
// Output: 2
// Explanation: 1's in-order successor node is 2. Note that both the node and the return value is of Node type.
// Example 2:
//
//
// Input: tree = [5,3,6,2,4,null,null,1], node = 6
// Output: null
// Explanation: There is no in-order successor of the current node, so the answer is null.
// Example 3:
//
//
// Input: tree = [15,6,18,3,7,17,20,2,4,null,13,null,null,null,null,null,null,null,null,9], node = 15
// Output: 17
// Example 4:
//
//
// Input: tree = [15,6,18,3,7,17,20,2,4,null,13,null,null,null,null,null,null,null,null,9], node = 13
// Output: 15
// Example 5:
//
// Input: tree = [0], node = 0
// Output: null
//
//
// Constraints:
//
// -10^5 <= Node.val <= 10^5
// 1 <= Number of Nodes <= 10^4
// All Nodes will have unique values.

/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func inorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}

	// larger one comes from parent
	if node.Right == nil {
		prev := node
		for node = node.Parent; node != nil && node.Left != prev; prev, node = node, node.Parent {
		}

		return node
	}

	// larger node comes from left-most node of right subtree
	for node = node.Right; node != nil && node.Left != nil; node = node.Left {
	}
	return node
}

func inorderSuccessor1(node *Node) *Node {
	var prev *Node
	for n := node; n != nil; prev, n = n, n.Parent {
		// node is left of parent
		if prev == n.Left && prev != nil {
			return n
		}

		// node with right child, find left-most node
		for nn := node.Right; nn != nil && n.Right != prev; nn = nn.Left {
			if nn.Left == nil {
				return nn
			}
		}
	}

	return nil
}

//	Notes
//	1.	if only node is root, I cannot always assume parent is not nil

//	2.	when I want to access two variables (a.b.c), need to care that it might
//		cause nil pointer problem

//	3.	refactor in a better recursive

//	4.	time complexity is O(log n)

//	5.	inspired from https://leetcode.com/problems/inorder-successor-in-bst-ii/discuss/231587/Java-find-in-parents-or-find-in-descendents

//		instead of trying to compare node, compare value which is more straight forward

package main

// Given a binary tree where all the right nodes are either leaf nodes with a sibling (a left node that shares the same parent node) or empty, flip it upside down and turn it into a tree where the original right nodes turned into left leaf nodes. Return the new root.
//
// Example:
//
// Input: [1,2,3,4,5]
//
//     1
//    / \
//   2   3
//  / \
// 4   5
//
// Output: return the root of the binary tree [4,5,2,#,#,3,1]
//
//    4
//   / \
//  5   2
//     / \
//    3   1
//
// Clarification:
//
// Confused what [4,5,2,#,#,3,1] means? Read more below on how binary tree is serialized on OJ.
//
// The serialization of a binary tree follows a level order traversal, where '#' signifies a path terminator where no node exists below.
//
// Here's an example:
//
//    1
//   / \
//  2   3
//     /
//    4
//     \
//      5
//
// The above binary tree is serialized as [1,2,3,#,#,4,#,#,5].

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	var next, prev, tmp *TreeNode
	cur := root

	for cur != nil {
		// preserve original tree order
		next = cur.Left

		// flip tree
		cur.Left = tmp

		// preserve right node for next round
		tmp = cur.Right

		// flip tree
		cur.Right = prev

		// next round
		prev = cur
		cur = next
	}

	return prev
}

func upsideDownBinaryTree4(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	newRoot := upsideDownBinaryTree4(root.Left)
	root.Left.Left = root.Right
	root.Left.Right = root
	root.Left, root.Right = nil, nil

	return newRoot
}

func upsideDownBinaryTree3(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	var node *TreeNode
	stack := make([]*TreeNode, 0)

	// find new root (the left most node from original tree)
	for node = root; node.Left != nil; node = node.Left {
		stack = append(stack, node)
	}
	newRoot := node

	for len(stack) > 0 {
		node = stack[len(stack)-1]

		if node.Right != nil {
			node.Left.Left = node.Right
		}
		node.Left.Right = node
		node.Left = nil
		node.Right = nil

		stack = stack[:len(stack)-1]
	}

	return newRoot
}

func upsideDownBinaryTree2(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	var newRoot *TreeNode
	for newRoot = root; newRoot.Left != nil; newRoot = newRoot.Left {
	}

	postOrder(root)

	return newRoot
}

func postOrder(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	l, r := postOrder(node.Left), postOrder(node.Right)

	if l != nil {
		l.Right = node
		l.Left = r
		node.Left, node.Right = nil, nil
	}

	return node
}

func upsideDownBinaryTree1(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	return traverse(root)
}

// return traverse(root)
func traverse(node *TreeNode) *TreeNode {
	// new root must be left most node
	if node.Left == nil {
		return node
	}

	newRoot := traverse(node.Left)

	if node.Right != nil {
		node.Left.Left = node.Right
	}

	node.Left.Right = node
	node.Left = nil
	node.Right = nil

	return newRoot
}

//	Notes
//	1.	beware of corner cases, especially when empty node, single node

//	2.	limitation on right, so left could be empty or single

//	3.	reference from https://www.geeksforgeeks.org/flip-binary-tree/
//		it's more clear explanation

//	4.	having problem updating function parameter of pointer, then found
//		explanation from https://github.com/golang/go/issues/25932

//	5.	wrong checking, when right child is nil, still need to flip tree

//	6.	inspired from https://leetcode.com/problems/binary-tree-upside-down/discuss/49412/Clean-Java-solution

//		no need to have additional newRoot as function parameter, can just
//		make it a return value. And flipping tree can use recursive only

//	7.	the most important information is that left-mode node of original
//		tree will be new tree's root. Because I didn't know this, my original
//		method starts from right-most node.

//	8.	reference from https://leetcode.com/problems/binary-tree-upside-down/discuss/49432/Easy-O(n)-iteration-solution-Java

//		visualize flow:
//			1
//		  2   3    =>     2 (next)    1 (prev)    3 (tmp)
//		4   5           4   5

//		=>         2 (prev)  4 (next)   5 (tmp)
//				 3   1

//		=>		4
//			  5    2
//				 3   1

//		I can understand what it wants to do, but cannot know how author
//		come up idea...

//		But some clue I can think ok:
//		- when flipping tree, right child be will changed into left node
//		- since original tree structure is tear apart, it needs to store
//		  next & prev nodes

//	9.	inspired from https://leetcode.com/problems/binary-tree-upside-down/discuss/49406/Java-recursive-(O(logn)-space)-and-iterative-solutions-(O(1)-space)-with-explanation-and-figure

//		very elegant solution, both in recursive & iterative way
//		the reason it's elegant is because author finds out a proper way to model
//		the problem in such simple way. I was still in the form of current & next,
//		but author sees this problem is node.left.left, which avoid some checkings

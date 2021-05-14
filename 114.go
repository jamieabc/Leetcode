package main

//Given a binary tree, flatten it to a linked list in-place.
//
//For example, given the following tree:
//
//    1
//   / \
//  2   5
// / \   \
//3   4   6
//
//The flattened tree should look like:
//
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	for node := root; node != nil; {
		if node.Left != nil {
			nodeWithoutRight := node.Left

			// find first node on left subtree without right child, so that this node
			// will be flattened in the future
			for ; nodeWithoutRight.Right != nil; nodeWithoutRight = nodeWithoutRight.Right {
			}

			nodeWithoutRight.Right = node.Right
			node.Right = node.Left
			node.Left = nil
		}

		node = node.Right
	}
}

func flatten5(root *TreeNode) {
	node := root

	for node != nil {
		if node.Left != nil {
			if node.Right != nil {
				for next := node.Left; true; next = next.Right {
					if next.Right == nil {
						next.Right = node.Right
						break
					}
				}
			}
			node.Right = node.Left
			node.Left = nil
		}

		node = node.Right
	}
}

func flatten4(root *TreeNode) {
	stack := []*TreeNode{root}

	// the point is to keep previous node, so that original tree can be transformed
	var prev *TreeNode

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// put in reverse order (right, left) such that get from stack with correct order
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}

		if cur.Left != nil {
			stack = append(stack, cur.Left)
			cur.Left = nil
		}

		if prev != nil {
			prev.Right = cur
		}
		prev = cur
	}
}

func flatten3(node *TreeNode) {
	var prev *TreeNode
	postOrder(node, &prev)
}

func postOrder(node *TreeNode, prev **TreeNode) {
	if node == nil {
		return
	}

	postOrder(node.Right, prev)
	postOrder(node.Left, prev)

	// very beautiful
	node.Right = *prev
	node.Left = nil
	*prev = node
}

func flatten2(root *TreeNode) {
	inOrder(root)
}

func inOrder(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	right := node.Right
	node.Right = nil

	if node.Left != nil {
		node.Right = inOrder(node.Left)
		node.Left = nil
	}

	cur := node
	for ; cur.Right != nil; cur = cur.Right {
	}

	cur.Right = inOrder(right)

	return node
}

func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
	convert(root)
}

func convert(node *TreeNode) (*TreeNode, *TreeNode) {
	if node.Left == nil && node.Right == nil {
		return node, node
	}

	var tmp *TreeNode

	if node.Left == nil && node.Right != nil {
		node.Right, tmp = convert(node.Right)
		return node, tmp
	}

	if node.Left != nil && node.Right == nil {
		node.Right, tmp = convert(node.Left)
		node.Left = nil
		return node, tmp
	}

	left, leftEnd := convert(node.Left)
	right, rightEnd := convert(node.Right)

	node.Right = left
	leftEnd.Right = right
	node.Left = nil

	return node, rightEnd
}

//	Notes
//	1.	becareful of boundary condition...cannot assume root is non-nil

//	2.	for recursion, better to set current node right tree to nil, such that
//		no infinite loop occurs

//	3.	inspired from solution, there's a very beautiful description of recursion:
//		Recursion is all about postponing decisions until something else is completed.

//		what a precise and exact description

//	4.	inspired from solution, morris traversal can reach tc O(n) & sp O(1)

//		very beautiful solution, especially when merging left node to right node,
//		it searches first node of left-subtree w/o right child

//	5.	inspired from https://leetcode.com/problems/flatten-binary-tree-to-linked-list/discuss/36977/My-short-post-order-traversal-Java-solution-for-share

//		author provides a very elegant recursion solution to this
//		the key point here is to solve right side first, then left side, and eventually
//		build a flatten list

//		the best part of it is no need to traverse when a new sub-tree going to right,
//		from #inOrder, the loop is about finding where to concat right node, so this way
//		reduces this part

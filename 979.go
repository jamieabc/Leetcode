package main

// Given the root of a binary tree with N nodes, each node in the tree has node.val coins, and there are N coins total.
//
// In one move, we may choose two adjacent nodes and move one coin from one node to another.  (The move may be from parent to child, or from child to parent.)
//
// Return the number of moves required to make every node have exactly one coin.
//
//
//
// Example 1:
//
// Input: [3,0,0]
// Output: 2
// Explanation: From the root of the tree, we move one coin to its left child, and one coin to its right child.
//
// Example 2:
//
// Input: [0,3,0]
// Output: 3
// Explanation: From the left child of the root, we move two coins to the root [taking two moves].  Then, we move one coin from the root of the tree to the right child.
//
// Example 3:
//
// Input: [1,0,2]
// Output: 2
//
// Example 4:
//
// Input: [1,0,0,null,3]
// Output: 4
//
//
//
// Note:
//
//     1<= N <= 100
//     0 <= node.val <= N

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
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
func distributeCoins(root *TreeNode) int {
	var steps int

	_ = traverse(root, &steps)

	return steps
}

// return value positive means transfer from child to parent
// return value negative means transfer from parent to child
func traverse(node *TreeNode, steps *int) int {
	if node == nil {
		return 0
	}

	var left, right int
	if node.Left != nil {
		left = traverse(node.Left, steps)
	}

	if node.Right != nil {
		right = traverse(node.Right, steps)
	}

	*steps += abs(left) + abs(right)

	return node.Val - 1 + left + right
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	problems
//	1.	needed steps are absolute value, if it's positive means transfer to
//		parent, if it's negative means parent transfer to child, so both
//		counts. But return value should be positive/negative to demonstrate
//		need/donate coins

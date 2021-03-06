package main

import "math"

// Given the root of a binary tree, find the maximum value V for which there exists different nodes A and B where V = |A.val - B.val| and A is an ancestor of B.
//
// (A node A is an ancestor of B if either: any child of A is equal to B, or any child of A is an ancestor of B.)
//
//
//
// Example 1:
//
// Input: [8,3,10,1,6,null,14,null,null,4,7,13]
// Output: 7
// Explanation:
// We have various ancestor-node differences, some of which are given below :
// |8 - 3| = 5
// |3 - 7| = 4
// |8 - 1| = 7
// |10 - 13| = 3
// Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.
//
//
//
// Note:
//
//     The number of nodes in the tree is between 2 and 5000.
//     Each node will have value between 0 and 100000.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return dfs(root, root.Val, root.Val)
}

func dfs(node *TreeNode, largest, smallest int) int {
	if node == nil {
		return largest - smallest
	}

	largest = max(largest, node.Val)
	smallest = min(smallest, node.Val)

	return max(dfs(node.Left, largest, smallest), dfs(node.Right, largest, smallest))
}

func maxAncestorDiff3(root *TreeNode) int {
	var maxDiff int

	dfs(root, &maxDiff)

	return maxDiff
}

func dfs3(node *TreeNode, maxDiff *int) (int, int) {
	if node == nil {
		return -1, -1
	}

	lMin, lMax := dfs3(node.Left, maxDiff)
	rMin, rMax := dfs3(node.Right, maxDiff)

	smallest, largest := node.Val, node.Val

	if lMin != -1 {
		*maxDiff = max(
			*maxDiff,
			max(
				abs(node.Val-lMin),
				abs(node.Val-lMax),
			),
		)
		smallest, largest = min(smallest, lMin), max(largest, lMax)
	}

	if rMin != -1 {
		*maxDiff = max(
			*maxDiff,
			max(
				abs(node.Val-rMin),
				abs(node.Val-rMax),
			),
		)
		smallest, largest = min(smallest, rMin), max(largest, rMax)
	}

	return smallest, largest
}

func maxAncestorDiff2(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func dfs(node *TreeNode, mx, mn int) int {
	if node == nil {
		return mx - mn
	}

	mx = max(mx, node.Val)
	mn = min(mn, node.Val)

	return max(dfs(node.Left, mx, mn), dfs(node.Right, mx, mn))
}

func maxAncestorDiff1(root *TreeNode) int {
	result := math.MinInt32
	postOrder(root, &result)

	return result
}

// return value is min, max of sub-tree
func postOrder(node *TreeNode, result *int) (int, int) {
	if node == nil {
		return 0, 0
	}

	l1, l2, r1, r2 := node.Val, node.Val, node.Val, node.Val

	if node.Left != nil {
		l1, l2 = postOrder(node.Left, result)
	}

	if node.Right != nil {
		r1, r2 = postOrder(node.Right, result)
	}

	v1 := min(node.Val, min(l1, r1))
	v2 := max(node.Val, max(l2, r2))

	*result = max(*result, max(abs(node.Val, min(l1, r1)), abs(node.Val, max(l2, r2))))

	return v1, v2
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func abs(i, j int) int {
	if i >= j {
		return i - j
	}
	return j - i
}

//	Notes
//	1.	forget about conditions that maximum value might come from max-min or
//		min-max

//	2.	add reference https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/discuss/274610/JavaC%2B%2BPython-Top-Down

//		lee uses dfs, passes max & min to leaf nodes to decide what's largest
//		difference

//	3.	use a complex way to solve this, I know problem wants (max, min) on
//		every path, so I try bottom-up approach. But it's really a waste,
//		because if only thing to know is (max, min), then can pass it down to
//		leaf node. If has less condition checking

//		also, it reduces check for abs, since max & min makes difference
//		positive

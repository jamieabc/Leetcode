package main

// The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.
//
// Determine the maximum amount of money the thief can rob tonight without alerting the police.
//
// Example 1:
//
// Input: [3,2,3,null,3,null,1]
//
// 3
// / \
// 2   3
// \   \
// 3   1
//
// Output: 7
// Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
// Example 2:
//
// Input: [3,4,5,1,3,null,1]
//
// 3
// / \
// 4   5
// / \   \
// 1   3   1
//
// Output: 9
// Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rob(root *TreeNode) int {
	v1, v2 := dfs(root)

	return max(v1, v2)
}

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	l1, l2 := dfs(node.Left)
	r1, r2 := dfs(node.Right)

	return max(node.Val+l2+r2, l1+r1), l1 + r1
}

func rob1(root *TreeNode) int {
	return dfs1(root)
}

func dfs1(node *TreeNode) int {
	if node == nil {
		return 0
	}

	var val int

	if node.Left != nil {
		val += dfs1(node.Left.Left) + dfs1(node.Left.Right)
	}

	if node.Right != nil {
		val += dfs1(node.Right.Left) + dfs1(node.Right.Right)
	}

	return max(node.Val+val, dfs1(node.Left)+dfs1(node.Right))
}

func max(i, j int) int {
	if i >= j {
		return i
	}

	return j
}

//	Notes
//	1.	parent-children cannot be selected at the same time, conditions meet
//		this not only be select one level and not select the other.

//		e.g.
//		4
//       \
//        2
//         \
//          1
//           \
//            3

//		4+1 == 2+3, but the other situation is to select 4+3

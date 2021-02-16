package main

// Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
//
// Example:
//
// Input: [1,2,3,null,5,null,4]
// Output: [1, 3, 4]
// Explanation:
//
//    1            <---
//  /   \
// 2     3         <---
//  \     \
//   5     4       <---

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)

	// becareful about this, because root put into queue, so it needs to be
	// non-nil
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	var i, idx int
	for len(queue) != 0 {
		for i, idx = 0, len(queue)-1; i <= idx; i++ {
			s := queue[i]
			if s.Left != nil {
				queue = append(queue, s.Left)
			}
			if s.Right != nil {
				queue = append(queue, s.Right)
			}
		}
		result = append(result, queue[idx].Val)
		queue = queue[idx+1:]
	}

	return result
}

func rightSideView1(root *TreeNode) []int {
	ans := make([]int, 0)

	dfs(root, 0, &ans)

	return ans
}

func dfs(node *TreeNode, level int, ans *[]int) {
	if node == nil {
		return
	}

	if len(*ans) < level+1 {
		*ans = append(*ans, make([]int, level+1-len(*ans))...)
	}

	(*ans)[level] = node.Val

	dfs(node.Left, level+1, ans)
	dfs(node.Right, level+1, ans)
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/binary-tree-right-side-view/discuss/56012/My-simple-accepted-solution(JAVA)

//		author put right child then left child, which means right most child will be first one

//		also, in order to determine the first right-mode node, author uses a number to check

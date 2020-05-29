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
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}

	var i, idx int
	for len(stack) != 0 {
		for i, idx = 0, len(stack)-1; i <= idx; i++ {
			s := stack[i]
			if s.Left != nil {
				stack = append(stack, s.Left)
			}
			if s.Right != nil {
				stack = append(stack, s.Right)
			}
		}
		result = append(result, stack[idx].Val)
		stack = stack[idx+1:]
	}

	return result
}

//	problems
//	1.	inspired from https://leetcode.com/problems/binary-tree-right-side-view/discuss/56012/My-simple-accepted-solution(JAVA)

//		author put right child then left child, which means right most child will be first one

//		also, in order to determine the first right-mode node, author uses a number to check

package main

// Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:
//
// The root is the maximum number in the array.
// The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
// The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
// Construct the maximum tree by the given array and output the root node of this tree.
//
// Example 1:
//
// Input: [3,2,1,6,0,5]
// Output: return the tree root node representing the following tree:
//
//       6
//     /   \
//    3     5
//     \    /
//      2  0
//        \
//         1
// Note:
//
// The size of the given array will be in the range [1,1000].

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := make([]*TreeNode, 0)

	var s *TreeNode
	for _, n := range nums {
		cur := &TreeNode{
			Val: n,
		}

		// make sure items in stack are decreasing order
		// if increasing is found, keep popping and build tree
		for len(stack) != 0 && stack[len(stack)-1].Val < n {
			s = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) > 0 && stack[len(stack)-1].Val < n {
				stack[len(stack)-1].Right = s
			} else {
				cur.Left = s
			}
		}

		stack = append(stack, cur)
	}

	for len(stack) > 0 {
		s = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) > 0 {
			stack[len(stack)-1].Right = s
		}
	}

	return s
}

func constructMaximumBinaryTree1(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, start, end int) *TreeNode {
	if start < 0 || end == len(nums) || start > end {
		return nil
	}

	if start == end {
		return &TreeNode{
			Val: nums[start],
		}
	}

	idx := findMax(nums, start, end)

	return &TreeNode{
		Val:   nums[idx],
		Left:  build(nums, start, idx-1),
		Right: build(nums, idx+1, end),
	}
}

func findMax(nums []int, start, end int) int {
	m, idx := nums[start], start

	for i := start + 1; i <= end; i++ {
		if m < nums[i] {
			idx = i
			m = nums[i]
		}
	}

	return idx
}

//	problems
//	1.	too slow, tc: O(n^2)

//	2.	inspired from https://leetcode.com/problems/maximum-binary-tree/discuss/258364/Python-O(n)-solution-with-explanation.

//		the problem can be solved in O(n). if next node is smaller, put next node into stack top
//		node's right child, and put next node into stack. if next node is larger, keep popping until
//		stack is empty or no-smaller. put last popped node as left child of next node

//	3.	add reference https://leetcode.com/problems/maximum-binary-tree/discuss/106147/C%2B%2B-8-lines-O(n-log-n)-map-plus-stack-with-binary-search

// 		author explains how O(n) is done

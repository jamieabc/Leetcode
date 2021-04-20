package main

// Given the root of a binary tree, return the postorder traversal of its nodes' values.
//
//
//
// Example 1:
//
// Input: root = [1,null,2,3]
// Output: [3,2,1]
//
// Example 2:
//
// Input: root = []
// Output: []
//
// Example 3:
//
// Input: root = [1]
// Output: [1]
//
// Example 4:
//
// Input: root = [1,2]
// Output: [2,1]
//
// Example 5:
//
// Input: root = [1,null,2]
// Output: [2,1]
//
//
//
// Constraints:
//
// The number of the nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
//
//
// Follow up:
//
// Recursive solution is trivial, could you do it iteratively?

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	if root == nil {
		return nums
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nums = append(nums, node.Val)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	// reverse order
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}

func postorderTraversal1(root *TreeNode) []int {
	nums := make([]int, 0)

	postOrder(root, &nums)

	return nums
}

func postOrder(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}

	postOrder(node.Left, nums)
	postOrder(node.Right, nums)

	*nums = append(*nums, node.Val)
}

//	Notes
//	1.	inspired from solution, post-order is similar to pre-order, just that
//		children traverse is right -> left

//	2.	inspired from https://leetcode.com/problems/binary-tree-postorder-traversal/discuss/45551/Preorder-Inorder-and-Postorder-Iteratively-Summarization

//		author provides all iterative traversing ways

//	3.	inspired from https://leetcode.com/problems/binary-tree-postorder-traversal/discuss/45559/My-Accepted-code-with-explaination.-Does-anyone-have-a-better-idea

//		author compare preorder and postorder, then come up with a solution,
//		I think this one is more suitable for me

//		pre-order:  NLR
//		post-order: LRN

//	4.	inspired from https://leetcode.com/problems/binary-tree-postorder-traversal/discuss/45550/C%2B%2B-Iterative-Recursive-and-Morris-Traversal

//		author says something about Morris traversal, uses constant space
//		not implement

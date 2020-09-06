package main

//Return the root node of a binary search tree that matches the given preorder traversal.
//
//(Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)
//
//
//
//Example 1:
//
//Input: [8,5,1,7,10,12]
//Output: [8,5,10,1,7,null,12]
//
//
//
//Note:
//
//    1 <= preorder.length <= 100
//    The values of preorder are distinct.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bstFromPreorder(nums []int) *TreeNode {
	root, _ := recursive(nums, 0, -1)
	return root
}

// tc: O(n)
func recursive(nums []int, idx, prevIdx int) (*TreeNode, int) {
	if len(nums) == 0 || idx >= len(nums) || (prevIdx != -1 && nums[idx] > nums[prevIdx]) {
		return nil, idx
	}

	node := &TreeNode{
		Val: nums[idx],
	}

	var right, next int
	node.Left, right = recursive(nums, idx+1, idx)
	node.Right, next = recursive(nums, right, prevIdx)

	return node, next
}

// tc: average is O(n log n), worst is O(n^2)
func bstFromPreorder2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	node := &TreeNode{
		Val: nums[0],
	}

	if len(nums) == 1 {
		return node
	}

	left := 1
	for ; left < len(nums); left++ {
		if nums[left] > nums[0] {
			break
		}
	}

	node.Left = bstFromPreorder(nums[1:left])
	node.Right = bstFromPreorder(nums[left:])

	return node
}

func bstFromPreorder1(preorder []int) *TreeNode {
	return construct(preorder)
}

func construct(preorder []int) *TreeNode {
	// terminate if index out of range
	length := len(preorder)
	if length == 0 {
		return nil
	}

	val := preorder[0]
	node := &TreeNode{
		Val: val,
	}

	left, right := -1, length

	if 1 < length && preorder[1] < val {
		left = 1
	}

	// find index of right node
	for i := 1; i < length; i++ {
		if preorder[i] > val {
			right = i
			break
		}
	}

	// found smaller number, those belongs to left
	if left != -1 {
		node.Left = construct(preorder[left:right])
	}

	// found larger number, those belongs to right
	if right != length {
		node.Right = construct(preorder[right:])
	}

	return node
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/discuss/252232/JavaC%2B%2BPython-O(N)-Solution

//		there exists O(n) solution

//	2.	inspired from https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/discuss/252722/Python-stack-solution-beats-100-on-runtime-and-memory

//		it reminds me that checking value greater than parent is like
//		stack operation

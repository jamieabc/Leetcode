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
func bstFromPreorder(preorder []int) *TreeNode {
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

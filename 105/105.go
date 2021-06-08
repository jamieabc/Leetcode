package main

// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
//
//
//
// Example 1:
//
// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// Output: [3,9,20,null,null,15,7]
//
// Example 2:
//
// Input: preorder = [-1], inorder = [-1]
// Output: [-1]
//
//
//
// Constraints:
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder and inorder consist of unique values.
// Each value of inorder also appears in preorder.
// preorder is guaranteed to be the preorder traversal of the tree.
// inorder is guaranteed to be the inorder traversal of the tree.


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    size := len(preorder)

    if size == 0 {
        return nil
    }

    // build node
    node := &TreeNode{
        Val: preorder[0],
    }

    // find location at inorder for first item of preoder
    var idx int
    for ; idx < size; idx++ {
        if inorder[idx] == preorder[0] {
            break
        }
    }

    // recursively generate left & right
    node.Left = buildTree(preorder[1:idx+1], inorder[:idx])
    node.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

    return node
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	return dfs1(preorder, inorder)
}

func dfs1(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	cur := &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return cur
	}

	var in int
	for i := range inorder {
		if inorder[i] == preorder[0] {
			in = i
			break
		}
	}

	cur.Left = dfs(preorder[1:1+in], inorder[:in])
	cur.Right = dfs(preorder[1+in:], inorder[in+1:])

	return cur
}

//	Notes
//	1.	inspired from solution, use hashmap to find index, reduce tc from
//		O(n^2) to O(n)

//		but to use hashmap, every value should be unique

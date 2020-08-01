package main

//Given the root node of a binary search tree (BST) and a value to be inserted into the tree, insert the value into the BST. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.
//
//Note that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.
//
//For example,
//
//Given the tree:
//        4
//       / \
//      2   7
//     / \
//    1   3
//And the value to insert: 5
//
//You can return this binary search tree:
//
//         4
//       /   \
//      2     7
//     / \   /
//    1   3 5
//
//This tree is also valid:
//
//         5
//       /   \
//      2     7
//     / \
//    1   3
//         \
//          4
//
//
//
//Constraints:
//
//    The number of nodes in the given tree will be between 0 and 10^4.
//    Each node will have a unique integer value from 0 to -10^8, inclusive.
//    -10^8 <= val <= 10^8
//    It's guaranteed that val does not exist in the original BST.

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	ptr := &root

	for *ptr != nil {
		if (*ptr).Val > val {
			ptr = &(*ptr).Left
		} else {
			ptr = &(*ptr).Right
		}
	}

	*ptr = &TreeNode{
		Val: val,
	}

	return root
}

func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	node := root
	for true {
		if val > node.Val {
			// go right
			if node.Right == nil {
				node.Right = &TreeNode{
					Val: val,
				}
				break
			}
			node = node.Right
		} else {
			// go left
			if node.Left == nil {
				node.Left = &TreeNode{
					Val: val,
				}
				break
			}
			node = node.Left
		}
	}

	return root
}

//	problems
//	1.	inspired from https://leetcode.com/problems/insert-into-a-binary-search-tree/discuss/222959/c++-iterative-Linus-style./324072

//		this solution depends on clear understanding of golang, for
//		example, if a data in struct is nil, golang still allocate
//		memory space for it, that's why pointer to pointer can work
//		here

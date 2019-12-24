package main

//Given a binary search tree (BST) with duplicates, find all the mode(s) (the most frequently occurred element) in the given BST.
//
//Assume a BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than or equal to the node's key.
//The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
//Both the left and right subtrees must also be binary search trees.
//
//
//For example:
//Given BST [1,null,2,2],
//
//   1
//    \
//     2
//    /
//   2
//
//
//return [2].
//
//Note: If a tree has more than one mode, you can return them in any order.
//
//Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type result struct {
	data []int
}

// BST with traversal of left-middle-right will have in-order output, so that it is possible to calculate max occurrence
func findMode(root *TreeNode) []int {
	r := &result{data: make([]int, 0)}

	if root == nil {
		return r.data
	}

	var max, prev, freq int

	traverse(root, &max, &prev, &freq, r)

	// for last element not processed
	if freq == max {
		r.data = append(r.data, prev)
	}

	if freq > max {
		r.data = []int{prev}
	}

	return r.data
}

func traverse(node *TreeNode, max, prev, freq *int, r *result) {
	if node == nil {
		return
	}

	traverse(node.Left, max, prev, freq, r)

	if *freq == 0 {
		*prev = node.Val
		*freq++
	} else if node.Val == *prev {
		*freq++
	} else if node.Val != *prev {
		if *freq > *max {
			*max = *freq
			r.data = make([]int, 0)
			r.data = append(r.data, *prev)
		} else if *freq == *max {
			r.data = append(r.data, *prev)
		}

		*freq = 1
		*prev = node.Val
	}

	traverse(node.Right, max, prev, freq, r)
}

// problems
// 1. not understanding language well, assign slice is a new one instead of reusing existing
// 2. condition check not correct, if freq > max, it will setup max, cause next if statement to be true
// 3. forget a condition that freq < max, which will cause error

package main

//Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Given the below binary tree and sum = 22,
//
//      5
//     / \
//    4   8
//   /   / \
//  11  13  4
// /  \    / \
//7    2  5   1
//
//Return:
//
//[
//   [5,4,11,2],
//   [5,8,4,5]
//]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	// traverse from root to node, check if sum is equal
	if root == nil {
		return [][]int{}
	}

	return traverse(root, sum, 0, 0)
}

func traverse(node *TreeNode, sum, current, level int) [][]int {
	// leaf
	tmp := current + node.Val
	if node.Left == nil && node.Right == nil {
		if tmp == sum {
			s := make([]int, level+1)
			s[level] = node.Val
			return [][]int{s}
		}
		return [][]int{}
	}

	// continue to reach leaf
	var left, right [][]int
	if node.Left != nil {
		left = traverse(node.Left, sum, tmp, level+1)
	}

	if node.Right != nil {
		right = traverse(node.Right, sum, tmp, level+1)
	}

	result := make([][]int, 0)

	if len(left) != 0 {
		for _, l := range left {
			l[level] = node.Val
			result = append(result, l)
		}
	}

	if len(right) != 0 {
		for _, r := range right {
			r[level] = node.Val
			result = append(result, r)
		}
	}

	return result
}

// problems
// 1. for any leaf node, it will duplicate correct result, since left and right will both return one
// 2. when using append, even if it's empty, still counts
// 3. wrong typo, append node.Val twice
// 4. when slice capacity is enough, it will reuse existing one, so that influence other branch result
// 5. use too much memory, since sum doesn't need to memorize all elements
// 6. still forget that slice might influence other branch
// 7. it needs to be elements order must be in root to leaf
// 8. forget adding to new slice
// 9. to reduce allocating memory, pre-allocate slice

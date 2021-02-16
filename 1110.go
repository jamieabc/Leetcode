package main

// Given the root of a binary tree, each node in the tree has a distinct value.
//
// After deleting all nodes with a value in to_delete, we are left with a forest (a disjoint union of trees).
//
// Return the roots of the trees in the remaining forest.  You may return the result in any order.
//
//
//
// Example 1:
//
//
//
// Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
// Output: [[1,2,null,4],[6],[7]]
//
//
// Constraints:
//
// The number of nodes in the given tree is at most 1000.
// Each node has a distinct value between 1 and 1000.
// to_delete.length <= 1000
// to_delete contains distinct values between 1 and 1000.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	table := make(map[int]bool)
	for _, i := range to_delete {
		table[i] = true
	}

	ans := make([]*TreeNode, 0)

	postOrder(root, table, &ans)

	// becareful about root, because nodes put into answer as long as parent
	// is removed, bur root has no parent
	if !table[root.Val] {
		ans = append(ans, root)
	}

	return ans
}

func postOrder(node *TreeNode, table map[int]bool, ans *[]*TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	// need to update tree when node is deleted
	node.Left, node.Right = postOrder(node.Left, table, ans), postOrder(node.Right, table, ans)

	if table[node.Val] {
		if node.Left != nil {
			*ans = append(*ans, node.Left)
		}

		if node.Right != nil {
			*ans = append(*ans, node.Right)
		}

		return nil
	}

	return node
}

func delNodes1(root *TreeNode, to_delete []int) []*TreeNode {
	result := make([]*TreeNode, 0)
	mapping := make(map[int]bool)
	for _, i := range to_delete {
		mapping[i] = true
	}

	removed := dfs(root, mapping, &result)

	if removed == 0 {
		result = append(result, root)
	}

	return result
}

// -1 means this node needs to be removed
func dfs(node *TreeNode, deletes map[int]bool, result *[]*TreeNode) int {
	if node == nil {
		return 0
	}

	if l := dfs(node.Left, deletes, result); l == -1 {
		node.Left = nil
	}

	if r := dfs(node.Right, deletes, result); r == -1 {
		node.Right = nil
	}

	if _, ok := deletes[node.Val]; ok {
		if node.Left != nil {
			*result = append(*result, node.Left)
		}

		if node.Right != nil {
			*result = append(*result, node.Right)
		}

		return -1
	}

	return 0
}

//  Notes
//  1.  be careful about root, it root is not removed, then root itself needs
//      to be put into answer

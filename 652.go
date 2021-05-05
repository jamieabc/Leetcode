package main

import (
	"fmt"
	"strconv"
)

// Given the root of a binary tree, return all duplicate subtrees.
//
// For each kind of duplicate subtrees, you only need to return the root node of any one of them.
//
// Two trees are duplicate if they have the same structure with the same node values.
//
//
//
// Example 1:
//
// Input: root = [1,2,3,4,null,2,4,null,null,4]
// Output: [[2,4],[4]]
//
// Example 2:
//
// Input: root = [2,1,1]
// Output: [[1]]
//
// Example 3:
//
// Input: root = [2,2,2,3,null,3,null]
// Output: [[2,3],[3]]
//
//
//
// Constraints:
//
// The number of the nodes in the tree will be in the range [1, 10^4]
// -200 <= Node.val <= 200

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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	id := 1
	uuidTable := make(map[string]int)
	counter := make(map[int]int)

	duplicates := make([]*TreeNode, 0)

	dfs(root, uuidTable, counter, &id, &duplicates)

	return duplicates
}

func dfs(node *TreeNode, uuidTable map[string]int, counter map[int]int, id *int, duplicates *[]*TreeNode) string {
	if node == nil {
		return "0"
	}

	left := dfs(node.Left, uuidTable, counter, id, duplicates)
	right := dfs(node.Right, uuidTable, counter, id, duplicates)

	current := left + "," + strconv.Itoa(node.Val) + "," + right

	var uuid int

	if val, ok := uuidTable[current]; !ok {
		uuid = *id
		uuidTable[current] = *id
		*id++
	} else {
		uuid = val
	}

	counter[uuid]++

	if counter[uuid] == 2 {
		*duplicates = append(*duplicates, node)
	}

	return strconv.Itoa(uuid)
}

func findDuplicateSubtrees1(root *TreeNode) []*TreeNode {
	table := make(map[string]int)

	duplicates := make([]*TreeNode, 0)

	dfs1(root, table, &duplicates)

	return duplicates
}

func dfs1(node *TreeNode, table map[string]int, duplicates *[]*TreeNode) string {
	if node == nil {
		return "-"
	}

	left := dfs1(node.Left, table, duplicates)

	if node.Left != nil && table[left] == 1 {
		*duplicates = append(*duplicates, node.Left)
	}
	table[left]++

	right := dfs1(node.Right, table, duplicates)

	if node.Right != nil && table[right] == 1 {
		*duplicates = append(*duplicates, node.Right)
	}
	table[right]++

	path := fmt.Sprintf("%s,%s,%s", strconv.Itoa(node.Val), left, right)

	return path
}

//	Notes
//	1.	some understanding of tree traversal
//		in-order: LNR, also know subtree structure
//		post-order: LRN, starts from leaves, know subtree structure
//		pre-order: NLR, starts from root, know level/depth

//		for this problem to know subtree, need in-order or post-order traversal,
//		but it wants to find most common node, so post-order is suitable

//	2.	to identify a tree, need a way to represent its structure => string with
//		special character to denote terminator

//	3.	when left subtree is traversed, also right subtree is unknown, still need
//		to add left subtree to table, such than when similar subtree exists in
//		right can be identified

//	4.	delay add string to table to parent node, because only parent node can
//		know if left subtree + right subtree + parent can form a more common
//		structure

//	5.	i over complicate the problem, shows all duplicate nodes

//	6.	inspired from solution, use map[string]int to store occurrence count
//		use map[string]*TreeNode to store duplicates

//		since any duplicate node is okay, just latter occurrence node overwrites
//		previous node

//	7.	inspired from https://leetcode.com/problems/find-duplicate-subtrees/discuss/106011/Java-Concise-Postorder-Traversal-Solution

//		author provides clear explanation

//		also, string manipulation takes additional O(n) to process, author provides
//		a very good way to generate unique id for each node

//		the explanation comes from https://leetcode.com/problems/find-duplicate-subtrees/discuss/106022/No-string-hash-Python-code-O(n)-time-and-space/108522

//		the goal is to make each subtree with signature, start from a simple tree

//				1
//			  /   \
//			2      3
//		   / 	  / \
//		  4		 2	 5
//				/
//			   4

//		initially, id is 1

//		for the left most node 4: signature (0, 4, 0) never exist, store it (0, 4, 0): 1,
//			and id increment to 2

//		for the left most node 2: signature (1, 2, 0) never exist, store it (1, 2, 0): 2,
//			and id increment to 3

//		for middle node 4: signature (0, 4, 0) already exist, in order to detect duplicate
//			sub-tree, use previous one (1), id is still 3

//		for middle node 2: signature (1, 2, 0) already exist,, to detect duplicate sub-tree,
//			use previous one (2), id is still 3

//		for right most node 5: signature (0, 5, 0) never exist, store it (0, 5, 0): 3,
//			and id increment to 4

//		in this way, never exist tree will have a uniq id, already exist tree shares
//		same id, so if there's a counter to store existence of tree structure, then
//		it's

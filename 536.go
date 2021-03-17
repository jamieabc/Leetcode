package main

// You need to construct a binary tree from a string consisting of parenthesis and integers.
//
// The whole input represents a binary tree. It contains an integer followed by zero, one or two pairs of parenthesis. The integer represents the root's value and a pair of parenthesis contains a child binary tree with the same structure.
//
// You always start to construct the left child node of the parent first if it exists.
//
//
//
// Example 1:
//
// Input: s = "4(2(3)(1))(6(5))"
// Output: [4,2,6,3,1,5]
//
// Example 2:
//
// Input: s = "4(2(3)(1))(6(5)(7))"
// Output: [4,2,6,3,1,5,7]
//
// Example 3:
//
// Input: s = "-4(2(3)(1))(6(5)(7))"
// Output: [-4,2,6,3,1,5,7]
//
//
//
// Constraints:
//
// 0 <= s.length <= 3 * 104
// s consists of digits, '(', ')', and '-' only.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func str2tree(s string) *TreeNode {
	var idx, val int
	mul := 1

	size := len(s)
	if size == 0 {
		return nil
	}

	if s[0] == '-' {
		mul = -1
		idx = 1
	}

	for idx < size && s[idx] != '(' {
		val *= 10
		val += int(s[idx] - '0')
		idx++
	}

	root := &TreeNode{
		Val: val * mul,
	}
	mul, val = 1, 0

	stack := []*TreeNode{root}

	for ; idx < size; idx++ {
		if s[idx] == '(' {
			idx++

			for ; idx < size && s[idx] != '(' && s[idx] != ')'; idx++ {
				if s[idx] == '-' {
					mul = -1
				} else {
					val *= 10
					val += int(s[idx] - '0')
				}
			}

			node := &TreeNode{
				Val: mul * val,
			}
			mul, val = 1, 0

			if stack[len(stack)-1].Left == nil {
				stack[len(stack)-1].Left = node
			} else {
				stack[len(stack)-1].Right = node
			}

			stack = append(stack, node)
			idx--
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return root
}

func str2tree1(s string) *TreeNode {
	size := len(s)
	if size == 0 {
		return nil
	}
	var idx int

	return recursive(s, &idx)
}

func recursive(s string, idx *int) *TreeNode {
	var val int
	mul := 1

	if s[*idx] == '-' {
		mul = -1
		*idx++
	}

	for ; *idx < len(s) && s[*idx] != '(' && s[*idx] != ')'; *idx++ {
		val *= 10
		val += int(s[*idx] - '0')
	}

	node := &TreeNode{
		Val: val * mul,
	}

	if *idx < len(s) && s[*idx] == '(' {
		*idx++
		node.Left = recursive(s, idx)
	}

	if *idx < len(s) && s[*idx] == '(' {
		*idx++
		node.Right = recursive(s, idx)
	}

	*idx++

	return node
}

//	Notes
//	1.	could be negative number...

//	2.	becareful boundary condition, no parenthesis

//	3.	inspired from https://leetcode.com/problems/construct-binary-tree-from-string/discuss/100422/Python-Straightforward-with-Explanation

//		awice has a very clean code, although it's a little waste of
//		computation

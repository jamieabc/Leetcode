package main

// Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.
//
// Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.
//
// Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.
//
//
//
// Example 1:
//
//
// Input: root = [1,2,3,null,null,4,5]
// Output: [1,2,3,null,null,4,5]
// Example 2:
//
// Input: root = []
// Output: []
// Example 3:
//
// Input: root = [1]
// Output: [1]
// Example 4:
//
// Input: root = [1,2]
// Output: [1,2]
//
//
// Constraints:
//
// The number of nodes in the tree is in the range [0, 104].
// -1000 <= Node.val <= 1000

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	ans := make([]byte, 0)
	table := make(map[*TreeNode]bool)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		n := stack[len(stack)-1]

		if _, ok := table[n]; !ok {
			ans = append(ans, '(')
			ans = append(ans, []byte(strconv.Itoa(n.Val))...)
			table[n] = true
		} else {
			ans = append(ans, ')')
			stack = stack[:len(stack)-1]
			continue
		}

		if n.Right != nil {
			stack = append(stack, n.Right)
		}

		if n.Left != nil {
			stack = append(stack, n.Left)
		} else if n.Right != nil {
			ans = append(ans, '(', ')')
		}
	}

	return string(ans[1 : len(ans)-1])
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	var idx int
	size := len(data)
	for ; idx < size && data[idx] != '('; idx++ {
	}

	num, _ := strconv.Atoi(data[:idx])
	root := &TreeNode{
		Val: num,
	}

	// 1(2)(3(4)(5))
	// 1(2()(4))
	stack := []*TreeNode{root}
	var to int
	var emptyLeft bool
	for idx < size {
		prev := stack[len(stack)-1]

		if data[idx] == ')' {
			stack = stack[:len(stack)-1]
			idx++
			continue
		}

		if data[idx] == '(' && data[idx+1] == ')' {
			idx += 2
			emptyLeft = true
		}

		for to = idx + 1; to < size && data[to] != '(' && data[to] != ')'; to++ {
		}

		num, _ := strconv.Atoi(data[idx+1 : to])
		node := &TreeNode{
			Val: num,
		}

		if prev.Left == nil && !emptyLeft {
			prev.Left = node
		} else {
			prev.Right = node
			emptyLeft = false
		}
		stack = append(stack, node)
		idx = to
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

//	Notes
//	1.	borrow idea from others

//	2.	inspired from solution, could be as simple as empty node denote as 'null'
//		when deserialize, split by ',' and can be processed like pre-order

package main

// Design an algorithm to encode an N-ary tree into a binary tree and decode the binary tree to get the original N-ary tree. An N-ary tree is a rooted tree in which each node has no more than N children. Similarly, a binary tree is a rooted tree in which each node has no more than 2 children. There is no restriction on how your encode/decode algorithm should work. You just need to ensure that an N-ary tree can be encoded to a binary tree and this binary tree can be decoded to the original N-nary tree structure.
//
// Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See following example).
//
// For example, you may encode the following 3-ary tree to a binary tree in this way:
//
//
//
// Input: root = [1,null,3,2,4,null,5,6]
// Note that the above is just an example which might or might not work. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.
//
//
//
// Constraints:
//
// The height of the n-ary tree is less than or equal to 1000
// The total number of nodes is between [0, 10^4]
// Do not use class member/global/static variables to store states. Your encode and decode algorithms should be stateless.

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
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

type MappingNodes struct {
	node  *Node
	tree  *TreeNode
	level int
}

type Codec struct {
	tree  *TreeNode
	node  *Node
	queue []MappingNodes
}

func Constructor() *Codec {
	return &Codec{
		queue: make([]MappingNodes, 0),
	}
}

// node:
//			1
//		2   3   4
//	   5 6

// tree:
//			1
//		2		1
//	  5   2   3   1
//	     6 2     4 1
//	left child is child, and right child are level then left-child's sibling
// 	if it's children, put into left child
// 	there's dummy node for level, but I didn't take time to modify it
func (this *Codec) encode(root *Node) *TreeNode {
	if root == nil {
		return nil
	}

	this.node = root
	this.tree = &TreeNode{
		Val: root.Val,
	}
	this.queue = append(this.queue, MappingNodes{
		node:  root,
		tree:  this.tree,
		level: 0,
	})
	for len(this.queue) != 0 {
		this.build()
	}

	return this.tree
}

func (this *Codec) build() {
	if len(this.queue) == 0 {
		return
	}

	m := this.queue[0]
	this.queue = this.queue[1:]

	for _, n := range m.node.Children {
		m.tree.Left = &TreeNode{
			Val: n.Val,
		}

		m.tree.Right = &TreeNode{
			Val: m.level + 1,
		}

		if len(n.Children) != 0 {
			this.queue = append(this.queue, MappingNodes{
				node:  n,
				tree:  m.tree.Left,
				level: m.level + 1,
			})
		}

		m.tree = m.tree.Right
	}
}

func (this *Codec) decode(root *TreeNode) *Node {
	if root == nil {
		return nil
	}

	this.node = &Node{
		Val:      root.Val,
		Children: make([]*Node, 0),
	}

	this.queue = append(this.queue, MappingNodes{
		tree:  root,
		node:  this.node,
		level: 0,
	})

	for len(this.queue) != 0 {
		this.rebuild()
	}

	return this.node
}

func (this *Codec) rebuild() {
	if len(this.queue) == 0 {
		return
	}

	m := this.queue[0]
	this.queue = this.queue[1:]

	n := m.tree

	for n.Right != nil {
		newNode := &Node{
			Val:      n.Left.Val,
			Children: make([]*Node, 0),
		}

		m.node.Children = append(m.node.Children, newNode)

		if n.Left.Left != nil {
			this.queue = append(this.queue, MappingNodes{
				node: newNode,
				tree: n.Left,
			})
		}

		n = n.Right
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * bst := obj.encode(root);
 * ans := obj.decode(bst);
 */

//	problems
//	1. inspired from https://leetcode.com/problems/encode-n-ary-tree-to-binary-tree/discuss/160687/Super-Simple-Java-Beats-99-2ms-with-clear-Explanation

//		siblings all at right child

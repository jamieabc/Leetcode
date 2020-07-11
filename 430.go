package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	dfs(root)
	return root
}

func dfs(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Child != nil {
		next := node.Next

		node.Next = node.Child
		node.Child.Prev = node
		node.Child = nil

		tail := dfs(node.Next)
		tail.Next = next

		if next != nil {
			next.Prev = tail
		}

		node = tail
	}

	if node.Next == nil {
		return node
	}

	return dfs(node.Next)
}

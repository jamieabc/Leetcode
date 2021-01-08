package main

// You are given all the nodes of an N-ary tree as an array of Node objects, where each node has a unique value.
//
// Return the root of the N-ary tree.
//
// Custom testing:
//
// An N-ary tree can be serialized as represented in its level order traversal where each group of children is separated by the null value (see examples).
//
// For example, the above tree is serialized as [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14].
//
// The testing will be done in the following way:
//
//     The input data should be provided as a serialization of the tree.
//     The driver code will construct the tree from the serialized input data and put each Node object into an array in an arbitrary order.
//     The driver code will pass the array to findRoot, and your function should find and return the root Node object in the array.
//     The driver code will take the returned Node object and serialize it. If the serialized value and the input data are the same, the test passes.
//
//
//
// Example 1:
//
// Input: tree = [1,null,3,2,4,null,5,6]
// Output: [1,null,3,2,4,null,5,6]
// Explanation: The tree from the input data is shown above.
// The driver code creates the tree and gives findRoot the Node objects in an arbitrary order.
// For example, the passed array could be [Node(5),Node(4),Node(3),Node(6),Node(2),Node(1)] or [Node(2),Node(6),Node(1),Node(3),Node(5),Node(4)].
// The findRoot function should return the root Node(1), and the driver code will serialize it and compare with the input data.
// The input data and serialized Node(1) are the same, so the test passes.
//
// Example 2:
//
// Input: tree = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
//
//
//
// Constraints:
//
//     The total number of nodes is between [1, 5 * 104].
//     Each node has a unique value.

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func findRoot(tree []*Node) *Node {
	var num int

	for _, i := range tree {
		num ^= i.Val

		for _, j := range i.Children {
			num ^= j.Val
		}
	}

	for i := range tree {
		if tree[i].Val == num {
			return tree[i]
		}
	}

	return nil
}

func findRoot1(tree []*Node) *Node {
	visited := make(map[int]bool)

	for i := range tree {
		dfs(tree[i].Children, visited)
	}

	for i := range tree {
		if _, ok := visited[tree[i].Val]; !ok {
			return tree[i]
		}
	}

	return nil
}

func dfs(tree []*Node, visited map[int]bool) {
	for i := range tree {
		if _, ok := visited[tree[i].Val]; ok {
			continue
		}
		visited[tree[i].Val] = true
		dfs(tree[i].Children, visited)
	}
}

//	Notes
//	1.	inspired from https://leetcode.com/explore/challenge/card/january-leetcoding-challenge-2021/580/week-2-january-8th-january-14th/3596/discuss/726453/Java-O(n)-time-with-O(n)-space-and-O(1)-space-follow-up

//		since every node appears in array, and all nodes except root will be in
//		some node's children, which means every node except root will be visited
//		twice

//		plus self value and minus children value, eventually gets result

//	2.	inspired from https://leetcode.com/problems/find-root-of-n-ary-tree/discuss/726536/c%2B%2B-bit-operation-1-pass-with-O(1)-space-solution

//		since every node except root appears twice, use XOR to find

//	3.	most key point to find easiest solution is to realize that all nodes except
//		root appears twice

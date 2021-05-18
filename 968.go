package main

//Given a binary tree, we install cameras on the nodes of the tree.
//
//Each camera at a node can monitor its parent, itself, and its immediate children.
//
//Calculate the minimum number of cameras needed to monitor all nodes of the tree.
//
//
//
//Example 1:
//
//
//Input: [0,0,null,0,0]
//Output: 1
//Explanation: One camera is enough to monitor all nodes if placed as shown.
//Example 2:
//
//
//Input: [0,0,null,0,null,0,null,null,0]
//Output: 2
//Explanation: At least two cameras are needed to monitor all nodes of the tree. The above image shows one of the valid configurations of camera placement.
//
//Note:
//
//The number of nodes in the given tree will be in the range [1, 1000].
//Every node has value 0.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	var count int

	val := dfs(root, &count)

	if val == -1 {
		return count + 1
	}

	return count
}

func dfs(node *TreeNode, count *int) int {
	if node == nil {
		return 0
	}

	l, r := dfs(node.Left, count), dfs(node.Right, count)

	if l == -1 || r == -1 {
		*count++
		return 1
	}

	return max(l, r) - 1
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	becareful about problem, self & children & parent only need one

//	2.	the point is to distinguish leaf node & non-leaf node with one side empty

//		empty node return -1 will cause a problem

//		e.g.			o
//				  	  /
//			 	 	o	<- left return 0, use this point, but it's un-necessary
//			   	  /
//			 	o		<- left child -2, right child -1, max is -1, return max(-1, -2)+1 = 0
//		   	  /
//		 	o			<- use this point, return -2 to parent
//		   /
//		 o

//		for a node with one side of empty cell, should return a value that will
//		be overwritten by the other side, use max so empty side should return -2

//		the only exception is that if both sides are -2, then return 0, because
//		it mean's it's parent should be marked

//	3.	root node is also an exception case, the reason return -1 no need to consider
//		is because it assumes current node will have parent, and use parent will
//		cover current node, and this assumption fails for root node because there's
//		no parent

//	4.	after re-think the problem, every node has 3 states: self covered by child,
//		child covered, child not covered (我靠別人罩，別人罩不了我，我要罩別人)

//		for recursion return back, each state degrade from self can cover other ->
//		child covered -> child not covered

//		for the empty node case, it's the same as child covered because no need to
//		consider empty child

//		whe child is not covered, the current node must be selected

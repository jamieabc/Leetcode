package main

// Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.
//
// Calling next() will return the next smallest number in the BST.
//
//
//
// Example:
//
// BSTIterator iterator = new BSTIterator(root);
// iterator.next();    // return 3
// iterator.next();    // return 7
// iterator.hasNext(); // return true
// iterator.next();    // return 9
// iterator.hasNext(); // return true
// iterator.next();    // return 15
// iterator.hasNext(); // return true
// iterator.next();    // return 20
// iterator.hasNext(); // return false
//
//
//
// Note:
//
//     next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
//     You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{
		stack: make([]*TreeNode, 0),
	}

	// loop until left most node
	for node := root; node != nil; node = node.Left {
		b.stack = append(b.stack, node)
	}

	return b
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	size := len(this.stack)
	n := this.stack[size-1]
	this.stack = this.stack[:size-1]

	for node := n.Right; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}

	return n.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

//	problems
//	1.	when initialising, I cannot assume root is non-nil

//	2.	inspired from https://leetcode.com/problems/binary-search-tree-iterator/discuss/52525/My-solutions-in-3-languages-with-Stack

//		initial pushing should check if node != nil, then loop to next left

//		this is a much more elegant way, since it's because BST traverse is
//		actually L-N-R (in-order), so the procedure is quite simple, if left
//		child exist, keep pushing until left child is empty. pop self. if
//		right child exist, put into stack, then put all left child into stack

//	3.	this is a good reference https://leetcode.com/problems/binary-search-tree-iterator/discuss/52647/Nice-Comparison-(and-short-Solution)

//		compares to traditional in-order iterative:

// 		TreeNode visit = root;
// 		Stack<TreeNode> stack = new Stack();
// 		while (visit != null || !stack.empty()) {
// 	    	while (visit != null) {
// 				stack.push(visit);
// 				visit = visit.left;
// 			}
// 			TreeNode next = stack.pop();
// 			visit = next.right;
// 			doSomethingWith(next.val);
// 		}

//		but it comes to me the other question, how to write post-order
//		traversal in iterative?

//		added 2021/1/30, put nodes in reverse order of L-R-N, which is N-R-L

//	4.	inspired from https://leetcode.com/problems/binary-search-tree-iterator/discuss/52584/My-java-accepted-solution

//		this is the transformation from traditional in-order traversal

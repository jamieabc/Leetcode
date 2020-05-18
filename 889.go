package main

// Return any binary tree that matches the given preorder and postorder traversals.
//
// Values in the traversals pre and post are distinct positive integers.
//
//
//
// Example 1:
//
// Input: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
// Output: [1,2,3,4,5,6,7]
//
//
//
// Note:
//
//    1 <= pre.length == post.length <= 30
//    pre[] and post[] are both permutations of 1, 2, ..., pre.length.
//    It is guaranteed an answer exists. If there exists multiple answers, you can return any of them.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
	length := len(pre)
	root := &TreeNode{
		Val: pre[0],
	}

	traverse(pre[1:], post[:length-1], root)

	return root
}

func traverse(pre, post []int, parent *TreeNode) {
	length := len(pre)

	if length == 0 {
		return
	}

	parent.Left = &TreeNode{
		Val: pre[0],
	}

	var leftEnd int
	for i := length - 1; i >= 0; i-- {
		if post[i] == pre[0] {
			leftEnd = i
			break
		}
	}

	traverse(pre[1:leftEnd+1], post[:leftEnd], parent.Left)

	if leftEnd < length-1 {
		// with right child
		parent.Right = &TreeNode{
			Val: pre[leftEnd+1],
		}
		traverse(pre[leftEnd+2:], post[leftEnd+1:length-1], parent.Right)
	}
}

func constructFromPrePost2(pre []int, post []int) *TreeNode {
	preLength := len(pre)
	if preLength == 0 {
		return nil
	}

	node := &TreeNode{
		Val: pre[0],
	}

	if preLength == 1 {
		return node
	}

	var leftLength int

	// find length of left subtree
	for i := 1; i < preLength; i++ {
		if pre[i] == post[preLength-2] {
			leftLength = i - 1
			break
		}
	}

	node.Left = constructFromPrePost(pre[1:1+leftLength], post[:leftLength])
	node.Right = constructFromPrePost(pre[1+leftLength:], post[leftLength:preLength-1])

	return node
}

// reference: https://www.youtube.com/watch?v=53aOi0Drp9I
// problems
// 1. cannot assume when length <= 3, it must be N-L-R, it could also be right is empty, when length is 3, possible situations are:
//     N             N
//   L   R         L
//               R
// 2. avoid slice index out of range, should use post length to decide
// 3. if some part is empty, just ignore it instead of using it
// 4. I kind of think it wrong, the condition should be if left of pre == right of post
// 5. optimize, length of pre & post are corresponded, so there's no need to calculate right index at post, just get length of left subtree, remaining elements are right subtree
// 6. optimize, at it comes to situation of one child is empty, no need to distinguish, it will put all nodes into right
// 7. optimize, no need to check when length is 2, because alogrithm will put all into right side

//	8. 	rewrite, forget about situation when length == 2, still with 2
//		possibilities

//	9.	reference from https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/discuss/161268/C%2B%2BJavaPython-One-Pass-Real-O(N)

//		this is really clever solution, I don't know how author comes up.
//		for pre-order, it travers left most node by lowest height to highest
//		height. post-order traverse directly to left-most last node

//		with this property, keep build tree using pre-order. when pre-order
//		number equals post-order means that sub tree is all built, then
//		starts to build right child tree

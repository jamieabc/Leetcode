package main

//Return any binary tree that matches the given preorder and postorder traversals.
//
//Values in the traversals pre and post are distinct positive integers.
//
//
//
//Example 1:
//
//Input: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
//Output: [1,2,3,4,5,6,7]
//
//
//
//Note:
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

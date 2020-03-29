package main

//A full binary tree is a binary tree where each node has exactly 0 or 2 children.
//
//Return a list of all possible full binary trees with N nodes.  Each element of the answer is the root node of one possible tree.
//
//Each node of each tree in the answer must have node.val = 0.
//
//You may return the final list of trees in any order.
//
//
//
//Example 1:
//
//Input: 7
//Output: [[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
//Explanation:
//
//
//
//Note:
//
//1 <= N <= 20

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(N int) []*TreeNode {
	if N == 0 || N%2 == 0 {
		return []*TreeNode{}
	}

	// initialise
	result := make([][]*TreeNode, N+1)
	for i := 1; i <= N; i += 2 {
		result[i] = make([]*TreeNode, 0)
	}

	// if N is 1, only 1 FBT
	result[1] = []*TreeNode{&TreeNode{}}

	if N == 1 {
		return result[1]
	}

	threeNode := &TreeNode{}
	threeNode.Left = &TreeNode{}
	threeNode.Right = &TreeNode{}
	result[3] = []*TreeNode{threeNode}

	// if N is 3, only 1 FBT
	if N == 3 {
		return result[3]
	}

	// dp - concat calculated tree to generate new tree
	for i := 5; i <= N; i += 2 {

		// every full binary tree can be composed of left - root - right
		// left & right can be constructed by already calculated trees
		// 5 = 1 + 1 + 3
		// 5 = 3 + 1 + 1
		for j := 1; j < i-1; j += 2 {
			remain := i - 1 - j

			// append existing results
			for _, k := range result[j] {
				for _, l := range result[remain] {
					newRoot := &TreeNode{}
					newRoot.Left = k
					newRoot.Right = l
					result[i] = append(result[i], newRoot)
				}
			}
		}
	}

	return result[N]
}

//	optimize
//	1.	takes too long time, I think most time spends on copy, I should remove
//		those copies

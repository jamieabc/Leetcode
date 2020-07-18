package main

// Given a binary tree where each path going from the root to any leaf form a valid sequence, check if a given string is a valid sequence in such binary tree.
//
// We get the given string from the concatenation of an array of integers arr and the concatenation of all values of the nodes along a path results in a sequence in the given binary tree.
//
//
//
// Example 1:
//
// Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,1,0,1]
// Output: true
// Explanation:
// The path 0 -> 1 -> 0 -> 1 is a valid sequence (green color in the figure).
// Other valid sequences are:
// 0 -> 1 -> 1 -> 0
// 0 -> 0 -> 0
//
// Example 2:
//
// Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,0,1]
// Output: false
// Explanation: The path 0 -> 0 -> 1 does not exist, therefore it is not even a sequence.
//
// Example 3:
//
// Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,1,1]
// Output: false
// Explanation: The path 0 -> 1 -> 1 is a sequence, but it is not a valid sequence.
//
//
//
// Constraints:
//
//     1 <= arr.length <= 5000
//     0 <= arr[i] <= 9
//     Each node's value is between [0 - 9].

func isValidSequence(root *TreeNode, arr []int) bool {
	return dfs(root, arr)
}

func dfs(node *TreeNode, arr []int) bool {
	if node == nil || len(arr) == 0 {
		return false
	}

	if node.Val != arr[0] {
		return false
	}

	if len(arr) == 1 {
		return node.Left == nil && node.Right == nil && node.Val == arr[0]
	}

	return dfs(node.Left, arr[1:]) || dfs(node.Right, arr[1:])
}

//	problems
//	1.	false could be two conditions:
//		- arr non-empty, node is empty
//		- arr is empty, node is non-empty

//	2.	to leaf, don't forget that

//	3.	inspired form sample code, instead of check for multiple condition,
//		author provides a clever of writing it:

//    	if node.Left == nil && node.Right == nil && len(arr) == 1 {
//			return node.Val == arr[0]
//		}

// 		above writing makes code going to another level deep, below writing
// 		stops program

//  	if len(arr) == 1 {
// 			return node.Left == nil && node.Right == nil && node.Val == arr[0]
// 		}

//	4.	the other way is to pass index, instead of passing smaller array

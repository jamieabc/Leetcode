package main

// Given a root of an N-ary tree, you need to compute the length of the diameter of the tree.
//
// The diameter of an N-ary tree is the length of the longest path between any two nodes in the tree. This path may or may not pass through the root.
//
// (Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value.)
//
//
//
// Example 1:
//
//
//
// Input: root = [1,null,3,2,4,null,5,6]
// Output: 3
// Explanation: Diameter is shown in red color.
// Example 2:
//
//
//
// Input: root = [1,null,2,null,3,4,null,5,null,6]
// Output: 4
// Example 3:
//
//
//
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: 7
//
//
// Constraints:
//
// The depth of the n-ary tree is less than or equal to 1000.
// The total number of nodes is between [0, 10^4].

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func diameter(root *Node) int {
	if root == nil {
		return 0
	}

	_, longest := recursive(root)

	return longest
}

func recursive(node *Node) (int, int) {
	if len(node.Children) == 0 {
		return 0, 0
	}

	var max1, max2, longest int

	for i := range node.Children {
		length, longestFromChild := recursive(node.Children[i])
		length++
		longest = max(longest, longestFromChild)

		if length > max1 {
			max1, max2 = length, max1
		} else if length > max2 {
			max2 = length
		}
	}

	return max1, max(max1+max2, longest)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	this problem is not hard, why didn't I solve it during contest?

//	2.	inspired from solution, two wordings need to be clarified

//		height of a node: longest length for a node to leaf
//		depth of a node: length from root to node

//				1
//			   /
//			  2                <-   height of node-2: 2 (2 -> 3 -> 5)
//			/   \                   depth of node-2: 1
//		   3    4
//        /
//       5

//	3.	inspired from solution, calculate longest distance for a node can also
//		obtained from depth

//		depth(node.leaf1) + depth(node.leaf2) - 2 * depth(node)

//		very brilliant solution

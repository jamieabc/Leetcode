package main

// Given a n * n matrix grid of 0's and 1's only. We want to represent the grid with a Quad-Tree.
//
// Return the root of the Quad-Tree representing the grid.
//
// Notice that you can assign the value of a node to True or False when isLeaf is False, and both are accepted in the answer.
//
// A Quad-Tree is a tree data structure in which each internal node has exactly four children. Besides, each node has two attributes:
//
// val: True if the node represents a grid of 1's or False if the node represents a grid of 0's.
// isLeaf: True if the node is leaf node on the tree or False if the node has the four children.
//
// class Node {
// public boolean val;
// public boolean isLeaf;
// public Node topLeft;
// public Node topRight;
// public Node bottomLeft;
// public Node bottomRight;
// }
//
// We can construct a Quad-Tree from a two-dimensional area using the following steps:
//
// If the current grid has the same value (i.e all 1's or all 0's) set isLeaf True and set val to the value of the grid and set the four children to Null and stop.
// If the current grid has different values, set isLeaf to False and set val to any value and divide the current grid into four sub-grids as shown in the photo.
// Recurse for each of the children with the proper sub-grid.
//
// If you want to know more about the Quad-Tree, you can refer to the wiki.
//
// Quad-Tree format:
//
// The output represents the serialized format of a Quad-Tree using level order traversal, where null signifies a path terminator where no node exists below.
//
// It is very similar to the serialization of the binary tree. The only difference is that the node is represented as a list [isLeaf, val].
//
// If the value of isLeaf or val is True we represent it as 1 in the list [isLeaf, val] and if the value of isLeaf or val is False we represent it as 0.
//
//
//
// Example 1:
//
// Input: grid = [[0,1],[1,0]]
// Output: [[0,1],[1,0],[1,1],[1,1],[1,0]]
// Explanation: The explanation of this example is shown below:
// Notice that 0 represnts False and 1 represents True in the photo representing the Quad-Tree.
//
// Example 2:
//
// Input: grid = [[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]
// Output: [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
// Explanation: All values in the grid are not the same. We divide the grid into four sub-grids.
// The topLeft, bottomLeft and bottomRight each has the same value.
// The topRight have different values so we divide it into 4 sub-grids where each has the same value.
// Explanation is shown in the photo below:
//
// Example 3:
//
// Input: grid = [[1,1],[1,1]]
// Output: [[1,1]]
//
// Example 4:
//
// Input: grid = [[0]]
// Output: [[1,0]]
//
// Example 5:
//
// Input: grid = [[1,1,0,0],[1,1,0,0],[0,0,1,1],[0,0,1,1]]
// Output: [[0,1],[1,1],[1,0],[1,0],[1,1]]
//
//
//
// Constraints:
//
// n == grid.length == grid[i].length
// n == 2^x where 0 <= x <= 6

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	return traverse(grid, len(grid), 0, 0)
}

func traverse(grid [][]int, length, x, y int) *Node {
	node := &Node{}

	if length == 1 {
		node.IsLeaf = true
		node.Val = grid[y][x] == 1

		return node
	}

	half := length / 2
	tl := traverse(grid, half, x, y)
	tr := traverse(grid, half, x+half, y)
	bl := traverse(grid, half, x, y+half)
	br := traverse(grid, half, x+half, y+half)

	if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf && tl.Val == tr.Val && tl.Val == bl.Val && tl.Val == br.Val {
		node.IsLeaf = true
		node.Val = tl.Val
	} else {
		node.TopLeft = tl
		node.TopRight = tr
		node.BottomLeft = bl
		node.BottomRight = br
	}

	return node
}

//	problems
//	1.	Too slow, I have many repeated calculations due to top-down
//		iteration

//	2.	reference from https://csc411.files.wordpress.com/2012/10/tutorial2.pdf

//		time complexity O(n^2)
//		the recursion formula is T(n) = 4T(n/2) + n

//		T(n) = 4T(n/2) + n			level sum: 4(n/2) = 2n
//		T(n/2) = 4T(n/4) + n		level sum: 4*4*(n/4) = 4n
//		T(n/4) = 4T(n/8) + n/4		level sum: 4*4*4*(n/8) = 8n

//		this video has a great explanation https://www.youtube.com/watch?v=1K9ebQJosvo

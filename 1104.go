package main

import "math"

// In an infinite binary tree where every node has two children, the nodes are labelled in row order.
//
// In the odd numbered rows (ie., the first, third, fifth,...), the labelling is left to right, while in the even numbered rows (second, fourth, sixth,...), the labelling is right to left.
//
// Given the label of a node in this tree, return the labels in the path from the root of the tree to the node with that label.
//
//
//
// Example 1:
//
// Input: label = 14
// Output: [1,3,4,14]
//
// Example 2:
//
// Input: label = 26
// Output: [1,2,6,10,26]
//
//
//
// Constraints:
//
//     1 <= label <= 10^6

func pathInZigZagTree(label int) []int {
	var level int
	for i := 2; i-1 < label; i = i << 1 {
		level++
	}
	ans := make([]int, level+1)
	ans[len(ans)-1] = label

	var idx, start int
	for ; level > 0; level-- {
		if level&1 > 0 {
			for start, idx = 1<<(level+1)-1, 0; start-idx != label; idx++ {
			}
			label = (start+1)/4 + idx/2
		} else {
			for start, idx = 1<<level, 0; start+idx != label; idx++ {
			}
			label = start - 1 - idx/2
		}

		ans[level-1] = label
	}

	return ans
}

func pathInZigZagTree1(label int) []int {
	level := 1
	var count int
	for count = 2; count-1 < label; count *= 2 {
		level++
	}

	result := make([]int, level)
	result[level-1] = label
	label = toNormal(level-1, label)
	label /= 2

	for i := level - 2; i >= 0; i-- {
		result[i] = toZigzag(i, label)
		label /= 2
	}

	return result
}

func toNormal(level, label int) int {
	if level%2 == 0 {
		return label
	}

	size := int(math.Pow(float64(2), float64(level)))
	return size - 1 + (size*2 - label)
}

func toZigzag(level, label int) int {
	if level%2 == 0 {
		return label
	}

	size := int(math.Pow(float64(2), float64(level)))
	return size*2 - 1 - (label - size)
}

//	Notes
//	1.	wrong checking for level, because label starts from 1, so when label
//		is power of 2, level is wrong

//	2.	inspired from https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/discuss/323293/C%2B%2B-O(log-n)
//		what I think is to find relative index from start of that level, and
//		find real zigzag value

//		author thinks differently, first if no zigzag, parent = current/2, so
//		he finds out the range of start & end of that level

//		there's slightly difference

//	3.	inspired from https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/discuss/324011/Python-O(logn)-time-and-space-with-readable-code-and-step-by-step-explanation

//		use original tree as basic, find a way to map zigzag to original

//             1
//           /   \
//         2       3
//       /  \     /  \
//     4     5   6     7
//   / |    /|   |\    | \
// 8   9  10 11 12 13  14  15

//             1
//           /   \
//         3       2  <- 3+2-3 = 2/2 = 1
//       /  \     /  \
//     4     5   6     7   <- 7+4-4 = 7/2 = 3
//   / |    /|   |\    | \
// 15 14  13 12 11 10  9  8   <- 15+8-14 = 9/2 = 4

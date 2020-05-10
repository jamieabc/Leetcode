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

//	problems
//	1.	wrong checking for level, because label starts from 1, so when label
//		is power of 2, level is wrong

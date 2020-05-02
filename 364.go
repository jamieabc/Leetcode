package main

import "fmt"

// Given a nested list of integers, return the sum of all integers in the list weighted by their depth.
//
// Each element is either an integer, or a list -- whose elements may also be integers or other lists.
//
// Different from the previous question where weight is increasing from root to leaf, now the weight is defined from bottom up. i.e., the leaf level integers have weight 1, and the root level integers have the largest weight.
//
// Example 1:
//
// Input: [[1,1],2,[1,1]]
// Output: 8
// Explanation: Four 1's at depth 1, one 2 at depth 2.
//
// Example 2:
//
// Input: [1,[4,[6]]]
// Output: 17
// Explanation: One 1 at depth 3, one 4 at depth 2, and one 6 at depth 1; 1*3 + 4*2 + 6*1 = 17.

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

func depthSumInverse(nestedList []*NestedInteger) int {
	var historySum, sum int
	var list *[]*NestedInteger
	list = &nestedList

	for len(*list) != 0 {
		tmp := make([]*NestedInteger, 0)
		for _, l := range *list {
			if l.IsInteger() {
				historySum += l.GetInteger()
			} else {
				tmp = append(tmp, l.GetList()...)
			}
		}
		sum += historySum
		list = &tmp
	}

	return sum
}

func depthSumInverse2(nestedList []*NestedInteger) int {
	rawSum, weighted, maxDepth := traverse(nestedList, 1)
	return rawSum*(maxDepth+1) - weighted
}

func traverse(list []*NestedInteger, level int) (int, int, int) {
	if len(list) == 0 {
		return 0, 0, 0
	}

	var rawSum, weighted int
	maxDepth := level

	for _, l := range list {
		if l.IsInteger() {
			rawSum += l.GetInteger()
			weighted += level * l.GetInteger()
		} else {
			flat, tmp, depth := traverse(l.GetList(), level+1)
			rawSum += flat
			weighted += tmp
			maxDepth = max(maxDepth, depth)
		}
	}

	return rawSum, weighted, maxDepth
}

//	problems
//	1.	This can be further improved by not saving traverse. Final result
//		only needs sum, so backtracking is not needed.

//		After seeing other's solution (https://leetcode.com/problems/nested-list-weight-sum-ii/discuss/114195/Java-one-pass-DFS-solution-mathematically), I found what I miss.

//		What is weight of a number in list? It's longest depth by child list.
//		A simple example as follows:
//			a
//		  /   \
//		 b     c
//			    \
//			     d

//		depth of a is 3 (a-c-d).

//		So it's dynamically calculated, for example of
//		[ [-1], [-2, [-3, [-4]]]
//		weight for every number -4: 1 z
//								-3: 2 y
//								-2: 3 x
//								-1: 3 w

//		But when processing -1, it can only know it's weight comes to 2 (
//		because it's from list). In brief, every element can only know it's
//		distance relative to root (not leaf)

//		Here comes the clever part, since longest depth can only be know
//		last, and for each element only know relative depth of a list.

//		I don't know how this come up, but it works. Consider previous
//		example, final result is 3w+2x+2y+z
//		it can be composed of 5(w+x+y+z) - (2w+2x+3y+4z)
//		be ware that coefficient of 2x+2x+3y+4z is distance from root,
//		which is easier to calculate.

//		I think it's somehow relates to program execution, -2 depth is 2,
//		-3 depth is 3, -4 depth is 4, the depth is known only when program
//		progress list, so this clever solution uses this property.

//	2. from another reference (https://leetcode.com/problems/nested-list-weight-sum-ii/discuss/83641/No-depth-variable-no-multiplication)

//		It uses a variable to store all previous sum from specific list,
//		the weighted means those sum are added multiple times. It's a BFS
//		algorithm.

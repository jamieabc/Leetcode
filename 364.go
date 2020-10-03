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
	var weighted, unweighted int
	queue := nestedList

	// similar to BFS, order of number doesn't matter, every time process
	// single number (same level in tree)
	for len(queue) > 0 {
		stop := len(queue)

		for i := 0; i < stop; i++ {
			if len(queue[i].GetList()) == 0 {
				unweighted += queue[i].GetInteger()
			} else {
				queue = append(queue, queue[i].GetList()...)
			}
		}

		weighted += unweighted
		queue = queue[stop:]
	}

	return weighted
}

func depthSumInverse1(nestedList []*NestedInteger) int {
	nums := make([]int, 0)
	dfs(nestedList, 0, &nums)

	var val int

	for i := range nums {
		val += nums[i] * (len(nums) - i)
	}

	return val
}

func dfs(nl []*NestedInteger, level int, nums *[]int) {
	for level > len(*nums)-1 {
		*nums = append(*nums, 0)
	}

	for _, l := range nl {
		if len(l.GetList()) == 0 {
			(*nums)[level] += l.GetInteger()
		} else {
			dfs(l.GetList(), level+1, nums)
		}
	}
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

//		this is a smart solution, initially I would like to get levels of each
//		number, then try to multiply by depth. But this is not necessary, since
//		going into next level # of times means what level that number sits in,
//		for all previous sum, add them whenever entering next level

//		the reason this works is because, for every number, order doesn't matter,
//		only level matters.

//		so, just like BFS, first process all single number, put them into a
//		variable (unweighted), and add unweighted to total sum (weighted), then
//		go to next level.

//		because next level means all previous level number should be added again,
//		unweighted is accumulated and added to total sum every level.

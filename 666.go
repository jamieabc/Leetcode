package main

import "sort"

// If the depth of a tree is smaller than 5, then this tree can be represented by an array of three-digit integers. For each integer in this array:
//
// The hundreds digit represents the depth d of this node where 1 <= d <= 4.
// The tens digit represents the position p of this node in the level it belongs to where 1 <= p <= 8. The position is the same as that in a full binary tree.
// The units digit represents the value v of this node where 0 <= v <= 9.
//
// Given an array of ascending three-digit integers nums representing a binary tree with a depth smaller than 5, return the sum of all paths from the root towards the leaves.
//
// It is guaranteed that the given array represents a valid connected binary tree.
//
//
//
// Example 1:
//
// Input: nums = [113,215,221]
// Output: 12
// Explanation: The tree that the list represents is shown.
// The path sum is (3 + 5) + (3 + 1) = 12.
//
// Example 2:
//
// Input: nums = [113,221]
// Output: 4
// Explanation: The tree that the list represents is shown.
// The path sum is (3 + 1) = 4.
//
//
//
// Constraints:
//
// 1 <= nums.length <= 15
// 110 <= nums[i] <= 489
// nums represents a valid binary tree with depth less than 5.

// tc: O(n), sc: O(n)
func pathSum(nums []int) int {
	var sum int
	arr := make([]int, 8)

	next := make([]int, 8)
	size := len(nums)

	for prev, i := size-1, size-1; i >= 0; i-- {
		// go up level
		if nums[i]/100 != nums[prev]/100 {
			prev = i

			// reset array
			copy(arr, next)
			for j := range next {
				next[j] = 0
			}
		}

		pos := nums[i]%100/10 - 1
		sum += nums[i] % 10 * (max(1, arr[pos]))
		next[pos/2] += max(1, arr[pos])
	}

	return sum
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

// tc: O(n log(n)), sc: O(n)
func pathSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		d1, d2 := nums[i]/100, nums[j]/100
		if d1 != d2 {
			return d1 > d2
		}

		return nums[i]%100 < nums[j]%100
	})

	var sum int
	arr := make([]int, 8)

	next := make([]int, 8)
	size := len(nums)

	for prev, i := 0, 0; i < size; i++ {
		// go up level
		if nums[i]/100 != nums[prev]/100 {
			prev = i

			// reset array
			copy(arr, next)
			for j := range next {
				next[j] = 0
			}
		}

		pos := nums[i]%100/10 - 1
		sum += nums[i] % 10 * (max(1, arr[pos]))
		next[pos/2] += max(1, arr[pos])
	}

	return sum
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	sort number by depth descending, same depth by position ascending

//	2.	there's a bug when i try to do recursive store # of leaves, initially
//		setup starting array to all position has 1 leaf, and do it recursively

//		e.g
//			1
//		  /	  \
//		 2	   3

//		it becomes pos[8] = [1, 1, 1, 1, 1, 1, 1, 1] means all leaves with 1 count
//		the back to parent, pos[0] = 1+1 = 2

//		but there's a bug when leaves not at deepest level

//		e.g.
//					1
//				  /	  \
//				3		1
//			   /
//			 5
//			/
//		  5

// 		right most node 1 is not count, because when recursively upward, pos[1] = 0

//	3.	inspired from solution, it's okay to do directly traversal

//		the reason to sort is to start from leaf nodes and traverse upward

//		the reason to start from leaf nodes is because it provides # of count for
//		each node (depending on # of leaf nodes in its sub-tree)

//		w/o sort, just do traversal backward, start from last node, because
//		array is already sorted in ascending level, which means leaf nodes are
//		places at last

//	4.	inspired from https://leetcode.com/problems/path-sum-iv/discuss/106887/C%2B%2B-Java-Clean-Code

//		can also use hash to store numbers at same level

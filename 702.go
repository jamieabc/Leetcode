package main

// Given an integer array sorted in ascending order, write a function to search target in nums.  If target exists, then return its index, otherwise return -1. However, the array size is unknown to you. You may only access the array using an ArrayReader interface, where ArrayReader.get(k) returns the element of the array at index k (0-indexed).
//
// You may assume all integers in the array are less than 10000, and if you access the array out of bounds, ArrayReader.get will return 2147483647.
//
//
//
// Example 1:
//
// Input: array = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4
//
// Example 2:
//
// Input: array = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1
//
//
//
// Note:
//
//     You may assume that all elements in the array are unique.
//     The value of each element in the array will be in the range [-9999, 9999].

/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
	start := 0
	end := 1

	for reader.get(end) < target {
		start = end
		end <<= 1
	}

	for start <= end {
		mid := start + (end-start)/2
		val := reader.get(mid)

		if val == target {
			return mid
		}

		if val < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

//	1. 	cannot use j directly because it could be the case that target value
//		resides beyond j

//	2.	when doing binary search, i == j could be the case, so need to include
//		that

//	3.	reference from https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/discuss/151685/Shortest-and-cleanest-Java-solution-so-far...

//		it's a pretty solution, because it treats out-of-bound same as actual
//		value, the point of binary search is that what number is too large
//		search mid, in this sense, out of bound value is actually indicating
//		value too big, choose other side.

//		the other way to think this problem is to view total array as
//		follows:
//		1, 2, 3, 4, 5, 7, 2147483647, 2147483647, 2147483647, ..., etc.

//		so out of bound value is same as others.

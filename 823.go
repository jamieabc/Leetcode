package main

import "sort"

// Given an array of unique integers, arr, where each integer arr[i] is strictly greater than 1.
//
// We make a binary tree using these integers, and each number may be used for any number of times. Each non-leaf node's value should be equal to the product of the values of its children.
//
// Return the number of binary trees we can make. The answer may be too large so return the answer modulo 109 + 7.
//
//
//
// Example 1:
//
// Input: arr = [2,4]
// Output: 3
// Explanation: We can make these trees: [2], [4], [4, 2, 2]
//
// Example 2:
//
// Input: arr = [2,4,5,10]
// Output: 7
// Explanation: We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
//
//
//
// Constraints:
//
// 1 <= arr.length <= 1000
// 2 <= arr[i] <= 109

func numFactoredBinaryTrees(arr []int) int {
	mod := int(1e9 + 7)
	sort.Ints(arr)

	table := make(map[int]int)
	for _, i := range arr {
		table[i] = 1
	}

	ans := len(arr)

	for i := range arr {
		for j := i - 1; j >= 0; j-- {
			if arr[i]%arr[j] == 0 {
				quo := arr[i] / arr[j]

				// number not exist
				if table[quo] == 0 {
					continue
				}

				tmp := table[arr[j]] * table[quo]
				table[arr[i]] += tmp
				ans = (ans + tmp) % mod
			}
		}
	}

	return ans
}

func numFactoredBinaryTrees1(arr []int) int {
	var ans int
	mod := int(1e9 + 7)
	table := make(map[int]int)

	for _, i := range arr {
		table[i] = 1
	}
	ans += table[arr[0]]

	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if div := arr[i] / arr[j]; arr[i]%arr[j] == 0 && table[div] > 0 {
				table[arr[i]] += table[arr[j]] * table[div]
			}
		}

		ans = (ans + table[arr[i]]) % mod
	}

	return ans
}

//	Notes
//	1.	need to guarantee when use child count, the child has correct total numbers
//		it means that when processing, need to guarantee when processing child count,
//		count is already correct

//		e.g. [2, 8, 16, 4]
//		start from 2, 2: 1, 8: 1, 16: 1, 4: 1
//		2 * 8 = 16, count[16] = count[2] * count[8] = 1 * 1 = 1
//		but actually, 8 not only from itself, but also from 2 * 4,
//		which means count[8] = 3 (8, 2*4, 4*2) = 3
//		further more, 4 not only from itself, but also from 2 * 2,
//		so count[4] = 2

//		count[16] should actually be count[2]*count[8] + count[4] = 1 * 9 + 2 = 11

//	2.	if children number is different, could swap left & right child, total count
//		should be doubled

//	3.	inspired from https://leetcode.com/problems/binary-trees-with-factors/discuss/1107268/JS-Python-Java-C%2B%2B-or-Fastest-Solution-w-Explanation

//		author explained very clear

//		the reason for dp is that, start from leaf node, a number can also
//		be non-leaf node, count when number is non-leaf comes from other leaf
//		node, which is form of sub-problems

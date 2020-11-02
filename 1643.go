package main

import "fmt"

// Bob is standing at cell (0, 0), and he wants to reach destination: (row, column). He can only travel right and down. You are going to help Bob by providing instructions for him to reach destination.
//
// The instructions are represented as a string, where each character is either:
//
//     'H', meaning move horizontally (go right), or
//     'V', meaning move vertically (go down).
//
// Multiple instructions will lead Bob to destination. For example, if destination is (2, 3), both "HHHVV" and "HVHVH" are valid instructions.
//
// However, Bob is very picky. Bob has a lucky number k, and he wants the kth lexicographically smallest instructions that will lead him to destination. k is 1-indexed.
//
// Given an integer array destination and an integer k, return the kth lexicographically smallest instructions that will take Bob to destination.
//
//
//
// Example 1:
//
// Input: destination = [2,3], k = 1
// Output: "HHHVV"
// Explanation: All the instructions that reach (2, 3) in lexicographic order are as follows:
// ["HHHVV", "HHVHV", "HHVVH", "HVHHV", "HVHVH", "HVVHH", "VHHHV", "VHHVH", "VHVHH", "VVHHH"].
//
// Example 2:
//
// Input: destination = [2,3], k = 2
// Output: "HHVHV"
//
// Example 3:
//
// Input: destination = [2,3], k = 3
// Output: "HHVVH"
//
//
//
// Constraints:
//
//     destination.length == 2
//     1 <= row, column <= 15
//     1 <= k <= nCr(row + column, row), where nCr(a, b) denotes a choose b.

func kthSmallestPath(destination []int, k int) string {
	h, v := destination[1], destination[0]
	str := make([]byte, h+v)

	for idx := 0; idx < len(str); idx++ {
		if h == 0 || v == 0 {
			for ; h > 0; idx, h = idx+1, h-1 {
				str[idx] = 'H'
			}

			for ; v > 0; idx, v = idx+1, v-1 {
				str[idx] = 'V'
			}
			break
		}

		// check if h is placed, maximum k
		count := binomial(h-1+v, h-1)

		if k <= count {
			str[idx] = 'H'
			h--
		} else {
			str[idx] = 'V'
			v--
			k -= count
		}
	}

	return string(str)
}

// total = 4, h = 2, v = 2
// count = 4!/(2!*2!)
func binomial(total, h int) int {
	val := float64(1)

	for i := float64(2); i <= float64(total); i++ {
		if i > float64(h) {
			val *= i
		}
	}

	for i := float64(2); i <= float64(total-h); i++ {
		val /= i
	}

	fmt.Println(total, h, val)

	return int(val)
}

//	Notes
//	1.	lexicographical order is not simply moving V to start, considering
//		example: HHHVV, lexicographical order: ["HHHVV", "HHVHV", "HHVVH",
//		"HVHHV", "HVHVH", "HVVHH", "VHHHV", "VHHVH", "VHVHH", "VVHHH"].

//		note that 4th order HVHHV, it's not moving V forward from 3rd
//		HHVVH -> HVHVH

//	2.	inspired from alex, for each position, H if k <= remaining permutations
//		e.g. HHHVV, for first H to be positioned, as long as k <= 4!/(2!*2!)

//	3.	overflow

//	4.	inspired from https://leetcode.com/problems/kth-smallest-instructions/discuss/918569/Java-Bottom-up-dp-O(m*n)-solution

//		use dp to compute ways, instead of combination, e.g. 3H2V

//		1 1 1 1
//		1 2 3 4
//		1 3 6 10

//		dp[2, 3] = 10, means 3H2V with 10 combinations,
//		to check for first character is H, means remaining 2H2V, combination
//		count is [2, 2] = 6

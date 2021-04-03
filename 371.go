package main

// Given two integers a and b, return the sum of the two integers without using the operators + and -.
//
//
//
// Example 1:
//
// Input: a = 1, b = 2
// Output: 3
// Example 2:
//
// Input: a = 2, b = 3
// Output: 5
//
//
// Constraints:
//
// -1000 <= a, b <= 1000

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := a & b
		a = sum
		b = carry << 1
	}

	return a
}

//	Notes
//	1.	stuck at the negative condition

//	2.	inspired from https://leetcode.com/problems/sum-of-two-integers/discuss/776952/Python-BEST-LeetCode-371-Explanation-for-Python

//		carry always apply to the left, xor does add
//		very clever solution

package main

// Given a positive integer K, you need to find the length of the smallest positive integer N such that N is divisible by K, and N only contains the digit 1.
//
// Return the length of N. If there is no such N, return -1.
//
// Note: N may not fit in a 64-bit signed integer.
//
//
//
// Example 1:
//
// Input: K = 1
// Output: 1
// Explanation: The smallest answer is N = 1, which has length 1.
//
// Example 2:
//
// Input: K = 2
// Output: -1
// Explanation: There is no such positive integer N divisible by 2.
//
// Example 3:
//
// Input: K = 3
// Output: 3
// Explanation: The smallest answer is N = 111, which has length 3.
//
//
//
// Constraints:
//
//     1 <= K <= 105

func smallestRepunitDivByK(K int) int {
	if K&1 == 0 || K%5 == 0 {
		return -1
	}

	if K == 1 {
		return 1
	}

	remainder := 1

	for i := 2; i <= K; i++ {
		remainder = (remainder*10 + 1) % K
		if remainder == 0 {
			return i
		}
	}

	return -1
}

//	Notes
//	1.	inspired from hint, 11 = 1*10+1, 111 = 11*10+1, and modulo can be stored
//		for next time calculating modulo

//	2.	any number ended with 5 or 0 cannot be divided

//	3.	inspired from solution, maximum iteration to calculate is k times
//		(https://en.wikipedia.org/wiki/Pigeonhole_principle)

//		if this principle is not think, use a set to store remainder value, once
//		a remainder value is repeated, then it's not possible to find an answer

//	4.	inspired from https://leetcode.com/problems/smallest-integer-divisible-by-k/discuss/261805/Very-short-and-clear-explanation-O(K)-(also-2-linesJavaC%2B%2BCJavaScriptPHP).

//		author provides a very good insight of modulo operator: map from 0 ~ K-1
//		to 0 ~ K-1, this also explains if a mapped number exist again, then
//		mapping will be a iteration and never finds an answer

//	5.	inspired from https://leetcode.com/problems/smallest-integer-divisible-by-k/discuss/260852/JavaC%2B%2BPython-O(1)-Space-with-Proves-of-Pigeon-Holes

//		use a valid bound

package main

// Every non-negative integer N has a binary representation.  For example, 5 can be represented as "101" in binary, 11 as "1011" in binary, and so on.  Note that except for N = 0, there are no leading zeroes in any binary representation.
//
// The complement of a binary representation is the number in binary you get when changing every 1 to a 0 and 0 to a 1.  For example, the complement of "101" in binary is "010" in binary.
//
// For a given number N in base-10, return the complement of it's binary representation as a base-10 integer.
//
//
//
// Example 1:
//
// Input: 5
// Output: 2
// Explanation: 5 is "101" in binary, with complement "010" in binary, which is 2 in base-10.
// Example 2:
//
// Input: 7
// Output: 0
// Explanation: 7 is "111" in binary, with complement "000" in binary, which is 0 in base-10.
// Example 3:
//
// Input: 10
// Output: 5
// Explanation: 10 is "1010" in binary, with complement "0101" in binary, which is 5 in base-10.
//
//
// Note:
//
// 0 <= N < 10^9
// This question is the same as 476: https://leetcode.com/problems/number-complement/

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}

	mask := N
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16

	return N ^ mask
}

func bitwiseComplement1(N int) int {
	var msb, num int

	for msb = 31; msb >= 0; msb-- {
		if (1<<msb)&N > 0 {
			break
		}
	}

	// 0
	if msb == -1 {
		return 1
	}

	for i := 0; i <= msb; i++ {
		if N&(1<<i) == 0 {
			num |= 1 << i
		}
	}

	return num
}

//	Notes
//	1.	for a number to be flipped, 0 is a boundary case

//	2.	flip a bit mean XOR 1
//		e.g. 0 ^ 1 = 1, 1 ^ 1 = 0

//	3.	if there's a number which has all one after msb, all 0 before msb,
//		xor that number gets answer

//		how to effectively generate bitmask, one clever way from solution is to
//		or original

//	4.	inspired from https://leetcode.com/problems/complement-of-base-10-integer/discuss/256740/JavaC%2B%2BPython-Find-111.....1111-greater-N

//		lee uses another way to calculate:
//		complement + original = 111....1
//		e.g 4 (100) + 3 (011) = 7 (111)

//		complement = 111...1 - original

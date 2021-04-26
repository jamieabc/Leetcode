package main

// Given an integer n, you must transform it into 0 using the following operations any number of times:
//
//     Change the rightmost (0th) bit in the binary representation of n.
//     Change the ith bit in the binary representation of n if the (i-1)th bit is set to 1 and the (i-2)th through 0th bits are set to 0.
//
// Return the minimum number of operations to transform n into 0.
//
//
//
// Example 1:
//
// Input: n = 0
// Output: 0
//
// Example 2:
//
// Input: n = 3
// Output: 2
// Explanation: The binary representation of 3 is "11".
// "11" -> "01" with the 2nd operation since the 0th bit is 1.
// "01" -> "00" with the 1st operation.
//
// Example 3:
//
// Input: n = 6
// Output: 4
// Explanation: The binary representation of 6 is "110".
// "110" -> "010" with the 2nd operation since the 1st bit is 1 and 0th through 0th bits are 0.
// "010" -> "011" with the 1st operation.
// "011" -> "001" with the 2nd operation since the 0th bit is 1.
// "001" -> "000" with the 1st operation.
//
// Example 4:
//
// Input: n = 9
// Output: 14
//
// Example 5:
//
// Input: n = 333
// Output: 393
//
//
//
// Constraints:
//
//     0 <= n <= 109

func minimumOneBitOperations(n int) int {
	if n <= 1 {
		return n
	}

	var bit int
	for i := 31; i >= 0; i-- {
		if n&(1<<i) > 0 {
			bit = i
			break
		}
	}

	return (1 << (bit + 1)) - 1 - minimumOneBitOperations(n^(1<<bit))
}

func minimumOneBitOperations2(n int) int {
	var count, turn int

	for i := 31; i >= 0; i-- {
		if n&(1<<i) > 0 {
			if turn&1 == 0 {
				count += (1 << (i + 1)) - 1
			} else {
				count -= (1 << (i + 1)) - 1
			}
			turn++
		}
	}

	return count
}

func minimumOneBitOperations1(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	var msb, lsb int

	for i := 2; i <= n; i++ {
		for msb = 31; msb >= 0; msb-- {
			if i&(1<<msb) > 0 {
				break
			}
		}

		for lsb = 0; lsb < 32; lsb++ {
			if i&(1<<lsb) > 0 {
				break
			}
		}

		if msb == lsb {
			dp[i] = 1 + 2*dp[i>>1]
		} else {
			dp[i] = dp[1<<msb] - dp[i-1<<msb]
		}
	}

	return dp[n]
}

//	Notes
//	1.	to find number of operation to convert a number into zero, operation count
//		is max when number is 2's power, e.g. 2, 4, 8, 16, etc.

//		because for any bit 1 at ith position to become 0, there needs to have
//		another bit 1 at (i-1)th position

//		e.g. 8 (1000), generate a bit close takes 7 (100) operations, update
//		(1100) to (0100) takes 1 operation, and convert 4 (100) to 0 takes 7
//		operations

//		overall, op(8) = 2 * op(4) + 1 = 2 * 7 + 1 = 15
//		op(4) = 2 * op(2) + 1 = 7
//		op(2) = 2 * op(1) + 1 = 3

//		but for any number in 9 ~ 15, operation count is less than 15, because
//		goal of whole operation is to generate 1 close to msb, thus any 1 in
//		between reduces operation count

//	2.	to clear a bit at ith index, it takes 2^(k+1)-1 steps

//		1: 2^1-1 = 1
//		2: 2^2-1 = 3
//		4: 2^3-1 = 7

//		the reason is because whole process to make power of 2 number become 0
//		with 3 parts:
//		- only msb and msb-1 bit become 1, others 0
//		- 110...0 -> 010...0
//		- msb-1 bit become 0

//		with this recursive relation, overall steps are op(kth bit 1) = 2^(k+1) -1

//	3.	bit sequence by this rule will traverse whole bits

//		10 -> 11 -> 01 -> 00
//		100 -> 101 -> 111 -> 110 -> 010 -> 011 -> 001 -> 000
//		1000 -> 1001 -> 1011 -> 1010 -> 1110 -> 1111 -> 1101 -> 1100 -> 0100

//		every sequence occurs only once, and more 1 means more steps can be
//		avoided

//		e.g. op(1010) = op(1000) - op(0010) = 15 - 3 = 12

//		if there are multiple 1, pattern will be plus and minus intervene
//		op(10000) - op(01000) + op(00100) - ...

//	4.	inspired from https://leetcode.com/problems/minimum-one-bit-operations-to-make-integers-zero/discuss/877708/PythonC%2B%2B-O(log-n)-with-Prove

//		author has similar technique but with detailed explanation

//	5.	inspired form https://leetcode.com/problems/minimum-one-bit-operations-to-make-integers-zero/discuss/877741/C%2B%2B-solution-with-explanation

//		- 1000...0 takes most steps
//		- more 1 means more reduction

//		(111) = (100) - (011) = (100) - ( (010) - (001) )
//			  = (100) - (010) + (001)

//		the sign changes each bit

//		author comes with recursion solution, which is really brilliant

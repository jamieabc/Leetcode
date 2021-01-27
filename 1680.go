package main

// Given an integer n, return the decimal value of the binary string formed by concatenating the binary representations of 1 to n in order, modulo 109 + 7.
//
//
//
// Example 1:
//
// Input: n = 1
// Output: 1
// Explanation: "1" in binary corresponds to the decimal value 1.
//
// Example 2:
//
// Input: n = 3
// Output: 27
// Explanation: In binary, 1, 2, and 3 corresponds to "1", "10", and "11".
// After concatenating them, we have "11011", which corresponds to the decimal value 27.
//
// Example 3:
//
// Input: n = 12
// Output: 505379714
// Explanation: The concatenation results in "1101110010111011110001001101010111100".
// The decimal value of that is 118505380540.
// After modulo 109 + 7, the result is 505379714.
//
//
//
// Constraints:
//
//    1 <= n <= 105
//

func concatenatedBinary(n int) int {
	num, digits := 1, 1
	var ans int
	mod := int(1e9 + 7)

	for ; num <= n; num++ {
		if num == 1<<digits {
			digits++
		}

		ans = ans << digits
		ans |= num

		ans = ans % mod
	}

	return ans
}

func concatenatedBinary1(n int) int {
	mod := int64(1e9 + 7)
	var ans int64

	for i := 1; i <= n; i++ {
		size := digits(i)
		ans = ((ans << size) | int64(i)) % mod
	}

	return int(ans)
}

func digits(i int) int {
	var shift int

	for shift = 31; (1<<shift)&i == 0; shift-- {
	}

	return shift + 1
}

//	Notes
//	1.	inspired from solution, add mean or operation

//	2.	inspired from solution, since # of digits is known, use that to check

//	3.	add https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/discuss/963549/python-(log(n))-2-approach

//		not implement it, but seems brilliant

package main

import "math"

// Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
//
// Return the quotient after dividing dividend by divisor.
//
// The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.
//
// Note:
//
// Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For this problem, assume that your function returns 231 − 1 when the division result overflows.
//
//
// Example 1:
//
// Input: dividend = 10, divisor = 3
// Output: 3
// Explanation: 10/3 = truncate(3.33333..) = 3.
// Example 2:
//
// Input: dividend = 7, divisor = -3
// Output: -2
// Explanation: 7/-3 = truncate(-2.33333..) = -2.
// Example 3:
//
// Input: dividend = 0, divisor = 1
// Output: 0
// Example 4:
//
// Input: dividend = 1, divisor = 1
// Output: 1
//
//
// Constraints:
//
// -231 <= dividend, divisor <= 231 - 1
// divisor != 0

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	var negative bool

	if dividend < 0 {
		dividend = -dividend
		negative = !negative
	}

	if divisor < 0 {
		divisor = -divisor
		negative = !negative
	}

	var ans int
	for current, mul := divisor, 1; dividend >= divisor; {
		if current < dividend {
			current = current << 1
			mul = mul << 1
		} else {
			// fraction, discard
			if current == divisor {
				ans += mul
				break
			}

			ans += mul >> 1
			mul = 1

			current = current >> 1
			dividend -= current
			current = divisor
		}
	}

	if negative {
		return -ans
	}
	return ans
}

//	Notes
//	1.	didn't think of way to solve it, got TLE

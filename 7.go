package main

import (
	"fmt"
	"math"
)

//Given a 32-bit signed integer, reverse digits of an integer.
//
//Example 1:
//
//Input: 123
//Output: 321
//Example 2:
//
//Input: -123
//Output: -321
//Example 3:
//
//Input: 120
//Output: 21
//Note:
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	y := int64(x)

	result := int64(0)
	var digit int64

	for y != 0 {
		digit = y % 10
		result += digit
		y /= 10
		if y != 0 {
			result *= 10
		}
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
	}
	return int(result)
}

func main() {
	x := 1534236469
	fmt.Printf("%d reversed: %d\n", x, reverse(x))
}

package main

import (
	"math"
)

//Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.
//
//Example 1:
//
//Input: 5
//Output: True
//Explanation: 1 * 1 + 2 * 2 = 5
//
//
//
//Example 2:
//
//Input: 3
//Output: False

func judgeSquareSum(c int) bool {
	sqrt := math.Sqrt(float64(c))

	if float64(int(sqrt)) == sqrt {
		return true
	}

	max := int(sqrt)
	i := 0
	j := max

	for {
		sum := i*i + j*j
		if sum == c {
			return true
		}

		if i == j {
			return false
		}

		if sum > c {
			j--
		} else {
			i++
		}
	}
}

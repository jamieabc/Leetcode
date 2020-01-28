package main

import "math"

//Write an algorithm to determine if a number is "happy".
//
//A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.
//
//Example:
//
//Input: 19
//Output: true
//Explanation:
//12 + 92 = 82
//82 + 22 = 68
//62 + 82 = 100
//12 + 02 + 02 = 1

func isHappy(n int) bool {
	limit := 20
	for limit > 0 {
		nums := digits(n)
		sum, overflow := squareSum(nums)
		if overflow {
			return false
		}
		if sum == 1 {
			return true
		}
		n = sum
		limit--
	}
	return false
}

func digits(n int) []int {
	result := make([]int, 0)
	for n != 0 {
		result = append(result, n%10)
		n /= 10
	}

	return result
}

func squareSum(nums []int) (int, bool) {
	result := 0
	for _, n := range nums {
		square := n * n
		if math.MaxInt32-square-result < 0 {
			return 0, false
		}
		result += square
	}

	return result, false
}

// problems
// 1. when using range, first return value is index

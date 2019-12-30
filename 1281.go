package main

//Given an integer number n, return the difference between the product of its digits and the sum of its digits.
//
//
//
//Example 1:
//
//Input: n = 234
//Output: 15
//Explanation:
//Product of digits = 2 * 3 * 4 = 24
//Sum of digits = 2 + 3 + 4 = 9
//Result = 24 - 9 = 15
//
//Example 2:
//
//Input: n = 4421
//Output: 21
//Explanation:
//Product of digits = 4 * 4 * 2 * 1 = 32
//Sum of digits = 4 + 4 + 2 + 1 = 11
//Result = 32 - 11 = 21
//
//
//
//Constraints:
//
//    1 <= n <= 10^5

func subtractProductAndSum(n int) int {
	digits := make([]int, 0)
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	sum := 0
	product := 1
	for _, i := range digits {
		sum += i
		product *= i
	}

	return product - sum
}

// problems
// 1. product - sum

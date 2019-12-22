package main

//Given a positive integer num, write a function which returns True if num is a perfect square else False.
//
//Note: Do not use any built-in library function such as sqrt.
//
//Example 1:
//
//Input: 16
//Output: true
//
//Example 2:
//
//Input: 14
//Output: false

func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}

	if num <= 1 {
		return true
	}

	var double int
	for i := 0; i < num; i++ {
		double = i * i
		if double == num {
			return true
		}
		if double > num {
			return false
		}
	}
	return false
}

// problem
// 1. double should be multiply

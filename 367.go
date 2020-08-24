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
	if num == 1 {
		return true
	}

	low, high := 1, num/2

	for low <= high {
		mid := low + (high-low)/2
		double := mid * mid

		if double == num {
			return true
		} else if double < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func isPerfectSquare1(num int) bool {
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

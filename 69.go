package main

//Implement int sqrt(int x).
//
//Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
//
//Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.
//
//Example 1:
//
//Input: 4
//Output: 2
//Example 2:
//
//Input: 8
//Output: 2
//Explanation: The square root of 8 is 2.82842..., and since
//the decimal part is truncated, 2 is returned.

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x < 4 {
		return 1
	}

	return recursive(2, x/2, x)
}

func recursive(start, end, target int) int {
	var multiple int

	if start >= end {
		multiple = start * start
		if multiple > target {
			return start - 1
		}
		return start
	}

	middle := (start + end) / 2
	multiple = middle * middle
	if multiple == target {
		return middle
	}

	if multiple < target {
		return recursive(middle+1, end, target)
	} else {
		return recursive(start, middle-1, target)
	}
}

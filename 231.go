package main

//Given an integer, write a function to determine if it is a power of two.
//
//Example 1:
//
//Input: 1
//Output: true
//Explanation: 20 = 1
//
//Example 2:
//
//Input: 16
//Output: true
//Explanation: 24 = 16
//
//Example 3:
//
//Input: 218
//Output: false

func isPowerOfTwo(n int) bool {
	// power of 2 means only exist one digit 1 in binary format

	if n <= 0 {
		return false
	}

	count1 := 0

	for i := 0; i < 32; i++ {
		digit := n & 1
		if digit == 1 {
			if count1 == 0 {
				count1++
			} else {
				return false
			}
		}
		n >>= 1
	}
	return true
}

// problems
// 1. forget about 0

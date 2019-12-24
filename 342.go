package main

//Given an integer (signed 32 bits), write a function to check whether it is a power of 4.
//
//Example 1:
//
//Input: 16
//Output: true
//
//Example 2:
//
//Input: 5
//Output: false
//
//Follow up: Could you solve it without loops/recursion?

// number of consecutive zeros from right to left
// power of 4 means count is even (2, 4, 6, 8, etc.)

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}

	if num == 1 {
		return true
	}

	var i, count1, count0 int
	tmp := num
	for i = 0; i < 32; i++ {
		if tmp&1 == 1 {
			count1++
		}
		tmp >>= 1
	}

	tmp = num
	for i = 0; i < 32; i++ {
		if tmp&1 == 0 {
			count0++
		} else {
			break
		}
		tmp >>= 1
	}

	if count1 == 1 && count0 > 0 && count0%2 == 0 {
		return true
	}

	return false
}

// problems
// 1. didn't understand problem clearly, it's power of 4, not able to be divided by 4
// 2. wrong consideration, power of 4 not only needs to have double zeros, but also only 1 in binary
// 3. forget that when second wrong of checking 0, numb is mutated (num >>= 1)

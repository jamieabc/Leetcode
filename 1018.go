package main

//Given an array A of 0s and 1s, consider N_i: the i-th subarray from A[0] to A[i] interpreted as a binary number (from most-significant-bit to least-significant-bit.)
//
//Return a list of booleans answer, where answer[i] is true if and only if N_i is divisible by 5.
//
//Example 1:
//
//Input: [0,1,1]
//Output: [true,false,false]
//Explanation:
//The input numbers in binary are 0, 01, 011; which are 0, 1, and 3 in base-10.  Only the first number is divisible by 5, so answer[0] is true.
//
//Example 2:
//
//Input: [1,1,1]
//Output: [false,false,false]
//
//Example 3:
//
//Input: [0,1,1,1,1,1]
//Output: [true,false,false,false,true,false]
//
//Example 4:
//
//Input: [1,1,1,0,1]
//Output: [false,false,false,false,false]
//
//
//
//Note:
//
//    1 <= A.length <= 30000
//    A[i] is 0 or 1

func prefixesDivBy5(A []int) []bool {
	length := len(A)
	result := make([]bool, length)

	var num int
	for i := 0; i < length; i++ {
		num = (num*2 + A[i]) % 5

		if num == 0 {
			result[i] = true
		}
	}

	return result
}

// problems
//	1.	might have overflow, length <= 3000, int max is 32 bits
//	2. 	check for last 3 digits of 1, 1, 0
//	3.	divided of 5 in binary form has no rules...use int64 instead
//	4. 	it passes when switching from int64 -> uint64
//	5. 	optimization, use mod value

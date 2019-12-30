package main

//For a non-negative integer X, the array-form of X is an array of its digits in left to right order.  For example, if X = 1231, then the array form is [1,2,3,1].
//
//Given the array-form A of a non-negative integer X, return the array-form of the integer X+K.
//
//
//
//Example 1:
//
//Input: A = [1,2,0,0], K = 34
//Output: [1,2,3,4]
//Explanation: 1200 + 34 = 1234
//
//Example 2:
//
//Input: A = [2,7,4], K = 181
//Output: [4,5,5]
//Explanation: 274 + 181 = 455
//
//Example 3:
//
//Input: A = [2,1,5], K = 806
//Output: [1,0,2,1]
//Explanation: 215 + 806 = 1021
//
//Example 4:
//
//Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
//Output: [1,0,0,0,0,0,0,0,0,0,0]
//Explanation: 9999999999 + 1 = 10000000000
//
//
//
//Note：
//
//    1 <= A.length <= 10000
//    0 <= A[i] <= 9
//    0 <= K <= 10000
//    If A.length > 1, then A[0] != 0

func addToArrayForm(A []int, K int) []int {
	if K == 0 {
		return A
	}

	length := len(A)
	carry := 0
	reversed := make([]int, 0)
	sum := 0

	for i := length - 1; K != 0 || i >= 0; i-- {
		digit := K % 10

		// make sure i is valid
		if i >= 0 {
			sum = A[i] + digit + carry
		} else {
			sum = digit + carry
		}

		// check if overflow
		if sum >= 10 {
			carry = 1
			reversed = append(reversed, sum-10)
		} else {
			carry = 0
			reversed = append(reversed, sum)
		}

		K /= 10
	}

	if carry == 1 {
		reversed = append(reversed, 1)
	}

	length = len(reversed)
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[length-1-i] = reversed[i]
	}

	return result
}

// problems
// 1. forget to return correct slice
// 2. A length to long, cannot use integer to add, need to do bit number addition
// 3. condition is k != 0, but forget to check that when k == 0, is A all traversed
// 4. forget about carry

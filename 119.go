package main

//Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
//
//Note that the row index starts from 0.
//
//
//In Pascal's triangle, each number is the sum of the two numbers directly above it.
//
//Example:
//
//Input: 3
//Output: [1,3,3,1]
//
//Follow up:
//
//Could you optimize your algorithm to use only O(k) extra space?

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1

	for limit := 1; limit <= rowIndex; limit++ {
		prev := 0
		for i := 0; i < limit; i++ {
			tmp := result[i]
			result[i] += prev
			prev = tmp
		}
		result[limit] = 1
	}

	return result
}

//	Notes
//	1.	inspired from solution, it's more efficient to start backward

//		e.g. row: 1 3 3 1
//		       => 1 3 3 1 1		append 1 at tail
//			   => 1 3 3 4 1		4 = previous 1+3
//             => 1 3 6 4 1		6 = previous 3+3
//             => 1 4 6 4 1		4 = previous 3+1

//		tc: O(k^2), cause each line will be calculated once
//		overall: 1 + 2 + 3 + 4 + 5 + ... + (k+1) = (k+1)*(k+2)/2 = O(k^2)

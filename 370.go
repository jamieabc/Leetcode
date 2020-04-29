package main

// Assume you have an array of length n initialized with all 0's and are given k update operations.
//
// Each operation is represented as a triplet: [startIndex, endIndex, inc] which increments each element of subarray A[startIndex ... endIndex] (startIndex and endIndex inclusive) with inc.
//
// Return the modified array after all k operations were executed.
//
// Example:
//
// Input: length = 5, updates = [[1,3,2],[2,4,3],[0,2,-2]]
// Output: [-2,0,3,5,3]
//
// Explanation:
//
// Initial state:
// [0,0,0,0,0]
//
// After applying operation [1,3,2]:
// [0,2,2,2,0]
//
// After applying operation [2,4,3]:
// [0,2,5,5,3]
//
// After applying operation [0,2,-2]:
// [-2,0,3,5,3]

func getModifiedArray(length int, updates [][]int) []int {
	result := make([]int, length)

	for _, u := range updates {
		if u[0] < length {
			result[u[0]] += u[2]
		}

		if u[1]+1 < length {
			result[u[1]+1] -= u[2]
		}
	}

	var sum int
	for i, r := range result {
		if r != 0 {
			sum += r
		}
		result[i] = sum
	}

	return result
}

//	problems
//	1.	too slow, many operations are repeated. Cannot find better solution,
//		reference: https://leetcode.com/problems/range-addition/discuss/343746/Python-event-based-idea-from-O(nlgn)-to-O(n)

//		the first solution is O(n log n), it converts each updates into events:
//		global add from index 1 & global minus after index 2, then sort events
//		by occur index. Suddenly, all updates are connected by a variable which
//		represents all operations before some index, so math operations are not
//		repeat again any more.

//		why a position's number cannot be decided by some interval? because
//		there exists overlap, which makes only look at one interval might be
//		wrong.

//	2.	previous method take O(n long n) because of sorting, if there's a array
//		that's corresponds with length, then each event can be put into array,
//		if multiple events into same position, merge it. With this way, sorting
//		complexity is reduced.

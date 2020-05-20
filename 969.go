package main

// Given an array A, we can perform a pancake flip: We choose some positive integer k <= A.length, then reverse the order of the first k elements of A.  We want to perform zero or more pancake flips (doing them one after another in succession) to sort the array A.
//
// Return the k-values corresponding to a sequence of pancake flips that sort A.  Any valid answer that sorts the array within 10 * A.length flips will be judged as correct.
//
//
//
// Example 1:
//
// Input: [3,2,4,1]
// Output: [4,2,4,3]
// Explanation:
// We perform 4 pancake flips, with k values 4, 2, 4, and 3.
// Starting state: A = [3, 2, 4, 1]
// After 1st flip (k=4): A = [1, 4, 2, 3]
// After 2nd flip (k=2): A = [4, 1, 2, 3]
// After 3rd flip (k=4): A = [3, 2, 1, 4]
// After 4th flip (k=3): A = [1, 2, 3, 4], which is sorted.
// Example 2:
//
// Input: [1,2,3]
// Output: []
// Explanation: The input is already sorted, so there is no need to flip anything.
// Note that other answers, such as [3, 3], would also be accepted.
//
//
// Note:
//
// 1 <= A.length <= 100
// A[i] is a permutation of [1, 2, ..., A.length]

func pancakeSort(A []int) []int {
	length := len(A)
	result := make([]int, 0)
	if length == 1 {
		return result
	}

	for i := length; i > 0; i-- {
		for j := range A {
			if A[j] == i {
				if j == i-1 {
					break
				}

				// move A[j] to first num
				if j != 0 {
					result = append(result, j+1)
					flip(A, j)
				}

				// reverse
				result = append(result, A[0])
				flip(A, A[0]-1)
			}
		}
	}

	return result
}

func flip(arr []int, idx int) {
	for i, j := 0, idx; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

//	problems
//	1.	add reference https://leetcode.com/problems/pancake-sorting/discuss/494417/Dew-It-or-True-O(n)-or-Explained-with-Diagrams

//		it has great diagram & explanation, but I didn't take time to understand it

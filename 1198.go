package main

// Given a matrix mat where every row is sorted in increasing order, return the smallest common element in all rows.
//
// If there is no common element, return -1.
//
//
//
// Example 1:
//
// Input: mat = [[1,2,3,4,5],[2,4,5,8,10],[3,5,7,9,11],[1,3,5,7,9]]
// Output: 5
//
//
//
// Constraints:
//
//     1 <= mat.length, mat[i].length <= 500
//     1 <= mat[i][j] <= 10^4
//     mat[i] is sorted in increasing order.

func smallestCommonElement(mat [][]int) int {
	yLength := len(mat)
	if yLength == 0 {
		return -1
	}
	xLength := len(mat[0])

	var found bool

	for i := 0; i < xLength; i++ {
		found = true

		for j := 1; j < yLength; j++ {
			found = found && binarySearch(mat[j], mat[0][i], 0, len(mat[j])-1)
			if !found {
				break
			}
		}

		if found {
			return mat[0][i]
		}
	}

	return -1
}

func binarySearch(nums []int, target, start, end int) bool {
	if start > end {
		return false
	}

	mid := start + (end-start)/2

	if nums[mid] == target {
		return true
	}

	if nums[mid] > target {
		return binarySearch(nums, target, start, mid-1)
	}
	return binarySearch(nums, target, mid+1, end)
}

//	problems
//	1.	each row is already sorted, so binary search can help to help candidates

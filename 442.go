package main

// Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
//
// Find all the elements that appear twice in this array.
//
// Could you do it without extra space and in O(n) runtime?
//
// Example:
//
// Input:
// [4,3,2,7,8,2,3,1]
//
// Output:
// [2,3]

func findDuplicates(nums []int) []int {
	result := make([]int, 0)

	for _, n := range nums {
		// index start from 0, but number starts from 1, so decrease 1 for
		// indirect accessing
		if n < 0 {
			n = -n - 1
		} else {
			n -= 1
		}

		// if the number appears twice, the other number referenced by the
		// number becomes negative
		if nums[n] < 0 {
			result = append(result, n+1)
		} else {
			nums[n] *= -1
		}
	}

	return result
}

//	Notes
//	1.	I store wrong number, it should be n not nums[n]

//	2.	inspired from solution, the wording is very precise, mark visited number
//		in input array itself, very elegant!!

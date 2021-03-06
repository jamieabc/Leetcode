package main

// Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.
//
// Note:
//
//     All numbers will be positive integers.
//     The solution set must not contain duplicate combinations.
//
// Example 1:
//
// Input: k = 3, n = 7
// Output: [[1,2,4]]
//
// Example 2:
//
// Input: k = 3, n = 9
// Output: [[1,2,6], [1,3,5], [2,3,4]]

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)

	combination(k, n, []int{}, &result)

	return result
}

func combination(k, n int, tmp []int, result *[][]int) {
	if n < 0 {
		return
	}

	if k == 0 {
		if n == 0 {
			*result = append(*result, tmp)
		}
		return
	}

	start := 1
	if len(tmp) > 0 {
		start = tmp[len(tmp)-1] + 1
	}

	for start <= 9 {
		t := append([]int{}, tmp...)
		t = append(t, start)
		combination(k-1, n-start, t, result)
		start++
	}
}

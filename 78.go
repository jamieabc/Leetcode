package main

// Given a set of distinct integers, nums, return all possible subsets (the power set).
//
// Note: The solution set must not contain duplicate subsets.
//
// Example:
//
// Input: nums = [1,2,3]
// Output:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

func subsets(nums []int) [][]int {
	result := [][]int{[]int{}}
	length := len(nums)

	if length == 0 {
		return result
	}

	bits := make([]bool, length)
	tmp := make([]int, 0)
	recursive(nums, bits, &tmp, &result, 0)

	return result
}

func recursive(nums []int, bits []bool, tmp *[]int, result *[][]int, idx int) {
	for i := idx; i < len(nums); i++ {
		// number is already used
		if bits[i] {
			continue
		}

		*tmp = append(*tmp, nums[i])
		prev := make([]int, len(*tmp))
		_ = copy(prev, *tmp)
		*result = append(*result, prev)

		bits[i] = true
		recursive(nums, bits, tmp, result, i)
		(*tmp) = (*tmp)[:len(*tmp)-1]
		bits[i] = false
	}
}

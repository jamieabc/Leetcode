package main

//Given a collection of distinct integers, return all possible permutations.
//
//Example:
//
//Input: [1,2,3]
//Output:
//[
//[1,2,3],
//[1,3,2],
//[2,1,3],
//[2,3,1],
//[3,1,2],
//[3,2,1]
//]

func permute(nums []int) [][]int {
	result := make([][]int, 0)

	recursive(&result, nums, []int{})

	return result
}

func recursive(result *[][]int, remain, existing []int) {
	// found
	if len(remain) == 0 {
		*result = append(*result, existing)
		return
	}

	// keep iterating
	for i := range remain {
		newExisting := append(existing, remain[i])

		// remove put in number
		newRemain := append([]int{}, remain[:i]...)
		newRemain = append(newRemain, remain[i+1:]...)

		recursive(result, newRemain, newExisting)
	}
}

func permute1(nums []int) [][]int {
	length := len(nums)

	if length == 0 {
		return [][]int{}
	}

	result := make([][]int, 1)
	result[0] = []int{nums[0]}

	for i := 1; i < length; i++ {
		tmp := make([][]int, 0)
		for j := range result {
			tmp = append(tmp, insert(result[j], nums[i])...)
		}
		result = tmp
	}

	return result
}

func insert(dest []int, target int) [][]int {
	result := make([][]int, 0)

	for i := range dest {
		tmp := make([]int, 0)
		tmp = append(tmp, dest[:i]...)
		tmp = append(tmp, target)
		tmp = append(tmp, dest[i:]...)
		result = append(result, tmp)
	}

	tmp := make([]int, 0)
	tmp = append(dest, target)
	result = append(result, tmp)

	return result
}

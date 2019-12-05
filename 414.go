package main

//Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).
//
//Example 1:
//
//Input: [3, 2, 1]
//
//Output: 1
//
//Explanation: The third maximum is 1.
//
//Example 2:
//
//Input: [1, 2]
//
//Output: 2
//
//Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
//
//Example 3:
//
//Input: [2, 2, 3, 1]
//
//Output: 1
//
//Explanation: Note that the third maximum here means the third maximum distinct number.
//Both numbers with value 2 are both considered as second maximum.

func thirdMax(nums []int) int {
	result := make([]int, 1)
	result[0] = nums[0]
	min := nums[0]
	index := 0

	for i := 1; i < len(nums); i++ {
		if same(result, nums[i]) {
			continue
		}

		// fill
		if len(result) < 3 {
			result = append(result, nums[i])
			index = findMinIndex(result)
			min = result[index]
			continue
		}

		// compare
		if nums[i] <= min {
			continue
		}

		result[index] = nums[i]
		index = findMinIndex(result)
		min = result[index]
	}

	if len(result) != 3 {
		return findMax(result)
	}

	return min
}

func same(result []int, target int) bool {
	for _, n := range result {
		if n == target {
			return true
		}
	}
	return false
}

func findMax(result []int) int {
	max := result[0]
	for _, n := range result {
		if n > max {
			max = n
		}
	}
	return max
}

func findMinIndex(result []int) int {
	index := 0
	min := result[index]

	for i, n := range result {
		if n < min {
			index = i
			min = n
		}
	}
	return index
}

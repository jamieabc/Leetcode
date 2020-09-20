package main

//An integer has sequential digits if and only if each digit in the number is one more than the previous digit.
//
//Return a sorted list of all the integers in the range [low, high] inclusive that have sequential digits.
//
//
//
//Example 1:
//
//Input: low = 100, high = 300
//Output: [123,234]
//Example 2:
//
//Input: low = 1000, high = 13000
//Output: [1234,2345,3456,4567,5678,6789,12345]
//
//
//Constraints:
//
//10 <= low <= high <= 10^9

func sequentialDigits(low int, high int) []int {
	nums := make([]int, 0)
	num := low
	for ; num > 0; num /= 10 {
		nums = append(nums, num%10)
	}

	// check if initial low can find a valid number
	result := make([]int, 0)
	if nums[len(nums)-1] <= 10-len(nums) {
		num = 0
		for i, next := 0, nums[len(nums)-1]; i < len(nums); i, next = i+1, next+1 {
			num *= 10
			num += next
		}

		if num >= low && num <= high {
			result = append(result, num)
		}
	}

	find(nums, high, &result)

	return result
}

func find(nums []int, high int, result *[]int) {
	if nums[len(nums)-1] >= 10-len(nums) {
		nums = append(nums, 1)
	} else {
		nums[len(nums)-1]++
	}

	// generate next valid number
	for i := len(nums) - 2; i >= 0; i-- {
		nums[i] = nums[i+1] + 1
	}

	// convert to number
	var num int
	for i := len(nums) - 1; i >= 0; i-- {
		num *= 10
		num += nums[i]
	}

	if num <= high {
		*result = append(*result, num)
		find(nums, high, result)
	}
}

//	Notes
//	1.	inspired from solution, sequential numbers are very limited, one way to
//		do it is pre-computed

//	2.	inspired from solution, sequential number must be part of order in
//		123456789

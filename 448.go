package main

// Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
//
// Find all the elements of [1, n] inclusive that do not appear in this array.
//
// Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
//
// Example:
//
// Input:
// [4,3,2,7,8,2,3,1]
//
// Output:
// [5,6]

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		val := nums[i]
		if val < 0 {
			val = -val
		}

		if nums[val-1] > 0 {
			nums[val-1] *= -1
		}
	}

	missing := make([]int, 0)
	for i := range nums {
		if nums[i] > 0 {
			missing = append(missing, i+1)
		}
	}

	return missing
}

func findDisappearedNumbers1(nums []int) []int {
	size := len(nums)

	for i := 0; i < size; {
		// index is the only unique
		if i != nums[i]-1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	missing := make([]int, 0)
	for i := range nums {
		if i != nums[i]-1 {
			missing = append(missing, i+1)
		}
	}

	return missing
}

//	problems
//	1.	inspired from solution
//		use index as an unique value, negate nums[index], for numbers still
//		positive, it's the solution

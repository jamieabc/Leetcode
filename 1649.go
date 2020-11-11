package main

// Given an integer array instructions, you are asked to create a sorted array from the elements in instructions. You start with an empty container nums. For each element from left to right in instructions, insert it into nums. The cost of each insertion is the minimum of the following:
//
//     The number of elements currently in nums that are strictly less than instructions[i].
//     The number of elements currently in nums that are strictly greater than instructions[i].
//
// For example, if inserting element 3 into nums = [1,2,3,5], the cost of insertion is min(2, 1) (elements 1 and 2 are less than 3, element 5 is greater than 3) and nums will become [1,2,3,3,5].
//
// Return the total cost to insert all elements from instructions into nums. Since the answer may be large, return it modulo 109 + 7
//
//
//
// Example 1:
//
// Input: instructions = [1,5,6,2]
// Output: 1
// Explanation: Begin with nums = [].
// Insert 1 with cost min(0, 0) = 0, now nums = [1].
// Insert 5 with cost min(1, 0) = 0, now nums = [1,5].
// Insert 6 with cost min(2, 0) = 0, now nums = [1,5,6].
// Insert 2 with cost min(1, 2) = 1, now nums = [1,2,5,6].
// The total cost is 0 + 0 + 0 + 1 = 1.
//
// Example 2:
//
// Input: instructions = [1,2,3,6,5,4]
// Output: 3
// Explanation: Begin with nums = [].
// Insert 1 with cost min(0, 0) = 0, now nums = [1].
// Insert 2 with cost min(1, 0) = 0, now nums = [1,2].
// Insert 3 with cost min(2, 0) = 0, now nums = [1,2,3].
// Insert 6 with cost min(3, 0) = 0, now nums = [1,2,3,6].
// Insert 5 with cost min(3, 1) = 1, now nums = [1,2,3,5,6].
// Insert 4 with cost min(3, 2) = 2, now nums = [1,2,3,4,5,6].
// The total cost is 0 + 0 + 0 + 0 + 1 + 2 = 3.
//
// Example 3:
//
// Input: instructions = [1,3,3,3,2,4,2,1,2]
// Output: 4
// Explanation: Begin with nums = [].
// Insert 1 with cost min(0, 0) = 0, now nums = [1].
// Insert 3 with cost min(1, 0) = 0, now nums = [1,3].
// Insert 3 with cost min(1, 0) = 0, now nums = [1,3,3].
// Insert 3 with cost min(1, 0) = 0, now nums = [1,3,3,3].
// Insert 2 with cost min(1, 3) = 1, now nums = [1,2,3,3,3].
// Insert 4 with cost min(5, 0) = 0, now nums = [1,2,3,3,3,4].
// ​​​​​​​Insert 2 with cost min(1, 4) = 1, now nums = [1,2,2,3,3,3,4].
// ​​​​​​​Insert 1 with cost min(0, 6) = 0, now nums = [1,1,2,2,3,3,3,4].
// ​​​​​​​Insert 2 with cost min(2, 4) = 2, now nums = [1,1,2,2,2,3,3,3,4].
// The total cost is 0 + 0 + 0 + 0 + 1 + 0 + 1 + 0 + 2 = 4.
//
//
//
// Constraints:
//
//     1 <= instructions.length <= 105
//     1 <= instructions[i] <= 105

// tc: average O(n log(n))
func createSortedArray(instructions []int) int {
	nums := make([]int, 0)
	used := make(map[int]int)
	mod := int(1e9 + 7)
	var cost int

	for _, i := range instructions {
		nums = append(nums, i)

		pivot := nums[len(nums)-1]
		store := 0

		for j := 0; j < len(nums); j++ {
			if nums[j] <= pivot {
				nums[j], nums[store] = nums[store], nums[j]
				store++
			}
		}

		cost += min(store-1-used[i], len(nums)-store)
		used[i]++

		if cost >= mod {
			cost -= mod
		}
	}

	return cost
}

// tc: O(n log(n)) if memory operation count as constant
// tc: O(n (log(n) + n)) if memory operation cost O(n)
func createSortedArray1(instructions []int) int {
	var cost int
	nums := make([]int, 0)

	for _, i := range instructions {
		if len(nums) == 0 || i >= nums[len(nums)-1] {
			nums = append(nums, i)
		} else if i <= nums[0] {
			nums = append([]int{i}, nums...)
		} else {
			cost += insert(&nums, i)
		}
	}

	return cost
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func insert(nums *[]int, i int) int {
	lower, higher := len(*nums)-1, 0

	for low, high := 0, len(*nums)-1; low <= high; {
		mid := low + (high-low)>>1

		if (*nums)[mid] == i {
			// find left & right boundary
			for lower = mid; lower >= 0; lower-- {
				if (*nums)[lower] != i {
					break
				}
			}

			for higher = mid; higher < len(*nums); higher++ {
				if (*nums)[higher] != i {
					break
				}
			}

			break
		} else if (*nums)[mid] > i {
			high = mid - 1
			higher = mid
		} else {
			lower = mid
			low = mid + 1
		}
	}

	cost := min(lower+1, len(*nums)-higher)

	tmp := append([]int{}, (*nums)[:lower+1]...)
	tmp = append(tmp, i)
	tmp = append(tmp, (*nums)[lower+1:]...)

	*nums = tmp
	return cost
}

//	Notes
//	1.	similar to problem 315, but this one needs to do twice of sorting, one
//		ascending and the other descending. ascending find number smaller on right,
//		descending find number larger on left

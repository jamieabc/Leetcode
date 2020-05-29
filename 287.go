package main

// Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.
//
// Example 1:
//
// Input: [1,3,4,2,2]
// Output: 2
//
// Example 2:
//
// Input: [3,1,3,4,2]
// Output: 3
//
// Note:
//
//     You must not modify the array (assume the array is read only).
//     You must use only constant, O(1) extra space.
//     Your runtime complexity should be less than O(n2).
//     There is only one duplicate number in the array, but it could be repeated more than once.

func findDuplicate(nums []int) int {
	length := len(nums)

	if length <= 1 {
		return -1
	}

	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	for slow = 0; slow != fast; {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func findDuplicate1(nums []int) int {
	length := len(nums)

	var i, j, count int
	for i, j = 1, length-1; i < j; {
		mid := i + (j-i)/2
		count = 0
		for k := range nums {
			if nums[k] <= mid {
				count++
			}
		}

		if count > mid {
			j = mid
		} else {
			i = mid + 1
		}
	}

	return i
}

//	problems
//	1.	add reference https://leetcode.com/problems/find-the-duplicate-number/discuss/72846/My-easy-understood-solution-with-O(n)-time-and-O(1)-space-without-modifying-the-array.-With-clear-explanation.

//		it's a brilliant solution, use slow & fast pointer, if there's a loop
//		eventually two pointers meet. the reason they meet is because
//		duplicates exist, to find that duplicate.

//		when 2 ptrs meet, from calculation that distance from start point
//		to loop == meet to loop, so it's easy to put either fast/slow back
//		to original and go each step until they meeet

//		time complexity is O(n)

package main

// Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.
//
// Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
//
// Note: You are not suppose to use the library's sort function for this problem.
//
// Example:
//
// Input: [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]
//
// Follow up:
//
//     A rather straight forward solution is a two-pass algorithm using counting sort.
//     First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
//     Could you come up with a one-pass algorithm using only constant space?

func sortColors(nums []int) {
	size := len(nums)
	for low, high, i := 0, size-1, 0; i <= high; {
		if nums[i] == 0 && i > low {
			nums[low], nums[i] = nums[i], nums[low]
			low++
		} else if nums[i] == 2 && i < high {
			nums[i], nums[high] = nums[high], nums[i]
			high--
		} else {
			i++
		}
	}
}

func sortColors1(nums []int) {
	size := len(nums)
	var low, high int

	// find next place to store 0
	for low < size && nums[low] == 0 {
		low++
	}

	// find next place to store 2
	for high = size - 1; high >= 0 && nums[high] == 2; {
		high--
	}

	// find next 0/2 to swap
	for j := low; j <= high; {
		if nums[j] == 0 && j != low {
			nums[low], nums[j] = nums[j], nums[low]
			low++
		} else if nums[j] == 2 && j != high {
			nums[high], nums[j] = nums[j], nums[high]
			high--
		} else {
			j++
		}
	}
}

//	problems
//	1.	didn't think of bucket sort, iterate through array to find each color
//		count, then fill into it

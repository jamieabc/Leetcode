package main

import "math"

// You are given an integer array nums and an integer x. In one operation, you can either remove the leftmost or the rightmost element from the array nums and subtract its value from x. Note that this modifies the array for future operations.
//
// Return the minimum number of operations to reduce x to exactly 0 if it's possible, otherwise, return -1.
//
//
//
// Example 1:
//
// Input: nums = [1,1,4,2,3], x = 5
// Output: 2
// Explanation: The optimal solution is to remove the last two elements to reduce x to zero.
//
// Example 2:
//
// Input: nums = [5,6,7,8,9], x = 4
// Output: -1
//
// Example 3:
//
// Input: nums = [3,2,20,1,1,3], x = 10
// Output: 5
// Explanation: The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.
//
//
//
// Constraints:
//
//     1 <= nums.length <= 105
//     1 <= nums[i] <= 104
//     1 <= x <= 109

func minOperations(nums []int, x int) int {
	size := len(nums)

	var sum int
	for _, num := range nums {
		sum += num
	}

	// need whole array
	if sum == x {
		return size
	}

	// assumes l-r is a range to deduct, which means sum should at least >= x
	if sum < x {
		return -1
	}

	shortest := math.MaxInt32
	for left, right := 0, 0; right < size; {
		sum -= nums[right]
		right++

		for left < right && sum < x {
			sum += nums[left]
			left++
		}

		if sum == x {
			shortest = min(shortest, size-(right-left))
		}
		//         for sum > x && right < size {
		//             sum -= nums[right]
		//             right++
		//         }

		//         if sum == x {
		//             shortest = min(shortest, size-(right - left))
		//         }

		//         sum += nums[left]
		//         left++
	}

	if shortest == math.MaxInt32 {
		return -1
	}
	return shortest
}

func minOperations2(nums []int, x int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	target := sum - x
	size := len(nums)

	// because sum is always increasing, total sum equals x means whole array are needed
	if target == 0 {
		return size
	}

	sum = 0
	var longest int

	for left, right := 0, 0; right < size; {
		sum += nums[right]
		right++

		for left < right && sum > target {
			sum -= nums[left]
			left++
		}

		if sum == target {
			longest = max(longest, right-left)
		}
	}

	if longest == 0 {
		return -1
	}
	return size - longest
}

func minOperations1(nums []int, x int) int {
	size := len(nums)
	rightSum := make([]int, size)
	for i := size - 1; i >= 0; i-- {
		rightSum[i] = nums[i] + rightSum[i+1]
	}

	var left, right int
	minLength := math.MaxInt32

	for ; left < size; left++ {
		if x <= 0 {
			break
		}
		x -= nums[left]
	}

	// not possible
	if x < 0 && left == size {
		return -1
	}

	if x == 0 {
		minLength = left
	}

	prev := size - 1

	for left--; left >= 0; left-- {
		x += nums[left]

		tmp := x

		for right = prev; right > max(0, left) && rightSum[right] < tmp; right-- {
		}

		if tmp == rightSum[right] {
			minLength = min(minLength, max(0, left)+size-right)
		}

		prev = right
	}

	if minLength == math.MaxInt32 {
		return -1
	}

	return minLength
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	during contest, I spend for an hour and tried for 6 times to get accepted.
//		I know this is a problem related to continuous sum of subarray, but I
//		didn't find a efficient way to solve it in the first time.

//	2.	inspired from solution, shortest subarray sum to x, it's similar to find
//		longest subarray sums to total_sum - x

//		because all numbers >= 1, which means total sum is always increasing, this
//		can reduce time complexity down to O(n) by using two pointers

//	3.	by directly calculating subarray sum, it's still possible to solve it in
//		time complexity O(n)

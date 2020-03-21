package main

//Given an array, rotate the array to the right by k steps, where k is non-negative.
//
//Example 1:
//
//Input: [1,2,3,4,5,6,7] and k = 3
//Output: [5,6,7,1,2,3,4]
//Explanation:
//rotate 1 steps to the right: [7,1,2,3,4,5,6]
//rotate 2 steps to the right: [6,7,1,2,3,4,5]
//rotate 3 steps to the right: [5,6,7,1,2,3,4]
//
//Example 2:
//
//Input: [-1,-100,3,99] and k = 2
//Output: [3,99,-1,-100]
//Explanation:
//rotate 1 steps to the right: [99,-1,-100,3]
//rotate 2 steps to the right: [3,99,-1,-100]
//
//Note:
//
//    Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
//    Could you do it in-place with O(1) extra space?

func rotate(nums []int, k int) {
	length := len(nums)
	if length == 0 {
		return
	}

	k = k % length
	if k == 0 {
		return
	}

	var count, start, prev, current, next int
	for start, count = 0, 0; count < length; start++ {
		current = start
		prev = nums[current]
		for next = -1; next != start; {
			next = (current + k) % length
			prev, nums[next] = nums[next], prev
			current = next
			count++
		}
	}
}

// [-1, -100, 3, 99] 2
// [-1, -2, -3, -4, -5, -6] 2 it's gonna cycle
// [1, 2, 3, 4, 5, 6, 7] 3
func rotate3(nums []int, k int) {
	length := len(nums)

	if length <= 1 {
		return
	}

	if k == 0 {
		return
	}

	k = k % length
	tmp := make([]int, k)
	for i := 0; i < k; i++ {
		tmp[i] = nums[length-k+i]
	}
	for i := length - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
}

func rotate2(nums []int, k int) {
	length := len(nums)

	if length <= 1 {
		return
	}

	k = k % length

	newNum := make([]int, len(nums))
	for i := 0; i < length; i++ {
		newNum[(i+k)%length] = nums[i]
	}

	for i := 0; i < length; i++ {
		nums[i] = newNum[i]
	}
}

func rotate1(nums []int, k int) {
	length := len(nums)

	if length <= 1 {
		return
	}

	k = k % length

	var i, prev, tmp int
	for ; k > 0; k-- {
		prev = nums[0]
		for i = 1; i < length-1; i++ {
			tmp = nums[i]
			nums[i] = prev
			prev = tmp
		}
		tmp = nums[i]
		nums[i] = prev
		nums[0] = tmp
	}
}

// problems
//	1.	length could be 0
//	2.	when k is 1, it's special case, it's behavior acts like mod == 1
//	3.	wrong condition, checking should be mod != 0, not mod == 1
//	4.	original algorithm is totally wrong, it cannot be separated into
//		fully divided, re-think
//		stop criteria is that every number should be moved at least once
//		if start index is repeated, proceed to next one

package main

//  Given a circular array (the next element of the last element is the first element of the array), print the Next Greater Number for every element. The Next Greater Number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, output -1 for this number.
//
// Example 1:
//
// Input: [1,2,1]
// Output: [2,-1,2]
// Explanation: The first 1's next greater number is 2;
// The number 2 can't find next greater number;
// The second 1's next greater number needs to search circularly, which is also 2.
//
// Note: The length of given array won't exceed 10000.

func nextGreaterElements(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	stack := make([]int, length)

	idx := -1
	for i := range nums {
		for idx > -1 {
			if nums[i] <= nums[stack[idx]] {
				break
			}
			result[stack[idx]] = nums[i]
			idx--
		}
		idx++
		stack[idx] = i
	}

	for i := range nums {
		for idx > -1 {
			if nums[i] <= nums[stack[idx]] {
				break
			}
			result[stack[idx]] = nums[i]
			idx--
		}
	}

	for idx > -1 {
		if result[stack[idx]] == 0 {
			result[stack[idx]] = -1
		}
		idx--
	}

	return result
}

func nextGreaterElements1(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	for i := range result {
		result[i] = -1
	}
	stack := make([]int, 0)

	for i := 0; i < length*2; i++ {
		for len(stack) > 0 {
			s := stack[len(stack)-1]
			if nums[i%length] <= nums[s] {
				break
			} else {
				result[s] = nums[i%length]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, i%length)
	}

	return result
}

//	problems
//	1.	add reference https://leetcode.com/problems/next-greater-element-ii/discuss/98270/JavaC%2B%2BPython-Loop-Twice

//		there's guarantee to have a maximum number in array (could be
//		multiple), when first pass looping, numbers in stack are those
//		could find larger number in normal order. loop second time will
//		eliminate numbers that are not maximum. but for the truthy maximum
//		number, it will always stay in stack, so initialise default to -1

//		because second loop statement is same as first loop, so 2 loops can
//		be merged into 1

//	2.	inspired from sample code, after 2 loops, numbers remain in stack
//		are maximum, so it's not needed to initialise default to -1

//		also, author uses an index to indicate stack operation, avoid
//		additional memory operation

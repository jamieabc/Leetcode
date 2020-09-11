package main

import "math"

//Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.
//
//Example 1:
//
//Input: [2,3,-2,4]
//Output: 6
//Explanation: [2,3] has the largest product 6.
//
//Example 2:
//
//Input: [-2,0,-1]
//Output: 0
//Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

func maxProduct(nums []int) int {
	maxProduct := math.MinInt32
	maxSoFar, minSoFar := 1, 1

	for _, n := range nums {
		if n == 0 {
			maxSoFar, minSoFar = 1, 1
			maxProduct = max(maxProduct, 0)
		} else if n > 0 {
			maxSoFar *= n
			maxProduct = max(maxProduct, maxSoFar)
			minSoFar = min(1, minSoFar*n)
		} else {
			tmp := minSoFar * n
			maxProduct = max(maxProduct, tmp)
			minSoFar = min(n, maxSoFar*n)
			maxSoFar = max(1, tmp)
		}
	}

	return maxProduct
}

func maxProduct1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	maxToCurrent := nums[0]
	minToCurrent := nums[0]
	result := nums[0]
	var tmp int

	// divide every number in 2 categories, total max to current number,
	// total min to current number, the max result could come from
	// current number, prev total max * current number,
	// prev total min * current number
	for i := 1; i < length; i++ {
		tmp = maxToCurrent
		maxToCurrent = max(nums[i], max(maxToCurrent*nums[i], minToCurrent*nums[i]))
		minToCurrent = min(nums[i], min(tmp*nums[i], minToCurrent*nums[i]))
		if maxToCurrent > result {
			result = maxToCurrent
		}
	}

	return result
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	when multiply, need 2 information, if current number is included, max
//		product could come from following current number conditions:
//		- positive: current number * all previous positive maximum product
//		- negative: current number * all previous negative minimum product

//		basic thinking: include this number or not:
//		if only use current number, max or min could only be previous with current
//		if current number not included, then product default to 1

//	2.	in order to have multiply number, set default positive & negative to 1,
//		but this number should not influence max product, in other words, if all
//		numbers are negative, max product should not be 1

//	3.	inspired from solution, variable namings are better: maxSoFar & minSoFar

//	4.	inspired from https://leetcode.com/problems/maximum-product-subarray/discuss/183483/JavaC%2B%2BPython-it-can-be-more-simple

//		lee provides a very brilliant solution, which comes from the fact that
//		max product must be continuous.

//		if there's no zero, odd negative numbers
//		_ _ _ _ _ -1 _ _ _ _ _ _ _ _ _
//      <-------->   <---------------->
//         A                B
//		max product comes from either A or B

//		if there's no zero, even negative numbers
//		_ _ _ _ _ -1 _ _ _ _ -1 _ _ _ _ _ _
//		<-------->  <------->  <----------->
//           A          B            C
//      <----------------------------------->
//					    D
//		max product comes from A, B, C, D

//		if there more than 1 odd negative numbers
//		_ _ _ _ _ -1 _ _ _ _ -1 _ _ _ _ _ _ -1 _ _ _ _ _ _
//		<-------->  <------->  <----------->  <----------->
//           A          B            C				D
//		<-------------------->
//				E
//					<-------------------------------------->
//										F
//		max product comes from A, B, C, D, E, F

//		overall, scan forward & backward, max product will always with in range

//		if encounter any zero, reset product to 1 and keep following rules

package main

//Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.
//
//Example:
//
//Given nums = [-2, 0, 3, -5, 2, -1]
//
//sumRange(0, 2) -> 1
//sumRange(2, 5) -> -1
//sumRange(0, 5) -> -3
//
//Note:
//
//    You may assume that the array does not change.
//    There are many calls to sumRange function.

type NumArray struct {
	accumulative []int
	length       int
}

func Constructor(nums []int) NumArray {
	length := len(nums)

	accumulative := make([]int, length)

	if length == 0 {
		return NumArray{length: length}
	}

	sum := 0
	for i := range nums {
		sum += nums[i]
		accumulative[i] = sum
	}

	return NumArray{
		accumulative: accumulative,
		length:       length + 1,
	}
}

// store first-half and second-half of sum, use it as a reference
func (this *NumArray) SumRange(i int, j int) int {
	if this.length == 0 || i >= this.length || j < 0 {
		return 0
	}

	// make sure i in range
	if i < 0 {
		i = 0
	}

	// make sure j is in range
	if j >= this.length {
		j = this.length - 1
	}

	if i == 0 {
		return this.accumulative[j]
	}

	return this.accumulative[j] - this.accumulative[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

// problems
// 1. forget about case that no element exists
// 2. forget to skip in constructor when array empty
// 3. too slow, use accumulative array, i to j is accumulative(j) - accumulative(i-1)

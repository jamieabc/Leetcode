package main

// Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.
//
//
//
// Example 1:
//
// Input: nums = [2,2,3,2]
// Output: 3
//
// Example 2:
//
// Input: nums = [0,1,0,1,0,1,99]
// Output: 99
//
//
//
// Constraints:
//
//     1 <= nums.length <= 3 * 104
//     -231 <= nums[i] <= 231 - 1
//     Each element in nums appears exactly three times except for one element which appears once.
//
//
//
// Follow up: Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

func singleNumber(nums []int) int {
	var ans, count int32

	for i := 31; i >= 0; i-- {
		count = 0

		for _, j := range nums {
			if j&(1<<i) > 0 {
				count = (count + 1) % 3
			}
		}

		if count == 1 {
			ans |= 1 << i
		}
	}

	return int(ans)
}

//	Notes
//	1.	didn't think of solution

//	2.	inspired from https://leetcode.com/problems/single-number-ii/discuss/43296/An-General-Way-to-Handle-All-this-sort-of-questions.

//		very interesting insight, the problem wants to design a counter that can
//		distinguish occurrence of bit by m or k times

//		if k = 3, need two bits, 2^2 = 4 can represent 4 states
//		in short, needs log(k) bits to represent k states

//	3.	inspired from https://leetcode.com/problems/single-number-ii/discuss/43297/Java-O(n)-easy-to-understand-solution-easily-extended-to-any-times-of-occurance

//		for 32 bits, iterate through all numbers, find 1's on each position,
//		when count = 3, reset to 0, thus the remain bit is that only bit

//		very brilliant and easy to understand solution

//		the other thing, (nums[j] >> i) & 1 == 1 to check if bit is set

//	5.	becareful about negative numbers (2's complement), for int go still
//		uses int64, which cause wrong number, need to specifically use int32

//	6.	add reference https://leetcode.com/problems/single-number-ii/discuss/43295/Detailed-explanation-and-generalization-of-the-bitwise-operation-method-for-single-numbers

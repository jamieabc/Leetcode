package main

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
//
//
//
// Example 1:
//
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
//
// Example 2:
//
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
//
//
//
// Constraints:
//
// 0 <= nums.length <= 104
// -109 <= nums[i] <= 109
//
//
// Follow up: Could you implement the O(n) solution?

func longestConsecutive(nums []int) int {
	table := make(map[int]bool)
	for _, n := range nums {
		table[n] = true
	}

	var longest int

	for _, n := range nums {
		if _, ok := table[n-1]; !ok {
			cur := 1

			for i := n + 1; table[i]; i++ {
				cur++
			}

			longest = max(longest, cur)
		}
	}

	return longest
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from solution, consecutive sequence has 2 ends, check n-1 exist
//		to only consider one side, which is very brilliant...

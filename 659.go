package main

// Given an array nums sorted in ascending order, return true if and only if you can split it into 1 or more subsequences such that each subsequence consists of consecutive integers and has length at least 3.
//
//
//
// Example 1:
//
// Input: [1,2,3,3,4,5]
// Output: True
// Explanation:
// You can split them into two consecutive subsequences :
// 1, 2, 3
// 3, 4, 5
//
// Example 2:
//
// Input: [1,2,3,3,4,4,5,5]
// Output: True
// Explanation:
// You can split them into two consecutive subsequences :
// 1, 2, 3, 4, 5
// 3, 4, 5
//
// Example 3:
//
// Input: [1,2,3,4,4,5]
// Output: False
//
//
//
// Constraints:
//
//     1 <= nums.length <= 10000

func isPossible(nums []int) bool {
	size := len(nums)
	minSize := 3
	if size < minSize {
		return false
	}

	// p1 - group of 1 consecutive number
	// p2 - group of 2 consecutive numbers
	// p3 - group of 3+ consecutive numbers
	var p1, p2, p3, count, prev int

	for i := 0; i < size; i++ {
		count = 1
		for i < size-1 && nums[i] == nums[i+1] {
			i, count = i+1, count+1
		}

		if nums[i] != prev+1 {
			// need to start new group, make sure previous groups
			// all meets criteria
			if p1 != 0 || p2 != 0 {
				return false
			}
			p1 = count
			p3 = 0
		} else {
			// not enough number to concat
			if count < p1+p2 {
				return false
			}

			// try to distribute all numbers to exiting groups
			p3 = p2 + min(p3, count-p1-p2)
			p2 = p1
			p1 = max(0, count-p2-p3)
		}
		prev = nums[i]
	}
	return p1 == 0 && p2 == 0
}

type group struct {
	max  int
	size int
}

func isPossible1(nums []int) bool {
	groups := make([]group, 0)

	for _, n := range nums {
		var found bool
		for i := len(groups) - 1; i >= 0; i-- {
			if groups[i].max == n-1 {
				groups[i].max = n
				groups[i].size++
				found = true
				break
			}
		}

		if !found {
			g := group{
				max:  n,
				size: 1,
			}
			groups = append(groups, g)
		}
	}

	// check each group
	for i := range groups {
		if groups[i].size < 3 {
			return false
		}
	}
	return true
}

//	problems
//	1.	inspired from https://leetcode.com/problems/split-array-into-consecutive-subsequences/discuss/106495/Java-O(n)-time-and-O(1)-space-solution-greedily-extending-shorter-subsequence

//		author uses a really clever way to store information. I actually
//		divide array into sub-arrays, but this is not necessary. When
//		distributing numbers, what I really doing is to put number into
//		1 consecutive sub-array, 2 consecutive sub-array, and 3+ consecutive
//		sub-array.

//		Once a number is put into 1 consecutive, it becomes 2 consecutive,
//		and 2 consecutive becomes 3 consecutive. But if there's too many
//		of the same number, extra count should be put into 1 consecutive.

//	2.	add reference https://leetcode.com/problems/split-array-into-consecutive-subsequences/discuss/106514/Python-Easy-Understand-Greedy

//		lee also provides a greedy algorithm to solve the problem, it's kind
//		of smart, but I didn't take time implement it.

//		basic idea is that every number is either in  some consecutive
//		order, or start of a sequence

//		another similar idea from https://leetcode.com/problems/split-array-into-consecutive-subsequences/discuss/106493/C%2B%2B-O(n)-solution-two-pass

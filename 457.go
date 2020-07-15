package main

// You are given a circular array nums of positive and negative integers. If a number k at an index is positive, then move forward k steps. Conversely, if it's negative (-k), move backward k steps. Since the array is circular, you may assume that the last element's next element is the first element, and the first element's previous element is the last element.
//
// Determine if there is a loop (or a cycle) in nums. A cycle must start and end at the same index and the cycle's length > 1. Furthermore, movements in a cycle must all follow a single direction. In other words, a cycle must not consist of both forward and backward movements.
//
//
//
// Example 1:
//
// Input: [2,-1,1,2,2]
// Output: true
// Explanation: There is a cycle, from index 0 -> 2 -> 3 -> 0. The cycle's length is 3.
// Example 2:
//
// Input: [-1,2]
// Output: false
// Explanation: The movement from index 1 -> 1 -> 1 ... is not a cycle, because the cycle's length is 1. By definition the cycle's length must be greater than 1.
// Example 3:
//
// Input: [-2,1,-1,-2,-2]
// Output: false
// Explanation: The movement from index 1 -> 2 -> 1 -> ... is not a cycle, because movement from index 1 -> 2 is a forward movement, but movement from index 2 -> 1 is a backward movement. All movements in a cycle must follow a single direction.
//
//
// Note:
//
// -1000 ≤ nums[i] ≤ 1000
// nums[i] ≠ 0
// 1 ≤ nums.length ≤ 5000
//
//
// Follow up:
//
// Could you solve it in O(n) time complexity and O(1) extra space complexity?

func circularArrayLoop(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	for i := range nums {
		if nums[i] == 0 {
			continue
		}

		var forward bool
		if nums[i] > 0 {
			forward = true
		}

		slow, fast := i, i
		for true {
			fast, slow = next(nums, next(nums, fast, forward), forward), next(nums, slow, forward)
			if fast == -1 || slow == -1 {
				break
			}

			if fast == slow {
				// check size
				size := 1
				slow = next(nums, slow, forward)

				for slow != fast {
					slow = next(nums, slow, forward)
					size++
				}

				if size > 1 {
					return true
				}
				break
			}
		}
	}

	return false
}

func next(nums []int, idx int, forward bool) int {
	if idx == -1 {
		return -1
	}

	i := idx + nums[idx]
	i = i % len(nums)

	if i < 0 {
		i += len(nums)
	}

	if (nums[i] > 0 && !forward) || (nums[i] < 0 && forward) {
		return -1
	}
	return i
}

//	problems
//	1.	need to check direction

//	2.	cycle size need to be careful check, not only does it has one jump,
//		but also jump size > 1

//	3.	inspired from https://leetcode.com/problems/circular-array-loop/discuss/94148/Java-SlowFast-Pointer-Solution

//		the problem is vague, I thought cycle length might means jump distance,
//		but it's not.

//		cycle could start form anywhere in the array, so cyclel size is just what
//		it means

//		also, to check same direction, just multiply them and get positive value

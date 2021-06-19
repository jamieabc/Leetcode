package main

// Given an integer array nums and two integers left and right, return the number of contiguous non-empty subarrays such that the value of the maximum array element in that subarray is in the range [left, right].

// The test cases are generated so that the answer will fit in a 32-bit integer.



// Example 1:

// Input: nums = [2,1,4,3], left = 2, right = 3
// Output: 3
// Explanation: There are three subarrays that meet the requirements: [2], [2, 1], [3].

// Example 2:

// Input: nums = [2,9,2,5,6], left = 2, right = 8
// Output: 7



// Constraints:

//     1 <= nums.length <= 105
//     0 <= nums[i] <= 109
//     0 <= left <= right <= 109

// tc: O(n)
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	var count, dist int
	n := len(nums)
	last := n // in case there are number in criteria at last

	for i := n-1; i >= 0; i-- {
		if nums[i] > right {
			dist = 0
			last = i
		} else if nums[i] >= left && nums[i] <= right {
			dist = last - i
		}

		count += dist
	}

	return count
}

// tc: O(n), sc: O(n)
func numSubarrayBoundedMax1(nums []int, left int, right int) int {
    n := len(nums)

    // in case array reaches end with intervals meet criteria, create
    // a sentinel at end
    if nums[n-1] <= right {
        nums = append(nums, right+1)
        n++
    }

    deque := make([]int, 0)
    var count, prev int

    for i := 0; i < n; i++ {
        if nums[i] >= left && nums[i] <= right {
            deque = append(deque, i)
        } else if nums[i] > right {
            for j := prev; j < i; j++ {
                if len(deque) == 0 {
                    break
                }

                if j < deque[0] {
                    count += i-deque[0]
                } else {
                    deque = deque[1:]
                    count += i-j
                }
            }
            prev = i+1
        }
    }

    return count
}

//	Notes
//	1.	takes 2 hours to think of tc O(n) & sc O(n) solution, could be improved
//
//		the idea is to count how many valid sub-arrays start from current index
//
//		e.g. [1, 2, 1, 3, 1, 4], range: 2 ~ 3
//
//		[1, 2, 1, 3, 1, 4]
//		 ^
//		 1 out of range, cannot start from self, total 4 sub-arrays:
//		 [1, 2], [1, 2, 1], [1, 2, 1, 3], [1, 2, 1, 3, 1]
//
//			^
//			2 in range, total 4 sub-arrays:
//			[2], [2, 1], [2, 1, 3], [2, 1, 3, 1]
//
//			   ^
//			   1 out of range, cannot start from self, total 2 sub-arrays:
//			   [1, 3], [1, 3, 1]
//
//				   ^
//				   3 in range, total 2 sub-arrays:
//				   [3], [3, 1]
//
//				      ^
//				      1 out of range, 0 sub-array
//
//		by above example, use a stack to store positions meets criteria,
//		when a out-of-range number occurs, scan interval to count valid
//		sub-arrays
//
//		tc: O(n), sc: O(n)

//	2.	based on previous observation, valid sub-array count is separted by
//		maximum number in range
//
//		e.g.
//		[1, 2, 1, 3, 1, 4]
//		 4  4  2  2
//
//		previous observation start from smallest index to largest index,
//		but if starts from largest index, it becomes linear solution w/o
//		extra space
//
//		[1, 2, 1, 3, 1, 4]
//					 ^
//					 not in range, add 0
//				  ^
//				  in range, distance to 4 is 2, add 2
//			   ^
//			   out of range, keep previous status, add 2
//			^
//			in range, update, distance to 4 is 4, add 4
//		 ^
//		 out of range, keep previous status, add 4

//	3.	inspired from https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/discuss/1278426/JS-Python-Java-C%2B%2B-or-Easy-Triangular-Number-Solution-w-Explanation
//
//		author provides very clear explanation

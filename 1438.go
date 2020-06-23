package main

// Given an array of integers nums and an integer limit, return the size of the longest non-empty subarray such that the absolute difference between any two elements of this subarray is less than or equal to limit.
//
//
//
// Example 1:
//
// Input: nums = [8,2,4,7], limit = 4
// Output: 2
// Explanation: All subarrays are:
// [8] with maximum absolute diff |8-8| = 0 <= 4.
// [8,2] with maximum absolute diff |8-2| = 6 > 4.
// [8,2,4] with maximum absolute diff |8-2| = 6 > 4.
// [8,2,4,7] with maximum absolute diff |8-2| = 6 > 4.
// [2] with maximum absolute diff |2-2| = 0 <= 4.
// [2,4] with maximum absolute diff |2-4| = 2 <= 4.
// [2,4,7] with maximum absolute diff |2-7| = 5 > 4.
// [4] with maximum absolute diff |4-4| = 0 <= 4.
// [4,7] with maximum absolute diff |4-7| = 3 <= 4.
// [7] with maximum absolute diff |7-7| = 0 <= 4.
// Therefore, the size of the longest subarray is 2.
//
// Example 2:
//
// Input: nums = [10,1,2,4,7,2], limit = 5
// Output: 4
// Explanation: The subarray [2,4,7,2] is the longest since the maximum absolute diff is |2-7| = 5 <= 5.
//
// Example 3:
//
// Input: nums = [4,2,2,2,4,4,2,2], limit = 0
// Output: 3
//
//
//
// Constraints:
//
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     0 <= limit <= 10^9

func longestSubarray(nums []int, limit int) int {
	var longest int
	maxDequeue, minDequeue := make([]int, 0), make([]int, 0)

	// [10,1,2,4,7,2]
	for left, right := 0, 0; right < len(nums); right++ {
		for len(maxDequeue) > 0 && maxDequeue[len(maxDequeue)-1] < nums[right] {
			maxDequeue = maxDequeue[:len(maxDequeue)-1]
		}

		for len(minDequeue) > 0 && minDequeue[len(minDequeue)-1] > nums[right] {
			minDequeue = minDequeue[:len(minDequeue)-1]
		}

		maxDequeue = append(maxDequeue, nums[right])
		minDequeue = append(minDequeue, nums[right])

		for maxDequeue[0]-minDequeue[0] > limit {
			if maxDequeue[0] == nums[left] {
				maxDequeue = maxDequeue[1:]
			}
			if minDequeue[0] == nums[left] {
				minDequeue = minDequeue[1:]
			}
			left++
		}

		longest = max(longest, right-left+1)
	}

	return longest
}

type MinPQ struct {
	Val   int
	Index int
}

type MinPQs []MinPQ

func (p MinPQs) Len() int           { return len(p) }
func (p MinPQs) Less(i, j int) bool { return p[i].Val < p[j].Val }
func (p MinPQs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p MinPQs) Peek() interface{} {
	if p.Len() > 0 {
		return p[0]
	}
	return nil
}
func (p *MinPQs) Push(x interface{}) { *p = append(*p, x.(MinPQ)) }
func (p *MinPQs) Pop() interface{} {
	popped := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return popped
}

// tc: O(n log n)
func longestSubarray1(nums []int, limit int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	minVals, maxVals := &MinPQs{}, &MinPQs{}
	heap.Init(minVals)
	heap.Init(maxVals)
	heap.Push(minVals, MinPQ{nums[0], 0})
	heap.Push(maxVals, MinPQ{-nums[0], 0})

	var longest, left, right int

	for left, right = 0, 0; left < size; {
		if right < size && (left == right || -maxVals.Peek().(MinPQ).Val-minVals.Peek().(MinPQ).Val <= limit) {
			longest = max(longest, right-left+1)
			right++

			if right < size {
				heap.Push(maxVals, MinPQ{-nums[right], right})
				heap.Push(minVals, MinPQ{nums[right], right})
			}
		} else {
			// pop items not in range
			for minVals.Len() > 0 {
				item := minVals.Peek().(MinPQ)

				if item.Index < left+1 {
					heap.Pop(minVals)
				} else {
					break
				}
			}

			for maxVals.Len() > 0 {
				item := maxVals.Peek().(MinPQ)

				if item.Index < left+1 {
					heap.Pop(maxVals)
				} else {
					break
				}
			}

			left++
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

//	problems
//	1.	too slow

//	2.	inspired from https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/discuss/609771/JavaC%2B%2BPython-Deques-O(N)

//		the O(n) solution uses two dequeue: one for maintaining increasing
//		sequence, the other for maintaining decreasing sequence.

//		the problem is about getting min/max in a range. I use pq is the
//		same idea.

//		however, I cannot understand why not updating longest using if,
//		through discussion, it has something relates to longest window
//		is always expanding instead of decreasing, but I cannot fully
//		understand the idea

//	3.	inspired from https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/discuss/609743/Java-Detailed-Explanation-Sliding-Window-Deque-O(N)

//		author has a more clear explanation

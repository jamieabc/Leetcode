package main

// Given an integer array nums and a positive integer k, return the most competitive subsequence of nums of size k.
//
// An array's subsequence is a resulting sequence obtained by erasing some (possibly zero) elements from the array.
//
// We define that a subsequence a is more competitive than a subsequence b (of the same length) if in the first position where a and b differ, subsequence a has a number less than the corresponding number in b. For example, [1,3,4] is more competitive than [1,3,5] because the first position they differ is at the final number, and 4 is less than 5.
//
//
//
// Example 1:
//
// Input: nums = [3,5,2,6], k = 2
// Output: [2,6]
// Explanation: Among the set of every possible subsequence: {[3,5], [3,2], [3,6], [5,2], [5,6], [2,6]}, [2,6] is the most competitive.
//
// Example 2:
//
// Input: nums = [2,4,3,3,5,4,9,6], k = 4
// Output: [2,3,3,4]
//
//
//
// Constraints:
//
//     1 <= nums.length <= 105
//     0 <= nums[i] <= 109
//     1 <= k <= nums.length

func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0)
	maxRemoved := len(nums) - k

	for i := range nums {
		for maxRemoved > 0 && len(stack) > 0 && stack[len(stack)-1] > nums[i] {
			stack = stack[:len(stack)-1]
			maxRemoved--
		}

		stack = append(stack, nums[i])
	}

	return stack[:k]
}

func mostCompetitive2(nums []int, k int) []int {
	size := len(nums)
	stack := make([]int, 0)
	toRemove := size - k

	for i := range nums {
		for len(stack) > 0 {
			if toRemove > 0 && stack[len(stack)-1] > nums[i] {
				toRemove--
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		stack = append(stack, nums[i])
	}

	return stack[:k]
}

type MinHeap [][]int

func (h MinHeap) Len() int { return len(h) }

// same number, smaller index first
func (h MinHeap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
}

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() []int   { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mostCompetitive1(nums []int, k int) []int {
	size := len(nums)

	h := &MinHeap{}
	heap.Init(h)

	// push size-k numbers into heap, make sure selecting
	// from smallest, and at least remaining numbers are
	// sufficient to form answers with k remaining numbers
	var i int
	for i = 0; i <= size-k; i++ {
		heap.Push(h, []int{nums[i], i})
	}

	ans := make([]int, 0)
	prev := -1

	for len(ans) < k {
		// remove numbers before selected ones
		// because smallest is selected, all previous
		// larger should be removed
		for h.Len() > 0 && h.Peek()[1] < prev {
			heap.Pop(h)
		}

		p := heap.Pop(h).([]int)

		if i < size {
			heap.Push(h, []int{nums[i], i})
			i++
		}

		ans = append(ans, p[0])
		prev = p[1]
	}

	return ans
}

//	Notes
//	1.	to meets rule, it's about finding non-decreasing subsequences, but
//		there's a problem if longest non-decreasing is insufficient long, what
//		to do. I use a recursive process to find next non-decreasing based on
//		previous search, but got TLE and stuck.

//	2.	inspired from https://www.youtube.com/watch?v=OTlusbuZX94

//		alex says something general about finding lexicographical order, always
//		choose smallest possible and can be solved by greedy technique because
//		as long as smallest is selected, numbers afterwards don't matter

//		this problem, is about finding lexicographical order about numbers, but
//		it might be longest non-decreasing numbers doesn't long enough, the way
//		to solve this is to use another variable to denote how many numbers can
//		be removed

//		e.g. nums = [2,4,3,3,5,4,9,6], k = 4
//		deque = [2, 4], toRemove = size - k = 8 - 4 = 4

//		because 2 is smallest now, and next number 3 <= 4, so 4 may be removed.
//		toRemove = 4, so 4 is removed

//		deque = [2, 3], toRemove = 3

//		deque = [2, 3, 3, 5], toRemove = 3
//		next number is 4 <= 5, and toRemove > 0, to 5 is removed

//		deque = [2, 3, 3, 4], toRemove = 2

//		deque = [2, 3, 3, 4, 6], toRemove = 1

//		another example nums = [8, 9, 5, 6, 7, 4], k = 3
//		toRemove = size - k = 6 - 3 = 3

//		deque = [8, 9], toRemove = 3
//		deque = [5], toRemove = 1

//		deque = [5, 6, 7], toRemove = 1

//		deque = [5, 6, 4], smallest lexicographical order

//	3.	inspired from https://discordapp.com/channels/612060087900438538/612685982411784245/782457644639453194

//		to find lexicographical order, select smallest as possible, but another
//		problem is: need to select at least k numbers. So, author preserves
//		size - k numbers to make sure there will be at least sufficient numbers
//		to choose.

//		then each time select smallest possible (heap) and remove cannot select
//		numbers

//		becareful, when values are same, sorting order should put smaller index
//		earlier

//	4.	inspired from solution, there are max numbers to be removed, make sure
//		stack length always larger than that value

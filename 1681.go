package main

import (
	"fmt"
	"math"
)

// You are given an integer array nums and an integer k. You are asked to distribute this array into k subsets of equal size such that there are no two equal elements in the same subset.
//
// A subset's incompatibility is the difference between the maximum and minimum elements in that array.
//
// Return the minimum possible sum of incompatibilities of the k subsets after distributing the array optimally, or return -1 if it is not possible.
//
// A subset is a group integers that appear in the array with no particular order.
//
//
//
// Example 1:
//
// Input: nums = [1,2,1,4], k = 2
// Output: 4
// Explanation: The optimal distribution of subsets is [1,2] and [1,4].
// The incompatibility is (2-1) + (4-1) = 4.
// Note that [1,1] and [2,4] would result in a smaller sum, but the first subset contains 2 equal elements.
//
// Example 2:
//
// Input: nums = [6,3,8,1,3,1,2,2], k = 4
// Output: 6
// Explanation: The optimal distribution of subsets is [1,2], [2,3], [6,8], and [1,3].
// The incompatibility is (2-1) + (3-2) + (8-6) + (3-1) = 6.
//
// Example 3:
//
// Input: nums = [5,3,3,6,3,3], k = 3
// Output: -1
// Explanation: It is impossible to distribute nums into 3 subsets where no two elements are equal in the same subset.
//
//
//
// Constraints:
//
//     1 <= k <= nums.length <= 16
//     nums.length is divisible by k
//     1 <= nums[i] <= nums.length

func minimumIncompatibility(nums []int, k int) int {
	size := len(nums)
	if size == k {
		return 0
	}

	// make sure no number appear more than k times
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
		if counter[n] > k {
			return -1
		}
	}

	maxMask := (1 << size) - 1
	used := make(map[int]int)

	val := dfs(nums, 0, maxMask, size/k, used)
	if val >= 9999 {
		return -1
	}

	return val
}

func dfs(nums []int, currentMask, maxMask, groupSize int, used map[int]int) int {
	if currentMask == maxMask {
		return 0
	}

	// used already computed value
	if val, ok := used[currentMask]; ok {
		return val
	}

	// find unique unused numbers
	uniq := make(map[int]bool)
	available := make([]int, 0)
	size := len(nums)
	for i := 0; i < size; i++ {
		if (1<<(size-1-i))&currentMask > 0 || uniq[nums[i]] {
			continue
		}
		uniq[nums[i]] = true

		available = append(available, i)
	}

	// insufficient to form a group
	if len(available) < groupSize {
		return 9999
	}

	// generate combinations from those numbers
	combs := findCombinations(available, groupSize)

	// count value
	incompatibility := 9999
	for _, c := range combs {
		nextMask := currentMask
		maxVal, minVal := math.MinInt32, math.MaxInt32

		for _, j := range c {
			maxVal = max(maxVal, nums[j])
			minVal = min(minVal, nums[j])
			nextMask |= 1 << (size - 1 - j)
		}

		incompatibility = min(incompatibility, maxVal-minVal+dfs(nums, nextMask, maxMask, groupSize, used))
	}
	used[currentMask] = incompatibility

	return incompatibility
}

type Comb struct {
	Idx  int
	Data []int
}

// BFS
func findCombinations(available []int, groupSize int) [][]int {
	ans := make([][]int, 0)
	stack := []Comb{
		{0, []int{}},
	}
	size := len(available)

	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		idx, data := s.Idx, s.Data

		if len(data) == groupSize {
			ans = append(ans, data)
			continue
		}

		for ; idx < size; idx++ {
			length := len(data)
			tmp := make([]int, length, length+1)
			copy(tmp, data)
			tmp = append(tmp, available[idx])
			stack = append(stack, Comb{
				Idx:  idx + 1,
				Data: tmp,
			})
		}
	}

	return ans
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
//	1.	use backtracking, but it's too slow....TLE

//	2.	inspired from https://www.youtube.com/watch?v=yZv5IUFt-Bo

//		alex uses bitmask dp to solve the problem.

//		first clue to this problem is that after selection group of numbers,
//		remaining numbers will share same minimum differences. And remaining
//		numbers is another sub-problem, which can be solved by recursion

//		e.g. nums = [6,3,8,1,3,1,2,2]
//			               <---> 1, 3, 1 are selected, remaining numbers are
//			[6, 3, 8, 2, 2] which will have same minimum diff, this implies a
//		dp solution

//		the other thing about bitmask is that, remaining numbers need to have
//		combinations, so it's best to use bitmask to denote combinations of
//		those numbers

//	3.	inspired from https://leetcode.com/problems/minimum-incompatibility/discuss/961969/Python-True-O(n*n*2n)-bit-dp-explained

//		author analogy this problem to TSP problem, a very good insight

//	4.	inspird from https://leetcode.com/problems/minimum-incompatibility/discuss/961731/Golang-and-JavaScript-Bitmask%2BDP%2BDFS

//		author provides detailed comments

//	5.	golang add to math.MaxInt32 won't cause any error, it's simply a value

//	6.	max/min value can only update cache if all combinations are computed

//	7.	TLE, it might because findCombinations written in recursion is not
//		efficient, change it to iteration

//		inspired from https://leetcode.com/problems/minimum-incompatibility/discuss/961731/Golang-and-JavaScript-Bitmask%2BDP%2BDFS

//		use iterative way to find combinations

//	8.	time complexity of this is similar to TSP problem

//		inspired from https://cs.stackexchange.com/questions/90149/analysis-of-time-complexity-of-travelling-salesman-problem

//		there could be n starting points, each starting points has 2^n combinations
//		to find min/min, and for group of numbers are selected takes n to find
//		incompatibility value

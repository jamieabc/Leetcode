package main

// A zero-indexed array A of length N contains all integers from 0 to N-1. Find and return the longest length of set S, where S[i] = {A[i], A[A[i]], A[A[A[i]]], ... } subjected to the rule below.
//
// Suppose the first element in S starts with the selection of element A[i] of index = i, the next element in S should be A[A[i]], and then A[A[A[i]]]… By that analogy, we stop adding right before a duplicate element occurs in S.
//
//
//
// Example 1:
//
// Input: A = [5,4,0,3,1,6,2]
// Output: 4
// Explanation:
// A[0] = 5, A[1] = 4, A[2] = 0, A[3] = 3, A[4] = 1, A[5] = 6, A[6] = 2.
//
// One of the longest S[K]:
// S[0] = {A[0], A[5], A[6], A[2]} = {5, 6, 2, 0}
//
//
//
// Note:
//
//     N is an integer within the range [1, 20,000].
//     The elements of A are all distinct.
//     Each element of A is an integer within the range [0, N-1].

func arrayNesting(nums []int) int {
	var ans int

	for i := range nums {
		idx := i
		var tmp int

		for nums[idx] != -1 {
			tmp++
			next := nums[idx]
			nums[idx] = -1
			idx = next
		}

		ans = max(ans, tmp)
	}

	return ans
}

func arrayNesting1(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	rank := make([]int, size)
	group := make([]int, size)
	for i := range group {
		group[i] = i
		rank[i] = 1
	}

	ans := 1

	for i := range nums {
		p1, p2 := find(group, i), find(group, nums[i])

		if p1 == p2 {
			continue
		}

		if rank[p1] >= p2 {
			rank[p1] += rank[p2]
			group[p2] = p1
			ans = max(ans, rank[p1])
		} else {
			rank[p2] += rank[p1]
			group[p1] = p2
			ans = max(ans, rank[p2])
		}
	}

	return ans
}

func find(groups []int, idx int) int {
	if groups[idx] != idx {
		groups[idx] = find(groups, groups[idx])
	}
	return groups[idx]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	becareful about boundary condition, ans will at least 1 if array not
//		empty

//	2.	cyclic sort could also work

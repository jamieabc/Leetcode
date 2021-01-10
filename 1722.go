package main

// You are given two integer arrays, source and target, both of length n. You are also given an array allowedSwaps where each allowedSwaps[i] = [ai, bi] indicates that you are allowed to swap the elements at index ai and index bi (0-indexed) of array source. Note that you can swap elements at a specific pair of indices multiple times and in any order.
//
// The Hamming distance of two arrays of the same length, source and target, is the number of positions where the elements are different. Formally, it is the number of indices i for 0 <= i <= n-1 where source[i] != target[i] (0-indexed).
//
// Return the minimum Hamming distance of source and target after performing any amount of swap operations on array source.
//
//
//
// Example 1:
//
// Input: source = [1,2,3,4], target = [2,1,4,5], allowedSwaps = [[0,1],[2,3]]
// Output: 1
// Explanation: source can be transformed the following way:
// - Swap indices 0 and 1: source = [2,1,3,4]
// - Swap indices 2 and 3: source = [2,1,4,3]
// The Hamming distance of source and target is 1 as they differ in 1 position: index 3.
//
// Example 2:
//
// Input: source = [1,2,3,4], target = [1,3,2,4], allowedSwaps = []
// Output: 2
// Explanation: There are no allowed swaps.
// The Hamming distance of source and target is 2 as they differ in 2 positions: index 1 and index 2.
//
// Example 3:
//
// Input: source = [5,1,2,4,3], target = [1,5,4,2,3], allowedSwaps = [[0,4],[4,2],[1,3],[1,4]]
// Output: 0
//
//
//
// Constraints:
//
//     n == source.length == target.length
//     1 <= n <= 105
//     1 <= source[i], target[i] <= 105
//     0 <= allowedSwaps.length <= 105
//     allowedSwaps[i].length == 2
//     0 <= ai, bi <= n - 1
//     ai != bi

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	size := len(source)
	groups := make([]int, size)
	for i := range groups {
		groups[i] = i
	}

	for _, s := range allowedSwaps {
		p1, p2 := find(groups, s[0]), find(groups, s[1])

		if p1 != p2 {
			groups[p1] = p2
		}
	}

	// find number candidates for each group
	// map[group]: available numbers
	counter := make(map[int]map[int]int)
	for i := range source {
		group := find(groups, i)

		if _, ok := counter[group]; !ok {
			counter[group] = make(map[int]int)
		}
		counter[group][source[i]]++
	}

	var distance int
	for i := range target {
		group := find(groups, i)

		if counter[group][target[i]] > 0 {
			counter[group][target[i]]--
		} else {
			distance++
		}
	}

	return distance
}

func find(groups []int, idx int) int {
	if groups[idx] != idx {
		groups[idx] = find(groups, groups[idx])
	}

	return groups[idx]
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func minimumHammingDistance1(source []int, target []int, allowedSwaps [][]int) int {
	graph := make(map[int][]int)
	for _, s := range allowedSwaps {
		graph[s[0]] = append(graph[s[0]], s[1])
		graph[s[1]] = append(graph[s[1]], s[0])
	}

	size := len(source)
	visited := make([]bool, size)
	distance := size

	for i := range source {
		if visited[i] {
			continue
		}

		idxes := make([]int, 0)
		dfs(graph, i, visited, &idxes)

		// in case no groups
		if len(idxes) == 0 {
			if source[i] == target[i] {
				distance--
			}
			continue
		}

		group1 := findNums(idxes, source)
		group2 := findNums(idxes, target)

		for num, count1 := range group1 {
			count := min(count1, group2[num])

			if count > 0 {
				group1[num] -= count
				group2[num] -= count
				distance -= count
			}
		}
	}

	return distance
}

func dfs(graph map[int][]int, num int, visited []bool, nums *[]int) {
	for _, i := range graph[num] {
		if !visited[i] {
			visited[i] = true
			*nums = append(*nums, i)
			dfs(graph, i, visited, nums)
		}
	}
}

func findNums(idx []int, nums []int) map[int]int {
	counter := make(map[int]int)
	for _, i := range idx {
		counter[nums[i]]++
	}

	return counter
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func minimumHammingDistance2(source []int, target []int, allowedSwaps [][]int) int {
	visited := make(map[string]bool)
	minDist := distance(source, target)
	dfs(source, target, allowedSwaps, visited, &minDist)
	visited[numsToStr(source)] = true

	return minDist
}

func dfs(cur, target []int, swaps [][]int, visited map[string]bool, minDist *int) {
	visited[numsToStr(cur)] = true

	for _, s := range swaps {
		cur[s[0]], cur[s[1]] = cur[s[1]], cur[s[0]]

		d := distance(cur, target)
		str := numsToStr(cur)

		if _, ok := visited[str]; !ok {
			*minDist = min(*minDist, d)
			dfs(cur, target, swaps, visited, minDist)
		}

		cur[s[0]], cur[s[1]] = cur[s[1]], cur[s[0]]
	}
}

func numsToStr(cur []int) string {
	b := make([]byte, 2*len(cur))
	for i := range cur {
		b[2*i] = byte(cur[i]) + '0'
		b[2*i+1] = '-'
	}
	return string(b)
}

func distance(cur, target []int) int {
	var diff int

	for i := range cur {
		if cur[i] != target[i] {
			diff++
		}
	}

	return diff
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	use dfs w/ memo to find all possible combinations, since I actually swap
//		slice, it's really slow

//	2.	go map key can also be array, as long as content is boolean, numeric,
//		string, pointer, channel, interfaces

//	3.	inspired from https://leetcode.com/problems/minimize-hamming-distance-after-swap-operations/discuss/1009771/Java-Detailed-Explanation-Union-Find-%2B-Greedy

//		key idea: for swaps w/ infinite operations, indexes combined together means
//		all indexes are able to appear at any position

//		e.g. swaps = [0, 1], [1, 2]

//		array = [0, 1, 2]
//		[0, 1, 2] swap [0, 1]: [1, 0, 2]
//		[1, 0, 2] swap [1, 2]: [1, 2, 0]

//		[0, 1, 2] swap [1, 2]: [0, 2, 1]
//		[0, 2, 1] swap [0, 1]: [2, 0, 1]
//		[2, 0, 1] swap [1, 2]: [2, 1, 0]

//		all combinations of [0, 1, 2] are valid

//		this problem can be viewed as: some numbers belong to some group, find
//		differences of source group & target group

//	4.	to find group, use find because there might be uncompressed path exist

//	5.	inspired from https://leetcode.com/problems/minimize-hamming-distance-after-swap-operations/discuss/1009867/Python-DFS-Solution

//		lee builds a graph to represents each index relationships (groups)
//		use visited to check if any index not checked

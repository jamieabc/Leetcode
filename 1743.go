package main

// There is an integer array nums that consists of n unique elements, but you have forgotten it. However, you do remember every pair of adjacent elements in nums.
//
// You are given a 2D integer array adjacentPairs of size n - 1 where each adjacentPairs[i] = [ui, vi] indicates that the elements ui and vi are adjacent in nums.
//
// It is guaranteed that every adjacent pair of elements nums[i] and nums[i+1] will exist in adjacentPairs, either as [nums[i], nums[i+1]] or [nums[i+1], nums[i]]. The pairs can appear in any order.
//
// Return the original array nums. If there are multiple solutions, return any of them.
//
//
//
// Example 1:
//
// Input: adjacentPairs = [[2,1],[3,4],[3,2]]
// Output: [1,2,3,4]
// Explanation: This array has all its adjacent pairs in adjacentPairs.
// Notice that adjacentPairs[i] may not be in left-to-right order.
// Example 2:
//
// Input: adjacentPairs = [[4,-2],[1,4],[-3,1]]
// Output: [-2,4,1,-3]
// Explanation: There can be negative numbers.
// Another solution is [-3,1,4,-2], which would also be accepted.
// Example 3:
//
// Input: adjacentPairs = [[100000,-100000]]
// Output: [100000,-100000]
//
//
// Constraints:
//
// nums.length == n
// adjacentPairs.length == n - 1
// adjacentPairs[i].length == 2
// 2 <= n <= 105
// -105 <= nums[i], ui, vi <= 105
// There exists some nums that has adjacentPairs as its pairs.

func restoreArray(adjacentPairs [][]int) []int {
	graph := make(map[int][]int)
	for _, pair := range adjacentPairs {
		u, v := pair[0], pair[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var start int
	for k, v := range graph {
		if len(v) == 1 {
			start = k
			break
		}
	}

	ans := []int{start, graph[start][0]}

	for len(graph[ans[len(ans)-1]]) == 2 {
		for _, p := range graph[ans[len(ans)-1]] {
			if p != ans[len(ans)-2] {
				ans = append(ans, p)
				break
			}
		}
	}

	return ans
}

func restoreArray1(pairs [][]int) []int {
	graph := make(map[int][]int)
	for _, pair := range pairs {
		u, v := pair[0], pair[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// start from first pair, pair[0] - pair[1]
	tmp := []int{pairs[0][0], pairs[0][1]}
	for len(graph[tmp[len(tmp)-1]]) == 2 {
		for _, i := range graph[tmp[len(tmp)-1]] {
			if i != tmp[len(tmp)-2] {
				tmp = append(tmp, i)
				break
			}
		}
	}

	additional := len(pairs) + 1 - len(tmp)
	ans := make([]int, additional)
	ans = append(ans, tmp...)

	for idx := additional - 1; idx >= 0; idx-- {
		for _, i := range graph[ans[idx+1]] {
			if i != ans[idx+2] {
				ans[idx] = i
				break
			}
		}
	}

	return ans
}

//	Notes
//	1.	inspired from https://www.youtube.com/watch?v=IOLi6XbolQQ

//		alex write it in more elegant way~

//	2.	inspired from sample code, just use append for slice

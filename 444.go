package main

import "math"

// Check whether the original sequence org can be uniquely reconstructed from the sequences in seqs. The org sequence is a permutation of the integers from 1 to n, with 1 ≤ n ≤ 104. Reconstruction means building a shortest common supersequence of the sequences in seqs (i.e., a shortest sequence so that all sequences in seqs are subsequences of it). Determine whether there is only one sequence that can be reconstructed from seqs and it is the org sequence.
//
//
//
// Example 1:
//
// Input: org = [1,2,3], seqs = [[1,2],[1,3]]
// Output: false
// Explanation: [1,2,3] is not the only one sequence that can be reconstructed, because [1,3,2] is also a valid sequence that can be reconstructed.
//
// Example 2:
//
// Input: org = [1,2,3], seqs = [[1,2]]
// Output: false
// Explanation: The reconstructed sequence can only be [1,2].
//
// Example 3:
//
// Input: org = [1,2,3], seqs = [[1,2],[1,3],[2,3]]
// Output: true
// Explanation: The sequences [1,2], [1,3], and [2,3] can uniquely reconstruct the original sequence [1,2,3].
//
// Example 4:
//
// Input: org = [4,1,5,2,6,3], seqs = [[5,2,6,3],[4,1,5,2]]
// Output: true
//
//
//
// Constraints:
//
//     1 <= n <= 10^4
//     org is a permutation of {1,2,...,n}.
//     1 <= segs[i].length <= 10^5
//     seqs[i][j] fits in a 32-bit signed integer.
//
//
//
// UPDATE (2017/1/8):
// The seqs parameter had been changed to a list of list of strings (instead of a 2d array of strings). Please reload the code definition to get the latest changes.

// time limit exceed
func sequenceReconstruction(org []int, seqs [][]int) bool {
	// no edge, no order, no unique sequence
	if len(seqs) == 0 && len(org) > 0 {
		return false
	}

	graph, inDegree := buildGraph(seqs, len(org))

	sorted := topologicalSort(graph, inDegree)

	if len(sorted) != len(org) {
		return false
	}

	// compare each element
	for i := range sorted {
		if org[i] != sorted[i] {
			return false
		}
	}

	return true
}

func topologicalSort(graph map[int][]int, inDegree []int) []int {
	todo := make([]int, 0)
	for i := range inDegree {
		if inDegree[i] == 0 {
			todo = append(todo, i)
		}
	}

	// should have only 1 start point
	if len(todo) == 0 || len(todo) > 1 {
		return []int{}
	}

	result := make([]int, 0)

	for len(todo) > 0 {
		end := len(todo)
		var count int

		for i := 0; i < end; i++ {
			next := todo[0]
			todo = todo[1:]
			result = append(result, next+1)

			// update children in-degree
			for _, child := range graph[next] {
				inDegree[child]--

				if inDegree[child] == 0 {
					count++
					if count > 1 {
						return []int{}
					}

					todo = append(todo, child)
				}
			}
		}
	}

	return result
}

func buildGraph(seqs [][]int, size int) (map[int][]int, []int) {
	graph := make(map[int][]int)
	inDegree := make([]int, size)

	for _, seq := range seqs {
		// single char can only be the org
		if len(seq) == 1 && seq[0] > size {
			inDegree[0] = math.MaxInt32
			continue
		}

		for i := range seq {
			for j := i + 1; j < len(seq); j++ {
				graph[seq[i]-1] = append(graph[seq[i]-1], seq[j]-1)
				inDegree[seq[j]-1]++
			}
		}
	}

	return graph, inDegree
}

//	problems
//	1.	don't know how to do it, it seems like a graph problem that shows
//		connections between nodes

//		org can be validated by seqs, but unique means seqs should also
//		guarantee one solution, I don't know how to check that

//	2.	inspired from https://leetcode.com/problems/sequence-reconstruction/discuss/92574/Very-short-solution-with-explanation

//		to guarantee unique, two conditions are met:
//		- seqs should be part of org
//		- every consecutive number in org should be relations in seqs

//		second point is really important, with this insight, this this problem
//		can be solved

//	3.	seqs should not be same number

//	4.	check seqs in org, and consecutive numbers in org has correct order in
//		seqs

//	5.	from problem description, it already says org is permutation of 1-N,
//		guarantee uniqueness, so it's better to build up relationships in
//		org and check correctness in seq

//	6.	so many corner cases, finally TLE, not keep doing it

//	7. 	didn't write pass solution

//	8. 	need to check if sequence is valid by order

//	9.	inspired from https://leetcode.com/problems/sequence-reconstruction/discuss/92572/Simple-Solution-%3A-one-pass-using-only-array-(C%2B%2B-92ms-Java-16ms)

//		not implement it

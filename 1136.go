package main

// You are given an integer n which indicates that we have n courses, labeled from 1 to n. You are also given an array relations where relations[i] = [a, b], representing a prerequisite relationship between course a and course b: course a has to be studied before course b.
//
// In one semester, you can study any number of courses as long as you have studied all the prerequisites for the course you are studying.
//
// Return the minimum number of semesters needed to study all courses. If there is no way to study all the courses, return -1.
//
//
//
// Example 1:
//
// Input: n = 3, relations = [[1,3],[2,3]]
// Output: 2
// Explanation: In the first semester, courses 1 and 2 are studied. In the second semester, course 3 is studied.
//
// Example 2:
//
// Input: n = 3, relations = [[1,2],[2,3],[3,1]]
// Output: -1
// Explanation: No course can be studied because they depend on each other.
//
//
//
// Constraints:
//
// 1 <= n <= 5000
// 1 <= relations.length <= 5000
// 1 <= a, b <= n
// a != b
// All the pairs [a, b] are unique.

func minimumSemesters(n int, relations [][]int) int {
	pre := make([][]int, n+1)
	inDegree := make([]int, n+1)

	for _, r := range relations {
		inDegree[r[1]]++
		pre[r[0]] = append(pre[r[0]], r[1])
	}

	queue := make([]int, 0)
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	remain := n
	var ans int

	for len(queue) > 0 {
		ans++
		size := len(queue)

		for i := 0; i < size; i++ {
			remain--

			for _, to := range pre[queue[i]] {
				inDegree[to]--

				if inDegree[to] == 0 {
					queue = append(queue, to)
				}
			}
		}

		queue = queue[size:]
	}

	if remain == 0 {
		return ans
	}
	return -1
}

//	Notes
//	1.	topological, tc: O(v+e), not n^2, because every node will be visited
//		connected edge times

//	2.	inspired from https://leetcode.com/problems/parallel-courses/discuss/363145/google-follow-up-questions

//		google follow-up: everytime can choose at most k courses, minimum turns
//		to accomplish all

//	3.	inspired from solution, semesters to finish all courses is the longest
//		path in graph

//		but need to check if cycle exist

//		this is an interesting solution, it provides another view to the problem

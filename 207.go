package main

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.
//
// Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]
//
// Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?
//
//
//
// Example 1:
//
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: true
// Explanation: There are a total of 2 courses to take.
//              To take course 1 you should have finished course 0. So it is possible.
//
// Example 2:
//
// Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: false
// Explanation: There are a total of 2 courses to take.
//              To take course 1 you should have finished course 0, and to take course 0 you should
//              also have finished course 1. So it is impossible.
//
//
//
// Constraints:
//
//     The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
//     You may assume that there are no duplicate edges in the input prerequisites.
//     1 <= numCourses <= 10^5

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	routes := make(map[int][]int)

	// calculate in-degree & possible routes
	for _, pre := range prerequisites {
		inDegree[pre[0]]++
		routes[pre[1]] = append(routes[pre[1]], pre[0])
	}

	// find basic course able to learn right away
	toLearn := make([]int, 0)
	for i := range inDegree {
		if inDegree[i] == 0 {
			toLearn = append(toLearn, i)
		}
	}

	// find available courses
	courses := make([]int, 0)
	for len(toLearn) > 0 {
		next := toLearn[0]
		toLearn = toLearn[1:]

		courses = append(courses, next)

		// reduce in-degree of courses
		for _, n := range routes[next] {
			inDegree[n]--
			if inDegree[n] == 0 {
				toLearn = append(toLearn, n)
			}
		}
	}

	return len(courses) == numCourses
}

//	problems
//	1.	inspired from https://leetcode.com/problems/course-schedule/discuss/58509/C%2B%2B-BFSDFS

//		dfs keeps track of visited & todo, and cyclic exist means todo ==
//		false && done == true

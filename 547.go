package main

// There are N students in a class. Some of them are friends, while some are not. Their friendship is transitive in nature. For example, if A is a direct friend of B, and B is a direct friend of C, then A is an indirect friend of C. And we defined a friend circle is a group of students who are direct or indirect friends.
//
// Given a N*N matrix M representing the friend relationship between students in the class. If M[i][j] = 1, then the ith and jth students are direct friends with each other, otherwise not. And you have to output the total number of friend circles among all the students.
//
// Example 1:
//
// Input:
// [[1,1,0],
//  [1,1,0],
//  [0,0,1]]
// Output: 2
// Explanation:The 0th and 1st students are direct friends, so they are in a friend circle.
// The 2nd student himself is in a friend circle. So return 2.
// Example 2:
//
// Input:
// [[1,1,0],
//  [1,1,1],
//  [0,1,1]]
// Output: 1
// Explanation:The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends,
// so the 0th and 2nd students are indirect friends. All of them are in the same friend circle, so return 1.
// Note:
//
// N is in range [1,200].
// M[i][i] = 1 for all students.
// If M[i][j] = 1, then M[j][i] = 1.

// worst tc: O(n^3), n^2 for edges, and worst case of union-find is n
func findCircleNum(isConnected [][]int) int {
	size := len(isConnected)

	parents, ranks := make([]int, size), make([]int, size)
	for i := range parents {
		parents[i] = i
		ranks[i] = 1
	}

	disconnected := size

	for i, c := range isConnected {
		for j := range c {
			if c[j] == 1 && i != j {
				p1, p2 := find(parents, i), find(parents, j)

				if p1 != p2 {
					disconnected--

					if ranks[p1] >= ranks[p2] {
						parents[p2] = p1
						ranks[p1]++
					} else {
						parents[p1] = p2
						ranks[p2]++
					}
				}
			}
		}
	}

	// another way to find connected component, there is only one root
	// var count int
	// for i := range parents {
	// 	if parents[i] == i {
	// 		count++
	// 	}
	// }

	return disconnected
}

func find(parents []int, idx int) int {
	if parents[idx] != idx {
		parents[idx] = find(parents, parents[idx])
	}

	return parents[idx]
}

// tc: O(n^2), all edges are visited once
func findCircleNum1(isConnected [][]int) int {
	size := len(isConnected)
	visited := make([]bool, size)
	var count int

	for i := range visited {
		if visited[i] {
			continue
		}
		count++

		bfs(isConnected, i, visited)
	}

	return count
}

func bfs(isConnected [][]int, start int, visited []bool) {
	queue := []int{start}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if visited[n] {
			continue
		}
		visited[n] = true

		for i, v := range isConnected[n] {
			if i != n && v == 1 && !visited[i] {
				queue = append(queue, i)
			}
		}
	}
}

//	Notes
//	1.	description is quite vague, but if b = a's direct, c = b's direct, d = c's direct, then
//		a, b, c, d all forms a circle

//	2.	worst tc of union-find is O(n^3), traverse all edges with O(n^2), and
//		worst cast of union-find is O(n)

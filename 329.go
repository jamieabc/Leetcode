package main

// Given an m x n integers matrix, return the length of the longest increasing path in matrix.
//
// From each cell, you can either move in four directions: left, right, up, or down. You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).
//
//
//
// Example 1:
//
// Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
// Output: 4
// Explanation: The longest increasing path is [1, 2, 6, 9].
//
// Example 2:
//
// Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
// Output: 4
// Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
//
// Example 3:
//
// Input: matrix = [[1]]
// Output: 1
//
//
//
// Constraints:
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 200
// 0 <= matrix[i][j] <= 231 - 1

type dir int

const (
	north dir = iota
	south
	east
	west
)

var (
	dirs = map[dir][]int{
		north: []int{-1, 0},
		south: []int{1, 0},
		east:  []int{0, 1},
		west:  []int{0, -1},
	}
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	var maxPath int
	memo := make(map[int]map[int]int) // max path from x, y
	for i := range matrix[0] {
		memo[i] = make(map[int]int)
	}

	for i := range matrix {
		for j := range matrix[0] {
			maxPath = max(maxPath, dfs(matrix, j, i, memo))
		}
	}

	return maxPath
}

func dfs(matrix [][]int, x, y int, memo map[int]map[int]int) int {
	maxPath := 1

	for _, arr := range dirs {
		newX, newY := x+arr[1], y+arr[0]
		if newX >= 0 && newX < len(matrix[0]) && newY >= 0 && newY < len(matrix) && matrix[newY][newX] > matrix[y][x] {
			if _, ok := memo[newX][newY]; !ok {
				memo[newX][newY] = dfs(matrix, newX, newY, memo)
			}
			maxPath = max(maxPath, 1+memo[newX][newY])
		}
	}

	return maxPath
}

// view it as graph, v: mn, e: 4mn
// tc: O(v+e) = O(mn + 4mn) = O(mn)
func longestIncreasingPath2(matrix [][]int) int {
	w, h := len(matrix[0]), len(matrix)
	memo := make([][]int, h)
	for i := range memo {
		memo[i] = make([]int, w)
	}

	graph := buildGraph(matrix)
	var longest int

	for i := range matrix {
		for j := range matrix[0] {
			longest = max(longest, 1+dfs2(graph, i, j, memo))
		}
	}

	return longest
}

func dfs2(graph [][][][2]int, y, x int, memo [][]int) int {
	if memo[y][x] != 0 {
		return memo[y][x]
	}

	var longest int

	for _, to := range graph[y][x] {
		longest = max(longest, 1+dfs2(graph, to[0], to[1]))
	}

	memo[y][x] = longest

	return longest
}

func buildGraph(matrix [][]int) [][][][2]int {
	w, h := len(matrix[0]), len(matrix)

	graph := make([][][][2]int, h)
	for i := range graph {
		graph[i] = make([][][2]int, w)
	}

	for i := range matrix {
		for j := range matrix[0] {
			for _, dir := range dirs {
				newY, newX := i+dir[0], j+dir[1]

				if newX >= 0 && newX < w && newY >= 0 && newY < h && matrix[newY][newX] > matrix[i][j] {
					graph[i][j] = append(graph[i][j], [2]int{newY, newX})
				}
			}
		}
	}

	return graph
}

// tc: O(mn * 3^k), m, n: width, height, k: longest path
// TLE
func longestIncreasingPath1(matrix [][]int) int {
	w, h := len(matrix[0]), len(matrix)
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	var longest int

	for i := range matrix {
		for j := range matrix[0] {
			visited[i][j] = true

			longest = max(longest, 1+backtracking(matrix, visited, i, j))

			visited[i][j] = true
		}
	}

	return longest
}

func backtracking(matrix [][]int, visited [][]bool, y, x int) int {
	w, h := len(matrix[0]), len(matrix)

	var longest int
	for _, dir := range dirs {
		newY, newX := y+dir[0], x+dir[1]

		if newX >= 0 && newX < w && newY >= 0 && newY < h && matrix[newY][newX] > matrix[y][x] {
			visited[newY][newX] = true

			longest = max(longest, 1+backtracking(matrix, visited, newY, newX))

			visited[newY][newX] = false
		}
	}

	return longest
}

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from solution, dfs can be seen as graph, each vertex/cell are
//		visited once, total tc is O(V+E), V is all points O(mn), E is 4
//		directions for a vertex, O(4V) = O(mn)

//	2.	not implement dp

//	3.	inspired from https://leetcode.com/problems/longest-increasing-path-in-a-matrix/discuss/288520/Longest-Path-in-DAG

//		this is DAG, and the condition makes there's won't be cycle

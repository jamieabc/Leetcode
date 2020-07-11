package main

import "math"

type pathInfo struct {
	x, y, dist int
}

type dir struct {
	x, y int
}

var dirs = []dir{
	{0, -1},  // north
	{1, 0},   // east
	{0, 1},   // south
	{-1, 0},  // west
	{1, -1},  // north-east
	{1, 1},   // south-east
	{-1, 1},  // south-west
	{-1, -1}, // north-west
}

func shortestPathBinaryMatrix(grid [][]int) int {
	// not reachable
	if len(grid) == 0 || grid[0][0] == 1 || grid[len(grid)-1][len(grid[0])-1] == 1 {
		return -1
	}

	queue := []pathInfo{
		{0, 0, 1},
	}
	grid[0][0] = 1

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[:1]

		if s.x == len(grid[0])-1 && s.y == len(grid) {
			return s.dist
		}

		for _, d := range dirs {
			newX, newY := s.x+d.x, s.y+d.y

			if validAxis(grid, newX, newY) && grid[newY][newX] == 0 {
				grid[newY][newX] = 1
				queue = append(queue, pathInfo{newX, newY, s.dist + 1})
			}
		}
	}

	return -1
}

func validAxis(grid [][]int, x, y int) bool {
	return !(x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]))
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	problems
//	1.	be careful about repeated path

//	2.	when using BFS, since every traverse is proportional to distance, so
//		it can be guarantee that first encounter bottom-right point with shortest
//		path

//	3. inspired from https://leetcode.com/problems/shortest-path-in-binary-matrix/discuss/312827/Python-Concise-BFS

//		bfs

//	4.	inspired from https://leetcode.com/problems/shortest-path-in-binary-matrix/discuss/312785/why-does-DFS-not-work

//		when using dfs to traverse, it needs to try every possible combinations
//		another post about dfs not working to find shortest in unweighted graph
//		https://cs.stackexchange.com/questions/4914/why-cant-dfs-be-used-to-find-shortest-paths-in-unweighted-graphs

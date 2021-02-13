package main

// In an N by N square grid, each cell is either empty (0) or blocked (1).
//
// A clear path from top-left to bottom-right has length k if and only if it is composed of cells C_1, C_2, ..., C_k such that:
//
// Adjacent cells C_i and C_{i+1} are connected 8-directionally (ie., they are different and share an edge or corner)
// C_1 is at location (0, 0) (ie. has value grid[0][0])
// C_k is at location (N-1, N-1) (ie. has value grid[N-1][N-1])
// If C_i is located at (r, c), then grid[r][c] is empty (ie. grid[r][c] == 0).
// Return the length of the shortest such clear path from top-left to bottom-right.  If such a path does not exist, return -1.
//
//
//
// Example 1:
//
// Input: [[0,1],[1,0]]
//
//
// Output: 2
//
// Example 2:
//
// Input: [[0,0,0],[1,1,0],[1,1,0]]
//
//
// Output: 4
//
//
//
// Note:
//
// 1 <= grid.length == grid[0].length <= 100
// grid[r][c] is 0 or 1

import "math"

func shortestPathBinaryMatrix(grid [][]int) int {
	w, h := len(grid[0]), len(grid)

	if grid[0][0] == 1 || grid[h-1][w-1] == 1 {
		return -1
	}

	// dp[i][j]: shortest path to [i, j]
	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, w)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	queue := [][]int{{0, 0}}
	dp[0][0] = 1

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			p := queue[i]

			for _, j := range []int{-1, 0, 1} {
				for _, k := range []int{-1, 0, 1} {
					if j == 0 && k == 0 {
						continue
					}

					newY, newX := p[0]+j, p[1]+k

					if newX >= 0 && newY >= 0 && newX < w && newY < h && grid[newY][newX] == 0 {
						if dp[p[0]][p[1]]+1 < dp[newY][newX] {
							dp[newY][newX] = dp[p[0]][p[1]] + 1
							queue = append(queue, []int{newY, newX})
						}
					}
				}
			}
		}

		queue = queue[size:]
	}

	if dp[h-1][w-1] == math.MaxInt32 {
		return -1
	}

	return dp[h-1][w-1]
}

func shortestPathBinaryMatrix2(grid [][]int) int {
	w, h := len(grid[0]), len(grid)

	// becareful about boundary condition, start point must be 0
	if grid[0][0] == 1 || grid[h-1][w-1] == 1 {
		return -1
	}

	if w == 1 && h == 1 {
		if grid[0][0] == 0 {
			return 1
		}
		return -1
	}

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	queue := [][]int{{0, 0}}
	steps := 1

	// bfs
	for len(queue) > 0 {
		size := len(queue)

		steps++
		for i := 0; i < size; i++ {
			p := queue[i]

			for _, j := range []int{-1, 0, 1} {
				for _, k := range []int{-1, 0, 1} {
					if j == 0 && k == 0 {
						continue
					}

					newY, newX := p[0]+j, p[1]+k

					if newY == h-1 && newX == w-1 {
						return steps
					}

					if newX >= 0 && newY >= 0 && newX < w && newY < h && grid[newY][newX] == 0 && !visited[newY][newX] {
						visited[newY][newX] = true
						queue = append(queue, []int{newY, newX})
					}
				}
			}
		}

		queue = queue[size:]
	}

	return -1
}

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

func shortestPathBinaryMatrix1(grid [][]int) int {
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

//	Notes
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

//	5.	inspired from solution

//		there's another approach: A* algorithm, which tries to guess most possible point that
//		can reach to destination by counting distance to destination. However, this number
//		doesn't guarantee do reach destination, so it's an improvement of BFS

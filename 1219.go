package main

// In a gold mine grid of size m * n, each cell in this mine has an integer representing the amount of gold in that cell, 0 if it is empty.
//
// Return the maximum amount of gold you can collect under the conditions:
//
// Every time you are located in a cell you will collect all the gold in that cell.
// From your position you can walk one step to the left, right, up or down.
// You can't visit the same cell more than once.
// Never visit a cell with 0 gold.
// You can start and stop collecting gold from any position in the grid that has some gold.
//
//
//
// Example 1:
//
// Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
// Output: 24
// Explanation:
// [[0,6,0],
// [5,8,7],
// [0,9,0]]
// Path to get the maximum gold, 9 -> 8 -> 7.
//
// Example 2:
//
// Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
// Output: 28
// Explanation:
// [[1,0,7],
// [2,0,6],
// [3,4,5],
// [0,3,0],
// [9,0,20]]
// Path to get the maximum gold, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.
//
//
//
// Constraints:
//
// 1 <= grid.length, grid[i].length <= 15
// 0 <= grid[i][j] <= 100
// There are at most 25 cells containing gold.

// tc: O(4*3^k)
func getMaximumGold(grid [][]int) int {
	var ans int
	w, h := len(grid[0]), len(grid)

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] > 0 {
				visited := make([][]bool, h)
				for i := range visited {
					visited[i] = make([]bool, w)
				}
				visited[i][j] = true

				ans = max(ans, dfs(grid, visited, i, j))
			}
		}
	}

	return ans
}

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func dfs(grid [][]int, visited [][]bool, y, x int) int {
	w, h := len(grid[0]), len(grid)

	var cur int
	for _, dir := range dirs {
		newY, newX := y+dir[0], x+dir[1]

		if newX >= 0 && newY >= 0 && newX < w && newY < h && !visited[newY][newX] && grid[newY][newX] > 0 {
			visited[newY][newX] = true
			cur = max(cur, dfs(grid, visited, newY, newX))
			visited[newY][newX] = false
		}
	}

	return grid[y][x] + cur
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	even for same position, start by different direction results in different value,
//		so cannot really use memo to cache value because same position with direction came
//		from results in different result

//		this problem becomes backtracking problem, try all possibilities

//	2.	inspired from https://leetcode.com/problems/path-with-maximum-gold/discuss/398388/JavaC%2B%2BPython-DFS-Backtracking-Clean-code-O(4*3k)

//		tc: O(4 * 3^k), k: number of cells

//	3.	inspired from https://leetcode.com/problems/path-with-maximum-gold/discuss/398184/C%2B%2B-Short-DFS

//		author mark visited point as 0 to save some space, only if input is allowed to modify

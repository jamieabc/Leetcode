package main

import "fmt"

//Given a 2D grid consisting of 1s (land) and 0s (water).  An island is a maximal 4-directionally (horizontal or vertical) connected group of 1s.
//
//The grid is said to be connected if we have exactly one island, otherwise is said disconnected.
//
//In one day, we are allowed to change any single land cell (1) into a water cell (0).
//
//Return the minimum number of days to disconnect the grid.
//
//
//
//Example 1:
//
//Input: grid = [[0,1,1,0],[0,1,1,0],[0,0,0,0]]
//Output: 2
//Explanation: We need at least 2 days to get a disconnected grid.
//Change land grid[1][1] and grid[0][2] to water and get 2 disconnected island.
//
//Example 2:
//
//Input: grid = [[1,1]]
//Output: 2
//Explanation: Grid of full water is also disconnected ([[1,1]] -> [[0,0]]), 0 islands.
//
//Example 3:
//
//Input: grid = [[1,0,1,0]]
//Output: 0
//
//Example 4:
//
//Input: grid = [[1,1,0,1,1],
//               [1,1,1,1,1],
//               [1,1,0,1,1],
//               [1,1,0,1,1]]
//Output: 1
//
//Example 5:
//
//Input: grid = [[1,1,0,1,1],
//               [1,1,1,1,1],
//               [1,1,0,1,1],
//               [1,1,1,1,1]]
//Output: 2
//
//
//
//Constraints:
//
//    1 <= grid.length, grid[i].length <= 30
//    grid[i][j] is 0 or 1.

func main() {
	fmt.Println(minDays([][]int{{1, 1, 1, 0}}))
}

func minDays(grid [][]int) int {
	island := countIsland(grid)

	if island != 1 {
		return 0
	}

	// try mark any point as 0, then count again to see if island count would
	// become 0 or > 1

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				grid[i][j] = 0

				if countIsland(grid) != 1 {
					return 1
				}

				grid[i][j] = 1
			}
		}
	}

	return 2
}

var dir = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func countIsland(grid [][]int) int {
	queue := make([][]int, 0)
	var island int

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 && !visited[i][j] {
				island++

				if island != 1 {
					return island
				}

				queue = append(queue, []int{i, j})

				for len(queue) > 0 {
					x, y := queue[0][1], queue[0][0]
					visited[y][x] = true
					queue = queue[1:]

					for _, d := range dir {
						point := []int{y + d[0], x + d[1]}

						if isLand(grid, visited, point) {
							queue = append(queue, point)
						}
					}
				}
			}
		}
	}

	return island
}

func isLand(grid [][]int, visited [][]bool, point []int) bool {
	x, y := point[1], point[0]
	return x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid) && grid[y][x] == 1 && !visited[y][x]
}

//	Notes
//	1.	while thinking this problem, I did think of cost to separate an island
//		is 2, because for any point at corer, maximum cost to separate it is
//		to mark it's adjacent to 1

//		e.g. 1 1 1			1 0 1
//			 1 1 1		=>	0 1 1
//			 1 1 1			1 1 1

//		however, I didn't realize that this is the most important clue to solve
//		the problem, is it reduces solution space down to 0, 1, or 2 only!!!!

//		I realized this when I peeking others solutions...QQ, it's a pity I
//		didn't figure it out myself, but at least I have notice details, it
//		could be better next time!!

//		Always focusing on any clues that might reduce computation!!

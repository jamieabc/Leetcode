package main

// Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
//
// Count the number of distinct islands. An island is considered to be the same as another if and only if one island can be translated (and not rotated or reflected) to equal the other.
//
// Example 1:
//
// 11000
// 11000
// 00011
// 00011
//
// Given the above grid map, return 1.
//
// Example 2:
//
// 11011
// 10000
// 00001
// 11011
//
// Given the above grid map, return 3.
//
// Notice that:
//
// 11
// 1
//
// and
//
//  1
// 11
//
// are considered different island shapes, because we do not consider reflection / rotation.
//
// Note: The length of each dimension in the given grid does not exceed 50.

func numDistinctIslands(grid [][]int) int {
	unique := make(map[string]bool)
	w, h := len(grid[0]), len(grid)

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for i := range grid {
		for j := range grid[0] {

			if grid[i][j] == 1 {
				queue := [][]int{{i, j}}
				enc := []byte{byte('o')}
				grid[i][j] = 0

				for len(queue) > 0 {
					p := queue[0]
					queue = queue[1:]

					for _, d := range dirs {
						newY, newX := p[0]+d[0], p[1]+d[1]

						if newY >= 0 && newX >= 0 && newY < h && newX < w && grid[newY][newX] == 1 {
							grid[newY][newX] = 0
							enc = append(enc, byte('d'+newY-i), byte('f'+newX-j))
							queue = append(queue, []int{newY, newX})
						}
					}
				}

				unique[string(enc)] = true
			}
		}
	}

	return len(unique)
}

//	Notes
//	1.	use encoding for denote each region

//	2.	to make sequence is unique, every position should be relative to
//		origin

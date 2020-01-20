package main

//In a given grid, each cell can have one of three values:
//
//    the value 0 representing an empty cell;
//    the value 1 representing a fresh orange;
//    the value 2 representing a rotten orange.
//
//Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.
//
//Return the minimum number of minutes that must elapse until no cell has a fresh orange.  If this is impossible, return -1 instead.
//
//
//
//Example 1:
//
//Input: [[2,1,1],[1,1,0],[0,1,1]]
//Output: 4
//
//Example 2:
//
//Input: [[2,1,1],[0,1,1],[1,0,1]]
//Output: -1
//Explanation:  The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
//
//Example 3:
//
//Input: [[0,2]]
//Output: 0
//Explanation:  Since there are already no fresh oranges at minute 0, the answer is just 0.
//
//
//
//Note:
//
//    1 <= grid.length <= 10
//    1 <= grid[0].length <= 10
//    grid[i][j] is only 0, 1, or 2.

type point struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	rowL := len(grid)
	if rowL == 0 {
		return -1
	}
	columnL := len(grid[0])

	if columnL == 0 {
		return -1
	}

	rotten := make([]point, 0)
	fresh := 0

	for r, row := range grid {
		for c, fruit := range row {
			if fruit == 1 {
				fresh++
			} else if fruit == 2 {
				rotten = append(rotten, point{c, r})
			}
		}
	}

	// no fresh, no need to run
	if fresh == 0 {
		return 0
	}

	// no rotten, never happen
	if len(rotten) == 0 {
		return -1
	}

	turn := 0
	var processing bool
	arr := make([]point, 0)
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	var p point

	for fresh > 0 {
		processing = false
		arr = arr[:0]
		for _, r := range rotten {
			for i := 0; i < 4; i++ {
				p = point{
					x: r.x + dx[i],
					y: r.y + dy[i],
				}
				if p.x < columnL && p.x >= 0 && p.y < rowL && p.y >= 0 {
					arr = append(arr, p)
				}
			}
		}

		rotten = rotten[:]

		for _, p := range arr {
			// rotten fresh one
			if grid[p.y][p.x] == 1 {
				grid[p.y][p.x] = 2
				fresh--
				rotten = append(rotten, p)
				processing = true
			}
		}

		if processing {
			turn++
		} else {
			// for condition means there exists some fresh, and no processing means cannot go any further
			return -1
		}
	}

	return turn
}

// problems
// 1. incorrect access variable
// 2. forget to check boundary condition: if no fresh exist, return 0, if not rotten, return -1
// 3. i mix x & y in reverse order
// 4. when checking for each fresh fruit, choose minimum distance among all distance to every rotten
// 5. wrong implementation, by 4, the order should be each fruit to all rotten
// 6. i have wrong algorithm, it cannot not simply use distance to decide, consider the situation that one rotten is alone and there's one fruit which is close to this alone rotten, if there's another rotten that is far but connected, my algorithm goes to wrong result.
// 7. use wrong value, when getting adjacents, i miss use limit as value
// 8. wrong understanding, when no fresh then terminates
// 9. refactor, use array to get adjacents, also, use only one variable of fresh to denote if any fresh exist

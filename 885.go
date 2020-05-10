package main

// On a 2 dimensional grid with R rows and C columns, we start at (r0, c0) facing east.
//
// Here, the north-west corner of the grid is at the first row and column, and the south-east corner of the grid is at the last row and column.
//
// Now, we walk in a clockwise spiral shape to visit every position in this grid.
//
// Whenever we would move outside the boundary of the grid, we continue our walk outside the grid (but may return to the grid boundary later.)
//
// Eventually, we reach all R * C spaces of the grid.
//
// Return a list of coordinates representing the positions of the grid in the order they were visited.
//
//
//
// Example 1:
//
// Input: R = 1, C = 4, r0 = 0, c0 = 0
// Output: [[0,0],[0,1],[0,2],[0,3]]
//
//
//
//
// Example 2:
//
// Input: R = 5, C = 6, r0 = 1, c0 = 4
// Output: [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]
//
//
//
//
// Note:
//
//     1 <= R <= 100
//     1 <= C <= 100
//     0 <= r0 < R
//     0 <= c0 < C

type direction int

const (
	east direction = iota
	south
	west
	north
)

// x, y steps
var steps = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	result := make([][]int, R*C)
	result[0] = []int{r0, c0}

	distance := 1
	dir := east
	for idx, x, y := 1, c0, r0; idx < len(result); {
		for tmp := distance; tmp > 0; tmp-- {
			x = x + steps[dir][0]
			y = y + steps[dir][1]

			if x >= 0 && x < C && y >= 0 && y < R {
				result[idx] = []int{y, x}
				idx++
			}
		}

		// spiral
		dir = (dir + 1) % 4
		if dir == west || dir == east {
			distance++
		}
	}

	return result
}

//	problems
//	1.	spiral walk distance wrong

//	2.	wrong walking, south means y is increased, north means y is decreased

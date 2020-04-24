package main

// A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum.
//
// Given an grid of integers, how many 3 x 3 "magic square" subgrids are there?  (Each subgrid is contiguous).
//
//
//
// Example 1:
//
// Input: [[4,3,8,4],
//         [9,5,1,9],
//         [2,7,6,2]]
// Output: 1
// Explanation:
// The following subgrid is a 3 x 3 magic square:
// 438
// 951
// 276
//
// while this one is not:
// 384
// 519
// 762
//
// In total, there is only one magic square inside the given grid.
//
// Note:
//
//     1 <= grid.length <= 10
//     1 <= grid[0].length <= 10
//     0 <= grid[i][j] <= 15

func numMagicSquaresInside(grid [][]int) int {
	yLength := len(grid)
	if yLength < 3 {
		return 0
	}

	xLength := len(grid[0])
	if xLength < 3 {
		return 0
	}

	var count int
	for y := 0; y <= yLength-3; y++ {
		for x := 0; x <= xLength-3; x++ {
			if isMagic(grid, x, y) {
				count++
			}
		}
	}

	return count
}

func isMagic(grid [][]int, x, y int) bool {
	nums := make([]bool, 10)
	sum := 15

	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			if grid[y+j][x+i] >= 10 || grid[y+j][x+i] <= 0 || nums[grid[y+j][x+i]] {
				return false
			}
			nums[grid[y+j][x+i]] = true
		}
	}

	if grid[y+1][x+1] != 5 {
		return false
	}

	if grid[y][x]+grid[y][x+1]+grid[y][x+2] != sum {
		return false
	}

	if grid[y+2][x]+grid[y+2][x+1]+grid[y+2][x+2] != sum {
		return false
	}

	if grid[y][x]+grid[y+1][x]+grid[y+2][x] != sum {
		return false
	}

	if grid[y][x+2]+grid[y+1][x+2]+grid[y+2][x+2] != sum {
		return false
	}

	if grid[y][x]+grid[y+1][x+1]+grid[y+2][x+2] != sum {
		return false
	}

	if grid[y][x+2]+grid[y+1][x+1]+grid[y+2][x] != sum {
		return false
	}

	return true
}

//	problems
//	1.	number can only be 1-9
//	2.	reference: https://leetcode.com/problems/magic-squares-in-grid/discuss/133874/Python-5-and-43816729
//		center must be 5, four corner must be even number, clockwise number should be 43816729

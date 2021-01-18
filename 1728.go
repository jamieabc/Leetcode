package main

import "fmt"

// A game is played by a cat and a mouse named Cat and Mouse.
//
// The environment is represented by a grid of size rows x cols, where each element is a wall, floor, player (Cat, Mouse), or food.
//
// Players are represented by the characters 'C'(Cat),'M'(Mouse).
// Floors are represented by the character '.' and can be walked on.
// Walls are represented by the character '#' and cannot be walked on.
// Food is represented by the character 'F' and can be walked on.
// There is only one of each character 'C', 'M', and 'F' in grid.
// Mouse and Cat play according to the following rules:
//
// Mouse moves first, then they take turns to move.
// During each turn, Cat and Mouse can jump in one of the four directions (left, right, up, down). They cannot jump over the wall nor outside of the grid.
// catJump, mouseJump are the maximum lengths Cat and Mouse can jump at a time, respectively. Cat and Mouse can jump less than the maximum length.
// Staying in the same position is allowed.
// Mouse can jump over Cat.
// The game can end in 4 ways:
//
// If Cat occupies the same position as Mouse, Cat wins.
// If Cat reaches the food first, Cat wins.
// If Mouse reaches the food first, Mouse wins.
// If Mouse cannot get to the food within 1000 turns, Cat wins.
// Given a rows x cols matrix grid and two integers catJump and mouseJump, return true if Mouse can win the game if both Cat and Mouse play optimally, otherwise return false.
//
//
//
// Example 1:
//
//
//
// Input: grid = ["####F","#C...","M...."], catJump = 1, mouseJump = 2
// Output: true
// Explanation: Cat cannot catch Mouse on its turn nor can it get the food before Mouse.
// Example 2:
//
//
//
// Input: grid = ["M.C...F"], catJump = 1, mouseJump = 4
// Output: true
// Example 3:
//
// Input: grid = ["M.C...F"], catJump = 1, mouseJump = 3
// Output: false
// Example 4:
//
// Input: grid = ["C...#","...#F","....#","M...."], catJump = 2, mouseJump = 5
// Output: false
// Example 5:
//
// Input: grid = [".M...","..#..","#..#.","C#.#.","...#F"], catJump = 3, mouseJump = 1
// Output: true
//
//
// Constraints:
//
// rows == grid.length
// cols = grid[i].length
// 1 <= rows, cols <= 8
// grid[i][j] consist only of characters 'C', 'M', 'F', '.', and '#'.
// There is only one of each character 'C', 'M', and 'F' in grid.
// 1 <= catJump, mouseJump <= 8

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	var cat, mouse, food []int

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 'M' {
				mouse = []int{i, j}
			} else if grid[i][j] == 'C' {
				cat = []int{i, j}
			} else if grid[i][j] == 'F' {
				food = []int{i, j}
			}
		}
	}

	// visited[i]: 0: not visited, 1: mouse, 2: cat, 3: both
	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
	}

	limit := 1000

	visited[mouse[0]][mouse[1]] += 1
	visited[cat[0]][cat[1]] += 2

	return dfs(grid, visited, cat, mouse, food, catJump, mouseJump, 0, &limit)
}

func dfs(grid []string, visited [][]int, cat, mouse, food []int, catJump, mouseJump, turns int, limit *int) bool {
	fmt.Println("cat", cat, "mouse", mouse)
	// mouse & cat at same location
	if samePosition(cat, food) || (samePosition(cat, mouse) && turns&1 == 0) {
		if turns < *limit {
			*limit = turns
		}
		return false
	}

	if turns >= *limit {
		return false
	}

	// mouse reaches food
	if samePosition(mouse, food) {
		return true
	}

	var win bool

	if turns&1 == 0 {
		// mouse
		nexts := possiblePosition(grid, visited, mouse, mouseJump, turns)
		fmt.Println("mouse", mouse, "cat", cat, nexts)

		for _, n := range nexts {
			win = win || dfs(grid, visited, cat, n, food, catJump, mouseJump, turns+1, limit)
		}
	} else {
		// cat
		nexts := possiblePosition(grid, visited, cat, catJump, turns)
		fmt.Println("cat", cat, "mouse", mouse, nexts)

		for _, n := range nexts {
			win = win || dfs(grid, visited, n, mouse, food, catJump, mouseJump, turns+1, limit)
		}
	}

	return win
}

func possiblePosition(grid []string, visited [][]int, orig []int, jump, turn int) [][]int {
	pos := make([][]int, 0)
	var steps, value int
	if turn&1 == 0 {
		steps, value = 1, 2
	} else {
		steps, value = 2, 1
	}

	for i := 1; i <= jump; i++ {
		for _, d := range dirs {
			newY, newX := orig[0]+i*d[0], orig[1]+i*d[1]

			if validPos(grid, newX, newY) && (visited[newY][newX] == 0 || visited[newY][newX] == value) {
				pos = append(pos, []int{newY, newX})
				visited[newY][newX] += steps
			}
		}
	}

	return pos
}

func validPos(grid []string, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid) && grid[y][x] != '#'
}

func samePosition(p1, p2 []int) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}

func main() {
	grid := []string{
		"C...#", "...#F", "....#", "M....",
	}

	catJump := 2
	mouseJump := 5

	fmt.Println(canMouseWin(grid, catJump, mouseJump))
}

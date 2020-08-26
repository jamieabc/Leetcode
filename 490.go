package main

//There is a ball in a maze with empty spaces and walls. The ball can go through empty spaces by rolling up, down, left or right, but it won't stop rolling until hitting a wall. When the ball stops, it could choose the next direction.
//
//Given the ball's start position, the destination and the maze, determine whether the ball could stop at the destination.
//
//The maze is represented by a binary 2D array. 1 means the wall and 0 means the empty space. You may assume that the borders of the maze are all walls. The start and destination coordinates are represented by row and column indexes.
//
//
//
//Example 1:
//
//Input 1: a maze represented by a 2D array
//
//0 0 1 0 0
//0 0 0 0 0
//0 0 0 1 0
//1 1 0 1 1
//0 0 0 0 0
//
//Input 2: start coordinate (rowStart, colStart) = (0, 4)
//Input 3: destination coordinate (rowDest, colDest) = (4, 4)
//
//Output: true
//
//Explanation: One possible way is : left -> down -> left -> down -> right -> down -> right.
//
//Example 2:
//
//Input 1: a maze represented by a 2D array
//
//0 0 1 0 0
//0 0 0 0 0
//0 0 0 1 0
//1 1 0 1 1
//0 0 0 0 0
//
//Input 2: start coordinate (rowStart, colStart) = (0, 4)
//Input 3: destination coordinate (rowDest, colDest) = (3, 2)
//
//Output: false
//
//Explanation: There is no way for the ball to stop at the destination.
//
//
//
//Note:
//
//    There is only one ball and one destination in the maze.
//    Both the ball and the destination exist on an empty space, and they will not be at the same position initially.
//    The given maze does not contain border (like the red rectangle in the example pictures), but you could assume the border of the maze are all walls.
//    The maze contains at least 2 empty spaces, and both the width and height of the maze won't exceed 100.

func hasPath(maze [][]int, start []int, destination []int) bool {
	if start[0] == destination[0] && start[1] == destination[1] {
		return true
	}

	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}

	queue := [][]int{start}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			x, y := pos[1], pos[0]

			for validPosition(maze, x+dir[1], y+dir[0]) {
				x += dir[1]
				y += dir[0]
			}

			// cannot go any further to new place
			if (x == pos[1] && y == pos[0]) || visited[y][x] {
				continue
			}

			visited[y][x] = true

			if x == destination[1] && y == destination[0] {
				return true
			}
			queue = append(queue, []int{y, x})
		}
	}

	return false
}

func hasPath1(maze [][]int, start []int, destination []int) bool {
	// start position equals destination position
	if start[0] == destination[0] && start[1] == destination[1] {
		return true
	}

	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}

	return backtracking(maze, visited, start, destination)
}

// in the form of [y, x]
var dirs = [][]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, 1},  // right
	{0, -1}, // left
}

func backtracking(maze [][]int, visited [][]bool, src, dst []int) bool {
	for _, dir := range dirs {
		// try until reach wall or limit
		x, y := src[1], src[0]

		for validPosition(maze, x+dir[1], y+dir[0]) {
			x, y = x+dir[1], y+dir[0]
		}

		// not able to move in this direction
		if y == src[0] && x == src[1] || visited[y][x] {
			continue
		}

		visited[y][x] = true

		if dst[0] == y && dst[1] == x {
			return true
		}

		found := backtracking(maze, visited, []int{y, x}, dst)
		if found {
			return true
		}
	}

	return false
}

func validPosition(maze [][]int, x, y int) bool {
	return x >= 0 && x < len(maze[0]) && y >= 0 && y < len(maze) && maze[y][x] == 0
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/the-maze/discuss/97071/Easy-understanding-Java-bfs-solution.

//		dont' forget boundary checking: if start position == destination
//		position, then just return true

//		also, write another version of bfs

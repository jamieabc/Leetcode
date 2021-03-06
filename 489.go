package main

import "fmt"

// Given a robot cleaner in a room modeled as a grid.
//
// Each cell in the grid can be empty or blocked.
//
// The robot cleaner with 4 given APIs can move forward, turn left or turn right. Each turn it made is 90 degrees.
//
// When it tries to move into a blocked cell, its bumper sensor detects the obstacle and it stays on the current cell.
//
// Design an algorithm to clean the entire room using only the 4 given APIs shown below.
//
// interface Robot {
//   // returns true if next cell is open and robot moves into the cell.
//   // returns false if next cell is obstacle and robot stays on the current cell.
//   boolean move();
//
//   // Robot will stay on the same cell after calling turnLeft/turnRight.
//   // Each turn will be 90 degrees.
//   void turnLeft();
//   void turnRight();
//
//   // Clean the current cell.
//   void clean();
// }
//
// Example:
//
// Input:
// room = [
//   [1,1,1,1,1,0,1,1],
//   [1,1,1,1,1,0,1,1],
//   [1,0,1,1,1,1,1,1],
//   [0,0,0,1,0,0,0,0],
//   [1,1,1,1,1,1,1,1]
// ],
// row = 1,
// col = 3
//
// Explanation:
// All grids in the room are marked by either 0 or 1.
// 0 means the cell is blocked, while 1 means the cell is accessible.
// The robot initially starts at the position of row=1, col=3.
// From the top left corner, its position is one row below and three columns right.
//
// Notes:
//
//     The input is only given to initialize the room and the robot's position internally. You must solve this problem "blindfolded". In other words, you must control the robot using only the mentioned 4 APIs, without knowing the room layout and the initial robot's position.
//     The robot's initial position will always be in an accessible cell.
//     The initial direction of the robot will be facing up.
//     All accessible cells are connected, which means the all cells marked as 1 will be accessible by the robot.
//     Assume all four edges of the grid are all surrounded by wall.

/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 *
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

var dirs = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func cleanRoom(robot *Robot) {
	visited := make(map[[2]int]struct{})
	dfs(robot, visited, 0, 0, 0)
}

func dfs(robot *Robot, visited map[[2]int]struct{}, x, y, dir int) {
	robot.Clean()
	visited[[2]int{y, x}] = struct{}{}

	for i, d := 0, dir; i < 4; i, d = i+1, (d+1)%4 {
		newY, newX := y+dirs[d][0], x+dirs[d][1]

		if _, ok := visited[[2]int{newY, newX}]; !ok && robot.Move() {
			dfs(robot, visited, newX, newY, d)
		}

		robot.TurnRight()
	}

	// finish trying all directions, back to previous state
	stepBack(robot)
}

func stepBack(robot *Robot) {
	robot.TurnRight()
	robot.TurnRight()

	robot.Move()

	robot.TurnRight()
	robot.TurnRight()
}

func cleanRoom2(robot *Robot) {
	visited := make(map[Point]bool)
	moves := [][]int{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}

	dfs(robot, Point{0, 0}, 0, visited, moves)
}

type Point struct {
	x, y int
}

func dfs(robot *Robot, p Point, dir int, visited map[Point]bool, moves [][]int) {
	robot.Clean()
	visited[p] = true

	for i := range moves {
		newDir := (dir + i) % 4
		newPoint := Point{
			x: p.x + moves[newDir][0],
			y: p.y + moves[newDir][1],
		}

		if !visited[newPoint] && robot.Move() {
			dfs(robot, newPoint, newDir, visited, moves)
			robot.Back()
		}
		robot.TurnRight()
	}
}

func (r *Robot) Back() {
	r.TurnRight()
	r.TurnRight()
	r.Move()
	r.TurnLeft()
	r.TurnLeft()
}
func cleanRoom1(robot *Robot) {
	visited := make(map[Point]bool)
	moves := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	dfs(robot, Point{0, 0}, 0, visited, moves)
}

type Point struct {
	x, y int
}

func dfs(robot *Robot, p Point, dir int, visited map[Point]bool, moves [][]int) {
	robot.Clean()
	visited[p] = true

	for i := range moves {
		newDir := (dir + i) % 4
		newPoint := Point{
			x: p.x + moves[i][1],
			y: p.y + moves[i][0],
		}

		if !visited[newPoint] && robot.Move() {
			dfs(robot, newPoint, newDir, visited, moves)
			robot.Back()
		}
		robot.TurnRight()
	}
}

func (r *Robot) Back() {
	r.TurnRight()
	r.TurnRight()
	r.Move()
	r.TurnLeft()
	r.TurnLeft()
}

//	Notes
//	1.	don't know how to solve this, at first I want to use a stack to store
//		all choices not done, e.g. store into stack of right, down, and left
//		because it defaults to go up. and when moving up, push a down into
//		stack. then I found it's really complicated

//	2.	add reference https://leetcode.com/problems/robot-room-cleaner/discuss/139057/Very-easy-to-understand-Java-solution

//		author uses a hash to store any visited place by "x-y", for every
//		position, try for 4 directions. if any direction is movable, go to
//		new position recursively.

//		when all directions are visited, do specific operations:
//		turn right * 2 (facing opposite direction), move (back to previous
//		position), turn left * 2 (facing to original direction), then continue
//		to the 4 direction checking

//	3.	from sample code, it adds more simpler way to write, and author uses
//		additional data structure in map to avoid string operation

//	4.	if processed correctly, robot won't try to go to where come from,
//		because that position should be already marked visited

//	5.	becareful, going north means y-1

//	6.	inspired form sample code, can use map[]struct{} for checking visited

//		also, author puts everything in the function, interesting

//	7.	inspired from another sample code, author uses map[[2]int]bool to check,
//		which is comparable by default

//	8.	inspired from https://leetcode.com/problems/robot-room-cleaner/discuss/153530/DFS-Logical-Thinking

//		need to track visited position, string of "x y" should also be fine

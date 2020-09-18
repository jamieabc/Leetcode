package main

//On an infinite plane, a robot initially stands at (0, 0) and faces north.  The robot can receive one of three instructions:
//
//"G": go straight 1 unit;
//"L": turn 90 degrees to the left;
//"R": turn 90 degress to the right.
//The robot performs the instructions given in order, and repeats them forever.
//
//Return true if and only if there exists a circle in the plane such that the robot never leaves the circle.
//
//
//
//Example 1:
//
//Input: "GGLLGG"
//Output: true
//Explanation:
//The robot moves from (0,0) to (0,2), turns 180 degrees, and then returns to (0,0).
//When repeating these instructions, the robot remains in the circle of radius 2 centered at the origin.
//Example 2:
//
//Input: "GG"
//Output: false
//Explanation:
//The robot moves north indefinitely.
//Example 3:
//
//Input: "GL"
//Output: true
//Explanation:
//The robot moves from (0, 0) -> (0, 1) -> (-1, 1) -> (-1, 0) -> (0, 0) -> ...
//
//
//Note:
//
//1 <= instructions.length <= 100
//instructions[i] is in {'G', 'L', 'R'}

func isRobotBounded(instructions string) bool {
	var dirs = [][]int{
		{1, 0},  // north
		{0, 1},  // east
		{-1, 0}, // south
		{0, -1}, // west
	}
	pos := []int{0, 0}
	var dir int

	for i := range instructions {
		if instructions[i] == 'L' {
			dir--
			if dir < 0 {
				dir = 3
			}
		} else if instructions[i] == 'R' {
			dir++

			if dir > 3 {
				dir = 0
			}
		} else {
			pos[0] += dirs[dir][1]
			pos[1] += dirs[dir][0]
		}
	}

	if dir == 1 || dir == 2 || dir == 3 {
		return true
	}

	if pos[0] != 0 || pos[1] != 0 {
		return false
	}

	return true
}

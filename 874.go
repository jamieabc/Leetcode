package main

//A robot on an infinite grid starts at point (0, 0) and faces north.  The robot can receive one of three possible types of commands:
//
//-2: turn left 90 degrees
//-1: turn right 90 degrees
//1 <= x <= 9: move forward x units
//Some of the grid squares are obstacles.
//
//The i-th obstacle is at grid point (obstacles[i][0], obstacles[i][1])
//
//If the robot would try to move onto them, the robot stays on the previous grid square instead (but still continues following the rest of the route.)
//
//Return the square of the maximum Euclidean distance that the robot will be from the origin.
//
//
//
//Example 1:
//
//Input: commands = [4,-1,3], obstacles = []
//Output: 25
//Explanation: robot will go to (3, 4)
//Example 2:
//
//Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
//Output: 65
//Explanation: robot will be stuck at (1, 4) before turning left and going to (1, 8)
//
//
//Note:
//
//0 <= commands.length <= 10000
//0 <= obstacles.length <= 10000
//-30000 <= obstacle[i][0] <= 30000
//-30000 <= obstacle[i][1] <= 30000
//The answer is guaranteed to be less than 2 ^ 31.

type position struct {
	p         point
	direction int // 0: north, 1: east, 2: south, 3: west
}

type point struct {
	x, y int
}

var (
	dx = [4]int{0, 1, 0, -1}
	dy = [4]int{1, 0, -1, 0}
)

func (p *position) forward(steps int, set map[point]bool, max *int, maxPoint *point) {
	if steps == 0 {
		return
	}

	x := p.p.x
	y := p.p.y
	var tmpX, tmpY int

	for i := 0; i < steps; i++ {
		tmpX = x + dx[p.direction]
		tmpY = y + dy[p.direction]
		if set[point{tmpX, tmpY}] {
			break
		} else {
			x = tmpX
			y = tmpY
		}
	}

	if abs(x) > abs(maxPoint.x) || abs(y) > abs(maxPoint.y) {
		newDistance := x*x + y*y
		if newDistance > *max {
			*max = newDistance
			maxPoint.x = x
			maxPoint.y = y
		}
	}

	p.p.x = x
	p.p.y = y
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (p *position) turn(i int) {
	if i == -1 {
		// turn right
		if p.direction == 3 {
			p.direction = 0
		} else {
			p.direction++
		}
	} else {
		// turn left
		if p.direction == 0 {
			p.direction = 3
		} else {
			p.direction--
		}
	}
}

func robotSim(commands []int, obstacles [][]int) int {
	p := &position{
		p:         point{0, 0},
		direction: 0,
	}

	max := 0
	maxPoint := &point{0, 0}
	set := make(map[point]bool)
	for _, o := range obstacles {
		set[point{o[0], o[1]}] = true
	}

	for _, c := range commands {
		if c < 0 {
			p.turn(c)
		} else {
			p.forward(c, set, &max, maxPoint)
		}
	}

	return max
}

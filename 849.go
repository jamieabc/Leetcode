package main

//In a row of seats, 1 represents a person sitting in that seat, and 0 represents that the seat is empty.
//
//There is at least one empty seat, and at least one person sitting.
//
//Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.
//
//Return that maximum distance to closest person.
//
//Example 1:
//
//Input: [1,0,0,0,1,0,1]
//Output: 2
//Explanation:
//If Alex sits in the second open seat (seats[2]), then the closest person has distance 2.
//If Alex sits in any other open seat, the closest person has distance 1.
//Thus, the maximum distance to the closest person is 2.
//
//Example 2:
//
//Input: [1,0,0,0]
//Output: 3
//Explanation:
//If Alex sits in the last seat, the closest person is 3 seats away.
//This is the maximum distance possible, so the answer is 3.
//
//Note:
//
//    1 <= seats.length <= 20000
//    seats contains only 0s or 1s, at least one 0, and at least one 1.

func maxDistToClosest(seats []int) int {
	// find max consecutive  zeros
	// edge cases: max consecutive zero at first or end
	// 0 0 0 1
	// 1 0 0 0
	// 1 0 0 1
	// 1 0 0 0 0 1
	// 1 0 0 0 1
	// 1 0 0 0 0 0 1

	var max, start, end int
	consecutive := false
	length := len(seats)
	for i, s := range seats {
		if s == 0 {
			if !consecutive {
				start = i
				consecutive = true
			}
		} else {
			if consecutive {
				consecutive = false
				end = i - 1
				dist := countDistance(start, end, length)
				if dist > max {
					max = dist
				}
			}
		}
	}

	// in case consecutive zeros at end
	if consecutive {
		consecutive = false
		end = length - 1
		dist := countDistance(start, end, length)
		if dist > max {
			return dist
		}
	}

	return max
}

func countDistance(start, end, length int) int {
	dist := end - start + 1
	if start == 0 || end == length-1 {
		return dist
	}

	if dist%2 == 0 {
		return dist / 2
	}
	return (dist + 1) / 2
}

// problems
// 1. forget about condition that when consecutive zeros at edge, it gets larger than in middle, wrong way of calculating max distance

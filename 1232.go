package main

//You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.
//
//
//
//
//
//Example 1:
//
//Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
//Output: true
//
//Example 2:
//
//Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
//Output: false
//
//
//
//Constraints:
//
//    2 <= coordinates.length <= 1000
//    coordinates[i].length == 2
//    -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
//    coordinates contains no duplicate point.

func checkStraightLine(coordinates [][]int) bool {
	length := len(coordinates)
	if length == 2 {
		return true
	}

	// vertical
	if coordinates[0][0] == coordinates[1][0] {
		for i := 2; i < length; i++ {
			if coordinates[i][0] != coordinates[0][0] {
				return false
			}
		}
		return true
	}

	// horizontal
	if coordinates[0][1] == coordinates[1][1] {
		for i := 2; i < length; i++ {
			if coordinates[i][1] != coordinates[0][1] {
				return false
			}
		}
		return true
	}

	// slope
	tmp := float32(coordinates[0][1]-coordinates[1][1]) / float32(coordinates[0][0]-coordinates[1][0])

	var slope int
	reversed := false
	if tmp < 1 {
		reversed = true
		slope = (coordinates[1][1] - coordinates[0][1]) / (coordinates[1][0] - coordinates[0][0])
	} else {
		slope = int(tmp)
	}

	var expected int
	for i := 2; i < length; i++ {
		if reversed {
			expected = coordinates[0][1] + slope*(coordinates[i][0]-coordinates[0][0])
		} else {
			expected = coordinates[0][1] - slope*(coordinates[0][0]-coordinates[i][0])
		}
		if expected != coordinates[i][1] {
			return false
		}
	}
	return true
}

//	problems
//	1.	optimize, avoid using float64
//	2.	optimize, use float32

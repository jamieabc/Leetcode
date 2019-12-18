package main

//A boomerang is a set of 3 points that are all distinct and not in a straight line.
//
//Given a list of three points in the plane, return whether these points are a boomerang.
//
//
//
//Example 1:
//
//Input: [[1,1],[2,3],[3,2]]
//Output: true
//
//Example 2:
//
//Input: [[1,1],[2,2],[3,3]]
//Output: false
//
//
//
//Note:
//
//    points.length == 3
//    points[i].length == 2
//    0 <= points[i][j] <= 100

func isBoomerang(points [][]int) bool {
	// check if any point in some direction is same
	if sameX(points[0], points[1]) {
		if sameX(points[0], points[2]) || sameY(points[0], points[1]) {
			return false
		}
		return true
	}

	if sameX(points[0], points[2]) {
		if sameX(points[0], points[1]) || sameY(points[0], points[2]) {
			return false
		}
		return true
	}

	if sameX(points[1], points[2]) {
		if sameX(points[0], points[1]) || sameY(points[1], points[2]) {
			return false
		}
		return true
	}

	if sameY(points[0], points[1]) {
		if sameY(points[0], points[2]) || sameX(points[0], points[1]) {
			return false
		}
		return true
	}

	if sameY(points[0], points[2]) {
		if sameY(points[0], points[1]) || sameX(points[0], points[2]) {
			return false
		}
		return true
	}

	if sameY(points[1], points[2]) {
		if sameY(points[0], points[1]) || sameX(points[1], points[2]) {
			return false
		}
		return true
	}

	// check points not is a line
	if float64(abs(points[0][1]-points[1][1]))/float64(abs(points[0][0]-points[1][0])) == float64(abs(points[0][1]-points[2][1]))/float64(abs(points[0][0]-points[2][0])) {
		return false
	}

	return true
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func sameX(p1, p2 []int) bool {
	return p1[0] == p2[0]
}

func sameY(p1, p2 []int) bool {
	return p1[1] == p2[1]
}

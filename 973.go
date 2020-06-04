package main

// We have a list of points on the plane.  Find the K closest points to the origin (0, 0).
//
// (Here, the distance between two points on a plane is the Euclidean distance.)
//
// You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.)
//
//
//
// Example 1:
//
// Input: points = [[1,3],[-2,2]], K = 1
// Output: [[-2,2]]
// Explanation:
// The distance between (1, 3) and the origin is sqrt(10).
// The distance between (-2, 2) and the origin is sqrt(8).
// Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
// We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].
// Example 2:
//
// Input: points = [[3,3],[5,-1],[-2,4]], K = 2
// Output: [[3,3],[-2,4]]
// (The answer [[-2,4],[3,3]] would also be accepted.)
//
//
// Note:
//
// 1 <= K <= points.length <= 10000
// -10000 < points[i][0] < 10000
// -10000 < points[i][1] < 10000

func kClosest(points [][]int, K int) [][]int {
	if K >= len(points) {
		return points
	}

	quickSelect(points, 0, len(points)-1, K)

	return points[:K]
}

func quickSelect(points [][]int, start, end, target int) {
	p := points[start][0]*points[start][0] + points[start][1]*points[start][1]

	var i, j int
	for i, j = start, end; i <= j; {
		if points[i][0]*points[i][0]+points[i][1]*points[i][1] <= p {
			i++
			continue
		}

		if points[j][0]*points[j][0]+points[j][1]*points[j][1] > p {
			j--
			continue
		}

		points[i], points[j] = points[j], points[i]
		i++
		j--
	}

	points[start], points[j] = points[j], points[start]

	if j == target {
		return
	} else if j < target {
		quickSelect(points, j+1, end, target)
	} else {
		quickSelect(points, start, j-1, target)
	}
}

//	problems
//	1.	there could exists equal distance points, since it's to find closest,
//		so put equal or larger to right

//	2.	too slow, because I got wrong about item search range... just do search

//	3.	add reference https://leetcode.com/problems/k-closest-points-to-origin/discuss/220235/Java-Three-solutions-to-this-classical-K-th-problem.

//		author adds some conclusion

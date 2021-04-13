package main

import "math"

// Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane, return the maximum number of points that lie on the same straight line.
//
//
//
// Example 1:
//
// Input: points = [[1,1],[2,2],[3,3]]
// Output: 3
//
// Example 2:
//
// Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// Output: 4
//
//
//
// Constraints:
//
// 1 <= points.length <= 300
// points[i].length == 2
// -104 <= xi, yi <= 104
// All the points are unique.

func maxPoints(points [][]int) int {

	var largest int
	for i := range points {
		slopes := make(map[float64]int)

		for j := i + 1; j < len(points); j++ {
			s := slope(points[j], points[i])
			slopes[s]++
			largest = max(largest, slopes[s])
		}
	}

	return largest + 1
}

func slope(p1, p2 []int) float64 {
	if p2[0] == p1[0] {
		return float64(math.MaxInt32)
	}

	if p2[1] == p1[1] {
		return 0
	}
	return float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/max-points-on-a-line/discuss/47113/A-java-solution-with-notes

//		find gcd, after division, use map[x]map[y] to denote slope
//		gcd(a, b) = gcd(b, a%b) until b == 0

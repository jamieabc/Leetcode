package main

import (
	"math"
	"sort"
)

// Given a set of points in the xy-plane, determine the minimum area of a rectangle formed from these points, with sides parallel to the x and y axes.
//
// If there isn't any rectangle, return 0.
//
//
//
// Example 1:
//
// Input: [[1,1],[1,3],[3,1],[3,3],[2,2]]
// Output: 4
//
// Example 2:
//
// Input: [[1,1],[1,3],[3,1],[3,3],[4,1],[4,3]]
// Output: 2
//
//
//
// Note:
//
//     1 <= points.length <= 500
//     0 <= points[i][0] <= 40000
//     0 <= points[i][1] <= 40000
//     All points are distinct.

// tc: O(n^2)
func minAreaRect(points [][]int) int {
	xy := make(map[int]map[int]bool)
	for _, point := range points {
		if _, ok := xy[point[0]]; !ok {
			xy[point[0]] = make(map[int]bool)
		}
		xy[point[0]][point[1]] = true
	}

	minArea := math.MaxInt32
	for i := range points {
		for j := range points {
			// cannot form a rectangle
			if points[j][0] == points[i][0] || points[j][1] == points[i][1] {
				continue
			}

			if xy[points[i][0]][points[j][1]] && xy[points[j][0]][points[i][1]] {
				minArea = min(minArea, abs(points[i][0]-points[j][0])*abs(points[i][1]-points[j][1]))
			}
		}
	}

	if minArea == math.MaxInt32 {
		return 0
	}

	return minArea
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

// tc: O(n^2 * m log m), n: all x size, m: average size of y on specific x
func minAreaRect1(points [][]int) int {
	xy := make(map[int]map[int]bool)

	for _, point := range points {
		if _, ok := xy[point[0]]; !ok {
			xy[point[0]] = make(map[int]bool)
		}
		xy[point[0]][point[1]] = true

	}

	allX := make([]int, 0)
	for key := range xy {
		allX = append(allX, key)
	}

	sort.Ints(allX)
	minArea := math.MaxInt32

	for i := range allX {
		for j := i + 1; j < len(allX); j++ {
			// intersection of y
			commonY := make([]int, 0)
			for y := range xy[allX[i]] {
				if xy[allX[j]][y] {
					commonY = append(commonY, y)
				}
			}

			if len(commonY) < 2 {
				continue
			}

			sort.Ints(commonY)
			minY := math.MaxInt32
			for k := 1; k < len(commonY); k++ {
				minY = min(minY, commonY[k]-commonY[k-1])
			}

			minArea = min(minArea, (allX[j]-allX[i])*minY)
		}
	}

	if minArea == math.MaxInt32 {
		return 0
	}
	return minArea
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	problems
//	1.	if no rectangle found, return 0

//	2.	inspired from https://leetcode.com/problems/minimum-area-rectangle/discuss/192025/Java-N2-Hashmap

//		no need to sort by x, for every point, iterate through all other
//		points to find possible formed rectangle

//		author uses a clever way to check rectangle existence: find a point
//		that is not in x or y axis same as the other point, the using map
//		to find the other 2 points exists or not

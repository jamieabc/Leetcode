package main

import (
	"math"
	"sort"
)

// You are given an array points, an integer angle, and your location, where location = [posx, posy] and points[i] = [xi, yi] both denote integral coordinates on the X-Y plane.
//
// Initially, you are facing directly east from your position. You cannot move from your position, but you can rotate. In other words, posx and posy cannot be changed. Your field of view in degrees is represented by angle, determining how wide you can see from any given view direction. Let d be the amount in degrees that you rotate counterclockwise. Then, your field of view is the inclusive range of angles [d - angle/2, d + angle/2].
//
// Your browser does not support the video tag or this video format.
//
// You can see some set of points if, for each point, the angle formed by the point, your position, and the immediate east direction from your position is in your field of view.
//
// There can be multiple points at one coordinate. There may be points at your location, and you can always see these points regardless of your rotation. Points do not obstruct your vision to other points.
//
// Return the maximum number of points you can see.
//
//
//
// Example 1:
//
// Input: points = [[2,1],[2,2],[3,3]], angle = 90, location = [1,1]
// Output: 3
// Explanation: The shaded region represents your field of view. All points can be made visible in your field of view, including [3,3] even though [2,2] is in front and in the same line of sight.
//
// Example 2:
//
// Input: points = [[2,1],[2,2],[3,4],[1,1]], angle = 90, location = [1,1]
// Output: 4
// Explanation: All points can be made visible in your field of view, including the one at your location.
//
// Example 3:
//
// Input: points = [[1,0],[2,1]], angle = 13, location = [1,1]
// Output: 1
// Explanation: You can only see one of the two points, as shown above.
//
//
//
// Constraints:
//
// 1 <= points.length <= 105
// points[i].length == 2
// location.length == 2
// 0 <= angle < 360
// 0 <= posx, posy, xi, yi <= 100

func visiblePoints(points [][]int, angle int, location []int) int {
	angles := make([]float64, 0)
	var common int

	for _, point := range points {
		if point[0] == location[0] && point[1] == location[1] {
			common++
		} else {
			angles = append(angles, calculateAngle(location, point))
		}
	}

	sort.Slice(angles, func(i, j int) bool {
		return angles[i] < angles[j]
	})

	size := len(angles)
	for i := 0; i < size; i++ {
		angles = append(angles, angles[i]+360)
	}

	var largest int
	for i, j := 0, 0; j < len(angles); {
		for ; i == j || (j < len(angles) && angles[j]-angles[i] <= float64(angle)); j++ {
		}
		largest = max(largest, j-i)
		i++
	}

	return largest + common
}

func angleInRange(src, dst, target float64) bool {
	if src+target >= 180 {
		if dst >= 0 {
			return false
		}

		return dst <= src+target-360
	}

	return dst >= src && dst <= min(src+target, float64(180))
}

func min(i, j float64) float64 {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/maximum-number-of-visible-points/discuss/879125/C%2B%2B-atan2-and-sliding-window

//		use atan2 to find angle

//	2.	two pointer, first round j head of i, second round, j before i (after
//		% operation)

//	3.	inspired form https://leetcode.com/problems/maximum-number-of-visible-points/discuss/877822/Python-clean-sliding-window-solution-with-explanation

//		first round is okay, but second round, convert original angle to angle
//		+ 360, to fit into two pointer way

//		original angle: [-30, 50, 100, 160]
//		converted: [-30, 50, 100, 160, 330, 410, 460, 520]

//		the other technique is not to convert to angle, use original meaning,
//		that value corresponds with normailzed angle

//	4.	inspired from https://leetcode.com/problems/maximum-number-of-visible-points/discuss/877735/C%2B%2B-Clean-with-Explanation

//		author provides a very good observation, if two points are identical,
//		the atan is 0 degree, same as horizontal situation, need to calculate
//		separately

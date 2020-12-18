package main

import "sort"

// Given n cuboids where the dimensions of the ith cuboid is cuboids[i] = [widthi, lengthi, heighti] (0-indexed). Choose a subset of cuboids and place them on each other.
//
// You can place cuboid i on cuboid j if widthi <= widthj and lengthi <= lengthj and heighti <= heightj. You can rearrange any cuboid's dimensions by rotating it to put it on another cuboid.
//
// Return the maximum height of the stacked cuboids.
//
//
//
// Example 1:
//
//
//
// Input: cuboids = [[50,45,20],[95,37,53],[45,23,12]]
// Output: 190
// Explanation:
// Cuboid 1 is placed on the bottom with the 53x37 side facing down with height 95.
// Cuboid 0 is placed next with the 45x20 side facing down with height 50.
// Cuboid 2 is placed next with the 23x12 side facing down with height 45.
// The total height is 95 + 50 + 45 = 190.
// Example 2:
//
// Input: cuboids = [[38,25,45],[76,35,3]]
// Output: 76
// Explanation:
// You can't place any of the cuboids on the other.
// We choose cuboid 1 and rotate it so that the 35x3 side is facing down and its height is 76.
// Example 3:
//
// Input: cuboids = [[7,11,17],[7,17,11],[11,7,17],[11,17,7],[17,7,11],[17,11,7]]
// Output: 102
// Explanation:
// After rearranging the cuboids, you can see that all cuboids have the same dimension.
// You can place the 11x7 side down on all cuboids so their heights are 17.
// The maximum height of stacked cuboids is 6 * 17 = 102.
//
//
// Constraints:
//
// n == cuboids.length
// 1 <= n <= 100
// 1 <= widthi, lengthi, heighti <= 100

func maxHeight(cuboids [][]int) int {
	for _, c := range cuboids {
		sort.Slice(c, func(i, j int) bool {
			return c[i] > c[j]
		})
	}

	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		}

		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}

		return cuboids[i][2] > cuboids[j][2]
	})

	size := len(cuboids)

	// dp[i]: maximum height up to i
	dp := make([]int, size)

	for i := range cuboids {
		dp[i] = max(dp[i], cuboids[i][0])

		for j := i + 1; j < size; j++ {
			if stackable(cuboids[i], cuboids[j]) {
				dp[j] = max(dp[j], dp[i]+cuboids[j][0])
			}
		}
	}

	var maxHeight int
	for _, h := range dp {
		maxHeight = max(maxHeight, h)
	}

	return maxHeight
}

func stackable(bottom, top []int) bool {
	for i := range top {
		if top[i] > bottom[i] {
			return false
		}
	}

	return true
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	becareful of sort.Slice, if values are same, it's not guarantee any order,
//		so it's better to specify how to sort when values are same

//	2.	inspired from https://www.youtube.com/watch?v=WlYCqB-Lz5s

//		errichto specifies the point: is there any way that sort order remains
//		coherent, if it is then it's possible to sort those tuples, otherwise,
//		need to find another way

//	3.	inspired from https://leetcode.com/problems/maximum-height-by-stacking-cuboids/discuss/970293/JavaC%2B%2BPython-DP-Prove-with-Explanation

//		lee also has a good explanation

//	4.	inspired from https://leetcode.com/problems/maximum-height-by-stacking-cuboids/discuss/970256/PythonC%2B%2B-O(n-*-n)-Fix-smallest-size

//		author also states that order matter when selecting bottom box, so it
//		should be in a good order

//	5.	inspired from https://leetcode.com/problems/maximum-height-by-stacking-cuboids/discuss/970394/Clean-Java

//		author provides another view of this problem: LIS

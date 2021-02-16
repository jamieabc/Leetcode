package main

// Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//
// Note: You may not slant the container and n is at least 2.
//
//
//
//
//
// The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
//
//
//
// Example:
//
// Input: [1,8,6,2,5,4,8,3,7]
// Output: 49

func maxArea(height []int) int {
	var result int

	for i, j := 0, len(height)-1; i < j; {
		result = max(result, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return result
}

func min(i, j int) int {
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
//	1.	inspired from https://leetcode.com/problems/container-with-most-water/discuss/6100/Simple-and-clear-proofexplanation

//		author concludes with some explanation, widest area always a candidate
//		if width is shrink, height need to be increased, thus it's safe to remove
//		shorter one

//		author doesn't check condition when height are same, so I think more about it. if height of
//		i, j are same, then there are two conditions:
//		- next is higher, next no matter which side I choose, it won't change any result because
//		  height is bounded by i or j
//		- next is lower, and since it's lower and width is smaller, it won't change result, so it's
//		  no problem to chose any side

//		also, for the minimum height * length condition, author didn't consider it, I think a little
//		more about it. because width is always 1, and the minimum difference to lowest is 1, but
//		due to width changes, it will not happen that min height * length is largest

//	2.	extremely hard for me to think, I though it was stack but fail to find
//		it pattern

//		to get most area, it only care about smaller one and don't care about
//		higher one, because area is bounded by lower one

//		so it's not just caring about the highest before/after self, width
//		also matters, that's the reason start from left most & right most bars

//	3.	inspired from https://leetcode.com/problems/container-with-most-water/discuss/200246/Proof-by-formula

//		author provides a proof why it's safe to remove shorter bar

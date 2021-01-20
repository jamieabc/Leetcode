package main

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
//
//
//
// Example 1:
//
//
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
// Example 2:
//
// Input: height = [4,2,0,3,2,5]
// Output: 9
//
//
// Constraints:
//
// n == height.length
// 0 <= n <= 3 * 104
// 0 <= height[i] <= 105

func trap(height []int) int {
	size := len(height)
	var leftHighest, rightHighest, area int

	for left, right := 0, size-1; left < right; {
		if height[left] < height[right] {
			if height[left] > leftHighest {
				leftHighest = height[left]
			} else {
				area += leftHighest - height[left]
			}

			left++
		} else {
			if height[right] > rightHighest {
				rightHighest = height[right]
			} else {
				area += rightHighest - height[right]
			}

			right--
		}
	}

	return area
}

// count area in vertical way
func trap3(height []int) int {
	size := len(height)
	left, right := make([]int, size), make([]int, size)

	left[0] = height[0]
	for i := 1; i < size; i++ {
		left[i] = max(left[i-1], height[i])
	}

	right[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	var area int

	for i := range left {
		area += min(left[i], right[i]) - height[i]
	}

	return area
}

func trap2(height []int) int {
	size := len(height)
	left, right := make([]int, size), make([]int, size)
	for i := range left {
		left[i] = -1
		right[i] = -1
	}

	stack := make([]int, 0)
	for i := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	var area int

	for i := range left {
		if left[i] > -1 && right[i] > -1 {
			area += (min(height[left[i]], height[right[i]]) - height[i]) * (right[i] - left[i] - 1)
		}
	}

	return area
}

func trap1(height []int) int {
	var area, prev int
	size := len(height)
	stack := make([]int, 0)

	//  l . . . . . r
	//      lower
	// (min(height[left], height[right]) - height[lower]) * (r - l)

	// initial ascending order cannot form any region
	var i int
	for i = 1; i < size; i++ {
		if height[i] < height[i-1] {
			stack = append(stack, i-1)
			break
		}
	}

	for ; i < size; i++ {
		// keep stack in non-decreasing order
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			prev = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				area += (min(height[stack[len(stack)-1]], height[i]) - height[prev]) * (i - stack[len(stack)-1] - 1)
			}
		}

		stack = append(stack, i)
	}

	return area
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	 to find area, needs height difference, higher left & right boundary

//		the most important thing is that, if one part of condition meets, e.g.
//		lower edge, higher right boundary, still cannot form an area

//	2.	inspired from solution, dp can also be used to solve this problem

//		To find left & right bar higher than self, stack can be used.
//		keep stack in descending order, for every stack element being popped, it
//		means right bar is found. When stack cannot be popped means it's in descending
//		order, left bar is found

//	3.	another way of calculating region is by bar, bar height is determined by
//		smaller of left highest or right highest

//				    _
//		_		  | 4			_
//   	3|		 |				|
//        _		|				|	this is the height contributed by height 1
//		  2|   |				|
//			___					_
//			 1

//		area of bar 1 can conduct is by min height of 3 & 4
//		so, the point is to find highest from left and from right, so dp can be
//		used, because highest is always increasing

//	4.	inspired from solution, when finding left highest & right highest in dp,
//		as long as right highest > left highest, it means region can be determined
//		by left highest

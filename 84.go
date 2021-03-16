package main

// Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.
//
//
//
//
// Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
//
//
//
//
// The largest rectangle is shown in the shaded area, which has area = 10 unit.
//
//
//
// Example:
//
// Input: [2,1,5,6,2,3]
// Output: 10

func largestRectangleArea(heights []int) int {
	// for each number, find left & right boundary such that in this boundary,
	// all numbers are greater or equal to self
	size := len(heights)
	if size == 0 {
		return 0
	}

	left, right := make([]int, size), make([]int, size)

	for i := range heights {
		j := i - 1
		for j >= 0 && heights[j] >= heights[i] {
			j = left[j]
		}
		left[i] = j
	}

	for i := size - 1; i >= 0; i-- {
		j := i + 1
		for j < size && heights[j] >= heights[i] {
			j = right[j]
		}
		right[i] = j
	}

	var maxArea int

	for i := range heights {
		maxArea = max(maxArea, heights[i]*(right[i]-left[i]-1))
	}

	return maxArea
}

func largestRectangleArea2(heights []int) int {
	size := len(heights)
	stack := make([]int, 0)

	var maxArea int
	for i := range heights {
		// keep stack in increasing order
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			start := 0
			if len(stack) > 0 {
				start = stack[len(stack)-1] + 1
			}

			maxArea = max(maxArea, heights[popped]*(i-start))
		}

		stack = append(stack, i)
	}

	// process remaining bars
	if len(stack) > 0 {
		maxArea = max(maxArea, heights[stack[0]]*size)
	}

	for i := 1; i < len(stack); i++ {
		maxArea = max(maxArea, heights[stack[i]]*(size-stack[i-1]-1))
	}

	return maxArea
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

type stack struct {
	data []int
}

func (s *stack) push(i int) {
	s.data = append(s.data, i)
}

func (s *stack) pop() int {
	length := len(s.data)
	if length == 0 {
		return -1
	}
	popped := s.data[length-1]
	s.data = s.data[:length-1]

	return popped
}

func (s *stack) peek() int {
	length := len(s.data)
	if length == 0 {
		return -1
	}
	return s.data[length-1]
}

func largestRectangleArea1(heights []int) int {
	s := stack{
		data: make([]int, 0),
	}

	length := len(heights)
	if length == 0 {
		return 0
	}

	s.push(-1)

	maxArea := 0
	var popped int
	for i, h := range heights {
		// make sure stack is increasing
		if s.peek() == -1 || heights[s.peek()] < h {
			maxArea = max(maxArea, h)
			s.push(i)
			continue
		}

		// encounter lower bar, do some calculation to make sure stack
		// height is still ascending
		for s.peek() != -1 && heights[s.peek()] >= h {
			popped = s.pop()
			// if popped is the last element in stack, then this height
			// is the minimum among specific range
			maxArea = max(maxArea, heights[popped]*(i-1-s.peek()))
		}
		s.push(i)
	}

	// for any ascending histogram
	startIndex := s.peek()
	for {
		popped = s.pop()
		if s.peek() == -1 {
			break
		}
		maxArea = max(maxArea, heights[popped]*(startIndex-s.peek()))
	}

	// the very last item in stack is the minimum among all bars
	if popped != -1 {
		maxArea = max(maxArea, heights[popped]*(startIndex-s.peek()))
	}

	return maxArea
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	when only one element, maxArea is not updated

//	2.	maxHeight can be <=

//	3.	time limit exceeds

//	4.	when finding left/right boundary, break after height is less

//	5.	use dp to reduce complexity, start from right most bar,
//		if heights[i-1] > heights[i], r[i-1] = i
// 	 	if heights[i-1] <= heights[i], find bar that is less height then
//	  	heights[i-1] at index p, r[i-1] = r[p]
//    	r is defaults to length

//	6.	list 5 is wrong, it still needs to calculate left part

// 	7.	optimize, use divide and conquer
//		maximum comes from 3 situations:
//	  	- lowest of all, multiply by all width
// 	  	- left part of lowest multiply by left width
// 	  	- right part of lowest multiply by right width

//	8.	bar can categorized in to 2 conditions:
// 		- increasing
//	 	- decreasing
// 	 	when bar is increasing, push into stack, can update max area value,
//	 	when bar is decreasing, keep popping from stack to calculate area until
// 	 	lower one is encountered.
//	 	finally, deal with stack remaining value that is in ascending order

//	9.	when dealing with last element in stack, that element height is the
//	  	minimum of all range.

//  10.	zero height is the special case, it terminates area calculation

// 	11. optimize, for the struct item, height is not necessary

//	12.	use stack to keep height increasing sequence, if sequence is
//		increasing, it's easier to calculate area

//	13.	inspired form https://leetcode.com/problems/largest-rectangle-in-histogram/discuss/28902/5ms-O(n)-Java-solution-explained-(beats-96)

//		author decompose problem into 2 indicators: for any height, find
//		out its right most & left most index that's >= self

//	14.	this cannot be solved by greedy, e.g. select highest bar, and expand
//		by adjacent bar which is higher...

//		i thought this is similar to 1793...but apparently i am wrong...

//		it fails at this test example: [4,2,0,3,2,4,3,4]

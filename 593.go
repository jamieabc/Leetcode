package main

// Given the coordinates of four points in 2D space, return whether the four points could construct a square.
//
// The coordinate (x,y) of a point is represented by an integer array with two integers.
//
// Example:
//
// Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
// Output: True
//
//
//
// Note:
//
//     All the input integers are in the range [-10000, 10000].
//     A valid square has four equal sides with positive length and four equal angles (90-degree angles).
//     Input points have no order.

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	l1 := lengthSquare(p1, p2)
	l2 := lengthSquare(p1, p3)
	l3 := lengthSquare(p2, p3)
	l4 := lengthSquare(p1, p4)
	l5 := lengthSquare(p2, p4)
	l6 := lengthSquare(p3, p4)

	if l1 == 0 || l2 == 0 || l3 == 0 || l4 == 0 {
		return false
	}

	if l1+l2 == l3 && l1 == l2 {
		return l5+l6 == l4 && l5 == l6
	}

	if l2+l3 == l1 && l2 == l3 {
		return l4+l5 == l6 && l4 == l5
	}

	if l1+l3 == l2 && l1 == l3 {
		return l6+l4 == l5 && l6 == l4
	}

	return false
}

func lengthSquare(p1, p2 []int) int {
	deltaX, deltaY := p1[0]-p2[0], p1[1]-p2[1]

	return deltaX*deltaX + deltaY*deltaY
}

//	problems
//	1.	don't forget to check length = 0

//	2.	inspired from https://leetcode.com/problems/valid-square/discuss/103442/C%2B%2B-3-lines-(unordered_set)

//		a square is that only 2 length exist in all points combinations, and
//		0 length not exist (2 points are same, then first rule applies)

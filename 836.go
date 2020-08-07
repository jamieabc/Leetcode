package main

//A rectangle is represented as a list [x1, y1, x2, y2], where (x1, y1) are the coordinates of its bottom-left corner, and (x2, y2) are the coordinates of its top-right corner.
//
//Two rectangles overlap if the area of their intersection is positive.  To be clear, two rectangles that only touch at the corner or edges do not overlap.
//
//Given two (axis-aligned) rectangles, return whether they overlap.
//
//Example 1:
//
//Input: rec1 = [0,0,2,2], rec2 = [1,1,3,3]
//Output: true
//Example 2:
//
//Input: rec1 = [0,0,1,1], rec2 = [1,0,2,1]
//Output: false
//Notes:
//
//Both rectangles rec1 and rec2 are lists of 4 integers.
//All coordinates in rectangles will be between -10^9 and 10^9.

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return intersect(rec1[0], rec1[2], rec2[0], rec2[2]) && intersect(rec1[1], rec1[3], rec2[1], rec2[3])
}

func intersect(s1, e1, s2, e2 int) bool {
	if s1 == s2 {
		return true
	} else if s1 < s2 {
		return !(s2 >= e1)
	}
	return !(s1 >= e2)
}

type point struct {
	x, y int
}

func isRectangleOverlap1(rec1 []int, rec2 []int) bool {
	p1, p2 := rectangle(rec1)
	p3, p4 := rectangle(rec2)
	o1 := overlap(p1, p2, p3, p4)
	o2 := overlap(p3, p4, p1, p2)

	return o1 || o2
}

func rectangle(data []int) (point, point) {
	return point{
			x: data[0],
			y: data[1],
		}, point{
			x: data[2],
			y: data[3],
		}
}

func overlap(p1, p2, p3, p4 point) bool {
	if p1.x <= p3.x && p3.x < p2.x {
		if p1.y <= p3.y && p3.y < p2.y {
			return true
		}

		if p4.y <= p2.y && p1.y < p4.y {
			return true
		}

		if p3.y <= p1.y && p2.y <= p4.y {
			return true
		}
	}
	return false
}

//	problems
//	1.	overlap is easier to check for non-overlap

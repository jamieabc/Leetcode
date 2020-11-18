package main

// There is a special square room with mirrors on each of the four walls.  Except for the southwest corner, there are receptors on each of the remaining corners, numbered 0, 1, and 2.
//
// The square room has walls of length p, and a laser ray from the southwest corner first meets the east wall at a distance q from the 0th receptor.
//
// Return the number of the receptor that the ray meets first.  (It is guaranteed that the ray will meet a receptor eventually.)
//
//
//
// Example 1:
//
// Input: p = 2, q = 1
// Output: 2
// Explanation: The ray meets receptor 2 the first time it gets reflected back to the left wall.
//
// Note:
//
// 1 <= p <= 1000
// 0 <= q <= p

func mirrorReflection(p int, q int) int {
	// find greatest common divisor (gcd)
	var a, b int
	for a, b = p, q; a != b; {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	p, q = p/a, q/a

	if p&1 > 0 && q&1 > 0 {
		return 1
	} else if p&1 == 0 {
		return 2
	}
	return 0
}

//	Notes
//	1.	inspired from https://www.cnblogs.com/grandyang/p/10646040.html

//		with graph explanation, corner position only relates to odd/even of p&q

//	2.	inspired from https://zh.wikipedia.org/zh-tw/最大公因數

//		some way to find gcd

//		gcd(a, a) = a
//		if a > b, gcd(a, b) = gcd(a-b, b)
//		if a < b, gcd(a, b) = gcd(a, b-a)

//	3.	inspired from https://leetcode.com/problems/mirror-reflection/discuss/141773

//		lee has a great solution, as long as p & q not both even, then anser can
//		be found

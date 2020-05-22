package main

// Given a picture consisting of black and white pixels, find the number of black lonely pixels.
//
// The picture is represented by a 2D char array consisting of 'B' and 'W', which means black and white pixels respectively.
//
// A black lonely pixel is character 'B' that located at a specific position where the same row and same column don't have any other black pixels.
//
// Example:
//
// Input:
// [['W', 'W', 'B'],
// ['W', 'B', 'W'],
// ['B', 'W', 'W']]
//
// Output: 3
// Explanation: All the three 'B's are black lonely pixels.
// Note:
//
// The range of width and height of the input 2D array is [1,500].

func findLonelyPixel(picture [][]byte) int {
	var count int
	y := len(picture)
	if y == 0 {
		return count
	}

	row := make([]int, y)
	column := make([]int, len(picture[0]))

	for i := range picture {
		for j := range picture[0] {
			if picture[i][j] == 'B' {
				row[i]++
				column[j]++
			}
		}
	}

	for i := range row {
		if row[i] != 1 {
			continue
		}
		for j := range column {
			if row[i] == 1 && column[j] == 1 && picture[i][j] == 'B' {
				count++
			}
		}
	}

	return count
}

//	problems
//	1.	in the case of all 'B', the algo is wrong because converts 'B' to 'W'

//	2.	boundary condition for only first line

//	3.	need to consider up row if it's B

//	4.	when checking for alone B, it's universal

//	5.	I didn't read careful about problem, it's same row & column w/o
//		other B. It's easy this way

//	6.	it's same row & column, I only consider same column...

//	7. 	I need another dp for source of column

//	8.	inspired from https://leetcode.com/problems/lonely-pixel-i/discuss/100018/Java-O(nm)-time-with-O(n%2Bm)-Space-and-O(1)-Space-Solutions

//		there's not necessary need source of column, I can scan row == 1 &&
//		column == 1 && picture[row][column] == B to make sure single B
//		comes from that place

//		also, when checking dp row, if its value is not 1, just skip it

//		the space O(1) solution is some nesty, I have come up similar idea
//		to use first row & column to record counts on row/column, but fail
//		because I think first row & column information will be lost. Author
//		just increment W -> X -> Y, B -> C -> D, etc. I think this is smart
//		but w/ some corner cases need to deal with, so didn't take time to
//		traverse.

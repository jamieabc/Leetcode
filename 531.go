package main

import "fmt"

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
	table := make(map[int]int)

	for i := range picture {
		for j := range picture[0] {
			if picture[i][j] == byte('B') {
				// if one union at x = 0, the other union at y = 0,
				// these two union operations will corrupt
				// so, j+1 to avoid error
				union(table, i, -(j + 1))
			}
		}
	}

	counter := make(map[int]int)
	for x := range table {
		// need to do path compress, avoid any number no updated
		counter[find(table, x)]++
	}

	var count int
	for _, v := range counter {
		// x: -y, -y: -y, two records means single B at row & column
		if v == 2 {
			count++
		}
	}

	return count
}

func union(table map[int]int, x, y int) {
	if _, ok := table[x]; !ok {
		table[x] = x
	}

	if _, ok := table[y]; !ok {
		table[y] = y
	}

	table[find(table, x)] = find(table, y)
}

func find(table map[int]int, x int) int {
	if table[x] != x {
		table[x] = find(table, table[x])
	}

	return table[x]
}

// tc: O(mn), sc: O(m+n)
func findLonelyPixel2(picture [][]byte) int {
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

// tc: O(mn), sc: O(m)
func findLonelyPixel1(picture [][]byte) int {
	w, h := len(picture[0]), len(picture)
	col := make([]bool, w)

	var count int

	for i := range picture {
		var b int
		idx := -1

		// check row
		for j := range picture[0] {
			if picture[i][j] == byte('B') {
				b++

				if idx == -1 {
					idx = j
				} else {
					col[idx] = true
					idx = j
				}
			}
		}

		if b == 0 {
			continue
		}

		if b > 1 || col[idx] {
			col[idx] = true
			continue
		}

		// check col
		for j := i + 1; j < h; j++ {
			if picture[j][idx] == byte('B') {
				col[idx] = true
				break
			}
		}

		if !col[idx] {
			count++
		}
		col[idx] = true
	}

	return count
}

//	Notes
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

//	9.	problem is not hard, but takes me 5 wrong submissions

//	10.	inspired from https://leetcode.com/problems/lonely-pixel-i/discuss/390101/Python-union-find-row-scanning

//		author used union-find to group x -> y (in hash, x > 0, y < 0)

package main

// In a N x N grid composed of 1 x 1 squares, each 1 x 1 square consists of a /, \, or blank space.  These characters divide the square into contiguous regions.
//
// (Note that backslash characters are escaped, so a \ is represented as "\\".)
//
// Return the number of regions.
//
//
//
// Example 1:
//
// Input:
// [
//   " /",
//   "/ "
// ]
// Output: 2
// Explanation: The 2x2 grid is as follows:
//
// Example 2:
//
// Input:
// [
//   " /",
//   "  "
// ]
// Output: 1
// Explanation: The 2x2 grid is as follows:
//
// Example 3:
//
// Input:
// [
//   "\\/",
//   "/\\"
// ]
// Output: 4
// Explanation: (Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.)
// The 2x2 grid is as follows:
//
// Example 4:
//
// Input:
// [
//   "/\\",
//   "\\/"
// ]
// Output: 5
// Explanation: (Recall that because \ characters are escaped, "/\\" refers to /\, and "\\/" refers to \/.)
// The 2x2 grid is as follows:
//
// Example 5:
//
// Input:
// [
//   "//",
//   "/ "
// ]
// Output: 3
// Explanation: The 2x2 grid is as follows:
//
//
//
// Note:
//
//     1 <= grid.length == grid[0].length <= 30
//     grid[i][j] is either '/', '\', or ' '.

func regionsBySlashes(grid []string) int {
	y := len(grid)
	if y == 0 {
		return 0
	}

	// split a cell into 4: top(0), right(1), bottom(2), left(3)
	arr := make([]int, y*y*4)
	for i := range arr {
		arr[i] = i
	}
	count := y * y * 4

	for i := range grid {
		for j := range grid[i] {
			// up
			if i > 0 {
				union(&arr, idx(y, i-1, j, 2), idx(y, i, j, 0), &count)
			}

			// left
			if j > 0 {
				union(&arr, idx(y, i, j-1, 1), idx(y, i, j, 3), &count)
			}

			if grid[i][j] != '/' {
				union(&arr, idx(y, i, j, 0), idx(y, i, j, 1), &count)
				union(&arr, idx(y, i, j, 2), idx(y, i, j, 3), &count)
			}

			if grid[i][j] != '\\' {
				union(&arr, idx(y, i, j, 0), idx(y, i, j, 3), &count)
				union(&arr, idx(y, i, j, 2), idx(y, i, j, 1), &count)
			}
		}
	}

	return count
}

func parent(arr *[]int, i int) int {
	if (*arr)[i] != i {
		(*arr)[i] = parent(arr, (*arr)[i])
	}
	return (*arr)[i]
}

func union(arr *[]int, dst, src int, count *int) {
	ps := parent(arr, src)
	pd := parent(arr, dst)
	if ps != pd {
		*count--
		(*arr)[pd] = ps
	}
}

func idx(length, i, j, k int) int {
	return (i*length+j)*4 + k
}

func regionsBySlashes1(grid []string) int {
	y := len(grid)
	if y == 0 {
		return 0
	}

	scaleFactor := 3
	scaled := make([][]int, scaleFactor*y)
	for i := range scaled {
		scaled[i] = make([]int, scaleFactor*y)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '/' {
				for k, l, m := i*scaleFactor, (j+1)*scaleFactor-1, scaleFactor; m > 0; k, l, m = k+1, l-1, m-1 {
					scaled[k][l] = -1
				}
			} else if grid[i][j] == '\\' {
				for k, l, m := i*scaleFactor, j*scaleFactor, scaleFactor; m > 0; k, l, m = k+1, l+1, m-1 {
					scaled[k][l] = -1
				}
			}
		}
	}

	var count int

	for i := range scaled {
		for j := range scaled {
			if scaled[i][j] == 0 {
				count++
				spread(&scaled, i, j, count)
			}
		}
	}

	return count
}

func spread(scaled *[][]int, i, j, count int) {
	if i < 0 || j < 0 || i == len(*scaled) || j == len(*scaled) {
		return
	}

	if (*scaled)[i][j] == 0 {
		(*scaled)[i][j] = count
		spread(scaled, i-1, j, count)
		spread(scaled, i, j-1, count)
		spread(scaled, i, j+1, count)
		spread(scaled, i+1, j, count)
	}
}

// problems
//	1.	totally no idea how to do it, see the first discussion w/ some idea,
//		https://leetcode.com/problems/regions-cut-by-slashes/discuss/205674/C%2B%2B-with-picture-DFS-on-upscaled-grid

//		the idea is to update scale 1 space to 3, so that I can union those
//		connected areas

//		the reason scale factor is 3 because if it's 2, in the case of
//		["//","/ "], there's not enough space to form a region (spread needs
//		at least one space to go forward)

//	2.	inspired from https://leetcode.com/problems/regions-cut-by-slashes/discuss/205680/JavaC%2B%2BPython-Split-4-parts-and-Union-Find

//		this guy is so brilliant, how come he think of such a way of union...
//		counts is started as maximum (n*n*4), each time a union is find, count
//		reduced

//	3.	when union, sequence matters

//	4.	the reason arr is []int and parent is recursive is to find the real parent
//		initially every element in arr is initialised to same number as it self,
//		but as union happens, value changes, it's important to track the real
//		actual parent to see two nearby region should be merged.

//		because unionned data of a cell depend on order of / or \
//		e.g if / is first processed, 0 1 2 3 => 3 1 1 3
//			if \ if first processed, 0 1 2 3 => 1 1 3 3

//		so it's important to find out what causes this triangle to be merged,
//		that's why author uses single length array, and use parent function to
//		find actual parent.

//		I spend 2 hours to figure out this... so clever

//		at first I try to union in 4 directions, but actually only half is
//		needed, then try to figure out this parent method...

//	5.	add reference https://leetcode.com/problems/regions-cut-by-slashes/discuss/205738/Using-Euler's-Formula-(V-E-%2B-F-2)

//		euler's formula: V - E + F = 2
//		didn't go through the solution

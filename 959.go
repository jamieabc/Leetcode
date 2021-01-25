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
	w, h := len(grid[0]), len(grid)

	arr := make([][]byte, h*3)
	for i := range arr {
		arr[i] = make([]byte, w*3)
	}

	// setup upscaled array (3*3 times larger)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {

			if grid[i][j] == '/' {
				for k, l := i*3, (j+1)*3-1; k < (i+1)*3; k, l = k+1, l-1 {
					arr[k][l] = '/'
				}
			} else if grid[i][j] == '\\' {
				for k, l := i*3, j*3; k < (i+1)*3; k, l = k+1, l+1 {
					arr[k][l] = '\\'
				}
			}
		}
	}

	var region int

	visited := make([][]bool, h*3)
	for i := range visited {
		visited[i] = make([]bool, w*3)
	}

	for i := range arr {
		for j := range arr[0] {
			if arr[i][j] == byte(0) && !visited[i][j] {
				region++

				// BFS
				bfs(arr, visited, j, i)
			}
		}
	}

	return region
}

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func bfs(arr [][]byte, visited [][]bool, x, y int) {
	stack := [][]int{{y, x}}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[p[0]][p[1]] {
			visited[p[0]][p[1]] = true

			for _, d := range dirs {
				newY, newX := d[0]+p[0], d[1]+p[1]

				if newX >= 0 && newY >= 0 && newX < len(arr[0]) && newY < len(arr) && arr[newY][newX] == byte(0) && !visited[newY][newX] {
					stack = append(stack, []int{newY, newX})
				}
			}
		}
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
//		counts is started as maximum (n*n*4), each time a union is found, count
//		reduced

//	3.	when union, sequence matters

//	4.	the reason arr is []int and parent is recursive is to find the real parent
//		initially every element in arr is initialised to same number as it self,
//		but as union happens, value changes, it's important to track the real
//		actual parent to see two nearby region should be merged.

//		because unionned data of a cell depend on order of / or \
//		e.g if / is first processed, 0 1 2 3 => 3 1 1 3
//			if \ if first processed, 0 1 2 3 => 1 1 3 3
//		 \0/
//	    3/2\ 1

//		so it's important to find out what causes this triangle to be merged,
//		that's why author uses single length array, and use parent function to
//		find actual parent.

//		I spend 2 hours to figure out this... so clever

//		at first I try to union in 4 directions, but actually only half is
//		needed, then try to figure out this parent method...

//	5.	add reference https://leetcode.com/problems/regions-cut-by-slashes/discuss/205738/Using-Euler's-Formula-(V-E-%2B-F-2)

//		euler's formula: V - E + F = 2
//		didn't go through the solution

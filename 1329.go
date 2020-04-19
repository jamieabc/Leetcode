package main

// Given a m * n matrix mat of integers, sort it diagonally in ascending order from the top-left to the bottom-right then return the sorted array.
//
//
//
// Example 1:
//
// Input: mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
// Output: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]
//
//
//
// Constraints:
//
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 100
//     1 <= mat[i][j] <= 100

func diagonalSort(mat [][]int) [][]int {
	y := len(mat)
	if y == 0 {
		return [][]int{}
	}
	x := len(mat[0])

	var nums [101]int

	for i := y; i >= 0; i-- {
		diagonal(mat, 0, i, &nums)
		replaceDiagonal(mat, 0, i, &nums)
	}

	for i := 1; i < x; i++ {
		diagonal(mat, i, 0, &nums)
		replaceDiagonal(mat, i, 0, &nums)
	}

	return mat
}

func diagonal(mat [][]int, x, y int, nums *[101]int) {
	for count := 0; y+count < len(mat) && x+count < len(mat[0]); count++ {
		nums[mat[y+count][x+count]]++
	}
}

func replaceDiagonal(mat [][]int, x, y int, nums *[101]int) {
	var count int
	for i := range nums {
		for nums[i] > 0 {
			mat[y+count][x+count] = i
			count++
			nums[i]--
		}
	}
}

//	problems
//	1.	Optimize, there's some additional operation for "sort", so I found a clever
//		way of doing sorting.

//		Because each element inside matrix is limit <= 100, so there's a way of
//		sorting: create an array of size 101, and index of the array means specific
//		number is in diagonal

//		e.g. 9
//				3
//					6

//		array index 0 ... 3 .... 6 .... 9 ....
//				    0	  1		 1   	1 ...

//		in this way, sorting might be faster
//	2.	replace number is determined by index, not by count (nums[i])

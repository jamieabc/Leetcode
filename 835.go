package main

import (
	"fmt"
	"math"
)

//Two images A and B are given, represented as binary, square matrices of the same size.  (A binary matrix has only 0s and 1s as values.)
//
//We translate one image however we choose (sliding it left, right, up, or down any number of units), and place it on top of the other image.  After, the overlap of this translation is the number of positions that have a 1 in both images.
//
//(Note also that a translation does not include any kind of rotation.)
//
//What is the largest possible overlap?
//
//Example 1:
//
//Input: A = [[1,1,0],
//            [0,1,0],
//            [0,1,0]]
//       B = [[0,0,0],
//            [0,1,1],
//            [0,0,1]]
//Output: 3
//Explanation: We slide A to right by 1 unit and down by 1 unit.
//
//Notes:
//
//    1 <= A.length = A[0].length = B.length = B[0].length <= 30
//    0 <= A[i][j], B[i][j] <= 1

func main() {
	A := [][]int{
		{0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {0, 0, 0, 1, 0}, {0, 1, 1, 0, 0}, {0, 0, 0, 1, 0},
	}

	B := [][]int{
		{0, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {1, 1, 1, 0, 1}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0},
	}

	fmt.Println(largestOverlap(A, B))
}

// tc: O(2n * 2n * (1+k)), n: length of array, k: average # of 1 in a row
func largestOverlap(A [][]int, B [][]int) int {
	rowsA, rowsB := make([]uint32, len(A)), make([]uint32, len(B))

	for i := range A {
		rowsA[i] = toNum(A[i])
	}

	for i := range B {
		rowsB[i] = toNum(B[i])
	}

	ones := math.MinInt32

	// check for every direction
	for x := -(len(A[0]) - 1); x < len(A[0]); x++ {
		for y := -(len(A) - 1); y < len(A); y++ {
			ones = max(ones, totalOnes(rowsA, rowsB, x, y))
		}
	}

	return ones
}

func rowOne(num uint32) int {
	var ones int

	for num > 0 {
		ones++
		num &= num - 1
	}

	return ones
}

func totalOnes(row1, row2 []uint32, x, y int) int {
	var ones int

	for i := range row1 {
		if (y < 0 && i+y < 0) || (y > 0 && i-y < 0) {
			continue
		}

		if x <= 0 {
			if y <= 0 {
				ones += rowOne(row1[i+y] & (row2[i] << (-x)))
			} else {
				ones += rowOne(row1[i] & (row2[i-y] << (-x)))
			}
		} else {
			if y <= 0 {
				ones += rowOne(row1[i+y] & (row2[i] >> (x)))
			} else {
				ones += rowOne(row1[i] & (row2[i-y] >> (x)))
			}
		}
	}

	return ones
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func toNum(row []int) uint32 {
	var num uint32

	for i := range row {
		if row[i] == 1 {
			num |= 1 << (31 - i)
		}
	}

	return num
}

//	Notes
//	1.	if not specify int32 type, then go might convert number into 64 bits
//		so it won't fits what I expect

//	2.	when align origin, moving up y lines means A[i-y] <-> B[i]
//						   moving down y lines means A[i] <-> B[i-y]

//	3.	inspired from https://leetcode.com/problems/image-overlap/discuss/130623/C%2B%2BJavaPython-Straight-Forward

//		when looping over, all 0 are compared many times, which is a waste
//		of computation, so one optimization is to store 1s index and only
//		compare 1s locations

//		the other thing is to store vectors of 1, since when moving, it's
//		actually to find count with same vectors

//		point a + vector = point b

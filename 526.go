package main

// Suppose you have N integers from 1 to N. We define a beautiful arrangement as an array that is constructed by these N numbers successfully if one of the following is true for the ith position (1 <= i <= N) in this array:
//
//     The number at the ith position is divisible by i.
//     i is divisible by the number at the ith position.
//
//
//
// Now given N, how many beautiful arrangements can you construct?
//
// Example 1:
//
// Input: 2
// Output: 2
// Explanation:
//
// The first beautiful arrangement is [1, 2]:
//
// Number at the 1st position (i=1) is 1, and 1 is divisible by i (i=1).
//
// Number at the 2nd position (i=2) is 2, and 2 is divisible by i (i=2).
//
// The second beautiful arrangement is [2, 1]:
//
// Number at the 1st position (i=1) is 2, and 2 is divisible by i (i=1).
//
// Number at the 2nd position (i=2) is 1, and i (i=2) is divisible by 1.
//
//
//
// Note:
//
//     N is a positive integer and will not exceed 15.

func countArrangement(N int) int {
	var count int
	flags := make([]bool, N)
	mapping := make(map[int][]int)

	// 1 is suitable for any number
	for i := 1; i <= N; i++ {
		mapping[1] = append(mapping[1], i)
	}

	var i, j int
	for i = 2; i <= N; i++ {
		tmp := []int{1, i}
		for j = 2; j <= N; j++ {
			if j != i && (i%j == 0 || j%i == 0) {
				tmp = append(tmp, j)
			}
		}
		mapping[i] = tmp
	}

	permutation(N, &flags, N, mapping, &count)

	return count
}

func permutation(n int, flags *[]bool, length int, mapping map[int][]int, count *int) {
	if length == 1 {
		*count++
		return
	}

	nums := mapping[length]
	for _, j := range nums {
		if !(*flags)[j-1] {
			(*flags)[j-1] = true
			permutation(n, flags, length-1, mapping, count)
			(*flags)[j-1] = false
		}
	}
}

//	problems
//	1.	too slow, due to memory operation

//	2.	no need to track all data, just record it's size

//	3. 	no need to do % operation every time, use a map to store it

//	4.	inspired from https://leetcode.com/problems/beautiful-arrangement/discuss/99711/Java-6ms-beats-98-back-tracking-(swap)-starting-from-the-back

//		start from big number cause it has less choices

//		it can further reduce cause length == 1 is always valid

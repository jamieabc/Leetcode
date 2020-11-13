package main

import "math"

// Given an array arr that represents a permutation of numbers from 1 to n. You have a binary string of size n that initially has all its bits set to zero.
//
// At each step i (assuming both the binary string and arr are 1-indexed) from 1 to n, the bit at position arr[i] is set to 1. You are given an integer m and you need to find the latest step at which there exists a group of ones of length m. A group of ones is a contiguous substring of 1s such that it cannot be extended in either direction.
//
// Return the latest step at which there exists a group of ones of length exactly m. If no such group exists, return -1.
//
//
//
// Example 1:
//
// Input: arr = [3,5,1,2,4], m = 1
// Output: 4
// Explanation:
// Step 1: "00100", groups: ["1"]
// Step 2: "00101", groups: ["1", "1"]
// Step 3: "10101", groups: ["1", "1", "1"]
// Step 4: "11101", groups: ["111", "1"]
// Step 5: "11111", groups: ["11111"]
// The latest step at which there exists a group of size 1 is step 4.
//
// Example 2:
//
// Input: arr = [3,1,5,4,2], m = 2
// Output: -1
// Explanation:
// Step 1: "00100", groups: ["1"]
// Step 2: "10100", groups: ["1", "1"]
// Step 3: "10101", groups: ["1", "1", "1"]
// Step 4: "10111", groups: ["1", "111"]
// Step 5: "11111", groups: ["11111"]
// No group of size 2 exists during any step.
//
// Example 3:
//
// Input: arr = [1], m = 1
// Output: 1
//
// Example 4:
//
// Input: arr = [2,1], m = 2
// Output: 2
//
//
//
// Constraints:
//
//     n == arr.length
//     1 <= n <= 10^5
//     1 <= arr[i] <= n
//     All integers in arr are distinct.
//     1 <= m <= arr.length

func findLatestStep(arr []int, m int) int {
	size := len(arr)
	length, counter := make([]int, size+2), make([]int, size+1)

	latest := -1

	for i, num := range arr {
		left, right := length[num-1], length[num+1]

		// new length is decided by left & right adjacent groups
		length[num] = left + right + 1

		// left & right most boundaries are places that could increase length
		length[num-left] = length[num]
		length[num+right] = length[num]

		// each number is inserted once, length increases one by one
		counter[left]--
		counter[right]--
		counter[length[num]]++

		if counter[m] > 0 {
			latest = i + 1
		}
	}

	return latest
}

func findLatestStep1(arr []int, m int) int {
	size := len(arr)
	groups := make([]int, size+1)
	for i := range groups {
		groups[i] = -1
	}
	counter := make([]int, size+1)
	var latest int

	operate(arr, groups, counter, m, &latest)

	if latest == 0 {
		return -1
	}

	return latest
}

func operate(nums, groups []int, counter []int, m int, latest *int) {
	size := len(groups)
	mapping := make(map[int]int)

	for i, num := range nums {
		counter[num]++
		groups[num] = num
		mapping[counter[num]]++

		left, right := math.MaxInt32, math.MaxInt32

		if num > 1 && groups[num-1] != -1 {
			left = find(groups, num-1)
		}

		if num < size-1 && groups[num+1] != -1 {
			right = find(groups, num+1)
		}

		if left == math.MaxInt32 && right == math.MaxInt32 {
			if mapping[m] > 0 {
				*latest = i + 1
			}
			continue
		}

		if left != math.MaxInt32 && right != math.MaxInt32 {
			mapping[counter[left]]--
			counter[left] += counter[num] + counter[right]
			mapping[counter[left]]++
			mapping[counter[num]]--
			mapping[counter[right]]--
			counter[num], counter[right] = 0, 0
			groups[right], groups[num] = left, left
		} else if left != math.MaxInt32 {
			mapping[counter[left]]--
			counter[left] += counter[num]
			mapping[counter[left]]++
			mapping[counter[num]]--
			counter[num] = 0
			groups[num] = left
		} else {
			mapping[counter[num]]--
			counter[num] += counter[right]
			mapping[counter[num]]++
			mapping[counter[right]]--
			counter[right] = 0
			groups[right] = num
		}

		if mapping[m] > 0 {
			*latest = i + 1
		}
	}
}

func find(groups []int, idx int) int {
	if groups[idx] != idx && groups[idx] != -1 {
		groups[idx] = find(groups, groups[idx])
	}

	return groups[idx]
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/find-latest-group-of-size-m/discuss/806716/C%2B%2B-Union-Find-(Count-groups-of-size-Reverse-mapping)

//		can use rank to choose nodes with less modifications

//	2.	inspired from https://leetcode.com/problems/find-latest-group-of-size-m/discuss/806786/JavaC%2B%2BPython-Count-the-Length-of-Groups-O(N)

//		lee has a brilliant solution, I don't know how he thinks of it

//		for every operation, left continuous ones and right continuous ones
//		are most important thing

//		since every number appears once, gap is bridged gradually, when intervals
//		are connected together in one operation, length increases, original smaller
//		length vanished

//		the point is: no need to memorize every position w/ correct length,
//		but to have correct total length on boundaries wit gap, which means in
//		the future, gap will be fulfilled

//		to avoid boundary checking (0 & size), use larger array (size + 2)

//		update length for farthest, because that's where next increment/decrement
//		happens

//		count should also be calculated

//	3.	inspired from https://leetcode.com/problems/find-latest-group-of-size-m/discuss/806871/C%2B%2B-No-other-algorithm-We-just-iterativly-do-it-from-back-to-front-Very-intuitive-and-clean-!!

//		start from last operation, each time a number cuts an interval

package main

import "sort"

// Given a sorted array arr, two integers k and x, find the k closest elements to x in the array. The result should also be sorted in ascending order. If there is a tie, the smaller elements are always preferred.
//
//
//
// Example 1:
//
// Input: arr = [1,2,3,4,5], k = 4, x = 3
// Output: [1,2,3,4]
// Example 2:
//
// Input: arr = [1,2,3,4,5], k = 4, x = -1
// Output: [1,2,3,4]
//
//
// Constraints:
//
// 1 <= k <= arr.length
// 1 <= arr.length <= 10^4
// Absolute value of elements in the array and x will not exceed 104

func findClosestElements(arr []int, k int, x int) []int {
	low, high := 0, len(arr)-k

	for low < high {
		mid := low + (high-low)/2

		if x-arr[mid] > arr[mid+k]-x {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return arr[low : low+k]
}

func findClosestElements2(arr []int, k int, x int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}

	idx := binarySearch(arr, x)
	start := max(0, idx-k)
	end := min(len(arr)-1, start+k-1)
	var minSum, sum int
	for i := start; i <= end; i++ {
		sum += dist(arr[i], x)
	}
	minSum = sum
	minStart, minEnd := start, end

	// sweep to right most range
	for i := 1; i <= k && end+i < len(arr); i++ {
		sum -= dist(x, arr[start+i-1])
		sum += dist(x, arr[end+i])

		if sum < minSum {
			minSum = sum
			minStart, minEnd = start+i, end+i
		}
	}

	return arr[minStart : minEnd+1]
}

func dist(i, j int) int {
	return abs(i - j)
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

// 1, 3, 6, 7, 9 => 5
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low < high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func findClosestElements1(arr []int, k int, x int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}

	result := make([]int, 0)
	idx := binarySearch(arr, x)
	var count, left, right int
	if arr[idx] == x {
		result = append(result, arr[idx])
		count++
		left, right = idx-1, idx+1
	} else {
		left, right = idx-1, idx
	}

	for next := idx; count < k; count++ {
		validLeft, validRight := validRange(arr, left), validRange(arr, right)
		if !validLeft && !validRight {
			break
		}

		if validLeft && validRight {
			if dist(arr[left], x) <= dist(arr[right], x) {
				next = left
				left--
			} else {
				next = right
				right++
			}
		} else if validLeft {
			next = left
			left--
		} else {
			next = right
			right++
		}

		result = append(result, arr[next])
	}

	sort.Ints(result)
	return result
}

func dist(i, j int) int {
	return abs(i - j)
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func validRange(arr []int, idx int) bool {
	return idx >= 0 && idx < len(arr)
}

// 1, 3, 6, 7, 9 => 5
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low < high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

//	problems
//	1.	first algorithm: use binary search to find number <= target, then choose
//		closest number w/ smallest distance, sort result again

//	2.	too slow, no need to sort result again, after finding index, it's a
//		problem of dp => start from left possible number to right possible number,
//		find smallest sum

//		tc: O(n)

//	3.	inspired from https://leetcode.com/problems/find-k-closest-elements/discuss/106426/JavaC%2B%2BPython-Binary-Search-O(log(N-K)-%2B-K)

//		lee uses a clever way to find starting point by binary search, some cases:

//		- ---- x ---- A[mid] -------- A[mid+k] ----   			move left
//		- ---- A[mid] -- x --------- A[mid+k] ----				move left
//		- ---- A[mid] ------------- x --- A[mid+k] ----			move right
//		- ---- A[mid] ------------------- A[mid+k] -- x ----	move right

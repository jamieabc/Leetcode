package main

import "sort"

// Given an integer array arr, and an integer target, return the number of tuples i, j, k such that i < j < k and arr[i] + arr[j] + arr[k] == target.
//
// As the answer can be very large, return it modulo 109 + 7.
//
//
//
// Example 1:
//
// Input: arr = [1,1,2,2,3,3,4,4,5,5], target = 8
// Output: 20
// Explanation:
// Enumerating by the values (arr[i], arr[j], arr[k]):
// (1, 2, 5) occurs 8 times;
// (1, 3, 4) occurs 8 times;
// (2, 2, 4) occurs 2 times;
// (2, 3, 3) occurs 2 times.
//
// Example 2:
//
// Input: arr = [1,1,2,2,2,2], target = 5
// Output: 12
// Explanation:
// arr[i] = 1, arr[j] = arr[k] = 2 occurs 12 times:
// We choose one 1 from [1,1] in 2 ways,
// and two 2s from [2,2,2,2] in 6 ways.
//
//
//
// Constraints:
//
// 3 <= arr.length <= 3000
// 0 <= arr[i] <= 100
// 0 <= target <= 300

// tc: O(n+m^2), n: arr length, m: range of numbers
func threeSumMulti(arr []int, target int) int {
	counter := make(map[int]int)
	for _, i := range arr {
		counter[i]++
	}

	var ans int
	mod := int(1e9 + 7)

	for i := 0; i <= 100; i++ {
		for j := i; j <= 100; j++ {
			// cannot find next number such that all 3 sums to target
			k := target - i - j
			if _, ok := counter[k]; !ok {
				continue
			}

			if i == j && j == k {
				ans = (ans + counter[j]*(counter[j]-1)*(counter[j]-2)/6) % mod
			} else if i == j && j != k {
				ans = (ans + counter[i]*(counter[i]-1)/2*counter[k]) % mod
			} else if i < j && j < k {
				ans = (ans + counter[i]*counter[j]*counter[k]) % mod
			}
		}
	}

	return ans
}

// tc: O(n^2), sc: O(1)
func threeSumMulti1(arr []int, target int) int {
	size := len(arr)
	sort.Ints(arr)
	var ans int
	mod := int(1e9 + 7)

	for i := range arr {
		goal := target - arr[i]

		for j, k := i+1, size-1; j < k; {
			sum := arr[j] + arr[k]

			if sum < goal {
				j++
			} else if sum > goal {
				k--
			} else {
				if arr[j] != arr[k] {
					left, right := j+1, k-1
					for ; left < k && arr[left] == arr[j]; left++ {
					}
					for ; right > j && arr[right] == arr[k]; right-- {
					}
					ans = (ans + (left-j)*(k-right)) % mod
					j, k = left, right
				} else {
					ans = (ans + (k-j+1)*(k-j)/2) % mod
					break
				}
			}
		}
	}

	return ans
}

// tc: O(n^2), sc: O(n)
func threeSumMulti1(arr []int, target int) int {
	mod := int(1e9 + 7)
	var ans int
	size := len(arr)

	for i := range arr {
		table := make(map[int]int)
		goal := target - arr[i]

		for j := i + 1; j < size; j++ {
			if count, ok := table[goal-arr[j]]; ok {
				ans = (ans + count) % mod
			}
			table[arr[j]]++
		}
	}

	return ans
}

//	Notes
//	1.	inspired from solution, since there are duplicates, can use sorted array
//		property to find, but need to deal with duplicate cases: two numbers are
//		same, two numbers are different with duplicates

//	2.	inspired form https://leetcode.com/problems/3sum-with-multiplicity/discuss/181131/C%2B%2BJavaPython-O(N-%2B-101-*-101)

//		lee uses number occurrence counter to find possible combination conditions:
//		- a == b && b == c
//		- a == b && b != c
//		- a < b && b < c

//		the last condition is for 3 numbers are different, a + b + c = target
//		take 1 + 2 + 3 = 6 as an example
//		(1) 1, 2, 3
//		(2) 1, 3, 2
//		(3) 2, 1, 3
//		(4) 2, 3, 1
//		(5) 3, 1, 2
//		(6) 3, 2, 1

//		because a != b && b != c && a != c, for above 6 conditions
//		select a > b && b < c condition, there will be 2: (3) & (6) results in repeated count
//		select a < b && b > c condition, there will be 2: (2) & (4) results in repeated count
//		select a < b && b < c condition, there will be 1: (1)
//		select a > b && b > c condition, there will be 1: (6)
//		so choose either (1) or (6) condition are fine, as long as not duplicate count

//	3.	inspired from https://leetcode.com/problems/3sum-with-multiplicity/discuss/181125/Knapsack-O(n-*-target)-or-Straightforward-O(n2)

//		author uses dp to solve the problem, not implement

package main

// You are given an array of positive integers nums and want to erase a subarray containing unique elements. The score you get by erasing the subarray is equal to the sum of its elements.
//
// Return the maximum score you can get by erasing exactly one subarray.
//
// An array b is called to be a subarray of a if it forms a contiguous subsequence of a, that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).
//
//
//
// Example 1:
//
// Input: nums = [4,2,4,5,6]
// Output: 17
// Explanation: The optimal subarray here is [2,4,5,6].
//
// Example 2:
//
// Input: nums = [5,2,1,2,5,2,1,2,5]
// Output: 8
// Explanation: The optimal subarray here is [5,2,1] or [1,2,5].
//
//
//
// Constraints:
//
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 104

// tc: O(n)
func maximumUniqueSubarray(nums []int) int {
	size := len(nums)

	var i, j, maximum, sum int
	table := make(map[int]int)

	for j < size {
		for ; j < size; j++ {
			if val, ok := table[nums[j]]; ok && val != -1 {
				// duplicates, shrink
				for ; i <= val; i++ {
					table[nums[i]] = -1
					sum -= nums[i]
				}

				break
			} else {
				sum += nums[j]
				maximum = max(maximum, sum)
				table[nums[j]] = j
			}
		}
	}

	return maximum
}

func max(i, j int) int {
	if i >= j {
		return i
	}

	return j
}

// tc: O(n)
func maximumUniqueSubarray1(nums []int) int {
	size := len(nums)
	sums := make([]int, size+1)
	var sum int

	for i := range nums {
		sums[i] = sum
		sum += nums[i]
	}
	sums[size] = sum

	var i, j, maximum int
	table := make(map[int]int)

	for j < size {
		for ; j < size; j++ {
			if val, ok := table[nums[j]]; ok && val != -1 {
				// duplicates
				maximum = max(maximum, sums[j]-sums[i])

				for ; i <= val; i++ {
					table[nums[i]] = -1
				}

				break
			} else {
				table[nums[j]] = j
			}
		}
	}

	return max(maximum, sums[size]-sums[i])
}

func max(i, j int) int {
	if i >= j {
		return i
	}

	return j
}

//	Notes
//	1.	this problem can be solved by 2-ptr, because for an sub-array
//		. . . i . . . dup. . j . . .
//		start from i, index dup & j are same, the maximum range (all numbers are
//		positive) are i ~ j-1, then for next iteration, start from dup+1

//		the mistake i have is to sum from dup~i only

//	2.	inspired form sample code, since i & j can shrink or expand one by one,
//		no need to pre calculate sums, just do it while doing 2-ptr check

//	3.	inspired from sample code, could use bucket to store number index, instead
//		of hash

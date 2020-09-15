package main

//You are given two sorted arrays of distinct integers nums1 and nums2.
//
//A valid path is defined as follows:
//
//Choose array nums1 or nums2 to traverse (from index-0).
//Traverse the current array from left to right.
//If you are reading any value that is present in nums1 and nums2 you are allowed to change your path to the other array. (Only one repeated value is considered in the valid path).
//Score is defined as the sum of uniques values in a valid path.
//
//Return the maximum score you can obtain of all possible valid paths.
//
//Since the answer may be too large, return it modulo 10^9 + 7.
//
//
//
//Example 1:
//
//
//
//Input: nums1 = [2,4,5,8,10], nums2 = [4,6,8,9]
//Output: 30
//Explanation: Valid paths:
//[2,4,5,8,10], [2,4,5,8,9], [2,4,6,8,9], [2,4,6,8,10],  (starting from nums1)
//[4,6,8,9], [4,5,8,10], [4,5,8,9], [4,6,8,10]    (starting from nums2)
//The maximum is obtained with the path in green [2,4,6,8,10].
//Example 2:
//
//Input: nums1 = [1,3,5,7,9], nums2 = [3,5,100]
//Output: 109
//Explanation: Maximum sum is obtained with the path [1,3,5,100].
//Example 3:
//
//Input: nums1 = [1,2,3,4,5], nums2 = [6,7,8,9,10]
//Output: 40
//Explanation: There are no common elements between nums1 and nums2.
//Maximum sum is obtained with the path [6,7,8,9,10].
//Example 4:
//
//Input: nums1 = [1,4,5,8,9,11,19], nums2 = [2,3,4,11,12]
//Output: 61
//
//
//Constraints:
//
//1 <= nums1.length <= 10^5
//1 <= nums2.length <= 10^5
//1 <= nums1[i], nums2[i] <= 10^7
//nums1 and nums2 are strictly increasing.
//Accepted

func maxSum(nums1 []int, nums2 []int) int {
	var i, j int
	var sum, s1, s2 int64

	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && (j == len(nums2) || nums1[i] < nums2[j]) {
			s1 += int64(nums1[i])
			i++
		} else if j < len(nums2) && (i == len(nums1) || nums1[i] > nums2[j]) {
			s2 += int64(nums2[j])
			j++
		} else {
			sum += max(s1, s2) + int64(nums1[i])
			s1, s2 = 0, 0
			i, j = i+1, j+1
		}
	}

	return int((sum + max(s1, s2)) % (1e9 + 7))
}

func max(i, j int64) int64 {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	one test case fail, inspired from https://leetcode.com/problems/get-the-maximum-score/discuss/768025/Python-Why-is-it-valid-to-only-mod-at-the-end

//		cannot use mod during adding intermediate value, because it will cause wrong
//		checks, e.g. 4 > 3, 4 % 4 < 3 % 4

//	2.	inspired from https://leetcode.com/problems/get-the-maximum-score/discuss/767987/JavaC%2B%2BPython-Two-Pointers-O(1)-Space

//		can put all adding logic into for loop, be careful about boundary
//		conditions, what if any pointer reaches end of array

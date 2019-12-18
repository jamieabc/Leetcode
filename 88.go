package main

//Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
//
//Note:
//
//    The number of elements initialized in nums1 and nums2 are m and n respectively.
//    You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
//
//Example:
//
//Input:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//Output: [1,2,2,3,5,6]

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	total := m + n
	// m = 3, n = 2, total = 5

	// move original nums1 element to end of nums1
	for i := 0; i < m; i++ {
		nums1[total-1-i] = nums1[m-1-i]
	}

	// compare nums1 & nums2 and fill into nums 1
	// i is original nums1 content index, start from n
	// j is nums2 content index
	// k is new nums1 index start from 0
	for i, j, k := n, 0, 0; i < n+m || j < n; k++ {
		if i < n+m && j < n {
			if nums1[i] <= nums2[j] {
				nums1[k] = nums1[i]
				i++
			} else {
				nums1[k] = nums2[j]
				j++
			}
			continue
		}

		if i == n+m {
			nums1[k] = nums2[j]
			j++
			continue
		}

		if j == n {
			nums1[k] = nums1[i]
			i++
			continue
		}
	}
}

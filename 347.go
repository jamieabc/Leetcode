package main

//Given a non-empty array of integers, return the k most frequent elements.
//
//Example 1:
//
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
//Example 2:
//
//Input: nums = [1], k = 1
//Output: [1]
//Note:
//
//You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
//Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

// iterate through all elements, create a map to store frequency
// create an array size of nums, to store frequency based on index
// e.g. index 0 means 0 times, index 5 means 5 times, if there exist
// multiple numbers with same frequency, store it in an array
func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)

	// map for frequency
	for _, n := range nums {
		if _, ok := frequency[n]; !ok {
			frequency[n] = 1
		} else {
			frequency[n]++
		}
	}

	mapping := make([][]int, len(nums)+1)
	for key, val := range frequency {
		mapping[val] = append(mapping[val], key)
	}

	result := make([]int, 0)

	for i := len(mapping) - 1; k > 0; i-- {
		length := len(mapping[i])
		if length > 0 {
			if length > k {
				result = append(result, mapping[i][:k]...)
			} else {
				result = append(result, mapping[i]...)
			}
			k -= length
		}
	}

	return result
}

//	problems
//	1.	should use k to search backward, it's more straight forward

//	2.	reference from https://leetcode.com/problems/top-k-frequent-elements/discuss/81635/3-Java-Solution-using-Array-MaxHeap-TreeMap

//		there are 3 ways to solve this, bucket, heap, treemap

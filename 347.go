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
	count := 0
	for count < k {
		for i := len(mapping) - 1; i >= 0; i-- {
			if len(mapping[i]) == 0 {
				continue
			}
			selectedCount := k - count
			if selectedCount >= len(mapping[i]) {
				result = append(result, mapping[i]...)
				count += len(mapping[i])
			} else {
				result = append(result, mapping[i][:selectedCount]...)
				count += selectedCount
			}
		}
	}
	return result
}

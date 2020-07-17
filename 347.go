package main

import "fmt"

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

type num struct {
	val, count int
}

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
	}

	counts := make([]num, 0)
	for n, count := range counter {
		counts = append(counts, num{n, count})
	}

	quickSelect(counts, k, 0, len(counts)-1)

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, counts[i].val)
	}

	return result
}

func quickSelect(nums []num, target, start, end int) {
	if start >= end {
		return
	}

	idx := partition(nums, target, start, end)
	if idx == target {
		return
	} else if idx < target {
		quickSelect(nums, target, idx+1, end)
	} else {
		quickSelect(nums, target, start, idx)
	}
}

func partition(nums []num, target, start, end int) int {
	if start >= end {
		return start
	}

	pivot := nums[start]
	nums[end], nums[start] = nums[start], nums[end]

	// valid items are all before store
	store := start
	for i := start; i <= end; i++ {
		if nums[i].count > pivot.count {
			nums[i], nums[store] = nums[store], nums[i]
			store++
		}
	}

	nums[store], nums[end] = nums[end], nums[store]
	return store
}

type nums []num

func (this nums) Len() int           { return len(this) }
func (this nums) Less(i, j int) bool { return this[i].count > this[j].count }
func (this nums) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }

func (this *nums) Push(x interface{}) {
	*this = append(*this, x.(num))
}

func (this *nums) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[:n-1]
	return x
}

func topKFrequent2(numbers []int, k int) []int {
	counter := make(map[int]int)
	for _, n := range numbers {
		counter[n]++
	}

	n := &nums{}
	heap.Init(n)
	for key, value := range counter {
		heap.Push(n, num{key, value})
	}

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(n).(num).val)
	}

	return result
}

func topKFrequent1(nums []int, k int) []int {
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

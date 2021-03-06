package main

import "container/heap"

// Given an array arr.  You can choose a set of integers and remove all the occurrences of these integers in the array.
//
// Return the minimum size of the set so that at least half of the integers of the array are removed.
//
//
//
// Example 1:
//
// Input: arr = [3,3,3,3,5,5,5,2,2,7]
// Output: 2
// Explanation: Choosing {3,7} will make the new array [5,5,5,2,2] which has size 5 (i.e equal to half of the size of the old array).
// Possible sets of size 2 are {3,5},{3,2},{5,2}.
// Choosing set {2,7} is not possible as it will make the new array [3,3,3,3,5,5,5] which has size greater than half of the size of the old array.
//
// Example 2:
//
// Input: arr = [7,7,7,7,7,7]
// Output: 1
// Explanation: The only possible set you can choose is {7}. This will make the new array empty.
//
// Example 3:
//
// Input: arr = [1,9]
// Output: 1
//
// Example 4:
//
// Input: arr = [1000,1000,3,7]
// Output: 1
//
// Example 5:
//
// Input: arr = [1,2,3,4,5,6,7,8,9,10]
// Output: 5
//
//
//
// Constraints:
//
//     1 <= arr.length <= 10^5
//     arr.length is even.
//     1 <= arr[i] <= 10^5

func minSetSize(arr []int) int {
	counter := make(map[int]int)
	var maxFreq int
	for _, num := range arr {
		counter[num]++
		maxFreq = max(maxFreq, counter[num])
	}

	// becareful, there could be more numbers with same frequency
	buckets := make([]int, maxFreq+1)
	for _, occurrence := range counter {
		buckets[occurrence]++
	}

	target := len(arr) >> 1
	cur := len(arr)
	var ans int

	for i := len(buckets) - 1; i >= 0 && cur > target; i-- {
		for j := 0; j < buckets[i] && cur > target; j++ {
			ans++
			cur -= i
		}
	}

	return ans
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int          { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	(*h) = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	popped := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return popped
}

// tc: O(n log(k)), k: number of unique numbers
// tc: O(k)
func minSetSize1(arr []int) int {
	// frequency counter
	freq := make(map[int]int)
	for _, n := range arr {
		freq[n]++
	}

	// remove from highest frequency
	mh := &MaxHeap{}
	heap.Init(mh)
	for _, occurCount := range freq {
		heap.Push(mh, occurCount)
	}

	var ans int
	target := len(arr) >> 1
	for cur := len(arr); mh.Len() > 0 && cur > target; {
		ans++
		cur -= mh.Peek()
		heap.Pop(mh)
	}

	return ans
}

//	Notes
//	1.	inspired from solution, when sort by heap. it could potentially use
//		bucket sort, it's space is decided by largest frequency

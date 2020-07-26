package main

// Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
//
// Example 1:
//
// Input: [3,2,1,5,6,4] and k = 2
// Output: 5
//
// Example 2:
//
// Input: [3,2,3,1,2,4,5,5,6] and k = 4
// Output: 4
//
// Note:
// You may assume k is always valid, 1 ≤ k ≤ array's length.
func findKthLargest(nums []int, k int) int {
	quickSelect(nums, 0, len(nums)-1, k)

	return nums[k-1]
}

func quickSelect(nums []int, start, end, target int) {
	if start >= end {
		return
	}

	pivot := partition(nums, start, end)
	if pivot == target-1 {
		return
	} else if pivot < target {
		quickSelect(nums, pivot+1, end, target)
	} else {
		quickSelect(nums, start, pivot-1, target)
	}
}

func partition(nums []int, start, end int) int {
	nums[start], nums[end] = nums[end], nums[start]
	storeAt := start

	for i := start; i < end; i++ {
		if nums[i] >= nums[end] {
			nums[i], nums[storeAt] = nums[storeAt], nums[i]
			storeAt++
		}
	}

	nums[storeAt], nums[end] = nums[end], nums[storeAt]

	return storeAt
}

func findKthLargest1(nums []int, k int) int {
	length := len(nums)
	start, end, idx := 0, length-1, length-k
	for start < end {
		tmp := partition(nums, start, end)
		if tmp > idx {
			end = tmp - 1
		} else if tmp < idx {
			start = tmp + 1
		} else {
			return nums[tmp]
		}
	}
	return nums[idx]
}

func partition(nums []int, start, end int) int {
	pivot := start
	for start <= end {
		for start <= end && nums[start] <= nums[pivot] {
			start++
		}
		for start <= end && nums[end] > nums[pivot] {
			end--
		}

		if start <= end {
			nums[start], nums[end] = nums[end], nums[start]
		}
	}
	nums[pivot], nums[end] = nums[end], nums[pivot]
	return end
}

// min heap
type heap struct {
	data []int
	size int
}

func (h *heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *heap) left(i int) int {
	l := (i * 2) + 1
	if l >= len(h.data) {
		return -1
	}
	return l
}

func (h *heap) right(i int) int {
	r := (i * 2) + 2
	if r >= len(h.data) {
		return -1
	}
	return r
}

func (h *heap) bubbleUp(i int) {
	if i == 0 {
		h.bubbleDown(0)
		return
	}
	p := h.parent(i)

	if h.data[p] > h.data[i] {
		h.swap(p, i)
		h.bubbleUp(p)
	} else {
		h.bubbleDown(i)
	}
}

func (h *heap) bubbleDown(i int) {
	l, r := h.left(i), h.right(i)

	if l == -1 && r == -1 {
		return
	} else if l != -1 && r != -1 {
		if h.data[i] <= h.data[l] && h.data[i] <= h.data[r] {
			return
		}

		if h.data[l] <= h.data[r] {
			h.swap(l, i)
			h.bubbleDown(l)
		} else {
			h.swap(r, i)
			h.bubbleDown(r)
		}
	} else if r == -1 {
		if h.data[l] <= h.data[i] {
			h.swap(l, i)
		}
	}
}

func (h *heap) swap(dst, src int) {
	h.data[src], h.data[dst] = h.data[dst], h.data[src]
}

func (h *heap) insert(n int) {
	if len(h.data) < h.size {
		// not full
		h.data = append(h.data, n)
		h.bubbleUp(len(h.data) - 1)
	} else {
		if n > h.data[0] {
			h.swap(0, len(h.data)-1)
			h.data[len(h.data)-1] = n
			h.bubbleDown(0)
			h.bubbleUp(len(h.data) - 1)
		}
	}
}

func findKthLargest1(nums []int, k int) int {
	h := &heap{
		data: make([]int, 0),
		size: k,
	}

	for _, n := range nums {
		h.insert(n)
	}

	return h.data[0]
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	problems
//	1.	when finding kth largest number, I should use min heap to make sure
//		all other number in the heap are larger, means root number is kth
//		largest

//	2.	inspired from https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/60300/Java-Quick-Select

//		quick-sort like mechanism can be used here
//		time complexity is O(n), it is reduced half each time,
//		n + n/2 + n/4 + n/8 + ... + 1 = n + n-1

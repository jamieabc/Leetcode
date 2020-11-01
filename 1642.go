package main

// You are given an integer array heights representing the heights of buildings, some bricks, and some ladders.
//
// You start your journey from building 0 and move to the next building by possibly using bricks or ladders.
//
// While moving from building i to building i+1 (0-indexed),
//
//     If the current building's height is greater than or equal to the next building's height, you do not need a ladder or bricks.
//     If the current building's height is less than the next building's height, you can either use one ladder or (h[i+1] - h[i]) bricks.
//
// Return the furthest building index (0-indexed) you can reach if you use the given ladders and bricks optimally.
//
//
//
// Example 1:
//
// Input: heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
// Output: 4
// Explanation: Starting at building 0, you can follow these steps:
// - Go to building 1 without using ladders nor bricks since 4 >= 2.
// - Go to building 2 using 5 bricks. You must use either bricks or ladders because 2 < 7.
// - Go to building 3 without using ladders nor bricks since 7 >= 6.
// - Go to building 4 using your only ladder. You must use either bricks or ladders because 6 < 9.
// It is impossible to go beyond building 4 because you do not have any more bricks or ladders.
//
// Example 2:
//
// Input: heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
// Output: 7
//
// Example 3:
//
// Input: heights = [14,3,19,3], bricks = 17, ladders = 0
// Output: 3
//
//
//
// Constraints:
//
//     1 <= heights.length <= 105
//     1 <= heights[i] <= 106
//     0 <= bricks <= 109
//     0 <= ladders <= heights.length

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int          { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &MaxHeap{}
	heap.Init(h)

	size := len(heights)
	if size == 1 {
		return 0
	}

	for i := 1; i < size; i++ {
		diff := heights[i] - heights[i-1]

		if diff > 0 {
			if bricks >= diff {
				bricks -= diff
				heap.Push(h, diff)
			} else {
				if ladders == 0 {
					return i - 1
				}

				ladders--
				heap.Push(h, diff)
				pop := heap.Pop(h).(int)

				// previous crossed gap larger, replace that with ladder and use
				// part of that bricks to cover next jump
				if pop != diff {
					bricks += pop - diff
				}
			}
		}
	}

	return size - 1
}

//	Notes
//	1.	ladders can be treat as single use of infinite gap, so better to use it
//		for as larger gap as possible

//		the whole behavior should use bricks first, if bricks is not enough,
//		replace bricks with ladders

//		what I first write is using deque, although passes tests, but it's not
//		correct, considering following test case:
//		heights = [2, 7, 9, 12]
//		bricks = 5
//		ladders = 1

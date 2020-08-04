package main

// Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing n - 1 elements by 1.

// Example:

// Input:
// [1,2,3]

// Output:
// 3

// Explanation:
// Only three moves are needed (remember each move increments two elements):

// [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]

func minMoves(nums []int) int {
	var sum, minCount int
	minimum := math.MaxInt32

	for _, n := range nums {
		if n < minimum {
			minimum = n
			minCount = 1
		} else if n == minimum {
			minCount++
		}

		sum += n
	}

	sum -= len(nums) * minimum

	return sum
}


type MaxHeap  []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int { return h[0] }

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

func minMoves2(nums []int) int {
	h := &MaxHeap{}
	heap.Init(h)
	minimum := math.MaxInt32

	for _, n := range nums {
		minimum = min(minimum, n)
		heap.Push(h, n)
	}

	var count int

	for h.Len() > 0 {
		if h.Peek() == minimum {
			break
		}

		count += h.Peek() - minimum
		heap.Pop(h)
	}

	return count
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func minMoves1(nums []int) int {
	var count int
	for true {
		tmp := rebalance(nums)
		if tmp == 0 {
			return count
		}

		count += tmp
	}

	return count
}

func rebalance(nums []int) int {
	maximum, minimum := nums[0], nums[0]
	var maxIdx int

	for i := 1; i < len(nums); i++ {
		if nums[i] > maximum {
			maximum = nums[i]
			maxIdx = i
		}

		if nums[i] < maximum {
			minimum = nums[i]
		}
	}

	if maximum == minimum {
		return 0
	}

	count := maximum - minimum
	nums[maxIdx] -= count

	return count
}

//	problems
//	1.	add other numbers 1 means subtract maximum number by 1, total
//		difference is still same, but operation can reduce a lot (from
//      n-1 to 1)

//		tc: O(n*m), m: count of number differ than minimum

//	2.	use max heap to keep track of maximum number, every time count
//		is incremented, pop maximum number

//		tc: O(n log n)

//	3.	since every iteration is to remove maximum number, I can just
//		iterate through every number and count difference to minimum

//		tc: O(n)

//	4.	no need to iterate 2 times, just iterate 1 times, have sum of
//		all numbers, and subtract from minimum number

//	5.	inspired from solution, sum of all numbers could overflow,
//      iterate two times to get answer

//	6.	inspired from solution, this operation can also use dp to solve

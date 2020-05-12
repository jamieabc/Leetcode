package main

// Given two 1d vectors, implement an iterator to return their elements alternately.
//
//
//
// Example:
//
// Input:
// v1 = [1,2]
// v2 = [3,4,5,6]
// Output: [1,3,2,4,5,6]
// Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,3,2,4,5,6].
//
//
//
// Follow up:
//
// What if you are given k 1d vectors? How well can your code be extended to such cases?
//
// Clarification for the follow up question:
// The "Zigzag" order is not clearly defined and is ambiguous for k > 2 cases. If "Zigzag" does not look right to you, replace "Zigzag" with "Cyclic". For example:
//
// Input:
// [1,2,3]
// [4,5,6,7]
// [8,9]
//
// Output: [1,4,8,2,5,9,3,6,7].

type ZigzagIterator struct {
	ptrs    []int
	values  [][]int
	size    int
	current int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	return &ZigzagIterator{
		ptrs:   []int{0, 0},
		values: [][]int{v1, v2},
		size:   2,
	}
}

func (this *ZigzagIterator) next() int {
	val := this.values[this.current][this.ptrs[this.current]]
	this.ptrs[this.current]++
	this.current = (this.current + 1) % this.size

	return val
}

func (this *ZigzagIterator) hasNext() bool {
	for i := 0; i < this.size; i++ {
		if len(this.values[this.current]) > this.ptrs[this.current] {
			return true
		} else {
			this.ptrs[this.current]++
			this.current = (this.current + 1) % this.size
		}
	}
	return false
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */

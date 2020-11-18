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
	Data                   [][]int
	Row, Column, MaxColumn int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	data := make([][]int, 0)
	if len(v1) > 0 {
		data = append(data, v1)
	}

	if len(v2) > 0 {
		data = append(data, v2)
	}

	z := &ZigzagIterator{
		Data:      data,
		MaxColumn: max(len(v1), len(v2)),
	}

	return z
}

func (this *ZigzagIterator) next() int {
	val := this.Data[this.Row][this.Column]
	this.findNext()

	return val
}

func (this *ZigzagIterator) findNext() {
	for this.Column < this.MaxColumn {
		if this.Row >= len(this.Data)-1 {
			this.Column++
			this.Row = 0
		} else {
			this.Row++
		}

		if len(this.Data[this.Row]) > this.Column {
			break
		}
	}
}

func (this *ZigzagIterator) hasNext() bool {
	return this.Column < this.MaxColumn
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

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

//	Notes
//	1.	cannot assume first array is always valid

//	2.	inspired from https://leetcode.com/problems/zigzag-iterator/discuss/71779/Simple-Java-solution-for-K-vector

//		only add array not empty, because empty list cannot be used anyway, it's
//		easier to deal with corner case, instead of add array and check later

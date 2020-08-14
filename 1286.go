package main

// Design an Iterator class, which has:
//
// A constructor that takes a string characters of sorted distinct lowercase English letters and a number combinationLength as arguments.
// A function next() that returns the next combination of length combinationLength in lexicographical order.
// A function hasNext() that returns True if and only if there exists a next combination.
//
//
// Example:
//
// CombinationIterator iterator = new CombinationIterator("abc", 2); // creates the iterator.
//
// iterator.next(); // returns "ab"
// iterator.hasNext(); // returns true
// iterator.next(); // returns "ac"
// iterator.hasNext(); // returns true
// iterator.next(); // returns "bc"
// iterator.hasNext(); // returns false
//
//
// Constraints:
//
// 1 <= combinationLength <= characters.length <= 15
// There will be at most 10^4 function calls per test.
// It's guaranteed that all calls of the function next are valid.

type CombinationIterator struct {
	Str   string
	Stack []int
	Size  int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	return CombinationIterator{
		Str:   characters,
		Size:  combinationLength,
		Stack: make([]int, 0),
	}
}

func (this *CombinationIterator) Next() string {
	result := make([]byte, this.Size)

	if len(this.Stack) == 0 {
		// first time
		for i := 0; i < this.Size; i++ {
			result[i] = this.Str[i]
			this.Stack = append(this.Stack, i)
		}
		return string(result)
	}

	for len(this.Stack) > 0 {
		q := this.Stack[len(this.Stack)-1]
		this.Stack = this.Stack[:len(this.Stack)-1]

		if q < len(this.Str)-1 {
			this.Stack = append(this.Stack, q+1)
			for i := q + 2; i < len(this.Str) && len(this.Stack) < this.Size; i++ {
				this.Stack = append(this.Stack, i)
			}

			if len(this.Stack) == this.Size {
				break
			}
		}
	}

	for i := 0; i < len(this.Stack); i++ {
		result[i] = this.Str[this.Stack[i]]
	}

	return string(result)
}

func (this *CombinationIterator) HasNext() bool {
	if len(this.Stack) == 0 {
		return true
	}

	// last permutation is in reversed order, e.g. 1 2 3 4 => 4 3 2 1
	for i, j := 0, len(this.Str)-this.Size; i < len(this.Stack); i, j = i+1, j+1 {
		if this.Stack[i] != j {
			return true
		}
	}

	return false
}

//	Notes
//	1.	forget to check boundary condition, if stack is empty

//	2.	inspired from solution, don't know what's knuth algorithm L,
//		searched a great explanation from http://guptamukul.blogspot.com/2009/12/understanding-algorithm-l_05.html

//		L algorithm comes with 4 parts:
//		- L1 Visit a1, a2, ..., an
//		- L2 j = n-1, find first position that aj is able to increment
//		- L3 l = n, find first position al > aj (j's minimal increment),
//		  swap aj, al
//		- L4 reverse aj+1, aj+2, ..., an

//		L1 is the original state

//		L2 I understand it as follows: original sorted lexical order
//		a1 <= a2 <= a3 <= ... <= an, after operation becomes
//		a1 >= a2 >= a3 >= ... >= an

//		so for any position still hold ascending lexical order
//		aj <= aj+1, that's where to do operation

//		L3 finds next smallest aj that can exchange
//		during visiting process, there exists some ascending &
//		descending mixed in the sequence,
//		e.g. aj+1 >= ... >= al-1 >= aj > al >= al+1 >= ... >= an

//		L4 updates order when whole combinations are visited at index j,
//		to make sure it follows lexical order sorting

//		e.g. 1 2 3 4, at some point visited order 1 4 3 2
//		L2: j = 0 (0 indexed)
//		L3: l = 3 (0 indexed), switch aj & al => 2 4 3 1
//		L4: 2 1 3 4

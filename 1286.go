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
	Stack          []int
	Str            string
	Cur            []byte
	Remain, Target int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	remain := 1
	size := len(characters)

	for i := 0; i < combinationLength; i++ {
		remain *= size - i
	}

	for i := 1; i <= combinationLength; i++ {
		remain /= i
	}

	return CombinationIterator{
		Stack:  make([]int, 0),
		Str:    characters,
		Cur:    make([]byte, 0),
		Remain: remain,
		Target: combinationLength,
	}
}

func (this *CombinationIterator) Next() string {
	if len(this.Cur) == 0 {
		for i := 0; i < this.Target; i++ {
			this.Stack = append(this.Stack, i)
			this.Cur = append(this.Cur, this.Str[i])
		}
	} else {
		// remove previous chars not able to meet target length
		for len(this.Stack) > 0 {
			this.Cur = this.Cur[:len(this.Cur)-1]
			last := this.Stack[len(this.Stack)-1]
			this.Stack = this.Stack[:len(this.Stack)-1]

			if last < len(this.Str)-(this.Target-len(this.Cur)) {
				this.Cur = append(this.Cur, this.Str[last+1])
				this.Stack = append(this.Stack, last+1)
				break
			}
		}

		// appending chars, each index increase by 1 from previous
		for len(this.Cur) < this.Target {
			i := this.Stack[len(this.Stack)-1]
			this.Cur = append(this.Cur, this.Str[i+1])
			this.Stack = append(this.Stack, i+1)
		}
	}

	this.Remain--
	return string(this.Cur)
}

func (this *CombinationIterator) HasNext() bool {
	return this.Remain > 0
}

type CombinationIterator2 struct {
	Data []string
}

func Constructor(characters string, combinationLength int) CombinationIterator2 {
	ans := make([]string, 0)
	size := len(characters)

	for i := (1 << size) - 1; i >= 0; i-- {
		if oneBits(i) == combinationLength {
			tmp := make([]byte, 0)

			for j := size - 1; j >= 0; j-- {
				if i&(1<<j) != 0 {
					// becareful, left shift is reverse order of
					// original string
					tmp = append(tmp, characters[size-1-j])
				}
			}

			ans = append(ans, string(tmp))
		}
	}

	return CombinationIterator2{
		Data: ans,
	}
}

func oneBits(i int) int {
	var count int

	for i > 0 {
		count++
		i = i & (i - 1)
	}

	return count
}

func (this *CombinationIterator2) Next() string {
	str := this.Data[0]
	this.Data = this.Data[1:]

	return str
}

func (this *CombinationIterator2) HasNext() bool {
	return len(this.Data) > 0
}

type CombinationIterator1 struct {
	Str   string
	Stack []int
	Size  int
}

func Constructor(characters string, combinationLength int) CombinationIterator1 {
	return CombinationIterator1{
		Str:   characters,
		Size:  combinationLength,
		Stack: make([]int, 0),
	}
}

func (this *CombinationIterator1) Next() string {
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

func (this *CombinationIterator1) HasNext() bool {
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

//	3.	inspired from solution, don't forget the scenario to precompute all
//		combinations, which results in tc O(1) & sc O(2^n)

//	4.	inspired from solution, instead of using backtracking, can use bitmask
//		bitmask start from 2^n-1, if 1's count == target, push it, very smart

//		count 1 in an integer can use a & (a-1) to do so

//		convert bitmask precompute into bitmark itegration is easier, the key is
//		to have bitmarks always the next valid number, so that hasNext can check

package main

// Given a nested list of integers, implement an iterator to flatten it.
//
// Each element is either an integer, or a list -- whose elements may also be integers or other lists.
//
// Example 1:
//
// Input: [[1,1],2,[1,1]]
// Output: [1,1,2,1,1]
// Explanation: By calling next repeatedly until hasNext returns false,
// the order of elements returned by next should be: [1,1,2,1,1].
//
// Example 2:
//
// Input: [1,[4,[6]]]
// Output: [1,4,6]
// Explanation: By calling next repeatedly until hasNext returns false,
// the order of elements returned by next should be: [1,4,6].

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	List  [][]*NestedInteger
	Index []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		List:  [][]*NestedInteger{nestedList},
		Index: []int{0},
	}
}

func (this *NestedIterator) Next() int {
	this.makeTopElementNumber()

	num := this.list()[this.idx()].GetInteger()
	this.Index[len(this.Index)-1]++

	return num
}

func (this *NestedIterator) list() []*NestedInteger {
	return this.List[len(this.List)-1]
}

func (this *NestedIterator) idx() int {
	return this.Index[len(this.Index)-1]
}

func (this *NestedIterator) makeTopElementNumber() {
	for len(this.List) > 0 {
		// remove list already used
		if this.idx() == len(this.list()) {
			this.Index = this.Index[:len(this.Index)-1]
			this.List = this.List[:len(this.List)-1]
			continue
		}

		top := this.list()[this.idx()]

		// terminate when it's integer
		if top.IsInteger() {
			return
		}

		// pushing list
		this.List = append(this.List, top.GetList())
		this.Index[len(this.Index)-1]++
		this.Index = append(this.Index, 0)
	}
}

func (this *NestedIterator) HasNext() bool {
	this.makeTopElementNumber()

	return len(this.List) > 0
}

type NestedIterator2 struct {
	Stack []*NestedInteger
}

func Constructor2(nestedList []*NestedInteger) *NestedIterator2 {
	stack := make([]*NestedInteger, 0)
	for i := len(nestedList) - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}

	return &NestedIterator2{
		Stack: stack,
	}
}

func (this *NestedIterator2) Next() int {
	this.makeStackTopAnInteger()

	s := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]

	return s.GetInteger()
}

func (this *NestedIterator2) HasNext() bool {
	this.makeStackTopAnInteger()

	return len(this.Stack) > 0
}

func (this *NestedIterator2) makeStackTopAnInteger() {
	for len(this.Stack) > 0 {
		s := this.Stack[len(this.Stack)-1]

		// do no pop stack, check if the element is integer or not
		if s.IsInteger() {
			return
		}

		// pop element, since it's a list, remove it and do further flatten
		this.Stack = this.Stack[:len(this.Stack)-1]

		list := s.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			this.Stack = append(this.Stack, list[i])
		}
	}
}

type NestedIterator1 struct {
	List  []*NestedInteger
	Idx   int
	Stack []*NestedIterator1
}

func Constructor1(nestedList []*NestedInteger) *NestedIterator1 {
	return &NestedIterator1{
		List: nestedList,
		Idx:  0,
	}
}

func (this *NestedIterator1) Next() int {
	s := this.Cleanup()

	if s != nil {
		return s.Next()
	}

	if this.List[this.Idx].IsInteger() {
		num := this.List[this.Idx].GetInteger()
		this.Idx++
		return num
	}

	this.Stack = append(this.Stack, &NestedIterator1{
		List: this.List[this.Idx].GetList(),
	})
	this.Idx++
	return this.Stack[0].Next()
}

func (this *NestedIterator1) Cleanup() *NestedIterator1 {
	var s *NestedIterator1

	for len(this.Stack) > 0 {
		s = this.Stack[len(this.Stack)-1]

		if s.Idx == len(s.List) && len(s.Stack) == 0 {
			this.Stack = this.Stack[:len(this.Stack)-1]
		} else {
			return s
		}
	}
	return nil
}

func (this *NestedIterator1) HasNext() bool {
	this.Cleanup()

	return !(this.Idx == len(this.List) && len(this.Stack) == 0)
}

//	Notes
//	1.	didn't think of a good solution

//	2.	inspired of solution, use a stack with reverse order of list, make
//		recursive call into linear

//		the beautiful is that every item will be called once and dropped, reduce
//		some un-necessary checks

//		recursion means something not done til some point, so use stack to break
//		something to do into list with reverse order, such that each call really
//		in a 'done' state

//		very clever

//	3.	becareful boundary condition, it's a pointer, could be nil

//	4.	becareful about boundary condition, [[]], i have to check it's nil pointer
//		or not, or by solution, the other way is to have another function to flatten
//		stack until last one is number

//		the thing to mention is that, loop terminates when top element of stack is a number,
//		so that's how a recursion converted to iteration. remember the technique, focus
//		on when to terminate loop

//	5.	inspired from solution, using two stacks to store state, there exist
//		cleanup function: makeTopElementInteger

//	6.	solution is very good, but I dind't take time to implement other methods

//	7.	inspired from https://leetcode.com/problems/flatten-nested-list-iterator/discuss/80146/Real-iterator-in-Python-Java-C%2B%2B

//		author said something about iterator: iterator shouldn't copy the
//		entire data, I think that's the reason to use iterative way to solve
//		this problem

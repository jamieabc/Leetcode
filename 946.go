package main

//Given two sequences pushed and popped with distinct values, return true if and only if this could have been the result of a sequence of push and pop operations on an initially empty stack.
//
//
//
//Example 1:
//
//Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
//Output: true
//Explanation: We might do the following sequence:
//push(1), push(2), push(3), push(4), pop() -> 4,
//push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//Example 2:
//
//Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
//Output: false
//Explanation: 1 cannot be popped before 2.
//
//
//Note:
//
//0 <= pushed.length == popped.length <= 1000
//0 <= pushed[i], popped[i] < 1000
//pushed is a permutation of popped.
//pushed and popped have distinct values.

// since every item should appear once
// the easier way is to follow operation of push/pop
// a pointer to popped item list, if this popped item is same as top of stack
// pop it and go to next popped item, otherwise push next pushed item
// operation is invalid when find stack is not empty

func validateStackSequences(pushed []int, popped []int) bool {
	i, j := 0, 0
	for _, n := range pushed {
		pushed[i] = n
		i++
		for i >= 0 && j < len(popped) && pushed[i-1] == popped[j] {
			i--
			j++
		}
	}

	return i == 0
}

func validateStackSequences3(pushed []int, popped []int) bool {
	stack := make([]int, 0)

	for i, j := 0, 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			j++
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func validateStackSequences2(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}

	if len(popped) <= 1 {
		return true
	}

	s := make([]int, len(pushed))
	loc := 0
	j := 0
	for i := 0; i < len(pushed); i++ {
		s[loc] = pushed[i]
		loc++
		for loc-1 >= 0 && s[loc-1] == popped[j] {
			loc--
			j++
		}
	}

	return loc == 0
}

// use a map to store item <-> index
// use another array to store if item pushed status
// a pointer to denote current location, if pointer moves left, mark path 1 as popped
// move pointer right, mark only the target item as popped
// if any pointer target location on left is marked as popped, then it's illegal
func validateStackSequences1(pushed []int, popped []int) bool {
	if len(popped) <= 1 {
		return true
	}

	m := make(map[int]int)
	for idx, item := range pushed {
		m[item] = idx
	}

	status := make([]bool, len(pushed))
	status[m[popped[0]]] = true

	for j := 1; j < len(popped); j++ {
		currentIndex := m[popped[j]]
		previousIndex := m[popped[j-1]]
		if status[currentIndex] == true {
			return false
		}

		if previousIndex > currentIndex {
			for tmp := currentIndex; tmp <= previousIndex; tmp++ {
				status[tmp] = true
			}
		} else {
			status[currentIndex] = true
		}
	}
	return true
}

//	problems
//	1.	inspired from https://leetcode.com/problems/validate-stack-sequences/discuss/197685/C%2B%2BJavaPython-Simulation-O(1)-Space

//		an elegant code

//		for O(1) space, uses pushed as stack, the point is to rewrite
//		stack so that no additional space is needed

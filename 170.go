package main

//Design and implement a TwoSum class. It should support the following operations: add and find.
//
//add - Add the number to an internal data structure.
//find - Find if there exists any pair of numbers which sum is equal to the value.
//
//Example 1:
//
//add(1); add(3); add(5);
//find(4) -> true
//find(7) -> false
//Example 2:
//
//add(3); add(1); add(2);
//find(3) -> true
//find(6) -> false

type TwoSum struct {
	mapping map[int]int
	keys    []int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		mapping: make(map[int]int),
		keys:    []int{},
	}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	if _, ok := this.mapping[number]; !ok {
		this.mapping[number] = 1
		this.keys = append(this.keys, number)
	} else {
		this.mapping[number]++
	}
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for _, key := range this.keys {
		remain := value - key
		if c, ok := this.mapping[remain]; ok {
			if remain != key {
				return true
			}

			// in case it's duplicates, e.g. add(4), Find(8)
			if remain == key && c >= 2 {
				return true
			}
		}
	}

	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */

//	problems
//	1.	wrong logic when it's same number, any add will always exceeds 2

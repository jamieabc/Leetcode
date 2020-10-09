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
	Nums map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		Nums: make(map[int]int),
	}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.Nums[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for num, count := range this.Nums {
		rest := value - num
		if (num == rest && count > 1) || (num != rest && this.Nums[rest] > 0) {
			return true
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

//	Notes
//	1.	wrong logic when it's same number, any add will always exceeds 2

//	2.	at least one operation tc needs to be O(n)

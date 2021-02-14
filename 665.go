package main

//Given an array with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.
//
//We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).
//
//Example 1:
//
//Input: [4,2,3]
//Output: True
//Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
//Example 2:
//
//Input: [4,2,1]
//Output: False
//Explanation: You can't get a non-decreasing array by modify at most one element.
//Note: The n belongs to [1, 10,000].

func checkPossibility(nums []int) bool {
	idx := -1
	size := len(nums)

	for i := 1; i < size; i++ {
		if nums[i] < nums[i-1] {
			if idx != -1 {
				return false
			}
			idx = i
		}
	}

	return idx == -1 || idx == 1 || idx == size-1 || nums[idx-2] <= nums[idx] || nums[idx-1] <= nums[idx+1]
}

func checkPossibility2(nums []int) bool {
	size := len(nums)
	var change, idx, prev int

	for i := 1; i < size; i++ {
		if nums[i] < nums[i-1] {
			if change == 0 {
				idx = i - 1
				prev = nums[i]
				nums[i] = nums[i-1]
				change++
			} else {
				change++
				break
			}
		}
	}

	if change <= 1 {
		return true
	}

	nums[idx+1] = prev
	nums[idx] = prev

	for i := 1; i < size; i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}

	return true
}

// [3, 4, 2, 3] false
// [2, 3, 3, 2, 4] true
func checkPossibility1(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	}

	var changes bool

	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			if changes {
				return false
			}

			if i > 0 {
				if nums[i-1] > nums[i+1] {
					nums[i+1] = nums[i]
				} else {
					nums[i] = nums[i+1]
				}
			} else {
				nums[i] = nums[i+1]
			}
			changes = true
		}
	}
	return true
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/non-decreasing-array/discuss/106826/JavaC%2B%2B-Simple-greedy-like-solution-with-explanation

//		the thinking is pretty good, the order relates to 3 numbers:
//		i-1, i, i+1

//		the best strategy is to modify num[i] cause the higher nums[i+1]
//		makes more risk to fail. but when nums[i-1] > nums[i+1] then it's
//		a must to change nums[i+1]

//	2.	add reference https://leetcode.com/problems/non-decreasing-array/discuss/106842/The-easiest-python-solution....

//		the other way is to check if there exist another descending for next number

//	3.	inspired from solution, there's a pretty good insight of the problem
//		non-decreasing sequence means chain of numbers, so as 3 numbers satisfies
//		this relationship, then first or last (depend on direction) no need to consider

//		if a <= b <= c, then a no need to consider
//		if last-2 <= last-1 <= last, then last no need to consider

//		by observation, original problem can be shrink into find range that doesn't
//		satisfy this rule

//		if sequence after prune length <= 2, then it's good to go; if length >= 6,
//		there's 2 disjoint set of length 3 sequence, which is impossible to make it
//		non-decreasing in just one move

//		length within 3 ~ 5 needs some additional checking

//	4.	inspired form solution, there's a much clever solution, it considers which property
//		makes one modification work

//		if there are 4 numbers a, b, c, d, to make this sequence non-decreasing,
//		either change b such that b >= a && b <= c or change c such that c >= b && c <= d
//		this means b & c value are determined by their neighbors

//		find position where a[i] < a[i-1], if there's more than 1 positions, then it's
//		impossible to make sequence non-decreasing in one modification

//		if i == 1 or size-1, change a[0] or a[size-1] because those positions doesn't need
//		consider previous or next number

//		if a[i-2], a[i-1], a[i], a[i+1] all exist, the only way one modification
//		works is that a[i-2] <= a[i] or a[i-1] <= a[i+1]

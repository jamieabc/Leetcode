package main

import "sort"

// Alice has a hand of cards, given as an array of integers.
//
// Now she wants to rearrange the cards into groups so that each group is size W, and consists of W consecutive cards.
//
// Return true if and only if she can.
//
//
//
// Example 1:
//
// Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
// Output: true
// Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8].
//
// Example 2:
//
// Input: hand = [1,2,3,4,5], W = 4
// Output: false
// Explanation: Alice's hand can't be rearranged into groups of 4.
//
//
//
// Constraints:
//
//     1 <= hand.length <= 10000
//     0 <= hand[i] <= 10^9
//     1 <= W <= hand.length
//
// Note: This question is the same as 1296: https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/

func isNStraightHand(hand []int, W int) bool {
	sort.Ints(hand)

	queue := make([][2]int, 0)
	var count, end int

	length := len(hand)
	for i := 0; i < length; {
		// find duplicate numbers
		for end = i + 1; end < length; end++ {
			if hand[end] != hand[end-1] {
				break
			}
		}
		tmp := end - i

		// check duplicate number count correct
		if tmp < count {
			return false
		} else if tmp > count {
			queue = append(queue, [2]int{hand[i] + W - 1, tmp - count})
			count = tmp
		}

		// update count
		if queue[0][0] == hand[i] {
			count -= queue[0][1]
			queue = queue[1:]
		}

		// go to next distinct number
		i = end
	}

	return len(queue) == 0
}

func isNStraightHand3(hand []int, W int) bool {
	if len(hand)%W != 0 {
		return false
	}

	counter := make(map[int]int)
	for _, i := range hand {
		counter[i]++
	}

	keys := make([]int, 0)
	for k := range counter {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	queue := make([]int, 0)
	var prev int

	// w/o updating counter, check for correctness of duplicate number count
	// queue means something is in group, so next number count >=
	// queue length, queue records when to reduce group count
	// next number count larger than queue length, means some new number
	// appears and group count is increased
	for i := range keys {
		if len(queue) > 0 && i == queue[0] {
			for len(queue) > 0 {
				if i == queue[0] {
					queue = queue[1:]
				} else {
					break
				}
			}
		}

		if len(queue) == 0 {
			for j := 0; j < counter[keys[i]]; j++ {
				queue = append(queue, i+W)
			}
			prev = keys[i]
		} else {
			// consecutive numbers
			if prev != keys[i]-1 {
				return false
			}

			if counter[keys[i]] > len(queue) {
				q := len(queue)
				for j := 0; j < counter[keys[i]]-q; j++ {
					queue = append(queue, i+W)
				}
			} else if counter[keys[i]] < len(queue) {
				return false
			}
			prev = keys[i]
		}
	}

	for i := range queue {
		if queue[i] != len(keys) {
			return false
		}
	}

	return true
}

func isNStraightHand2(hand []int, W int) bool {
	if len(hand)%W != 0 {
		return false
	}

	counter := make(map[int]int)
	for _, i := range hand {
		counter[i]++
	}

	// unique sorted numbers
	keys := make([]int, 0)
	for k := range counter {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var j int
	for i := 0; i < len(keys); {
		// if a min number is found, num, num+1, .... num+W-1 is absolutely
		// into same group, remove those number count by 1
		for j = keys[i]; j < keys[i]+W; j++ {
			if counter[j] == 0 {
				return false
			} else {
				counter[j]--
			}
		}

		// find next smallest number
		for j = i; j < len(keys); j++ {
			if counter[keys[j]] > 0 {
				break
			}
		}
		i = j
	}

	return true
}

func isNStraightHand1(hand []int, W int) bool {
	length := len(hand)
	if length == 1 {
		return true
	}

	if length%W != 0 {
		return false
	}

	sort.Ints(hand)
	count, max := make([]int, 0), make([]int, 0)
	var idx int

	// try to put every incoming number into specific group
	// if cannot find any, crate new group, otherwise, put that number in
	// and go on to next number
	for _, i := range hand {
		var found bool
		for j := idx; j < len(count); j++ {
			if max[j] == i-1 && count[j] < W {
				found = true
				max[j] = i
				count[j]++

				if count[j] == W {
					idx = j + 1
				}
				break
			}
		}

		if !found {
			count = append(count, 1)
			max = append(max, i)

			if count[len(count)-1] == W {
				idx++
			}
		}
	}

	return idx == len(count)
}

//	problems
//	1.	wrong understanding of the problem, it's to group consecutive
//		numbers in n length, there's no limit on group count

//	2.	too slow, use 2 arrays to store each group count & max, since count
//		is used to track if group is full, it needs to be checked in every
//		statement

//	3.	inspired from https://leetcode.com/problems/hand-of-straights/discuss/135655/Python-O(nlgn)-simple-solution-with-intuition

//		use a map to store each number occurrence, then loop through
//		counter keys, when every key count > 0, deduct count following
//		W-1 numbers

//		tc: O(n log n + nW), n: # unique cards

//	4.	inspired from sample code, no counter is need, can iterate through
//		inputs

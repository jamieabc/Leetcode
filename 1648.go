package main

import "sort"

// You have an inventory of different colored balls, and there is a customer that wants orders balls of any color.
//
// The customer weirdly values the colored balls. Each colored ball's value is the number of balls of that color you currently have in your inventory. For example, if you own 6 yellow balls, the customer would pay 6 for the first yellow ball. After the transaction, there are only 5 yellow balls left, so the next yellow ball is then valued at 5 (i.e., the value of the balls decreases as you sell more to the customer).
//
// You are given an integer array, inventory, where inventory[i] represents the number of balls of the ith color that you initially own. You are also given an integer orders, which represents the total number of balls that the customer wants. You can sell the balls in any order.
//
// Return the maximum total value that you can attain after selling orders colored balls. As the answer may be too large, return it modulo 109 + 7.
//
//
//
// Example 1:
//
// Input: inventory = [2,5], orders = 4
// Output: 14
// Explanation: Sell the 1st color 1 time (2) and the 2nd color 3 times (5 + 4 + 3).
// The maximum total value is 2 + 5 + 4 + 3 = 14.
//
// Example 2:
//
// Input: inventory = [3,5], orders = 6
// Output: 19
// Explanation: Sell the 1st color 2 times (3 + 2) and the 2nd color 4 times (5 + 4 + 3 + 2).
// The maximum total value is 3 + 2 + 5 + 4 + 3 + 2 = 19.
//
// Example 3:
//
// Input: inventory = [2,8,4,10,6], orders = 20
// Output: 110
//
// Example 4:
//
// Input: inventory = [1000000000], orders = 1000000000
// Output: 21
// Explanation: Sell the 1st color 1000000000 times for a total value of 500000000500000000. 500000000500000000 modulo 109 + 7 = 21.
//
//
//
// Constraints:
//
//     1 <= inventory.length <= 105
//     1 <= inventory[i] <= 109
//     1 <= orders <= min(sum(inventory[i]), 109)

// tc: O(n log(m)), n: # inventory, m: maximum number in inventory count
func maxProfit(inventory []int, orders int) int {
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i] > inventory[j]
	})

	bound := binarySearch(inventory, orders) // not including bound itself

	var profit int64
	mod := int64(1e9 + 7)

	// add profit to lower bound
	for idx := 0; idx < len(inventory) && inventory[idx] > bound; idx++ {
		size := inventory[idx] - bound
		profit += int64(inventory[idx]+bound+1) * int64(size) / int64(2) % mod
		profit = profit % mod
	}

	count := sumLarger(inventory, bound)
	if count < orders {
		profit += int64(bound) * int64(orders-count) % mod
	}

	return int(profit % mod)
}

func sumLarger(inv []int, criteria int) int {
	var sum int

	for _, i := range inv {
		sum += max(0, i-criteria)
	}

	return sum
}

func binarySearch(inv []int, orders int) int {
	var largest int
	for _, i := range inv {
		largest = max(largest, i)
	}

	var ans int

	for low, high := 0, largest; low <= high; {
		mid := low + (high-low)>>1
		count := sumLarger(inv, mid)

		if count == orders {
			return mid
		} else if count > orders {
			low = mid + 1
		} else {
			ans = mid
			high = mid - 1
		}
	}

	return ans
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int          { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxProfit1(inventory []int, orders int) int {
	// sell balls with maximum length
	h := &MaxHeap{}
	heap.Init(h)

	var profit int64
	mod := int64(1e9 + 7)
	heap.Push(h, 0) // always exist one number 0

	for _, i := range inventory {
		heap.Push(h, i)
	}

	for i := 0; i < orders; i++ {
		tmp := heap.Pop(h).(int)
		diff := tmp - h.Peek()

		for j := 0; j <= diff && tmp > 0 && i < orders; i, j = i+1, j+1 {
			profit += int64(tmp)
			if profit > mod {
				profit -= mod
			}
			tmp--
		}
		i--

		heap.Push(h, tmp)
	}

	return int(profit)
}

//	Notes
//	1.	max heap can solve the problem, but not most efficient way

//	2.	inspired from https://leetcode.com/problems/sell-diminishing-valued-colored-balls/discuss/927560/C%2B%2B-Binary-Answer

//		inspired from alex video https://www.youtube.com/watch?v=ONeUL17hfUU

//		sold sequence will be non-decreasing sequence:

//		e.g. inv: [2, 4, 6, 8, 10], orders = 20

//		if all inventories are sold, sold sequence will be
//		1, 1, 1, 1, 1,
//		2, 2, 2, 2, 2,
//		3, 3, 3, 3
//		4, 4, 4, 4
//		5, 5, 5
//		6, 6, 6
//		7, 7
//		8, 8
//		9,
//		10

//		it's not necessary to list all sequence numbers to find which number is
//		the lower bound meets criteria, instead, binary search can be used to
//		find, the key point is guess a number and count what if this number is
//		lower bound (the smaller cutoff number is, the larger inventory sold)

//		e.g. guess sold inventory prices above 3
//		inventory 10 sold 10 - 3 = 7 (10, 9, 8, 7, 6, 5, 4)
//		inventory 8 sold 8 - 3 = 5 (8, 7, 6, 5, 4)
//		inventory 6 sold 6 - 3 = 3 (6, 5, 4)
//		inventory 4 sold 1
//		inventory 2 sold 0

//		total sold inventories = 7 + 5 + 3 + 1 = 16

//		this is the hardest part to understand while trying to solve the problem

//	3. mod operator should be applied in the last

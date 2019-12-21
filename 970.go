package main

//Given two positive integers x and y, an integer is powerful if it is equal to x^i + y^j for some integers i >= 0 and j >= 0.
//
//Return a list of all powerful integers that have value less than or equal to bound.
//
//You may return the answer in any order.  In your answer, each value should occur at most once.
//
//
//
//Example 1:
//
//Input: x = 2, y = 3, bound = 10
//Output: [2,3,4,5,7,9,10]
//Explanation:
//2 = 2^0 + 3^0
//3 = 2^1 + 3^0
//4 = 2^0 + 3^1
//5 = 2^1 + 3^1
//7 = 2^2 + 3^1
//9 = 2^3 + 3^0
//10 = 2^0 + 3^2
//
//Example 2:
//
//Input: x = 3, y = 5, bound = 15
//Output: [2,4,6,8,10,14]
//
//
//
//Note:
//
//    1 <= x <= 100
//    1 <= y <= 100
//    0 <= bound <= 10^6

func powerfulIntegers(x int, y int, bound int) []int {
	mapping := make(map[int]bool)

	main := x
	sub := y
	if y > x {
		main = y
		sub = x
	}

	base := 1
	for i := 0; ; i++ {
		if base >= bound {
			break
		}

		added := 1
		for j := 0; ; j++ {
			sum := base + added
			if sum > bound {
				break
			}

			if _, ok := mapping[sum]; !ok {
				mapping[sum] = true
			}

			if sub == 1 {
				break
			}
			added *= sub
		}

		if main == 1 {
			break
		}
		base *= main
	}

	keys := make([]int, 0)

	for k, _ := range mapping {
		keys = append(keys, k)
	}

	return keys
}

// problems
// 1. I assume every power of number will increasing, but there's an exception: 1, always be 1
// 2. forget to consider which number to be basis, e.g. 1, 2 should use 2 as basis

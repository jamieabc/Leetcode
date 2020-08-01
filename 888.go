package main

//Alice and Bob have candy bars of different sizes: A[i] is the size of the i-th bar of candy that Alice has, and B[j] is the size of the j-th bar of candy that Bob has.
//
//Since they are friends, they would like to exchange one candy bar each so that after the exchange, they both have the same total amount of candy.  (The total amount of candy a person has is the sum of the sizes of candy bars they have.)
//
//Return an integer array ans where ans[0] is the size of the candy bar that Alice must exchange, and ans[1] is the size of the candy bar that Bob must exchange.
//
//If there are multiple answers, you may return any one of them.  It is guaranteed an answer exists.
//
//
//
//Example 1:
//
//Input: A = [1,1], B = [2,2]
//Output: [1,2]
//
//Example 2:
//
//Input: A = [1,2], B = [2,3]
//Output: [1,2]
//
//Example 3:
//
//Input: A = [2], B = [1,3]
//Output: [2,3]
//
//Example 4:
//
//Input: A = [1,2,5], B = [2,4]
//Output: [5,4]
//
//
//
//Note:
//
//    1 <= A.length <= 10000
//    1 <= B.length <= 10000
//    1 <= A[i] <= 100000
//    1 <= B[i] <= 100000
//    It is guaranteed that Alice and Bob have different total amounts of candy.
//    It is guaranteed there exists an answer.

func fairCandySwap(A []int, B []int) []int {
	counter := make(map[int]bool)
	var diff int
	for _, n := range A {
		counter[n] = true
		diff += n
	}

	for _, n := range B {
		diff -= n
	}

	// guarantee to have answer
	diff /= 2
	for _, n := range B {
		if counter[diff+n] {
			return []int{diff + n, n}
		}
	}

	return []int{}
}

func fairCandySwap2(A []int, B []int) []int {
	counter := make(map[int]bool)
	var sumA, sumB int
	for _, n := range A {
		counter[n] = true
		sumA += n
	}

	for _, n := range B {
		sumB += n
	}

	// guarantee to have an answer
	diff := (sumA + sumB) / 2

	for _, n := range B {
		target := diff - (sumB - n)
		if counter[target] {
			return []int{target, n}
		}
	}

	return []int{0, 0}
}

func fairCandySwap1(A []int, B []int) []int {
	counter := make([]bool, 100001)
	var sumA, sumB int
	for _, n := range A {
		counter[n] = true
		sumA += n
	}

	for _, n := range B {
		sumB += n
	}

	// guarantee to have an answer
	diff := (sumA + sumB) / 2

	for _, n := range B {
		target := diff - (sumB - n)
		if target > 0 && target <= 100001 && counter[target] {
			return []int{target, n}
		}
	}

	return []int{0, 0}
}

//	problems
//	1.	since it's array accessing, make sure index in range

//	2.	try to use hashmap, since if input array size not very large,
//		it's not efficient to allocate such big array

//	3.	inspired from sample code, absolution sum value is not important,
//		most important is difference

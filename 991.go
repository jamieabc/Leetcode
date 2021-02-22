package main

// On a broken calculator that has a number showing on its display, we can perform two operations:
//
//     Double: Multiply the number on the display by 2, or;
//     Decrement: Subtract 1 from the number on the display.
//
// Initially, the calculator is displaying the number X.
//
// Return the minimum number of operations needed to display the number Y.
//
//
//
// Example 1:
//
// Input: X = 2, Y = 3
// Output: 2
// Explanation: Use double operation and then decrement operation {2 -> 4 -> 3}.
//
// Example 2:
//
// Input: X = 5, Y = 8
// Output: 2
// Explanation: Use decrement and then double {5 -> 4 -> 8}.
//
// Example 3:
//
// Input: X = 3, Y = 10
// Output: 3
// Explanation:  Use double, decrement and double {3 -> 6 -> 5 -> 10}.
//
// Example 4:
//
// Input: X = 1024, Y = 1
// Output: 1023
// Explanation: Use decrement operations 1023 times.
//
//
//
// Note:
//
//     1 <= X <= 10^9
//     1 <= Y <= 10^9

// half Y until Y < X, so tc would be how many times to reach difference
// Y / X
// tc: O(log(Y/X))
func brokenCalc(X int, Y int) int {
	var steps int

	// (Y+2)/2 takes 3 steps
	// Y/2 + 1 takes 2 steps
	// it turns out that divide is always best solution, if possible
	for Y > X {
		steps++

		if Y&1 > 0 {
			Y++
		} else {
			Y = Y >> 1
		}
	}

	return steps + X - Y
}

func brokenCalc1(X int, Y int) int {
	queue := []int{X}
	var steps int

	for true {
		size := len(queue)
		steps++

		for i := 0; i < size; i++ {
			n := queue[0]
			queue = queue[1:]

			if n == Y {
				return steps - 1
			}

			queue = append(queue, n-1, n<<1)
		}
	}

	return 0
}

//	Notes
//	1.	start from BFS, but turns out to be timeout

//	2.	couldn't think of ways to do it

//		inspired form solution, if start from Y then it's more easily

//		basic proof like (Y+2)/2 takes 3 steps
//						 Y/2+1   takes 2 steps
//		above two ways with same result, but divide first with less steps
//		so it can be treat greedily

//	3.	inspired form https://leetcode.com/problems/broken-calculator/discuss/234484/JavaC%2B%2BPython-Change-Y-to-X-in-1-Line

//		pretty good explanation and tc: O(log(Y/X))

//	4.	inspired from https://leetcode.com/problems/broken-calculator/discuss/236565/Detailed-Proof-Of-Correctness-Greedy-Algorithm

//		another good proof

//	5.	it's kind of similar to minimum deviation, while half with multiple
//		times (solution space is pretty big), and double only once (smaller
//		solution space)

//		this problem with similar technique

package main

//Write a program to check whether a given number is an ugly number.
//
//Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
//
//Example 1:
//
//Input: 6
//Output: true
//Explanation: 6 = 2 × 3
//Example 2:
//
//Input: 8
//Output: true
//Explanation: 8 = 2 × 2 × 2
//Example 3:
//
//Input: 14
//Output: false
//Explanation: 14 is not ugly since it includes another prime factor 7.
//Note:
//
//1 is typically treated as an ugly number.
//Input is within the 32-bit signed integer range: [−231,  231 − 1].

func isUgly(num int) bool {
	// check positive
	// 1, 2, 3, 5 is true

	if num <= 0 {
		return false
	}

	if num <= 5 {
		return true
	}

	factors := []int{2, 3, 5}

	for _, n := range factors {
		for num != 1 {
			if num%n == 0 {
				num /= n
			} else {
				break
			}
		}
	}

	return num == 1
}

// problems
// 1. this is not finding prime number, it should check if all factors only be composed of 2, 3, 5
// 2. wrong end condition, divisible means at last, number is 1
// 3. too slow, use prime factor to get all prime, check if any prime is a factor
// 4. wrong when iterate prime, n is value, but n needs to be divided by index(prime)
// 5. using range will start from 0, and then num will be divide by 0....
// 6. oom

package main

//Count the number of prime numbers less than a non-negative number, n.
//
//Example:
//
//Input: 10
//Output: 4
//Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

func countPrimes(n int) int {
	// for 0 and 1, don't count
	if n == 0 || n == 1 {
		return 0
	}

	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			continue
		}

		for j := i * 2; j < n; j += i {
			isPrime[j] = false
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}

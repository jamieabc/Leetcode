package main

//We have n buildings numbered from 0 to n - 1. Each building has a number of employees. It's transfer season, and some employees want to change the building they reside in.
//
//You are given an array requests where requests[i] = [fromi, toi] represents an employee's request to transfer from building fromi to building toi.
//
//All buildings are full, so a list of requests is achievable only if for each building, the net change in employee transfers is zero. This means the number of employees leaving is equal to the number of employees moving in. For example if n = 3 and two employees are leaving building 0, one is leaving building 1, and one is leaving building 2, there should be two employees moving to building 0, one employee moving to building 1, and one employee moving to building 2.
//
//Return the maximum number of achievable requests.
//
//
//
//Example 1:
//
//Input: n = 5, requests = [[0,1],[1,0],[0,1],[1,2],[2,0],[3,4]]
//Output: 5
//Explantion: Let's see the requests:
//From building 0 we have employees x and y and both want to move to building 1.
//From building 1 we have employees a and b and they want to move to buildings 2 and 0 respectively.
//From building 2 we have employee z and they want to move to building 0.
//From building 3 we have employee c and they want to move to building 4.
//From building 4 we don't have any requests.
//We can achieve the requests of users x and b by swapping their places.
//We can achieve the requests of users y, a and z by swapping the places in the 3 buildings.
//
//Example 2:
//
//Input: n = 3, requests = [[0,0],[1,2],[2,1]]
//Output: 3
//Explantion: Let's see the requests:
//From building 0 we have employee x and they want to stay in the same building 0.
//From building 1 we have employee y and they want to move to building 2.
//From building 2 we have employee z and they want to move to building 1.
//We can achieve all the requests.
//
//Example 3:
//
//Input: n = 4, requests = [[0,3],[3,1],[1,2],[2,0]]
//Output: 4
//
//
//
//Constraints:
//
//    1 <= n <= 20
//    1 <= requests.length <= 16
//    requests[i].length == 2
//    0 <= fromi, toi < n

// tc: O(2^r * (n+r))
func maximumRequests(n int, requests [][]int) int {
	var maxRequest int
	size := len(requests)
	balance := make([]int, n)

	// try all possibilities
	for i := 0; i < 1<<size; i++ {
		if check(requests, balance, i) {
			maxRequest = max(maxRequest, oneCount(i))
		}
	}

	return maxRequest
}

func check(requests [][]int, balance []int, counter int) bool {
	size := len(requests)

	for i := 0; i < size; i++ {
		if counter&(1<<i) > 0 {
			balance[requests[i][0]]--
			balance[requests[i][1]]++
		}
	}

	// check all zero and cleanup
	allZero := true
	for i := range balance {
		if balance[i] != 0 {
			allZero = false
			balance[i] = 0
		}
	}

	return allZero
}

func oneCount(i int) int {
	var count int

	for i > 0 {
		count++
		i = i & (i - 1)
	}

	return count
}

func maximumRequests1(n int, requests [][]int) int {
	flags := make([]bool, len(requests))
	return recursive(n, 0, flags, requests)
}

func recursive(n, start int, flags []bool, requests [][]int) int {
	// check if current combinations are valid
	var count int
	if start != 0 {
		counter := make(map[int]int)
		for i := range flags {
			if flags[i] {
				counter[requests[i][0]]--
				counter[requests[i][1]]++
				count++
			}
		}

		// exist non-balanced building, reset count to 0
		for _, val := range counter {
			if val != 0 {
				count = 0
				break
			}
		}
	}

	for i := start; i < len(requests); i++ {
		flags[i] = true
		count = max(count, recursive(n, i+1, flags, requests))
		flags[i] = false
	}

	return count
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	didn't solve during contest

//	2.	after watch alex video https://www.youtube.com/watch?v=5QGk9NGmiNI

//		one clue is that problem range is quite small, which implies a brute
//		force solution

//		but I still cannot figure out if it's a dp or pure brute force

//		tc: O((n+r) * 2^r), all combinations are 2^r, every combinations needs to
//		update r request then check n buildings balance

//	3.	inspired from https://leetcode.com/problems/maximum-number-of-achievable-transfer-requests/discuss/866456/Python-Check-All-Combinations

//		this is np problem, try all combinations takes 2^r, each combination
//		take n+r to check

//	4.	inspired form https://leetcode.com/problems/maximum-number-of-achievable-transfer-requests/discuss/868403/C%2B%2BPython-knapsack-01

//		voturbac also mention to identity this is search problem

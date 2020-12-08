package main

//In a list of songs, the i-th song has a duration of time[i] seconds.
//
//Return the number of pairs of songs for which their total duration in seconds is divisible by 60.  Formally, we want the number of indices i < j with (time[i] + time[j]) % 60 == 0.
//
//
//
//Example 1:
//
//Input: [30,20,150,100,40]
//Output: 3
//Explanation: Three pairs have a total duration divisible by 60:
//(time[0] = 30, time[2] = 150): total duration 180
//(time[1] = 20, time[3] = 100): total duration 120
//(time[1] = 20, time[4] = 40): total duration 60
//
//Example 2:
//
//Input: [60,60,60]
//Output: 3
//Explanation: All three pairs have a total duration of 120, which is divisible by 60.
//
//
//
//Note:
//
//    1 <= time.length <= 60000
//    1 <= time[i] <= 500

func numPairsDivisibleBy60(time []int) int {
	if len(time) <= 1 {
		return 0
	}

	arr := make([]int, 60)

	var remain, modular int
	count := 0
	for _, num := range time {
		modular = num % 60
		remain = 60 - modular

		if remain == 60 {
			remain = 0
		}

		count += arr[remain]
		arr[modular]++
	}

	return count
}

func numPairsDivisibleBy60_1(time []int) int {
	validNums := make([]int, 0)
	for i := 60; i < 1000; i += 60 {
		validNums = append(validNums, i)
	}

	counter := make(map[int]int)
	var ans int

	for _, t := range time {
		for _, j := range validNums {
			if j > t+500 {
				break
			}

			if count, ok := counter[j-t]; ok {
				ans += count
			}
		}
		counter[t]++
	}

	return ans
}

// Notes

//	1.	wrong return of true/false
//	2.	wrong variable name
//	3.	it should be total count, not just exist
//	4.	count should add by remain, not modular
//	5.	remain number exist, doesn't mean modular number exist
//	6.	slow, using array of length 60 to do, see if it's faster

//	7.	at first glance, I think it's a two sum problem, but then I found it's
//		% not + operator, so I list all numbers that is divisible by 60, tc:
//		O(16n)

//	8.	inspired from https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/discuss/256738/JavaC%2B%2BPython-Two-Sum-with-K-60

//		since after %, number ranges from 0 ~ 59, and then only possible sum that
//		meets critiera is a+b=60, convert it to two-sum problem

//	9.	% operator limits outcome

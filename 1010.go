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

// problems
// 1. wrong return of true/false
// 2. wrong variable name
// 3. it should be total count, not just exist
// 4. count should add by remain, not modular
// 5. remain number exist, doesn't mean modular number exist
// 6. slow, using array of length 60 to do, see if it's faster

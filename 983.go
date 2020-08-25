package main

//In a country popular for train travel, you have planned some train travelling one year in advance.  The days of the year that you will travel is given as an array days.  Each day is an integer from 1 to 365.
//
//Train tickets are sold in 3 different ways:
//
//    a 1-day pass is sold for costs[0] dollars;
//    a 7-day pass is sold for costs[1] dollars;
//    a 30-day pass is sold for costs[2] dollars.
//
//The passes allow that many days of consecutive travel.  For example, if we get a 7-day pass on day 2, then we can travel for 7 days: day 2, 3, 4, 5, 6, 7, and 8.
//
//Return the minimum number of dollars you need to travel every day in the given list of days.
//
//
//
//Example 1:
//
//Input: days = [1,4,6,7,8,20], costs = [2,7,15]
//Output: 11
//Explanation:
//For example, here is one way to buy passes that lets you travel your travel plan:
//On day 1, you bought a 1-day pass for costs[0] = $2, which covered day 1.
//On day 3, you bought a 7-day pass for costs[1] = $7, which covered days 3, 4, ..., 9.
//On day 20, you bought a 1-day pass for costs[0] = $2, which covered day 20.
//In total you spent $11 and covered all the days of your travel.
//
//Example 2:
//
//Input: days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
//Output: 17
//Explanation:
//For example, here is one way to buy passes that lets you travel your travel plan:
//On day 1, you bought a 30-day pass for costs[2] = $15 which covered days 1, 2, ..., 30.
//On day 31, you bought a 1-day pass for costs[0] = $2 which covered day 31.
//In total you spent $17 and covered all the days of your travel.
//
//
//
//Note:
//
//    1 <= days.length <= 365
//    1 <= days[i] <= 365
//    days is in strictly increasing order.
//    costs.length == 3
//    1 <= costs[i] <= 1000

func mincostTickets(days []int, costs []int) int {
	days = append([]int{0}, days...)

	// dp[i] means min cost from day 1 - i
	dp := make([]int, len(days))
	week, month := 0, 0

	for i := 1; i < len(days); i++ {
		// in case calendar day out of boundary
		if days[i] > 365 {
			days[i] = days[i-1]
		}

		for days[i]-days[week] >= 7 {
			week++
		}

		for days[i]-days[month] >= 30 {
			month++
		}

		dp[i] = min(dp[i-1]+costs[0], min(dp[max(0, week-1)]+costs[1], dp[max(0, month-1)]+costs[2]))
	}

	return dp[len(dp)-1]
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from solution, 365 day is a hard limit, otherwise it will reach to next year, even pass test cases,
//		don't forget to check that during interview

//	2.	inspired from https://leetcode.com/problems/minimum-cost-for-tickets/discuss/226659/Two-DP-solutions-with-pictures

//		author provides a very clear explanation

//		although I didn't use 365 day solution, but there's a pretty good
//		technique of using rolling array with size up to 30
//		(because it's only up to 30 days)

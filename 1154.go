package main

import (
	"strconv"
	"strings"
)

//Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.
//
//
//
//Example 1:
//
//Input: date = "2019-01-09"
//Output: 9
//Explanation: Given date is the 9th day of the year in 2019.
//
//Example 2:
//
//Input: date = "2019-02-10"
//Output: 41
//
//Example 3:
//
//Input: date = "2003-03-01"
//Output: 60
//
//Example 4:
//
//Input: date = "2004-03-01"
//Output: 61
//
//
//
//Constraints:
//
//    date.length == 10
//    date[4] == date[7] == '-', and all other date[i]'s are digits
//    date represents a calendar date between Jan 1st, 1900 and Dec 31, 2019.

func dayOfYear(date string) int {
	data := strings.Split(date, "-")
	year, _ := strconv.Atoi(data[0])
	month, _ := strconv.Atoi(data[1])
	day, _ := strconv.Atoi(data[2])

	days := daysOfYear(year)
	var result int
	for i := 0; i+1 < month; i++ {
		result += days[i]
	}

	return result + day
}

func daysOfYear(year int) []int {
	result := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeapYear(year) {
		result[1] = 29
	}

	return result
}

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

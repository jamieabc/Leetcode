package main

import (
	"strconv"
	"strings"
)

// You are given several logs that each log contains a unique id and timestamp. Timestamp is a string that has the following format: Year:Month:Day:Hour:Minute:Second, for example, 2017:01:01:23:59:59. All domains are zero-padded decimal numbers.
//
// Design a log storage system to implement the following functions:
//
// void Put(int id, string timestamp): Given a log's unique id and timestamp, store the log in your storage system.
//
// int[] Retrieve(String start, String end, String granularity): Return the id of logs whose timestamps are within the range from start to end. Start and end all have the same format as timestamp. However, granularity means the time level for consideration. For example, start = "2017:01:01:23:59:59", end = "2017:01:02:23:59:59", granularity = "Day", it means that we need to find the logs within the range from Jan. 1st 2017 to Jan. 2nd 2017.
//
// Example 1:
//
// put(1, "2017:01:01:23:59:59");
// put(2, "2017:01:01:22:59:59");
// put(3, "2016:01:01:00:00:00");
// retrieve("2016:01:01:01:01:01","2017:01:01:23:00:00","Year"); // return [1,2,3], because you need to return all logs within 2016 and 2017.
// retrieve("2016:01:01:01:01:01","2017:01:01:23:00:00","Hour"); // return [1,2], because you need to return all logs start from 2016:01:01:01 to 2017:01:01:23, where log 3 is left outside the range.
//
// Note:
//
//     There will be at most 300 operations of Put or Retrieve.
//     Year ranges from [2000,2017]. Hour ranges from [00,23].
//     Output for Retrieve has no order required.

type LogSystem struct {
	mapping map[string]int
}

func Constructor() LogSystem {
	return LogSystem{
		mapping: make(map[string]int),
	}
}

func (this *LogSystem) Put(id int, timestamp string) {
	this.mapping[timestamp] = id
}

func (this *LogSystem) Retrieve(s string, e string, gra string) []int {
	start, end := truncateStart(s, gra), truncateEnd(e, gra)
	arr1, arr2 := toInt(start), toInt(end)
	result := make([]int, 0)

	for k, v := range this.mapping {
		arr := toInt(k)
		found := true

		for i := range arr {
			if arr[i] < arr1[i] || arr[i] > arr2[i] {
				found = false
				break
			}
		}

		if found {
			result = append(result, v)
		}
	}

	return result
}

func toInt(str string) []int {
	s := strings.Split(str, ":")
	result := make([]int, len(s))

	for i := range s {
		result[i], _ = strconv.Atoi(s[i])
	}

	return result
}

func truncateEnd(str string, gra string) string {
	switch gra {
	case "Year":
		return str[:4] + ":12:31:23:59:59"
	case "Month":
		return str[:7] + ":31:23:59:59"
	case "Day":
		return str[:10] + ":23:59:59"
	case "Hour":
		return str[:13] + ":59:59"
	case "Minute":
		return str[:16] + ":59"
	default:
		return str
	}
}

func truncateStart(str string, gra string) string {
	switch gra {
	case "Year":
		return str[:4] + ":00:00:00:00:00"
	case "Month":
		return str[:7] + ":01:00:00:00"
	case "Day":
		return str[:10] + ":00:00:00"
	case "Hour":
		return str[:13] + ":00:00"
	case "Minute":
		return str[:16] + ":00"
	default:
		return str
	}
}

/**
 * Your LogSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(id,timestamp);
 * param_2 := obj.Retrieve(s,e,gra);
 */

//	problems
//	1.	first thinking about using timestamp, but after referring to others,
//		change mind

//	2.	add reference https://leetcode.com/problems/design-log-storage-system/discuss/105016/Python-Straightforward-with-Explanation

//		author uses map, and for different granularity, truncate string to
//		compare, it's a good solution, so I implement this

//	3.	still fail, because I didn't distinguish situation of same day
//		& different day condition, but I will pass this

package main

//Design a hit counter which counts the number of hits received in the past 5 minutes.
//
//Each function accepts a timestamp parameter (in seconds granularity) and you may assume that calls are being made to the system in chronological order (ie, the timestamp is monotonically increasing). You may assume that the earliest timestamp starts at 1.
//
//It is possible that several hits arrive roughly at the same time.
//
//Example:
//
//HitCounter counter = new HitCounter();
//
//// hit at timestamp 1.
//counter.hit(1);
//
//// hit at timestamp 2.
//counter.hit(2);
//
//// hit at timestamp 3.
//counter.hit(3);
//
//// get hits at timestamp 4, should return 3.
//counter.getHits(4);
//
//// hit at timestamp 300.
//counter.hit(300);
//
//// get hits at timestamp 300, should return 4.
//counter.getHits(300);
//
//// get hits at timestamp 301, should return 3.
//counter.getHits(301);
//
//Follow up:
//What if the number of hits per second could be very large? Does your design scale?

type Counter struct {
	start, end int
}

type HitCounter struct {
	data        map[int]*Counter
	index, prev int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		data: make(map[int]*Counter),
	}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	// not exist in hash, create one
	if _, ok := this.data[timestamp]; !ok {
		this.data[timestamp] = &Counter{
			start: this.index,
			end:   this.index,
		}

		if this.prev != 0 {
			this.data[this.prev].end = this.index - 1
		}
		this.prev = timestamp
	} else {
		this.data[timestamp].end = this.index
	}
	this.index++
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	// past 5 minutes, should be 1 ~ 300, 2 ~ 301, etc.
	start := timestamp - 300 + 1
	if start < 1 {
		start = 1
	}

	var startIdx, endIdx int
	for start <= timestamp {
		if tmp, ok := this.data[start]; !ok {
			start++
		} else {
			startIdx = tmp.start
			break
		}
	}

	if start > timestamp {
		return 0
	}

	for timestamp > start {
		if tmp, ok := this.data[timestamp]; !ok {
			timestamp--
		} else {
			endIdx = tmp.end
			break
		}
	}

	if start == timestamp {
		return this.data[timestamp].end - this.data[timestamp].start + 1
	}

	return endIdx - startIdx + 1
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */

// problems
//  1.	I thought monotonically means no duplicates, but there exists
//  2.	when finding numbers, start range should be first occurrence, end
//	  	range should be last occurrence
//  3.	wrong about next search range, if mid less than target, next search
//	   	range should be larger part
//  4.	not thinking clear about boundary, e.g. when searching for end of
//	   	range, the criteria to go right is <=, which means even if value
//	   	is equal still goes to right, it could potentially have wrong answer
//	   	by this.
//  5.	the way of choosing could potentially be larger, so needs to check
//	   	if reach maximum index
//	6.  when there's always a return value, I forget to check if return
//		value is valid, e.g. start and end index is invalid
//	7.	optimize, I forget that input of every number is chronological,
//		which means every input is larger than previous one. This could
//		reduce the complexity, since no longer to save data more than
//		5 minutes.
//
//		And considering if the input coming rate is pretty high, I would
//		use hash to store data, for quickly query.
//
//		This problem doesn't care about mid status, which means no need to
//		store every data, just count. So I would use a mono increasing
//		number as index, and hash value stores a number's start & end index.

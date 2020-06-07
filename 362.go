package main

// Design a hit counter which counts the number of hits received in the past 5 minutes.
//
// Each function accepts a timestamp parameter (in seconds granularity) and you may assume that calls are being made to the system in chronological order (ie, the timestamp is monotonically increasing). You may assume that the earliest timestamp starts at 1.
//
// It is possible that several hits arrive roughly at the same time.
//
// Example:
//
// HitCounter counter = new HitCounter();
//
// hit at timestamp 1.
// counter.hit(1);
//
// hit at timestamp 2.
// counter.hit(2);
//
// hit at timestamp 3.
// counter.hit(3);
//
// get hits at timestamp 4, should return 3.
// counter.getHits(4);
//
// hit at timestamp 300.
// counter.hit(300);
//
// get hits at timestamp 300, should return 4.
// counter.getHits(300);
//
// get hits at timestamp 301, should return 3.
// counter.getHits(301);
//
// Follow up:
// What if the number of hits per second could be very large? Does your design scale?

type HitCounter struct {
	hits  []int
	times []int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		hits:  make([]int, 300),
		times: make([]int, 300),
	}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	idx := timestamp % 300
	if this.times[idx] != timestamp {
		this.hits[idx] = 1
		this.times[idx] = timestamp
	} else {
		this.hits[idx]++
	}
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	var sum int
	for i := range this.hits {
		if timestamp-this.times[i] < 300 {
			sum += this.hits[i]
		}
	}

	return sum
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

//	8.	optimize, since time is always mono-increasing, there's actually only
//		needs to store 300 seconds of data. And if data exceeds current limit,
//		cleanup old ones

//	9.	too slow, because array clean up should only happens when necessary,
//		e.g. if timestamp is 100, 200, 100000, 100001, then cleanup of 100001
//		is not necessary

//	10.	reference uses another array to store each hits corresponds time, since
//		time is mono-increasing, it's unique, so another array is used to denotes
//		if time has elapsed more than 300 seconds.

//		reference: https://leetcode.com/problems/design-hit-counter/discuss/83483/Super-easy-design-O(1)-hit()-O(s)-getHits()-no-fancy-data-structure-is-needed!

//		this method is thread safe only if every incoming calls are time
//		monotonically increasing.

//		the other thing to concern is that in real world, hit is usually
//		higher frequency than getHit, so when designing api, cannot put
//		cleanup in get function because it might called after a long
//		period of time

//		to make sure it's thread safe, use atomic.AddInt32

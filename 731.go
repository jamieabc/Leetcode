package main

import "sort"

// Implement a MyCalendarTwo class to store your events. A new event can be added if adding the event will not cause a triple booking.
//
// Your class will have one method, book(int start, int end). Formally, this represents a booking on the half open interval [start, end), the range of real numbers x such that start <= x < end.
//
// A triple booking happens when three events have some non-empty intersection (ie., there is some time that is common to all 3 events.)
//
// For each call to the method MyCalendar.book, return true if the event can be added to the calendar successfully without causing a triple booking. Otherwise, return false and do not add the event to the calendar.
// Your class will be called like this: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)
//
// Example 1:
//
// MyCalendar();
// MyCalendar.book(10, 20); // returns true
// MyCalendar.book(50, 60); // returns true
// MyCalendar.book(10, 40); // returns true
// MyCalendar.book(5, 15); // returns false
// MyCalendar.book(5, 10); // returns true
// MyCalendar.book(25, 55); // returns true
// Explanation:
// The first two events can be booked.  The third event can be double booked.
// The fourth event (5, 15) can't be booked, because it would result in a triple booking.
// The fifth event (5, 10) can be booked, as it does not use time 10 which is already double booked.
// The sixth event (25, 55) can be booked, as the time in [25, 40) will be double booked with the third event;
// the time [40, 50) will be single booked, and the time [50, 55) will be double booked with the second event.
//
//
//
// Note:
//
//     The number of calls to MyCalendar.book per test case will be at most 1000.
//     In calls to MyCalendar.book(start, end), start and end are integers in the range [0, 10^9].

type Event struct {
	Start, End int
}

type MyCalendarTwo struct {
	Booked   []Event
	Overlaps []Event
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		Booked:   make([]Event, 0),
		Overlaps: make([]Event, 0),
	}
}

func (this *MyCalendarTwo) Book(start, end int) bool {
	for _, event := range this.Overlaps {
		if max(start, event.Start) < min(end, event.End) {
			return false
		}
	}

	for _, event := range this.Booked {
		overlapStart, overlapEnd := max(start, event.Start), min(end, event.End)

		// overlap
		if overlapStart < overlapEnd {
			this.Overlaps = append(this.Overlaps, Event{
				Start: overlapStart,
				End:   overlapEnd,
			})
		}
	}

	// not found any partial overlap
	this.Booked = append(this.Booked, Event{
		Start: start,
		End:   end,
	})

	return true
}

type Node struct {
	Start, End  int
	Overlap     bool
	Left, Right *Node
}

type MyCalendarTwo struct {
	Root *Node
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

// worst case: O(n^2)
func (this *MyCalendarTwo) Book(start, end int) bool {
	if this.Root == nil {
		this.Root = &Node{
			Start: start,
			End:   end,
		}
		return true
	}

	// check if it's insertable
	if !Insertable(this.Root, start, end) {
		return false
	}

	// insert
	this.Root.Insert(start, end)

	return true
}

func (this *Node) Insert(start, end int) {
	if start == end {
		return
	}

	// non-overlap
	if end <= this.Start {
		if this.Left == nil {
			this.Left = &Node{
				Start: start,
				End:   end,
			}
			return
		}
		this.Left.Insert(start, end)
		return
	} else if start >= this.End {
		if this.Right == nil {
			this.Right = &Node{
				Start: start,
				End:   end,
			}
			return
		}
		this.Right.Insert(start, end)
		return
	}

	this.Overlap = true
	overlapStart, overlapEnd := max(this.Start, start), min(this.End, end)

	// insert left non-overlap interval
	l, r := min(start, this.Start), max(start, this.Start)
	if this.Left == nil && start != this.Start {
		this.Left = &Node{
			Start: l,
			End:   r,
		}
	} else {
		this.Left.Insert(l, r)
	}

	// insert right non-overlap interval
	l, r = min(end, this.End), max(end, this.End)
	if this.Right == nil && end != this.End {
		this.Right = &Node{
			Start: l,
			End:   r,
		}
	} else {
		this.Right.Insert(l, r)
	}

	// update self range
	this.Start, this.End = overlapStart, overlapEnd
}

func Insertable(n *Node, start, end int) bool {
	if n == nil || start == end {
		return true
	}

	// non-overlap
	if end <= n.Start || start >= n.End {
		return Insertable(n.Left, start, end) && Insertable(n.Right, start, end)
	}

	// already overlap
	if n.Overlap {
		return false
	}

	// check children insertable
	return Insertable(n.Left, min(n.Start, start), max(n.Start, start)) && Insertable(n.Right, min(n.End, end), max(n.End, end))
}

type MyCalendarTwo struct {
	Booked map[int]int
	Limit  int
	Events []int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		Booked: make(map[int]int),
		Limit:  2,
		Events: make([]int, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	if _, ok := this.Booked[start]; !ok {
		this.Events = append(this.Events, start)
	}

	if _, ok := this.Booked[end]; !ok {
		this.Events = append(this.Events, end)
	}

	this.Booked[start]++
	this.Booked[end]--

	sort.Ints(this.Events)
	var count int
	for _, event := range this.Events {
		count += this.Booked[event]

		if count > this.Limit {
			this.Booked[start]--
			this.Booked[end]++

			return false
		}
	}

	return true
}

func binarySearch(data []int, target int) int {
	var i, j int
	for i, j = 0, len(data)-1; i < j; {
		mid := i + (j-i)/2

		if data[mid] < target {
			i = mid + 1
		} else if data[mid] > target {
			j = mid - 1
		} else {
			return mid
		}
	}

	return i
}

type event struct {
	start, end int
}

type MyCalendarTwo struct {
	events       []event
	limit        int
	starts, ends []int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		events: make([]event, 0),
		limit:  2,
		starts: make([]int, 0),
		ends:   make([]int, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	if len(this.events) == 0 {
		this.events = append(this.events, event{start, end})
	} else {
		if len(this.starts) == 0 {
			// initialize data
			for _, e := range this.events {
				this.starts = append(this.starts, e.start)
				this.ends = append(this.ends, e.end)
			}
		}

		// add newly event
		this.starts = append(this.starts, start)
		this.ends = append(this.ends, end)

		sort.Ints(this.starts)
		sort.Ints(this.ends)

		var i, j, count int
		for i, j = 0, 0; i < len(this.starts); {
			if this.starts[i] < this.ends[j] {
				count++
				i++

				// clean-up
				if count > this.limit {
					this.starts = this.starts[:0]
					this.ends = this.ends[:0]
					return false
				}
			} else if this.starts[i] > this.ends[j] {
				count--
				j++
			} else {
				// no change
				i++
				j++
			}
		}

		// pass test
		this.events = append(this.events, event{start, end})
	}

	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

//	problems
//	1.	when using map to store events, this.Events add start/end when no
//		key exist. And no need to remove keys, since it won't affect result

//	2.	inspired from https://leetcode.com/problems/my-calendar-ii/discuss/114882/Java-Binary-Search-Tree-method-clear-and-easy-to-undertand-beats-99

//		author uses binary search, but add some changes:
//		- for each node, add attributes of overlap to denote if an interval
//		  is already used
//		- when incoming interval overlap w/ origina, split incoming
//		  interval into overlap one and non- overlap one, update original
//		  interval start & end & overlap, then add non-overlap interval
//		  into tree

//		I think this is the right way to do it, not only because it's fast,
//		but also it utilizes original binary search concept, though it's
//		not scalable.

//	3.	beware of range segment update, if two points not same, e.g.
//		start != this.Start then there definitely exists a non-overlap
//		region

//	4.	inspiref from https://leetcode.com/problems/my-calendar-ii/discuss/109528/nlogd-Java-solution-using-segment-tree-with-lazy-propagation-(for-the-general-case-of-K-booking)

//		author writes about segment tree

//	5.	inspired from https://leetcode.com/problems/my-calendar-ii/discuss/109530/N2-Python-Short-and-Elegant

//		uses another array to store overlap intervals

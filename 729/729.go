package main

// Implement a MyCalendar class to store your events. A new event can be added if adding the event will not cause a double booking.
//
// Your class will have the method, book(int start, int end). Formally, this represents a booking on the half open interval [start, end), the range of real numbers x such that start <= x < end.
//
// A double booking happens when two events have some non-empty intersection (ie., there is some time that is common to both events.)
//
// For each call to the method MyCalendar.book, return true if the event can be added to the calendar successfully without causing a double booking. Otherwise, return false and do not add the event to the calendar.
// Your class will be called like this: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)
//
// Example 1:
//
// MyCalendar();
// MyCalendar.book(10, 20); // returns true
// MyCalendar.book(15, 25); // returns false
// MyCalendar.book(20, 30); // returns true
// Explanation:
// The first event can be booked.  The second can't because time 15 is already booked by another event.
// The third event can be booked, as the first event takes every time less than 20, but not including 20.
//
//
//
// Note:
//
//     The number of calls to MyCalendar.book per test case will be at most 1000.
//     In calls to MyCalendar.book(start, end), start and end are integers in the range [0, 10^9].

type Node struct {
	Left, Right *Node
	Start, End  int
}

func (n *Node) insert(start, end int) bool {
	if start < n.Start {
		if end <= n.Start {
			if n.Left == nil {
				n.Left = &Node{
					Start: start,
					End:   end,
				}
				return true
			}

			return n.Left.insert(start, end)
		}

		return false
	} else if start >= n.End {
		if n.Right == nil {
			n.Right = &Node{
				Start: start,
				End:   end,
			}
			return true
		}

		return n.Right.insert(start, end)
	}

	return false
}

type MyCalendar struct {
	root *Node
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if this.root == nil {
		this.root = &Node{
			Start: start,
			End:   end,
		}
		return true
	}
	return this.root.insert(start, end)
}

type MyCalendar struct {
	Booked [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		Booked: make([][2]int, 0),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	// insert time
	this.Booked = append(this.Booked, [2]int{start, end})

	// find proper location based on start time
	var i, j int
	if len(this.Booked) > 1 {
		for i, j = 0, len(this.Booked)-2; i <= j; {
			if this.Booked[i][0] <= start {
				i++
			} else if this.Booked[j][0] > start {
				j--
			} else {
				this.Booked[i], this.Booked[j] = this.Booked[j], this.Booked[i]
				i++
				j--
			}
		}
		this.Booked[i], this.Booked[len(this.Booked)-1] = this.Booked[len(this.Booked)-1], this.Booked[i]
	}

	// check end time of previous slots no larger than start time
	for j = 0; j < i; j++ {
		if this.Booked[j][1] > start {
			// cleanup
			this.Booked = append(this.Booked[:i], this.Booked[i+1:]...)
			return false
		}
	}

	// check for start time of later slots no smaller than end time
	for j = i + 1; j < len(this.Booked); j++ {
		if end > this.Booked[j][0] {
			// cleanup
			this.Booked = append(this.Booked[:i], this.Booked[i+1:]...)
			return false
		}
	}

	return true
}

type MyCalendar struct {
	Booked map[int]bool
}

func Constructor() MyCalendar {
	return MyCalendar{
		Booked: make(map[int]bool),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	var i int
	for i = start; i < end; i++ {
		if this.Booked[i] {
			i--
			break
		} else {
			this.Booked[i] = true
		}
	}

	if i != end {
		// cleanup
		for ; i >= start; i-- {
			this.Booked[i] = false
		}
		return false
	}

	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

//	problems
//	1.	cannot use key exist or not, it could be invalid in one query but
//		valid in others, but key is already inserted

//	2.	too slow, use array instead

//	3.	out of memory

//	4.	inspired from https://leetcode.com/problems/my-calendar-i/discuss/109475/JavaC%2B%2B-Clean-Code-with-Explanation

//		author provides a very elegant way to checking overlay:
//		max(start) < min(end)

//	5.	inspired from sample code, use binary search technique to store
//		trees.

//		every node stores start, end, left node, and right node.
//		valid situations:
//		- new start <= current start && new end <= current start
//		- new start >= current end

//		the concept that binary search tree can be used here is because
//		there's no range overlap within, every range can be determined by
//		smaller or larger, no equal exists

//	6.	inspired from https://leetcode.com/problems/my-calendar-i/discuss/109463/Short-brute-force-python-solution-and-binary-search-based-version

//		author has another approach to decide range overlap:
//		new start < current end && new end > current start
//		comes from !(new start >= current end || new end <= current start)

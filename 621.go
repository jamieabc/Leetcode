package main

import (
	"container/heap"
	"math"
)

// You are given a char array representing tasks CPU need to do. It contains capital letters A to Z where each letter represents a different task. Tasks could be done without the original order of the array. Each task is done in one unit of time. For each unit of time, the CPU could complete either one task or just be idle.
//
// However, there is a non-negative integer n that represents the cooldown period between two same tasks (the same letter in the array), that is that there must be at least n units of time between any two same tasks.
//
// You need to return the least number of units of times that the CPU will take to finish all the given tasks.
//
//
//
// Example 1:
//
// Input: tasks = ["A","A","A","B","B","B"], n = 2
// Output: 8
// Explanation:
// A -> B -> idle -> A -> B -> idle -> A -> B
// There is at least 2 units of time between any two same tasks.
// Example 2:
//
// Input: tasks = ["A","A","A","B","B","B"], n = 0
// Output: 6
// Explanation: On this case any permutation of size 6 would work since n = 0.
// ["A","A","A","B","B","B"]
// ["A","B","A","B","A","B"]
// ["B","B","B","A","A","A"]
// ...
// And so on.
// Example 3:
//
// Input: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
// Output: 16
// Explanation:
// One possible solution is
// A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> idle -> idle -> A -> idle -> idle -> A
//
//
// Constraints:
//
// The number of tasks is in the range [1, 10000].
// The integer n is in the range [0, 100].

type Task struct {
	Name  byte
	Count int
}

type Tasks []Task

func (h Tasks) Len() int { return len(h) }

// no need to have fixed order, since this problem only wants count
func (h Tasks) Less(i, j int) bool {
	return h[i].Count > h[j].Count
	// if h[i].Count > h[j].Count {
	// 	return true
	// } else if h[i].Count < h[j].Count {
	// 	return false
	// }
	// return h[i].Name < h[j].Name
}

func (h Tasks) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Tasks) Peek() Task    { return h[0] }

func (h *Tasks) Push(x interface{}) {
	*h = append(*h, x.(Task))
}

func (h *Tasks) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// tc: O(m log k), m: n * len(tasks), k: # of distinct task
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	h := &Tasks{}
	heap.Init(h)

	// find task counts
	counter := make([]int, 256)
	for i := range tasks {
		counter[tasks[i]]++
	}

	// put tasks into max heap
	for i := range counter {
		if counter[i] > 0 {
			heap.Push(h, Task{
				Name:  byte(i),
				Count: counter[i],
			})
		}
	}

	var duration, i int
	var popped Task
	queue := make([]Task, 0)

	for h.Len() > 0 {
		// fill task from max count
		for i = 0; i <= n; i++ {
			if h.Len() == 0 {
				break
			}

			popped = heap.Pop(h).(Task)
			popped.Count--

			if popped.Count > 0 {
				queue = append(queue, popped)
			}
		}

		for len(queue) > 0 {
			heap.Push(h, queue[0])
			queue = queue[1:]
		}

		if h.Len() > 0 {
			duration += n + 1
		} else {
			duration += i
		}
	}

	return duration
}

// tc: O(256 + n)
func leastInterval1(tasks []byte, n int) int {
	counter := make([]int, 26)
	for i := range tasks {
		counter[tasks[i]-'A']++
	}

	validNextIndex := make([]int, 26)

	var duration int

	for true {
		idx, available := getNextTask(counter, validNextIndex, duration)
		if available == 0 {
			break
		}
		duration++

		// cannot find any valid task, expand duration
		if idx == -1 {
			continue
		}

		validNextIndex[idx] = (duration - 1) + n + 1
		counter[idx]--
	}

	return duration
}

func getNextTask(counter, validNextIndex []int, idx int) (int, int) {
	maxCount := math.MinInt32
	var available int
	next := -1

	// find next max count task
	for i := range counter {
		if counter[i] > 0 {
			available++
			if counter[i] > maxCount && validNextIndex[i] <= idx {
				maxCount = counter[i]
				next = i
			}
		}
	}

	return next, available
}

//	problems
//	1.	inspired from https://leetcode.com/problems/task-scheduler/discuss/104500/Java-O(n)-time-O(1)-space-1-pass-no-sorting-solution-with-detailed-explanation

//		author thinks in a way that if all max count characters are arranges, then
//		rest spaces are for less-repeated characters

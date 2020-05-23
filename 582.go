package main

// Given n processes, each process has a unique PID (process id) and its PPID (parent process id).
//
// Each process only has one parent process, but may have one or more children processes. This is just like a tree structure. Only one process has PPID that is 0, which means this process has no parent process. All the PIDs will be distinct positive integers.
//
// We use two list of integers to represent a list of processes, where the first list contains PID for each process and the second list contains the corresponding PPID.
//
// Now given the two lists, and a PID representing a process you want to kill, return a list of PIDs of processes that will be killed in the end. You should assume that when a process is killed, all its children processes will be killed. No order is required for the final answer.
//
// Example 1:
//
// Input:
// pid =  [1, 3, 10, 5]
// ppid = [3, 0, 5, 3]
// kill = 5
// Output: [5,10]
// Explanation:
//            3
//          /   \
//         1     5
//              /
//             10
// Kill 5 will also kill 10.
//
// Note:
//
//     The given kill id is guaranteed to be one of the given PIDs.
//     n >= 1.

// iterative
func killProcess(pid []int, ppid []int, kill int) []int {
	if len(pid) == 0 || len(ppid) == 0 || kill == 0 {
		return []int{}
	}

	mapping := make(map[int][]int)
	for i := range ppid {
		mapping[ppid[i]] = append(mapping[ppid[i]], pid[i])
	}

	stack := []int{kill}
	killed := make([]int, 0)

	for len(stack) > 0 {
		target := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		killed = append(killed, target)
		if tmp, ok := mapping[target]; ok {
			stack = append(stack, tmp...)
		}
	}

	return killed
}

// recursive
func killProcess1(pid []int, ppid []int, kill int) []int {
	if len(pid) == 0 || len(ppid) == 0 || kill == 0 {
		return []int{}
	}

	mapping := make(map[int][]int) // parent -> children

	for i := range ppid {
		mapping[ppid[i]] = append(mapping[ppid[i]], pid[i])
	}

	killed := []int{kill}
	traverse(mapping, kill, &killed)

	return killed
}

func traverse(mapping map[int][]int, target int, killed *[]int) {
	*killed = append(*killed, mapping[target]...)
	for i := range mapping[target] {
		traverse(mapping, mapping[target][i], killed)
	}
}

//	problems
//	1.	too slow, use map to search

package main

import (
	"fmt"
	"math"
)

//  You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.
//
// The lock initially starts at '0000', a string representing the state of the 4 wheels.
//
// You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.
//
// Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.
//
// Example 1:
//
// Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
// Output: 6
// Explanation:
// A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
// Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
// because the wheels of the lock become stuck after the display becomes the dead end "0102".
//
// Example 2:
//
// Input: deadends = ["8888"], target = "0009"
// Output: 1
// Explanation:
// We can turn the last wheel in reverse to move from "0000" -> "0009".
//
// Example 3:
//
// Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
// Output: -1
// Explanation:
// We can't reach the target without getting stuck.
//
// Example 4:
//
// Input: deadends = ["0000"], target = "8888"
// Output: -1
//
// Note:
//
//     The length of deadends will be in the range [1, 500].
//     target will not be in the list deadends.
//     Every string in deadends and the string target will be a string of 4 digits from the 10,000 possibilities '0000' to '9999'.

func openLock(deadends []string, target string) int {
    invalid := make(map[string]bool)
    for _, dead := range deadends {
        invalid[dead] = true
    }

    if invalid["0000"] || invalid[target] {
        return -1
    }

    if target == "0000" {
        return 0
    }

    visited1, visited2 := make(map[string]bool), make(map[string]bool)
    queue1, queue2 := []string{"0000"}, []string{target}

	// in case first step reaches target, because loop assumes
	// destination already in hash map (visited), so need to firt
	// make sure start/end strings are valid
    visited1["0000"] = true
    visited2[target] = true

    var steps int

    // bfs
    for len(queue1) > 0 && len(queue2) > 0 {
        // select bfs w/ less items
        if len(queue1) > len(queue2) {
            queue1, queue2 = queue2, queue1
            visited1, visited2 = visited2, visited1
        }

        size := len(queue1)
        steps++

        for i := 0; i < size; i++ {
            str := queue1[0]
            queue1 = queue1[1:]

            for _, lock := range rotate([]byte(str)) {
                if !visited1[lock] && !invalid[lock] {
                    queue1 = append(queue1, lock)
                    visited1[lock] = true

                    if visited2[lock] {
                        return steps
                    }
                }
            }
        }
    }

    return -1
}

var nextStep = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var prevStep = []byte{'9', '0', '1', '2', '3', '4', '5', '6', '7', '8'}

func rotate(bytes []byte) []string {
    ans := make([]string, 0)

    for i := range bytes {
        orig := bytes[i]

        bytes[i] = nextStep[orig-'0']
        ans = append(ans, string(bytes))

        bytes[i] = prevStep[orig-'0']
        ans = append(ans, string(bytes))

        bytes[i] = orig
    }

    return ans
}

func openLock2(deadends []string, target string) int {
	invalid := make(map[string]bool)
	for _, deadend := range deadends {
		invalid[deadend] = true
	}

	if target == "0000" {
		return 0
	}

	if invalid["0000"] {
		return -1
	}

	visited := make(map[string]int)
	visited["0000"] = 0
	minCost := math.MaxInt32

	dfs([]byte{'0', '0', '0', '0'}, target, 0, &minCost, invalid, visited)

	if minCost == math.MaxInt32 {
		return -1
	}

	return minCost
}

func dfs(current []byte, target string, cost int, minCost *int, invalid map[string]bool, visited map[string]int) {
	for i := range current {
		possibilities := [][]byte{nextNum(current, i), prevNum(current, i)}

		for _, bytes := range possibilities {
			str := string(bytes)

			if str == target {
				*minCost = min(*minCost, cost+1)
				return
			}

			if !invalid[str] {
				if steps, ok := visited[str]; !ok {
					visited[str] = cost + 1
					dfs(bytes, target, cost+1, minCost, invalid, visited)
				} else {
					if steps > cost+1 {
						visited[str] = cost + 1
						dfs(bytes, target, cost+1, minCost, invalid, visited)
					}
				}
			}
		}
	}
}

func openLock1(deadends []string, target string) int {
	invalid := make(map[string]bool)
	for _, d := range deadends {
		invalid[d] = true
	}

	if target == "0000" {
		return 0
	}

	// initial state in dead end
	if invalid["0000"] {
		return -1
	}

	queue := [][]byte{
		{'0', '0', '0', '0'},
	}

	visited := make(map[string]bool)
	var str string
	steps := 1

	for len(queue) > 0 {
		end := len(queue)

		for i := 0; i < end; i++ {
			popped := queue[i]

			for i := 0; i < len(target); i++ {
				possibilities := [][]byte{nextNum(popped, i), prevNum(popped, i)}

				for _, bytes := range possibilities {
					str = string(bytes)

					if str == target {
						return steps
					}

					if !invalid[str] && !visited[str] {
						visited[str] = true
						queue = append(queue, bytes)
					}

				}
			}
		}
		queue = queue[end:]
		steps++
	}

	return -1
}

func nextNum(current []byte, idx int) []byte {
	next := make([]byte, len(current))
	copy(next, current)

	if next[idx] == '9' {
		next[idx] = '0'
	} else {
		next[idx]++
	}

	return next
}

func prevNum(current []byte, idx int) []byte {
	next := make([]byte, len(current))
	copy(next, current)

	if next[idx] == '0' {
		next[idx] = '9'
	} else {
		next[idx]--
	}

	return next
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

// 	Notes
//	1.	why can't I write DFS solution? I don't know how to stop recursion.
//		because there are 2 ways of iterating: increase number or decrease
//		number.

//		I don't know how to write a program that can detect increment,
//		and when encounter any deadend, stop it.

//	2.	no need to detect, just iterate through all possibilities

//	3.	when implementing dfs, I encounter a problem: dfs needs to store
//		map[string]int, due to nature of dfs, sequence will like
//		0000 -> 0001 -> 0002 -> 0003

//		so even if a number is visited, there's still possibility other
//		path exists smaller outcome

//	4.	inspired from https://leetcode.com/problems/open-the-lock/discuss/110230/BFS-solution-C%2B%2B

//		bi-directional search starts from 0000 and target, and each time
//		choose sets with smaller size

//	5.	inspired from solution, for every number, there are 2 directions(
//		increment & decrement), there arn n digits, and when generating
//		new number copy take O(2n^2)

//		there are total 10^4 possible combinations, and first need iterate
//		invalid words

//		tc: O(n^2 * a^n + d), n: 4 digits, a: 10 digits, d: size of deadends

//	6.	inspired from https://leetcode.com/problems/open-the-lock/discuss/1250580/C%2B%2BJavaPython-BFS-Level-Order-Traverse-Clean-and-Concise
//
//		there are 10^4 numbers, each number needs to scan whole string and
//		did the string replacement (n*n)
//
//		overall tc would be (10^4 * n^2 + d)

//	7.	inspired from solution, use two sided BFS to run faster, the point is to
//		choose BFS with less items

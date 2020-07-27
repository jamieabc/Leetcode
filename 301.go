package main

// Remove the minimum number of invalid parentheses in order to make the input string valid. Return all possible results.
//
// Note: The input string may contain letters other than the parentheses ( and ).
//
// Example 1:
//
// Input: "()())()"
// Output: ["()()()", "(())()"]
//
// Example 2:
//
// Input: "(a)())()"
// Output: ["(a)()()", "(a())()"]
//
// Example 3:
//
// Input: ")("
// Output: [""]

func removeInvalidParentheses(s string) []string {
	// original s is valid, no need to do anything
	if isValid(s) {
		return []string{s}
	}

	valids := make(map[string]bool)
	seen := make(map[string]bool)
	queue := []string{s}

	stop := false

	for len(queue) > 0 && !stop {
		end := len(queue)

		for i := 0; i < end; i++ {
			old := queue[0]
			queue = queue[1:]

			if seen[old] {
				continue
			} else {
				seen[old] = true
			}

			for j := 0; j < len(old); j++ {
				// remove one ( or ) each time, not including normal char
				if old[j] != '(' && old[j] != ')' {
					continue
				}

				str := old[:j] + old[j+1:]

				if isValid(str) {
					stop = true
					valids[str] = true
				} else {
					queue = append(queue, str)
				}
			}
		}
	}

	result := make([]string, 0)
	for s := range valids {
		result = append(result, s)
	}

	return result
}

func isValid(str string) bool {
	var count int
	for i := range str {
		if str[i] == '(' {
			count++
		} else if str[i] == ')' {
			count--
		}

		if count < 0 {
			return false
		}
	}

	return count == 0
}

//	problems
//	1.	contain normal character, need to ignore it

//	2.	bfs might allocates too much memory, need to prune it

//	3.	inspired form https://leetcode.com/problems/remove-invalid-parentheses/discuss/75032/Share-my-Java-BFS-solution

//		first level total n possibilities
//		second level (n-1) * c(n, 2)
//		total: n + (n-1) * c(n, 2) + (n-3) * c(n, 3) + ...
//		tc: O(n * 2^(n-1))

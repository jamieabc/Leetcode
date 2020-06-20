package main

import "fmt"

// Equations are given in the format A / B = k, where A and B are variables represented as strings, and k is a real number (floating point number). Given some queries, return the answers. If the answer does not exist, return -1.0.
//
// Example:
// Given a / b = 2.0, b / c = 3.0.
// queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? .
// return [6.0, 0.5, -1.0, 1.0, -1.0 ].
//
// The input is: vector<pair<string, string>> equations, vector<double>& values, vector<pair<string, string>> queries , where equations.size() == values.size(), and the values are positive. This represents the equations. Return vector<double>.
//
// According to the example above:
//
// equations = [ ["a", "b"], ["b", "c"] ],
// values = [2.0, 3.0],
// queries = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"] ].
//
//
//
// The input is always valid. You may assume that evaluating the queries will result in no division by zero and there is no contradiction.

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	parents := make(map[string]string)
	factors := make(map[string]float64)

	for i, eq := range equations {
		// find
		r1 := find(parents, factors, eq[0])
		r2 := find(parents, factors, eq[1])

		// union
		// make latest divisor to be parent of dividend
		parents[r1] = r2
		// update factor, since factor[r2] = 1, factor[r1] is no longer 1
		// factor[r1] is get from relative relationships
		factors[r1] = values[i] * factors[eq[1]] / factors[eq[0]]
	}

	results := make([]float64, len(queries))
	for i := range results {
		results[i] = -1
	}

	for i, query := range queries {
		dividend, divisor := query[0], query[1]
		if _, ok := parents[dividend]; !ok {
			continue
		}

		if _, ok := parents[divisor]; !ok {
			continue
		}

		r1 := find(parents, factors, dividend)
		r2 := find(parents, factors, divisor)

		if r1 != r2 {
			continue
		}

		results[i] = factors[dividend] / factors[divisor]
	}

	return results
}

func find(parents map[string]string, factors map[string]float64, str string) string {
	if _, ok := parents[str]; !ok {
		parents[str] = str
		factors[str] = float64(1)
		return str
	}

	parent := parents[str]
	if parent == str {
		return str
	}

	root := find(parents, factors, parent)
	parents[str] = root

	// update factor according parent, recursive process
	factors[str] *= factors[parent]

	return parents[str]
}

type query struct {
	dividend string
	divisor  string
}

func calcEquation1(equations [][]string, values []float64, queries [][]string) []float64 {
	// create a hashmap to store possible equations, map[string]map[string]int
	eqs := make(map[string]map[string]float64)
	for i, eq := range equations {
		addPath(eq[0], eq[1], values[i], eqs)
	}

	// for any query, make sure both string exist
	// check if dividend string == divisor string
	// merge [a,b] & [b,c] = [a,c]
	// bfs with visited map

	result := make([]float64, len(queries))
	for i := range result {
		result[i] = -1
	}
	for i, q := range queries {
		dividend, divisor := q[0], q[1]
		if _, ok := eqs[dividend]; !ok {
			result[i] = float64(-1)
			continue
		}

		if _, ok := eqs[divisor]; !ok {
			result[i] = float64(-1)
			continue
		}

		if dividend == divisor {
			result[i] = float64(1)
			continue
		}

		visited := make(map[string]map[string]bool)
		queue := make([]query, 0)

		// start from dividend
		for d := range eqs[dividend] {
			addVisited(dividend, d, visited)

			queue = append(queue, query{
				dividend: dividend,
				divisor:  d,
			})
		}

		for len(queue) > 0 {
			end := len(queue)

			for j := 0; j < end; j++ {
				popped := queue[j]

				for d, value := range eqs[popped.divisor] {
					if popped.dividend == dividend && d == divisor {
						result[i] = eqs[popped.dividend][popped.divisor] * value
						j = end
						break
					}

					if visited[popped.dividend][d] {
						continue
					}

					addVisited(dividend, d, visited)
					addPath(popped.dividend, d, eqs[popped.dividend][popped.divisor]*value, eqs)

					queue = append(queue, query{
						dividend: popped.dividend,
						divisor:  d,
					})
				}
			}

			queue = queue[end:]
		}
	}

	return result
}

func addPath(dividend, divisor string, result float64, eqs map[string]map[string]float64) {
	if _, ok := eqs[dividend]; !ok {
		eqs[dividend] = make(map[string]float64)
	}

	if _, ok := eqs[dividend][divisor]; !ok {
		eqs[dividend][divisor] = result
	} else {
		return
	}

	if _, ok := eqs[divisor]; !ok {
		eqs[divisor] = make(map[string]float64)
	}

	if _, ok := eqs[divisor][dividend]; !ok {
		eqs[divisor][dividend] = 1 / result
	}
}

func addVisited(dividend, divisor string, visited map[string]map[string]bool) {
	if _, ok := visited[dividend]; !ok {
		visited[dividend] = make(map[string]bool)
	}

	visited[dividend][divisor] = true

	if _, ok := visited[divisor]; !ok {
		visited[divisor] = make(map[string]bool)
	}
	visited[divisor][dividend] = true
}

//	problems
//	1.	in worst case, every query visits all equations once, tc: O(m*n^2),
//		n: # of variables in equations, m: size of query, and in worst case
//		every variable has at most n-1 connections to all other variables

//		sc: O(n*k), k: average size for a specific string can extend

//	2.	inspired from https://leetcode.com/problems/evaluate-division/discuss/88275/Python-fast-BFS-solution-with-detailed-explantion

//		this is a graph problem, a/b = 2 can be viewed as a -> b w/ edge
//		weight 2, and b/c = 4 can be viewed as b -> c w/ edge 4,
//		and a -> c can be traveled by a -> b -> c w/ edge 2*4=8

//		also, a compression can be done by saving all intermediate values
//		to speed up further calculations

//		if compress path, then sc can grow to O(n^2), and query takes O(1)

//		also, create a function to add transitions(edges)

//	3.	add reference https://leetcode.com/problems/evaluate-division/discuss/147281/Java-Union-Find-solution-faster-than-99

//		author uses union-find to solve the problem, interesting but not
//		implement

//		when using union-find, divisor is always treated as root, occurrence
//		order doesn't matter, as long as factor relationships is correct

//		it's a pretty elegant solution

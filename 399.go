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
	parents, factors := buildRelation(equations, values)

	result := make([]float64, len(queries))
	for i, query := range queries {
		if _, ok := parents[query[0]]; !ok {
			result[i] = float64(-1)
			continue
		}

		if _, ok := parents[query[1]]; !ok {
			result[i] = float64(-1)
			continue
		}

		r1 := find(parents, factors, query[0])
		r2 := find(parents, factors, query[1])

		if r1 != r2 {
			result[i] = float64(-1)
		} else {
			result[i] = factors[query[0]] / factors[query[1]]
		}
	}

	return result
}

func buildRelation(equations [][]string, values []float64) (map[string]string, map[string]float64) {
	parents, factors := make(map[string]string), make(map[string]float64)

	for i, eq := range equations {
		r1 := find(parents, factors, eq[0])
		r2 := find(parents, factors, eq[1])

		// a = k*r1, b = j*r2, a = m*b
		// k*r1 = m*j*r2, r1/r2 = m*j/k

		// a/b = 2, a = 2*b, so b is a's parent
		parents[r1] = r2
		factors[r1] = values[i] * factors[eq[1]] / factors[eq[0]]
	}

	return parents, factors
}

func find(parents map[string]string, factors map[string]float64, target string) string {
	if _, ok := parents[target]; !ok {
		parents[target] = target
		factors[target] = float64(1)
		return target
	}

	if parents[target] == target {
		return target
	}

	p := find(parents, factors, parents[target])

	// factors need to be updated if parent has been changed,
	// otherwise, factor is not correct
	if p != parents[target] {
		factors[target] *= factors[parents[target]]
		parents[target] = p
	}

	return parents[target]
}

// tc: O(mn), m: size of equations, n: size of queries, sc: O(m)
func calcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {
	table := graph(equations, values)
	result := make([]float64, len(queries))

	for i := range queries {
		result[i] = query(table, queries[i])
	}

	return result
}

func graph(equations [][]string, values []float64) map[string]map[string]float64 {
	table := make(map[string]map[string]float64)

	for i, eq := range equations {
		if _, ok := table[eq[0]]; !ok {
			table[eq[0]] = make(map[string]float64)
		}

		if _, ok := table[eq[1]]; !ok {
			table[eq[1]] = make(map[string]float64)
		}

		table[eq[0]][eq[1]] = values[i]
		table[eq[1]][eq[0]] = float64(1) / values[i]

		// implies self / self = 1
		table[eq[0]][eq[0]] = float64(1)
		table[eq[1]][eq[1]] = float64(1)
	}

	return table
}

type EQ struct {
	Val float64
	Str string
}

func query(table map[string]map[string]float64, query []string) float64 {
	if _, ok := table[query[0]]; !ok {
		return float64(-1)
	}

	if _, ok := table[query[1]]; !ok {
		return float64(-1)
	}

	queue := []EQ{{float64(1), query[1]}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		if visited[q.Str] {
			continue
		}

		visited[q.Str] = true

		for str, val := range table[q.Str] {
			if str == query[0] {
				return float64(1) / (val * q.Val)
			}

			queue = append(queue, EQ{
				Val: val * q.Val,
				Str: str,
			})
		}
	}

	return float64(-1)
}

// tc: O(m*n^2), but it could be a wast of computation, since values should be
// evaluated if query exist
func calcEquation1(equations [][]string, values []float64, queries [][]string) []float64 {
	table := buildGraph(equations, values)

	result := make([]float64, len(queries))

	for i, q := range queries {
		if val, ok := table[q[0]][q[1]]; ok {
			result[i] = val
		} else {
			result[i] = float64(-1)
		}
	}

	return result
}

func buildGraph(equations [][]string, values []float64) map[string]map[string]float64 {
	table := make(map[string]map[string]float64)

	for i, eq := range equations {
		if _, ok := table[eq[0]]; !ok {
			table[eq[0]] = make(map[string]float64)
		}

		if _, ok := table[eq[1]]; !ok {
			table[eq[1]] = make(map[string]float64)
		}

		table[eq[0]][eq[1]] = values[i]
		table[eq[1]][eq[0]] = float64(1) / values[i]

		table[eq[0]][eq[0]] = float64(1)
		table[eq[1]][eq[1]] = float64(1)
	}

	for i := range table {
		for j := range table {
			if _, ok := table[i][j]; ok && i != j {
				for k := range table {
					if _, ok := table[j][k]; ok {
						table[i][k] = table[i][j] * table[j][k]
						table[k][i] = float64(1) / table[i][k]
					}
				}
			}
		}
	}

	return table
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

//		author uses union-find to solve the problem, interesting

//		for each new coming equation, use new equations to update root relations

//		when using union-find, divisor is always treated as root, occurrence
//		order doesn't matter, as long as factor relationships is correct
//		e.g. a / b = 2, a = 2 * b, that why treat divisor as root

//		it's a pretty elegant solution

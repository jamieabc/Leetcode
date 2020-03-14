package main

//Given a directed, acyclic graph of N nodes.  Find all possible paths from node 0 to node N-1, and return them in any order.
//
//The graph is given as follows:  the nodes are 0, 1, ..., graph.length - 1.  graph[i] is a list of all nodes j for which the edge (i, j) exists.
//
//Example:
//Input: [[1,2], [3], [3], []]
//Output: [[0,1,3],[0,2,3]]
//Explanation: The graph looks like this:
//0--->1
//|    |
//v    v
//2--->3
//There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
//
//Note:
//
//    The number of nodes in the graph will be in the range [2, 15].
//    You can print different paths in any order, but you should keep the order of nodes inside one path.

func allPathsSourceTarget(graph [][]int) [][]int {
	length := len(graph)
	if length == 0 {
		return [][]int{}
	}

	result := make([][]int, 0)
	paths := []int{0}

	dfs(graph, &result, paths)

	return result
}

func dfs(graph [][]int, result *[][]int, paths []int) {
	length := len(paths)
	last := paths[length-1]
	if len(graph[last]) == 0 {
		if paths[length-1] == len(graph)-1 {
			tmp := append([]int{}, paths...)
			*result = append(*result, tmp)
		} else {
			return
		}
	}

	for _, j := range graph[last] {
		tmp := append([]int{}, paths...)
		dfs(graph, result, append(tmp, j))
	}
}

func bfs(graph [][]int, tracing *[][]int) {
	var changed bool
	tmp := make([][]int, 0)
	length := len(graph)

	for i := range *tracing {
		l := len((*tracing)[i])
		last := (*tracing)[i][l-1]
		if len(graph[last]) == 0 && last == length-1 {
			tmp = append(tmp, (*tracing)[i])
			continue
		}

		changed = true

		for _, j := range graph[last] {
			t := make([]int, l+1)
			copy(t, (*tracing)[i])
			t[l] = j
		}
	}

	*tracing = tmp

	if changed {
		bfs(graph, tracing)
	}
}

// problems
//	1.	no need another 2D array, original array is able to trace routes
// 	2.	use new array to store existing paths, avoid changes by content
//	3.	when using slice, be ware of slice is a descriptor to array, so if
//		slice is not exceeding capacity during append, then underlying array
//		is still the same, this causes bug that change slice in the loop
//		might end up with all data are identical
//	4.	use dfs to search
//	5.	optimize, currently needs to traverse twice, first round down to
//		end, then up, the second round can be reduced

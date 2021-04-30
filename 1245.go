package main

// Given an undirected tree, return its diameter: the number of edges in a longest path in that tree.
//
// The tree is given as an array of edges where edges[i] = [u, v] is a bidirectional edge between nodes u and v.  Each node has labels in the set {0, 1, ..., edges.length}.
//
//
//
// Example 1:
//
// Input: edges = [[0,1],[0,2]]
// Output: 2
// Explanation:
// A longest path of the tree is the path 1 - 0 - 2.
//
// Example 2:
//
// Input: edges = [[0,1],[1,2],[2,3],[1,4],[4,5]]
// Output: 4
// Explanation:
// A longest path of the tree is the path 3 - 2 - 1 - 4 - 5.
//
//
//
// Constraints:
//
// 0 <= edges.length < 10^4
// edges[i][0] != edges[i][1]
// 0 <= edges[i][j] <= edges.length
// The given edges form an undirected tree.

// tc: O(n)
func treeDiameter(edges [][]int) int {
	size := len(edges)
	adjList := make([][]int, size+1)
	for i := range adjList {
		adjList[i] = make([]int, 0)
	}

	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	// start from node 0
	visited := make([]bool, size+1)
	visited[0] = true
	var largest int

	dfs(adjList, visited, 0, &largest)

	return largest
}

func dfs(adjList [][]int, visited []bool, cur int, largest *int) int {
	var longest1, longest2 int

	for _, to := range adjList[cur] {
		if !visited[to] {
			visited[to] = true

			tmp := dfs(adjList, visited, to, largest)

			if tmp > longest1 {
				longest2 = longest1
				longest1 = tmp
			} else if tmp > longest2 {
				longest2 = tmp
			}
		}
	}

	*largest = max(*largest, longest1+longest2)

	return 1 + longest1
}

// tc: O(n)
func treeDiameter2(edges [][]int) int {
	size := len(edges)
	inDegree := make([]int, size+1)

	adjList := make([][]int, size+1)
	for i := range adjList {
		adjList[i] = make([]int, 0)
	}

	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
		inDegree[edge[0]]++
		inDegree[edge[1]]++
	}

	leaves := make([]int, 0)
	count := size + 1
	for i := range inDegree {
		if inDegree[i] == 1 {
			leaves = append(leaves, i)
			inDegree[i] = 0
		}
	}

	var dist int

	// topological sort
	for count > 2 {
		dist++

		length := len(leaves)
		count -= length

		for i := 0; i < length; i++ {
			for _, to := range adjList[leaves[i]] {
				inDegree[to]--
				if inDegree[to] == 1 {
					leaves = append(leaves, to)
				}
			}
		}

		leaves = leaves[length:]
	}

	if count == 1 {
		return dist * 2
	}

	return dist*2 + 1
}

// tc: O(n)
func treeDiameter1(edges [][]int) int {
	adjList := make(map[int][]int)

	for _, edge := range edges {
		if _, ok := adjList[edge[0]]; !ok {
			adjList[edge[0]] = make([]int, 0)
		}

		if _, ok := adjList[edge[1]]; !ok {
			adjList[edge[1]] = make([]int, 0)
		}

		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	var start int
	for n, to := range adjList {
		if len(to) == 1 {
			start = n
			break
		}
	}

	leaf, longest := bfs(adjList, start)

	_, longest = bfs(adjList, leaf)

	return longest - 1
}

func bfs(adjList map[int][]int, start int) (int, int) {
	// start from leaf node
	visited := make(map[int]bool)

	var longest, last int
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		size := len(queue)
		longest++

		for i := 0; i < size; i++ {
			for _, to := range adjList[queue[i]] {
				if !visited[to] {
					visited[to] = true

					queue = append(queue, to)
					last = to
				}
			}
		}

		queue = queue[size:]
	}

	return last, longest
}

//	Notes
//	1.	longest path can do two times of BFS to find longest

//	2.	inspired from solution, borrow MHT concept can find the diameter

//		there only exists 2 conditions of MHT, which has different calculation

//		also, strip from outer, cna use topological sort technique

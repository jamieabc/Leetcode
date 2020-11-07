package main

// A tree is an undirected graph in which any two vertices are connected by exactly one path. In other words, any connected graph without simple cycles is a tree.
//
// Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where edges[i] = [ai, bi] indicates that there is an undirected edge between the two nodes ai and bi in the tree, you can choose any node of the tree as the root. When you select a node x as the root, the result tree has height h. Among all possible rooted trees, those with minimum height (i.e. min(h))  are called minimum height trees (MHTs).
//
// Return a list of all MHTs' root labels. You can return the answer in any order.
//
// The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.
//
//
//
// Example 1:
//
//
// Input: n = 4, edges = [[1,0],[1,2],[1,3]]
// Output: [1]
// Explanation: As shown, the height of the tree is 1 when the root is the node with label 1 which is the only MHT.
// Example 2:
//
//
// Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
// Output: [3,4]
// Example 3:
//
// Input: n = 1, edges = []
// Output: [0]
// Example 4:
//
// Input: n = 2, edges = [[0,1]]
// Output: [0,1]
//
//
// Constraints:
//
// 1 <= n <= 2 * 104
// edges.length == n - 1
// 0 <= ai, bi < n
// ai != bi
// All the pairs (ai, bi) are distinct.
// The given input is guaranteed to be a tree and there will be no repeated edges.

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make(map[int][]int)
	for _, edge := range edges {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = make([]int, 0)
		}

		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = make([]int, 0)
		}

		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	longest1 := make([]int, 0)
	visited1 := make([]bool, n)
	dfs(0, graph, []int{0}, &longest1, visited1)

	longest2 := make([]int, 0)
	visited2 := make([]bool, n)
	dfs(longest1[len(longest1)-1], graph, []int{longest1[len(longest1)-1]}, &longest2, visited2)

	return longest2[(len(longest2)-1)/2 : len(longest2)/2+1]
}

func dfs(node int, graph map[int][]int, current []int, longest *[]int, visited []bool) {
	if visited[node] {
		if len(current) > len(*longest) {
			*longest = (*longest)[:0]
			*longest = append(*longest, current[:len(current)-1]...)
		}
		return
	}
	visited[node] = true

	for _, to := range graph[node] {
		tmp := append([]int{}, current...)
		tmp = append(tmp, to)
		dfs(to, graph, tmp, longest, visited)
	}
}

// tc: O(n)
func findMinHeightTrees2(n int, edges [][]int) []int {
	if n <= 1 {
		ans := make([]int, 0)
		for i := 0; i < n; i++ {
			ans = append(ans, i)
		}
		return ans
	}

	graph := make(map[int][]int)
	inDegree := make([]int, n)

	for _, edge := range edges {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = make([]int, 0)
		}

		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = make([]int, 0)
		}

		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		inDegree[edge[0]]++
		inDegree[edge[1]]++
	}

	queue := make([]int, 0)

	// find leaf nodes
	for i, degree := range inDegree {
		if degree == 1 {
			queue = append(queue, i)
		}
	}

	for n > 2 {
		size := len(queue)
		n -= size

		for i := 0; i < size; i++ {
			inDegree[queue[i]]--

			for _, to := range graph[queue[i]] {
				inDegree[to]--

				// next round leaf node
				if inDegree[to] == 1 {
					queue = append(queue, to)
				}
			}
		}

		queue = queue[size:]
	}

	return queue
}

// tc: O(nh), n: # nodes, h: height of tree
func findMinHeightTrees1(n int, edges [][]int) []int {
	graph := buildGraph(edges)

	visited := make([]bool, n)
	stack := make([]int, 0)

	// removing leaf until remaining nodes <= 2
	for n > 2 {
		stack = stack[:0]

		// find not visited leaf nodes
		// this is a waste of computation, no need to iterate all nodes to find
		// leaf, should come from previous iteration leaf nodes
		for from, to := range graph {
			if !visited[from] && len(to) == 1 {
				stack = append(stack, from)
			}
		}

		// remove leaf and connected edges
		for i := 0; i < len(stack); i++ {
			for n := range graph[stack[i]] {
				delete(graph[n], stack[i])
			}
			delete(graph, stack[i])
			visited[stack[i]] = true
			n--
		}
	}

	ans := make([]int, 0)
	for i := range visited {
		if !visited[i] {
			ans = append(ans, i)
		}
	}

	return ans
}

func buildGraph(edges [][]int) map[int]map[int]bool {
	table := make(map[int]map[int]bool)

	for _, edge := range edges {
		if _, ok := table[edge[0]]; !ok {
			table[edge[0]] = make(map[int]bool)
		}

		if _, ok := table[edge[1]]; !ok {
			table[edge[1]] = make(map[int]bool)
		}

		table[edge[0]][edge[1]] = true
		table[edge[1]][edge[0]] = true
	}

	return table
}

//	Notes
//	1.	there could exist at most 2 MHT

//		e.g.           1
//					/  \   \
//                 2   3    4
//                      \    \
//                      5    6

//		h(1) = 2
//		h(3) = 3
//		h(4) = 3

//		the reason height of a tree can be reduced is because moving a node from
//		longest path to other paths not longest to reduce height

//		but this only works when there's one longest path, or all longest paths
//		share same common node(s)

//		consider this in another way, treat every node with same distant to leaf
//		as redundant nodes, after each deletion, there could be 3 conditions:

//		1         1         1        1
//       \                /   \  =>
//        2             2      3

//		there can exist at most 2 MHT because finally, only left and middle
//		conditions holds

//	2.	to find shortest distance, use BFS

//		if there are more than 2 nodes with same distance, there must be a parent
//		to all those nodes

//	3.	if there are more than 3 nodes, need to check again

//		   0
//       /   \
//     1      2
//          /   \
//        3      et4
//             /
//           5

//		both nodes 0, 2, and 4 has distance 1 to leaf, but 2 is MHT because
//		2 can reach both 0 & 4

//	4.	a node can only be marked visited if all connected are processed

//	5.	a node can be marked as visited when round ends, since every round
//		shares same cost, node should be marked visited when this round ends to
//		avoid order dependent situation

//	6.	inspired from solution, each iteration find leaf nodes: out-degree -
//		in-degree = 1, which is topological sort

//		I didn't think of this relates to topological sort

//	7.	from solution, distance of two nodes is the number of edges connect two
//		nodes

//	8.	inspired from https://leetcode.com/problems/minimum-height-trees/discuss/76055/Share-some-thoughts

//		author provides a great way of think process: simplify problem into
//		one way tree (each node w/ only 1 branch).

//		if n is unknown, and no random access to nodes, only way to find MHT
//		is traversing from two sides, when they meet, it's node with MHT

//	9.	inspired from https://leetcode.com/problems/minimum-height-trees/discuss/76052/Two-O(n)-solutions

//		MHT is middle point(s) on overall tree longest path

//		to find overall longest path, randomly select a node x and find longest
//		path ends ay y. this y must be on tree overall longest path (if not,
//		then longest path should be updated, which is not longest anymore)

//		start from y, find longest path ends at z, middle points of y - z is
//		MHT

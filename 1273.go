package main

// A tree rooted at node 0 is given as follows:
//
//     The number of nodes is nodes;
//     The value of the i-th node is value[i];
//     The parent of the i-th node is parent[i].
//
// Remove every subtree whose sum of values of nodes is zero.
//
// After doing so, return the number of nodes remaining in the tree.
//
//
//
// Example 1:
//
// Input: nodes = 7, parent = [-1,0,0,1,2,2,2], value = [1,-2,4,0,-2,-1,-1]
// Output: 2
//
//
//
// Constraints:
//
//     1 <= nodes <= 10^4
//     -10^5 <= value[i] <= 10^5
//     parent.length == nodes
//     parent[0] == -1 which indicates that 0 is the root.

func deleteTreeNodes(n int, parent []int, value []int) int {
	children := make([][]int, n)

	// construct tree
	var idx int
	for i := range parent {
		if parent[i] == -1 {
			idx = i
		} else {
			if len(children[parent[i]]) == 0 {
				children[parent[i]] = []int{i}
			} else {
				children[parent[i]] = append(children[parent[i]], i)
			}
		}
	}

	// traverse tree
	count := n
	traverse(children, idx, value, &count)
	return count
}

// return sum of subtree, # of node in subtree
func traverse(children [][]int, node int, value []int, count *int) (int, int) {
	sum, sub := value[node], 1
	var tmp1, tmp2 int

	for _, c := range children[node] {
		tmp1, tmp2 = traverse(children, c, value, count)
		sum += tmp1
		sub += tmp2
	}

	if sum == 0 {
		*count -= sub
		sub = 0
	}

	return sum, sub
}

//	problems
//	1.	inspired from https://leetcode.com/problems/delete-tree-nodes/discuss/440829/JavaC%2B%2BPython-One-pass

//		there's no need to have whole tree constructed, what I need is
//		parent -> children relationships, so another array is enough

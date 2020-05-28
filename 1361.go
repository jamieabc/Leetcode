package main

// You have n binary tree nodes numbered from 0 to n - 1 where node i has two children leftChild[i] and rightChild[i], return true if and only if all the given nodes form exactly one valid binary tree.
//
// If node i has no left child then leftChild[i] will equal -1, similarly for the right child.
//
// Note that the nodes have no values and that we only use the node numbers in this problem.
//
//
//
// Example 1:
//
// Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
// Output: true
//
// Example 2:
//
// Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]
// Output: false
//
// Example 3:
//
// Input: n = 2, leftChild = [1,0], rightChild = [-1,-1]
// Output: false
//
// Example 4:
//
// Input: n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1]
// Output: false
//
//
//
// Constraints:
//
//     1 <= n <= 10^4
//     leftChild.length == rightChild.length == n
//     -1 <= leftChild[i], rightChild[i] <= n - 1
//
//

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	nodes := make([]bool, n)

	for _, c := range leftChild {
		if c == -1 {
			continue
		}

		if !nodes[c] {
			nodes[c] = true
		} else {
			return false
		}
	}

	for _, c := range rightChild {
		if c == -1 {
			continue
		}

		if !nodes[c] {
			nodes[c] = true
		} else {
			return false
		}
	}

	// root will not appear in any child
	root := -1
	for i := range nodes {
		if !nodes[i] {
			if root == -1 {
				root = i
				continue
			} else {
				return false
			}
		}
	}

	if root == -1 || (n > 1 && leftChild[root] == -1 && rightChild[root] == -1) {
		return false
	}

	return true
}

//	problems
//	1.	when checking parent <-> child loop, cannot assume parent's parent
//		always exist (root), but other situation shoule be false

//	2.	no any child's parent should be root

//	3.	root cannot be child

//	4.	it has no order, so root may not necessary be 0

//	5.	if root has 2 children, then there could 2 checks from child, so
//		when root != -1, still need to check if root == parent

//	6.	dont' forget to skip -1 node

//	7.	too slow, I think it's because creating map

//	8.	every node can appear in either left or right only once, the problem
//		is circular route, how to detect that?

//	9.	when validating use every node appears at least one, there could
//		exist one possibility that one node has no left or child, and other
//		nodes form a circular route.

//		e.g. n = 4
//			 left = [1, 2, 0, -1]
//			 right = [-1, -1, -1, -1]

//	10.	another problem is that when using left & right child of root is
//		non-nil, the boundary condition one node only(single root)

//	11.	too slow, since map is used as a set, use array instead
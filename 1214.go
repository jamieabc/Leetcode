package main

// Given two binary search trees, return True if and only if there is a node in the first tree and a node in the second tree whose values sum up to a given integer target.
//
//
//
// Example 1:
//
// Input: root1 = [2,1,4], root2 = [1,0,3], target = 5
// Output: true
// Explanation: 2 and 3 sum up to 5.
//
// Example 2:
//
// Input: root1 = [0,-10,10], root2 = [5,1,7,0,2], target = 18
// Output: false
//
//
//
// Constraints:
//
//     Each tree has at most 5000 nodes.
//     -10^9 <= target, node.val <= 10^9

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	if root1 == nil || root2 == nil {
		return false
	}

	// traverse like sorted array, start from smallest + largest
	n1, n2 := root1, root2
	stack1, stack2 := make([]*TreeNode, 0), make([]*TreeNode, 0)

	for (len(stack1) != 0 || n1 != nil) && (len(stack2) != 0 || n2 != nil) {
		for n1 != nil {
			stack1 = append(stack1, n1)
			n1 = n1.Left
		}

		n1 = stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]

		for n2 != nil {
			stack2 = append(stack2, n2)
			n2 = n2.Right
		}

		n2 = stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]

		sum := n1.Val + n2.Val

		if sum == target {
			return true
		} else if sum > target {
			stack1 = append(stack1, n1)
			n1 = nil
			n2 = n2.Left
		} else {
			stack2 = append(stack2, n2)
			n2 = nil
			n1 = n1.Right
		}
	}

	return false
}

func twoSumBSTs2(root1 *TreeNode, root2 *TreeNode, target int) bool {
	if root1 == nil || root2 == nil {
		return false
	}

	stack1, stack2 := make([]*TreeNode, 0), make([]*TreeNode, 0)

	for s1, s2 := root1, root2; true; {
		// traverse left-most
		for ; s1 != nil; s1 = s1.Left {
			stack1 = append(stack1, s1)
		}

		// traverse right-mode
		for ; s2 != nil; s2 = s2.Right {
			stack2 = append(stack2, s2)
		}

		if len(stack1) == 0 || len(stack2) == 0 {
			return false
		}

		t1, t2 := stack1[len(stack1)-1], stack2[len(stack2)-1]
		sum := t1.Val + t2.Val

		if sum == target {
			return true
		}

		if sum < target {
			// traverse root1
			stack1 = stack1[:len(stack1)-1]
			s1 = t1.Right
		} else {
			// traverse root2
			stack2 = stack2[:len(stack2)-1]
			s2 = t2.Left
		}
	}

	return false
}

func twoSumBSTs1(root1 *TreeNode, root2 *TreeNode, target int) bool {
	mapping := make(map[int]bool)
	traverse(root2, &mapping)

	return inOrder(root1, root2, target, mapping)
}

func inOrder(node, root2 *TreeNode, target int, mapping map[int]bool) bool {
	if node == nil {
		return false
	}

	if _, ok := mapping[target-node.Val]; ok {
		return true
	}

	return inOrder(node.Left, root2, target, mapping) || inOrder(node.Right, root2, target, mapping)
}

func traverse(node *TreeNode, mapping *map[int]bool) {
	if node == nil {
		return
	}

	(*mapping)[node.Val] = true

	traverse(node.Right, mapping)
	traverse(node.Left, mapping)
}

//	problems
//	1.	time complexity is (n log m), it can use additional map to speed-up

//	2. 	when using iterative, it's time limit exceed, so I think there's too
//		many memory operation on stack, the fail test case for root1 probably
//		w/ 1000+ nodes, and for root2 only 1

//	3.	add reference from https://leetcode.com/problems/two-sum-bsts/discuss/397624/Simple-Stack-Solution

//		author uses strategy of finding largest + smallest in two trees

//		it's a little weird that need additional intermediate variables
//		t1 & t2 in order to avoid memory problem

//	4.	add reference https://leetcode.com/problems/two-sum-bsts/discuss/397796/Searching-in-a-BST-is-not-O(-log(n)-)

//		for balance BST, tc is O(n log n), but if it's un-balanced,
//		tc might be O(h) and worst case is O(n) means all items line up in
//		a row, so worst case scenario tc is O(mn)

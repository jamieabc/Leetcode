package main

// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
//
//
//
// But the following [1,2,2,null,3,null,3] is not:
//
//     1
//    / \
//   2   2
//    \   \
//    3    3

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack1, stack2 := []*TreeNode{root.Left}, []*TreeNode{root.Right}

	for len(stack1) > 0 && len(stack2) > 0 {
		n1, n2 := stack1[len(stack1)-1], stack2[len(stack2)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = stack2[:len(stack2)-1]

		if n1 == nil && n2 == nil {
			continue
		}

		if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) || n1.Val != n2.Val {
			return false
		}

		stack1 = append(stack1, n1.Left, n1.Right)
		stack2 = append(stack2, n2.Right, n2.Left)
	}

	return len(stack1) == 0 && len(stack2) == 0
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recursive(root.Left, root.Right)
}

func recursive(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		return false
	}

	if n1.Val != n2.Val {
		return false
	}

	return recursive(n1.Left, n2.Right) && recursive(n1.Right, n2.Left)
}

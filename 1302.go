package main

//Given a binary tree, return the sum of values of its deepest leaves.
//
//
//
//Example 1:
//
//Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
//Output: 15
//
//
//
//Constraints:
//
//    The number of nodes in the tree is between 1 and 10^4.
//    The value of nodes is between 1 and 100.

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var levelSum int
	var added bool
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		added = false

		for i := 0; i < size; i++ {
			n := queue[i]

			if n.Left != nil {
				added = true
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				added = true
				queue = append(queue, n.Right)
			}
		}

		if !added {
			break
		} else {
			queue = queue[size:]
		}
	}

	for _, q := range queue {
		levelSum += q.Val
	}

	return levelSum
}

func deepestLeavesSum2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]Info, 0)
	cur := Info{
		Node:  root,
		Level: 0,
	}
	var maxLevel, maxLevelSum int

	for cur.Node != nil || len(stack) > 0 {
		for cur.Node != nil {
			stack = append(stack, cur)
			cur = Info{
				Node:  cur.Node.Left,
				Level: cur.Level + 1,
			}
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if n.Node.Left == nil && n.Node.Right == nil {
			if n.Level > maxLevel {
				maxLevel = n.Level
				maxLevelSum = n.Node.Val
			} else if n.Level == maxLevel {
				maxLevelSum += n.Node.Val
			}
		}

		cur = Info{
			Node:  n.Node.Right,
			Level: n.Level + 1,
		}
	}

	return maxLevelSum
}

func deepestLeavesSum1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_, sum := dfs(root, 0)

	return sum
}

func dfs(node *TreeNode, level int) (int, int) {
	if node.Left == nil && node.Right == nil {
		return level, node.Val
	}

	var lLevel, lSum, rLevel, rSum int
	if node.Left != nil {
		lLevel, lSum = dfs(node.Left, level+1)
	}

	if node.Right != nil {
		rLevel, rSum = dfs(node.Right, level+1)
	}

	if lLevel == rLevel {
		return lLevel, lSum + rSum
	} else if lLevel > rLevel {
		return lLevel, lSum
	}
	return rLevel, rSum
}

//	problems
//	1.	for iterative BFS, adding same level number is not optimal
//		solution, can just check if new child is added

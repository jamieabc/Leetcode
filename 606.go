package main

import (
	"fmt"
	"strconv"
)

//You need to construct a string consists of parenthesis and integers from a binary tree with the preorder traversing way.
//
//The null node needs to be represented by empty parenthesis pair "()". And you need to omit all the empty parenthesis pairs that don't affect the one-to-one mapping relationship between the string and the original binary tree.
//
//Example 1:
//
//Input: Binary tree: [1,2,3,4]
//       1
//     /   \
//    2     3
//   /
//  4
//
//Output: "1(2(4))(3)"
//
//Explanation: Originallay it needs to be "1(2(4)())(3()())",
//but you need to omit all the unnecessary empty parenthesis pairs.
//And it will be "1(2(4))(3)".
//
//Example 2:
//
//Input: Binary tree: [1,2,3,null,4]
//       1
//     /   \
//    2     3
//     \
//      4
//
//Output: "1(2()(4))(3)"
//
//Explanation: Almost the same as the first example,
//except we can't omit the first parenthesis pair to break the one-to-one mapping relationship between the input and the output.

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	stack := []*TreeNode{t}
	result := make([]byte, 0)
	visited := make(map[*TreeNode]bool)

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		if visited[n] {
			result = append(result, ')')
			stack = stack[:len(stack)-1]
			continue
		}

		visited[n] = true
		result = append(result, '(')
		result = append(result, []byte(strconv.Itoa(n.Val))...)

		if n.Right != nil {
			stack = append(stack, n.Right)
		}

		if n.Left != nil {
			stack = append(stack, n.Left)
		} else {
			if n.Right != nil {
				result = append(result, '(', ')')
			}
		}
	}

	// because root rule is different from child rule, remove first &
	// last ()
	return string(result[1 : len(result)-1])
}

func tree2str3(t *TreeNode) string {
	if t == nil {
		return ""
	} else if t.Left == nil && t.Right == nil {
		return fmt.Sprintf("%d", t.Val)
	} else if t.Left == nil && t.Right != nil {
		return fmt.Sprintf("%d()(%s)", t.Val, tree2str(t.Right))
	} else if t.Left != nil && t.Right == nil {
		return fmt.Sprintf("%d(%s)", t.Val, tree2str(t.Left))
	}
	return fmt.Sprintf("%d(%s)(%s)", t.Val, tree2str(t.Left), tree2str(t.Right))
}

func tree2str2(t *TreeNode) string {
	if t == nil {
		return ""
	} else if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	} else if t.Left == nil && t.Right != nil {
		return strconv.Itoa(t.Val) + "()" + "(" + tree2str(t.Right) + ")"
	} else if t.Left != nil && t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	}
	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")" + "(" + tree2str(t.Right) + ")"
}

func tree2str1(t *TreeNode) string {
	if t == nil {
		return ""
	}

	result := make([]byte, 0)
	preOrder(t, 0, &result)

	return string(result)
}

func preOrder(node *TreeNode, level int, result *[]byte) {
	if node == nil {
		return
	}

	if level > 0 {
		*result = append(*result, '(')
	}

	*result = append(*result, []byte(strconv.Itoa(node.Val))...)

	if node.Left != nil {
		preOrder(node.Left, level+1, result)
	} else {
		if node.Right != nil {
			*result = append(*result, '(', ')')
		}
	}

	if node.Right != nil {
		preOrder(node.Right, level+1, result)
	}

	if level > 0 {
		*result = append(*result, ')')
	}
}

//	problems
//	1.	inspired from sample code, use fmt.Sprintf instead of string concatenation

//	2.	inspired from solution, iterative is a very beautiful solution,
//		uses stack to do pre-order traversal.

//		then the problem becomes: how to know # of closing parenthesis
//		to add? A very clever way is to use hash map to store visited
//		node, and node keeps in stack even left & right child is already
//		added to stack. The reason of doing this is to keep order of
//		traversal, when a node popped from stack and is already visited,
//		then it's time to add )

//		additional caution here is that, root node rule is a little
//		different from child rule

//		child rule: (val ...
//		root rule: val(

//		to solve this difference, just tree root as normal child, then
//		rip off first & last ()

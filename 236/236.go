package main

// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”



// Example 1:

// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// Output: 3
// Explanation: The LCA of nodes 5 and 1 is 3.

// Example 2:

// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// Output: 5
// Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.

// Example 3:

// Input: root = [1,2], p = 1, q = 2
// Output: 1



// Constraints:

//     The number of nodes in the tree is in the range [2, 105].
//     -109 <= Node.val <= 109
//     All Node.val are unique.
//     p != q
//     p and q will exist in the tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// self matches, pre-order means all previous visited nodes has no match, so self is the answer
	if root == nil || root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l == nil {
		return r
	}
	return l
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
     common, _ := dfs1(root, p, q)

     if common == nil {
         return root
     }
     return common
}

func dfs1(cur, p, q *TreeNode) (*TreeNode, int) {
    if cur == nil {
        return nil, 0
    }

    l, lCount := dfs1(cur.Left, p, q)
    if lCount == 2 {
        return l, 2
    }

    r, rCount := dfs1(cur.Right, p, q)
    if rCount == 2 {
        return r, 2
    }

    if lCount + rCount == 2 {
        return cur, 2
    }

    if cur == p || cur == q {
        if lCount == 1 || rCount == 1 {
            return cur, 2
        }
        return nil, 1
    }

    return nil, lCount + rCount
}

//	Notes
//	1.	takes a while to think of solution, this is recursion: self cannot decide until some further
//		information is provided, which is post-order traversal (LRN)

//	2.	inspired from sample code, author uses pre-order traversal (NLR), as long as self matches p or q,
//		it means the other match is in children (pre-order means all previous visited nodes doesn't exist
//		any matching), very smart...

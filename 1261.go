package main

// Given a binary tree with the following rules:
//
//     root.val == 0
//     If treeNode.val == x and treeNode.left != null, then treeNode.left.val == 2 * x + 1
//     If treeNode.val == x and treeNode.right != null, then treeNode.right.val == 2 * x + 2
//
// Now the binary tree is contaminated, which means all treeNode.val have been changed to -1.
//
// You need to first recover the binary tree and then implement the FindElements class:
//
//     FindElements(TreeNode* root) Initializes the object with a contamined binary tree, you need to recover it first.
//     bool find(int target) Return if the target value exists in the recovered binary tree.
//
//
//
// Example 1:
//
// Input
// ["FindElements","find","find"]
// [[[-1,null,-1]],[1],[2]]
// Output
// [null,false,true]
// Explanation
// FindElements findElements = new FindElements([-1,null,-1]);
// findElements.find(1); // return False
// findElements.find(2); // return True
//
// Example 2:
//
// Input
// ["FindElements","find","find","find"]
// [[[-1,-1,-1,-1,-1]],[1],[3],[5]]
// Output
// [null,true,true,false]
// Explanation
// FindElements findElements = new FindElements([-1,-1,-1,-1,-1]);
// findElements.find(1); // return True
// findElements.find(3); // return True
// findElements.find(5); // return False
//
// Example 3:
//
// Input
// ["FindElements","find","find","find","find"]
// [[[-1,null,-1,-1,null,-1]],[2],[3],[4],[5]]
// Output
// [null,true,false,false,true]
// Explanation
// FindElements findElements = new FindElements([-1,null,-1,-1,null,-1]);
// findElements.find(2); // return True
// findElements.find(3); // return False
// findElements.find(4); // return False
// findElements.find(5); // return True
//
//
//
// Constraints:
//
//     TreeNode.val == -1
//     The height of the binary tree is less than or equal to 20
//     The total number of nodes is between [1, 10^4]
//     Total calls of find() is between [1, 10^4]
//     0 <= target <= 10^6

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
	mapping map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	mapping := make(map[int]bool)
	traverse(root, mapping, 0)

	return FindElements{
		mapping: mapping,
	}
}

func traverse(node *TreeNode, mapping map[int]bool, val int) {
	if node == nil {
		return
	}

	node.Val = val
	mapping[val] = true

	traverse(node.Left, mapping, 2*node.Val+1)
	traverse(node.Right, mapping, 2*node.Val+2)
}

func (this *FindElements) Find(target int) bool {
	return this.mapping[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */

//	problems
//	1.	too slow, some search path is not needed

//	2.	use its rule, any number can be decomposed. If it's odd, left node,
//		if it's even, right node. And its parent can be decided by it's
//		position, e.g. if it's value is 51, since it's odd, it's left node
//		from parent, and its parent is (51-1)/2=25

//		keep iterate this, then a path from root is found, try to traverse
//		from root, it all nodes exists, then the value exists

//	3.	the problem needs to recover tree, so I still need to traverse whole
//		tree, then a map is needed

//	4.	reference from https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/discuss/431107/JavaPython-3-DFS-clean-code-w-analysis.

//		pass next value into recursion, to remove redundant if checking

//	5.	reference from https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/discuss/431229/Python-Special-Way-for-find()-without-HashSet-O(1)-Space-O(logn)-Time

//		The rule is actually bit-wise operation, *2 means left shift 1 bit,
//		left side plus 1 means last bit is 1, right side plus 2 means last
//		bit is 0

//		With this property, I can find out it's left or right child of parent
//		I have noticed that no need to recover whole tree and find path to a
//		node in reverse order:

//		paths := make([]bool, 0)    // left: true, right: false
//		for target != 0 {
//			if target % 2 == 0 {
//				paths = append(paths, false)
//				target = (target-2)/2
//			} else {
//				paths = append(paths, true)
//				target = (target-1)/2
//			}
//		}

//		My way can still find path, but it needs to traverse whole to find
//		left/right child of root. which means complexity is O(log n).

//		Above reference is much better and insightful because it can decide
//		left/right of parent by it's current value.

//		Although this method is fast, but problem description says tree
//		needs to be recovered, so this method is more suitable when tree
//		doesn't need to be reconstructed.

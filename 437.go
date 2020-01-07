package main

//You are given a binary tree in which each node contains an integer value.
//
//Find the number of paths that sum to a given value.
//
//The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).
//
//The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.
//
//Example:
//
//root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//
//      10
//     /  \
//    5   -3
//   / \    \
//  3   2   11
// / \   \
//3  -2   1
//
//Return 3. The paths that sum to 8 are:
//
//1.  5 -> 3
//2.  5 -> 2 -> 1
//3. -3 -> 11

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	count := 0
	cache := make(map[int]int)
	cache[0] = 1
	traverse(root, cache, sum, 0, &count)

	return count
}

func traverse(node *TreeNode, cache map[int]int, sum, sumToCurrent int, count *int) {
	// for every node, it has possibilities of choose and not choose
	if node == nil {
		return
	}

	sumToCurrent += node.Val
	diff := sumToCurrent - sum

	if val, ok := cache[diff]; ok {
		*count += val
	}

	cache[sumToCurrent]++
	traverse(node.Left, cache, sum, sumToCurrent, count)
	traverse(node.Right, cache, sum, sumToCurrent, count)
	cache[sumToCurrent]--
}

// problems
// 1. when calculating combinations of left & right, remain should deduct self value
// 2. as long as sum matches, stop, doesn't need to further do calculation
// 3. wrong understanding, it can go one way, but not both, e.g. left-mid-right is now allowed
// 4. forget that any node matches target
// 5. try to reduce temporarily memory
// 6. use map to reduce memory & operation time
// 7. map will de-duplicate, but i should still count those subtree with same sum
// 8. fix how times calculated
// 9. forget the situation that initial count is 0
// 10. from program, always add value to map value is correct
// 11. wrong understanding of golang
// 12. the way adding new number, means every path from root to leaf is used, it shouldn't be replace
// 13. forget the case that when right child is zero, map is nil
// 14. not think clear of array range, it's not map range
// 15. forget the situation when already sum to target, skip that element
// 16. cannot use map, because differ subtree sum will be viewed as same
// 17. wrong parameter order
// 18. wrong logic, sum to current should be added first
// 19. remove initial condition of 0
// 20. diff is sum - sum to current
// 21. wrong cache key value, it should keep from top to current node sum
// 22. wrong key of cache, it should be sumToCurrent - sum
// 23. when exactly match, the cache needs an entry of 0
// 24. if list is empty, result is 0 because there's not such path

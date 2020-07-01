package main

// In a row of trees, the i-th tree produces fruit with type tree[i].
//
// You start at any tree of your choice, then repeatedly perform the following steps:
//
//     Add one piece of fruit from this tree to your baskets.  If you cannot, stop.
//     Move to the next tree to the right of the current tree.  If there is no tree to the right, stop.
//
// Note that you do not have any choice after the initial choice of starting tree: you must perform step 1, then step 2, then back to step 1, then step 2, and so on until you stop.
//
// You have two baskets, and each basket can carry any quantity of fruit, but you want each basket to only carry one type of fruit each.
//
// What is the total amount of fruit you can collect with this procedure?
//
//
//
// Example 1:
//
// Input: [1,2,1]
// Output: 3
// Explanation: We can collect [1,2,1].
//
// Example 2:
//
// Input: [0,1,2,2]
// Output: 3
// Explanation: We can collect [1,2,2].
// If we started at the first tree, we would only collect [0, 1].
//
// Example 3:
//
// Input: [1,2,3,2,2]
// Output: 4
// Explanation: We can collect [2,3,2,2].
// If we started at the first tree, we would only collect [1, 2].
//
// Example 4:
//
// Input: [3,3,3,1,2,1,1,2,3,3,4]
// Output: 5
// Explanation: We can collect [1,2,1,1,2].
// If we started at the first tree or the eighth tree, we would only collect 4 fruits.
//
//
//
// Note:
//
//     1 <= tree.length <= 40000
//     0 <= tree[i] < tree.length

func totalFruit(tree []int) int {
	var count, maxCollections int
	first, second := 0, -1

	for i := range tree {
		if tree[i] == tree[first] {
			count++
			first = i
		} else if second == -1 || tree[i] == tree[second] {
			count++
			second = i
		} else {
			maxCollections = max(maxCollections, count)
			count = abs(first-second) + 1
			first, second = i-1, i
		}
	}

	// for last two fruits
	maxCollections = max(maxCollections, count)

	return maxCollections
}

func totalFruit2(tree []int) int {
	size := len(tree)
	if size == 0 {
		return 0
	}

	var left, right, theOtherFruit, maxCollections int

	// find first index with two types of fruits
	for right = 1; right < size; right++ {
		if tree[right] != tree[left] {
			theOtherFruit = tree[right]
			maxCollections = right + 1
			right++
			break
		}
	}

	// already reach end of tree, all fruits are collectable
	if right == size {
		return size
	}

	for ; right < size; right++ {
		if tree[right] != tree[left] && tree[right] != theOtherFruit {
			maxCollections = max(maxCollections, right-left)

			// find continuous start of left tree fruit
			left = right - 1
			for tmp := right - 2; tmp >= 0; tmp-- {
				if tree[tmp] != tree[right-1] {
					left = tmp + 1
					break
				}
			}

			// update attributes
			theOtherFruit = tree[right]
		}
	}

	maxCollections = max(maxCollections, size-left)

	return maxCollections
}

func totalFruit1(tree []int) int {
	deque := make([]int, 0)
	var maxCollections int

	for i, fruit := range tree {
		if len(deque) == 0 || (len(deque) == 1 && fruit != tree[deque[0]]) {
			deque = append(deque, i)
			continue
		}

		if tree[deque[0]] != fruit && tree[deque[1]] != fruit {
			// calculate max substring length
			maxCollections = max(maxCollections, i-deque[0])
			deque[0] = i - 1
			deque[1] = i

			// update fruit start index, make sure they are continuous same
			// for the case of 1, 2, 1, 2, 3
			for j := i - 2; j >= 0; j-- {
				if tree[j] != tree[i-1] {
					deque[0] = j + 1
					break
				}
			}
		}
	}

	// process remaining fruits
	if len(deque) > 0 {
		maxCollections = max(maxCollections, len(tree)-deque[0])
	}

	return maxCollections
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	when one type of fruit is removed, the only possible solution is fruit
//		on left

//	2.	this is actually a sliding window problem, what I use deque of size 2
//		is same as using two pointers.

//	3.	inspired from https://leetcode.com/problems/fruit-into-baskets/discuss/171954/Java-Very-simple-solution-few-lines-Time-O(n)-Space-O(1)

//		author provides a smarter way to decide existing count when new fruit
//		added. it mains two pointer first & second which are latest index of
//		fruit 1 and fruit 2. and i is the tree current encountered.

//		when fruit 3 comes, count of i-1 is abs(first - second), and first will
//		be updated to i-1, second is updated to i

//		it's really brilliant solution, compares to my solution which needs to
//		traverse back is really slow

//	4.	add reference https://leetcode.com/problems/fruit-into-baskets/discuss/170745/Problem%3A-Longest-Subarray-With-2-Elements

//		lee also posts an interesting solution, didn't take time to understand

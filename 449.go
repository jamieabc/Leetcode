package main

import (
	"strconv"
	"strings"
)

// Serialization is converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.
//
// Design an algorithm to serialize and deserialize a binary search tree. There is no restriction on how your serialization/deserialization algorithm should work. You need to ensure that a binary search tree can be serialized to a string, and this string can be deserialized to the original tree structure.
//
// The encoded string should be as compact as possible.
//
//
//
// Example 1:
//
// Input: root = [2,1,3]
// Output: [2,1,3]
//
// Example 2:
//
// Input: root = []
// Output: []
//
//
//
// Constraints:
//
//     The number of nodes in the tree is in the range [0, 104].
//     0 <= Node.val <= 104
//     The input tree is guaranteed to be a binary search tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	Data []byte
}

func Constructor() Codec {
	return Codec{
		Data: make([]byte, 0),
	}
}

// 4 bytes for each number, each number range from 0 ~ 10^4
func (this *Codec) Add(num int) {
	var b byte
	for i := 4; i > 0; i-- {
		b = byte(0)
		for j := 0; j < 4; j++ {
			if num&(1<<(4*i-1-j)) > 0 {
				b |= 1 << (3 - j)
			}
		}
		this.Data = append(this.Data, b)
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	postOrder(root, this)
	return string(this.Data)
}

func postOrder(node *TreeNode, codec *Codec) {
	if node == nil {
		return
	}

	postOrder(node.Left, codec)
	postOrder(node.Right, codec)

	codec.Add(node.Val)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	// convert string back to post order sequence
	nums := make([]int, len(data)/4)
	var num int
	for i := 0; 4*i < len(data); i++ {
		num = 0

		for j := 0; j < 4 && 4*i+j < len(data); j++ {
			num |= int(data[4*i+j]) << (4 * (3 - j))
		}

		nums[i] = num
	}

	return traverse(nums, 0, len(nums)-1)
}

func traverse(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	node := &TreeNode{
		Val: nums[end],
	}

	var boundary int

	for boundary = start; nums[boundary] < nums[end]; boundary++ {
	}

	node.Left = traverse(nums, start, boundary-1)
	node.Right = traverse(nums, boundary, end-1)

	return node
}

type Codec1 struct {
	Bytes     []byte
	Delimiter byte
}

func Constructor() Codec1 {
	return Codec1{
		Bytes:     make([]byte, 0),
		Delimiter: ',',
	}
}

func (this *Codec11) Add(str string) {
	this.Bytes = append(this.Bytes, this.Delimiter)
	this.Bytes = append(this.Bytes, []byte(str)...)
}

// Serializes a tree to a single string.
func (this *Codec1) serialize(root *TreeNode) string {
	postOrder(root, this)
	return string(this.Bytes[1:])
}

func postOrder(node *TreeNode, codec *Codec1) {
	if node == nil {
		return
	}

	postOrder(node.Left, codec)
	postOrder(node.Right, codec)

	codec.Add(strconv.Itoa(node.Val))
}

// Deserializes your encoded data to tree.
//    2
//      3
//        4
// [4, 3, 2]
func (this *Codec1) deserialize(data string) *TreeNode {
	strs := strings.Split(data, string(this.Delimiter))
	nums := make([]int, len(strs))
	for i := range strs {
		nums[i], _ = strconv.Atoi(strs[i])
	}

	return build(nums, 0, len(nums)-1)
}

func build(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	node := &TreeNode{
		Val: nums[end],
	}

	var left int
	for left = start; nums[left] < node.Val; left++ {
	}

	node.Left = build(nums, start, left-1)
	node.Right = build(nums, left, end-1)

	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

//	Notes
//	1. in-order traversal of BST generates ascending sequence => not unique

//	2.	inspired from solution, serialization of tree means 3 things to contain:
//		- tree structure
//		- node value
//		- delimiters to separate numbers in string
//		(need to split a big problem into smaller ones, this teaches me how
//		to tackle problem, I didn't figure it out this so I am not able to solve)

//	3.	becareful about post order traversal

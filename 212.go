package main

import "sort"

// Given an m x n board of characters and a list of strings words, return all words on the board.
//
// Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.
//
//
//
// Example 1:
//
//
// Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
// Output: ["eat","oath"]
// Example 2:
//
//
// Input: board = [["a","b"],["c","d"]], words = ["abcb"]
// Output: []
//
//
// Constraints:
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 12
// board[i][j] is a lowercase English letter.
// 1 <= words.length <= 3 * 104
// 1 <= words[i].length <= 10
// words[i] consists of lowercase English letters.
// All the strings of words are unique.

type Position struct {
	X, Y int
}

type Trie struct {
	Children [26]*Trie
}

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func (t *Trie) Build(board [][]byte, positions []Position, length int) {
	w, h := len(board[0]), len(board)

	for _, pos := range positions {
		// dfs
		visited := make([][]bool, h)
		for i := range visited {
			visited[i] = make([]bool, w)
		}

		idx := int(board[pos.Y][pos.X] - 'a')
		if t.Children[idx] == nil {
			t.Children[idx] = &Trie{}
		}

		t.Children[idx].dfs(board, pos, length-1, visited)
	}
}

func (t *Trie) dfs(board [][]byte, pos Position, length int, visited [][]bool) {
	w, h := len(board[0]), len(board)

	if length == 0 {
		return
	}

	for _, dir := range dirs {
		newY, newX := pos.Y+dir[0], pos.X+dir[1]

		if newX >= 0 && newY >= 0 && newX < w && newY < h && !visited[newY][newX] {
			if t.Children[board[newY][newX]-'a'] == nil {
				t.Children[board[newY][newX]-'a'] = &Trie{}
			}

			node := t.Children[board[newY][newX]-'a']
			visited[newY][newX] = true

			node.dfs(board, Position{newX, newY}, length-1, visited)

			// it's important to mark false when leave dfs
			// otherwise, other paths are restricted
			visited[newY][newX] = false
		}
	}
}

func (t *Trie) Search(word string) bool {
	node := t

	for i := range word {
		idx := int(word[i] - 'a')

		if node.Children[idx] == nil {
			return false
		}

		node = node.Children[idx]
	}

	return true
}

func findWords(board [][]byte, words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		if words[i][0] != words[j][0] {
			return words[i] < words[j]
		}

		return len(words[i]) > len(words[j])
	})

	pos := make([][]Position, 26)
	for i := range board {
		for j := range board[i] {
			pos[board[i][j]-'a'] = append(pos[board[i][j]-'a'], Position{j, i})
		}
	}

	existStrings := make([]string, 0)
	roots := make([]*Trie, 26)

	for _, word := range words {
		idx := int(word[0] - 'a')

		// build
		if roots[idx] == nil {
			root := &Trie{}

			root.Build(board, pos[idx], len(word))

			roots[idx] = root
		}

		// search
		if roots[idx].Search(word) {
			existStrings = append(existStrings, word)
		}
	}

	return existStrings
}

//	Notes
//	1.	inspired from solution, unlike my implementation, solution stores target
//		words as trie, and use this trie to indicate traverse on board

//		I think it's better because worst tc comes from board backtracking

//	2.	https://leetcode.com/problems/word-search-ii/discuss/59780/Java-15ms-Easiest-Solution-(100.00)

//		some points of optimization

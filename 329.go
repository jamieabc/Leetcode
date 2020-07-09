package main

type dir int

const (
	north dir = iota
	south
	east
	west
)

var (
	dirs = map[dir][]int{
		north: []int{-1, 0},
		south: []int{1, 0},
		east:  []int{0, 1},
		west:  []int{0, -1},
	}
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	var maxPath int
	memo := make(map[int]map[int]int) // max path from x, y
	for i := range matrix[0] {
		memo[i] = make(map[int]int)
	}

	for i := range matrix {
		for j := range matrix[0] {
			maxPath = max(maxPath, dfs(matrix, j, i, memo))
		}
	}

	return maxPath
}

func dfs(matrix [][]int, x, y int, memo map[int]map[int]int) int {
	maxPath := 1

	for _, arr := range dirs {
		newX, newY := x+arr[1], y+arr[0]
		if newX >= 0 && newX < len(matrix[0]) && newY >= 0 && newY < len(matrix) && matrix[newY][newX] > matrix[y][x] {
			if _, ok := memo[newX][newY]; !ok {
				memo[newX][newY] = dfs(matrix, newX, newY, memo)
			}
			maxPath = max(maxPath, 1+memo[newX][newY])
		}
	}

	return maxPath
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	inspired from solution, dfs can be seen as graph, each vertex/cell are
//		visited once, total tc is O(V+E), V is all points O(mn), E is 4
//		directions for a vertex, O(4V) = O(mn)

//	2.	not implement dp

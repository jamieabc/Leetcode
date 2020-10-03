package main

// Given an m x n matrix of non-negative integers representing the height of each unit cell in a continent, the "Pacific ocean" touches the left and top edges of the matrix and the "Atlantic ocean" touches the right and bottom edges.
//
// Water can only flow in four directions (up, down, left, or right) from a cell to another one with height equal or lower.
//
// Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.
//
// Note:
//
//     The order of returned grid coordinates does not matter.
//     Both m and n are less than 150.
//
//
//
// Example:
//
// Given the following 5x5 matrix:
//
//   Pacific ~   ~   ~   ~   ~
//        ~  1   2   2   3  (5) *
//        ~  3   2   3  (4) (4) *
//        ~  2   4  (5)  3   1  *
//        ~ (6) (7)  1   4   5  *
//        ~ (5)  1   1   2   4  *
//           *   *   *   *   * Atlantic
//
// Return:
//
// [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (positions with parentheses in above matrix).

func pacificAtlantic(matrix [][]int) [][]int {
	// later code assume matrix[0] exist
	if len(matrix) == 0 {
		return nil
	}

	pacific, atlantic := make([][]bool, len(matrix)), make([][]bool, len(matrix))
	for i := range pacific {
		pacific[i] = make([]bool, len(matrix[0]))
		atlantic[i] = make([]bool, len(matrix[0]))
	}

	// pacific reachable points
	q1 := make([][]int, 0)
	for j := range matrix[0] {
		q1 = append(q1, []int{0, j})
	}
	for i := 1; i < len(matrix); i++ {
		q1 = append(q1, []int{i, 0})
	}
	bfs(matrix, q1, pacific)

	// atlantic reachable points
	q2 := make([][]int, 0)
	for j := range matrix[0] {
		q2 = append(q2, []int{len(matrix) - 1, j})
	}
	for i := 0; i < len(matrix)-1; i++ {
		q2 = append(q2, []int{i, len(matrix[0]) - 1})
	}
	bfs(matrix, q2, atlantic)

	ans := make([][]int, 0)

	for i := range pacific {
		for j := range pacific[0] {
			if pacific[i][j] && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}

func bfs(matrix [][]int, queue [][]int, visited [][]bool) {
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		if visited[q[0]][q[1]] {
			continue
		}

		for _, dir := range dirs {
			x, y := q[1]+dir[1], q[0]+dir[0]

			if validPoint(matrix, x, y) && matrix[y][x] >= matrix[q[0]][q[1]] {
				queue = append(queue, []int{y, x})
			}
		}

		visited[q[0]][q[1]] = true
	}
}

func validPoint(matrix [][]int, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(matrix[0]) && y < len(matrix)
}

func pacificAtlantic1(matrix [][]int) [][]int {
	ans := make([][]int, 0)
	added := make([][]bool, len(matrix))
	invalid := make([][]bool, len(matrix))
	for i := range added {
		added[i] = make([]bool, len(matrix[0]))
		invalid[i] = make([]bool, len(matrix[0]))
	}

	for i := range matrix {
		for j := range matrix[0] {
			if added[i][j] || invalid[i][j] {
				continue
			}

			if touchable(matrix, j, i, added, invalid) {
				// all points with same or higher height are
				// possible points to be origin of water
				findEqualOrHigher(matrix, j, i, added, &ans)
			}
		}
	}

	return ans
}

var dirs [][]int = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func touchable(matrix [][]int, x, y int, added, invalid [][]bool) bool {
	var pacific, atlantic bool
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}

	queue := [][]int{{x, y}}

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		// for any point already added, it must be a valid path
		if added[q[1]][q[0]] {
			return true
		}

		// for any point already tries invalid, skip
		// for any visited point, skip
		if invalid[q[1]][q[0]] || visited[q[1]][q[0]] {
			continue
		}
		visited[q[1]][q[0]] = true

		updateCoast(matrix, q[0], q[1], &pacific, &atlantic)

		if pacific && atlantic {
			return true
		}

		for _, dir := range dirs {
			newX, newY := q[0]+dir[0], q[1]+dir[1]

			if newX >= 0 && newY >= 0 && newX < len(matrix[0]) && newY < len(matrix) && !visited[newY][newX] && matrix[newY][newX] <= matrix[q[1]][q[0]] {
				queue = append(queue, []int{newX, newY})
			}
		}
	}

	// only if a point cannot reach any ocean, mark as invalid
	if !pacific && !atlantic {
		invalid[y][x] = true
	}

	return false
}

func updateCoast(matrix [][]int, x, y int, pacific, atlantic *bool) {
	if x == 0 || y == 0 {
		*pacific = true
	}

	if x == len(matrix[0])-1 || y == len(matrix)-1 {
		*atlantic = true
	}
}

func findEqualOrHigher(matrix [][]int, x, y int, added [][]bool, ans *[][]int) {
	queue := [][]int{{x, y}}

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		if added[q[1]][q[0]] {
			continue
		}
		added[q[1]][q[0]] = true
		*ans = append(*ans, []int{q[1], q[0]})

		for _, dir := range dirs {
			newX, newY := q[0]+dir[0], q[1]+dir[1]

			if newX >= 0 && newX < len(matrix[0]) && newY >= 0 && newY < len(matrix) && !added[q[1]][q[0]] && matrix[newY][newX] >= matrix[y][x] {
				queue = append(queue, []int{newX, newY})
			}
		}
	}
}

//	Notes
//	1.	for every point, try if start from this point, both pacific &
//		atlantic are touchable, if yes then find equal or higher points near
//		starting point

//	2.	during check, some path is know to be not possible to reach both oceans,
//		for this kind of points should avoid trying again

//	3.	I think in the way problem describes, go for every point, check if this
//		point can go to both pacific & atlantic, ts: O(m^2 * n^2), but in this
//		ways, a point can go to pacific & atlantic will be calculated many times.

//		one way to improve this is to memoize a point can go to pacific/atlantic
//		or not.

//		once a point is found be reach both pacific & atlantic, other points
//		might also reachable as long as adjacent points are equal or higher.

//		there exists two computations:
//		- find equal or lower: check if a points is reachable to oceans
//		- find equal or higher: for a point  is reachable to oceans, adjacent
//		  points might also reachable

//		when I see this, I know there's something wrong for my thinking, because
//		it's a tangled solution, messy and complicated

//	4.	inspired from https://leetcode.com/problems/pacific-atlantic-water-flow/discuss/90733/Java-BFS-and-DFS-from-Ocean

//		another way to view this problem is by ocean point of view, a point can
//		reach oceans means oceans can also reach this point in reverse way

//		the problem wants to find a points reaches to both oceans, is same as
//		intersections of pacific reachable points & atlantic reachable points

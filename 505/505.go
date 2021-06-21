package main

// There is a ball in a maze with empty spaces (represented as 0) and walls (represented as 1). The ball can go through the empty spaces by rolling up, down, left or right, but it won't stop rolling until hitting a wall. When the ball stops, it could choose the next direction.

// Given the m x n maze, the ball's start position and the destination, where start = [startrow, startcol] and destination = [destinationrow, destinationcol], return the shortest distance for the ball to stop at the destination. If the ball cannot stop at destination, return -1.

// The distance is the number of empty spaces traveled by the ball from the start position (excluded) to the destination (included).

// You may assume that the borders of the maze are all walls (see examples).



// Example 1:

// Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [4,4]
// Output: 12
// Explanation: One possible way is : left -> down -> left -> down -> right -> down -> right.
// The length of the path is 1 + 1 + 3 + 1 + 2 + 2 + 2 = 12.

// Example 2:

// Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [3,2]
// Output: -1
// Explanation: There is no way for the ball to stop at the destination. Notice that you can pass through the destination but you cannot stop there.

// Example 3:

// Input: maze = [[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]], start = [4,3], destination = [0,1]
// Output: -1



// Constraints:

//     m == maze.length
//     n == maze[i].length
//     1 <= m, n <= 100
//     maze[i][j] is 0 or 1.
//     start.length == 2
//     destination.length == 2
//     0 <= startrow, destinationrow <= m
//     0 <= startcol, destinationcol <= n
//     Both the ball and the destination exist in an empty space, and they will not be in the same position initially.
//     The maze contains at least 2 empty spaces.


func shortestDistance(maze [][]int, start []int, destination []int) int {
    m := len(maze)
    if m == 0 {
        return -1
    }

    n := len(maze[0])
    if n == 0 {
        return -1
    }

    // 0: not visited
    visited := make([][]int, m)
    for i := range visited {
        visited[i] = make([]int, n)
    }

    dfs(maze, visited, start, destination)

    if visited[start[0]][start[1]] == int(1e9+7) {
        return -1
    }
    return visited[start[0]][start[1]]
}

var dirs = [][]int{
    {0, 1},
    {0, -1},
    {1, 0},
    {-1, 0},
}

func dfs(maze, visited [][]int, cur, dest []int) int {
    if cur[0] == dest[0] && cur[1] == dest[1] {
        return 0
    }

    m, n := len(maze), len(maze[0])
    result := int(1e9+7)
    visited[cur[0]][cur[1]] = result

    for _, dir := range dirs {
        newY, newX := cur[0], cur[1]
        var dist int

        for newY >= 0 && newX >= 0 && newY < m && newX < n && maze[newY][newX] == 0 {
            newY += dir[0]
            newX += dir[1]
            dist++
        }

        // loop terminates when point is invalid, need to revert back
        newY -= dir[0]
        newX -= dir[1]
        dist--

        // forward any step
        if newY != cur[0] || newX != cur[1] {
            if visited[newY][newX] == 0 {
                // not visited
                result = min(result, dist+dfs(maze, visited, []int{newY, newX}, dest))
            } else {
                // find shortest distance
                result = min(result, dist+visited[newY][newX])
            }
        }
    }

    if result != int(1e9+7) {
        visited[cur[0]][cur[1]] = result
    }

    return visited[cur[0]][cur[1]]
}

func min(i, j int) int {
    if i <= j {
        return i
    }
    return j
}

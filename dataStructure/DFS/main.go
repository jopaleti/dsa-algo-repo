package main

import "fmt"

/*
Question
Given an m x n 2D binary grid grid which represents a map of
 '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed
by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.
*/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	// Initializing the size of the grid ... m x n
	m, n := len(grid), len(grid[0]) 
	numIslands := 0

	var dfs func(int, int)

	dfs = func(i, j int) {
		if i < 0 || j<0 || i>=m || j>=n || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0' // Mark the cell as visited
		dfs(i-1, j) // up
		dfs(i+1, j) // down
		dfs(i, j-1) // left
		dfs(i, j+1) // right
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				numIslands++
				dfs(i, j)
			}
		}
	}

	return numIslands
}

func main() {
	grid := [][]byte{
        {'1', '1', '1', '1', '0'},
        {'1', '1', '0', '1', '0'},
        {'1', '1', '0', '0', '0'},
        {'0', '0', '0', '0', '0'},
    }

    fmt.Println(numIslands(grid)) // Output: 1
}
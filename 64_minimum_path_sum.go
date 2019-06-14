package leetcode

import "log"

func minPathSum(grid [][]int) int {
	cache := make(map[int]map[int]int)
	return minPathSumWithLimit(grid, 0, 0, cache)
}

func minPathSumWithLimit(grid [][]int, y, x int, cache map[int]map[int]int) (result int) {
	if cache[y] == nil {
		cache[y] = make(map[int]int)
	}
	if cache[y][x] != 0 {
		return cache[y][x]
	}

	defer func() {
		log.Println(result)
		cache[y][x] = result
	}()

	base := grid[y][x]

	if y+1 == len(grid) && x+1 == len(grid[0]) {
		result = base
		return
	}
	if y+1 == len(grid) {
		result = base + minPathSumWithLimit(grid, y, x+1, cache)
		return
	}
	if x+1 == len(grid[0]) {
		result = base + minPathSumWithLimit(grid, y+1, x, cache)
		return
	}

	result = base + min(minPathSumWithLimit(grid, y, x+1, cache), minPathSumWithLimit(grid, y+1, x, cache))
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// https://leetcode.com/problems/minimum-path-sum/discuss/23610/Go-Golang-solution
func minPathSum2(grid [][]int) int {
	var m int = len(grid)
	var n int = len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min64(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]

}
func min64(a, b int) int {
	if a < b {
		return a
	}
	return b
}

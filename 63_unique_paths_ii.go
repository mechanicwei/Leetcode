package leetcode

var cache map[int]map[int]int

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	cache = make(map[int]map[int]int)
	return uniquePathsWithObstacles2(obstacleGrid, 0, 0)
}

func uniquePathsWithObstacles2(obstacleGrid [][]int, x, y int) int {
	if cache[y] == nil {
		cache[y] = make(map[int]int)
	}
	if val, ok := cache[y][x]; ok {
		return val
	}

	if obstacleGrid[y][x] == 1 {
		cache[y][x] = 0
		return 0
	}
	if x == len(obstacleGrid[0])-1 && y == len(obstacleGrid)-1 {
		cache[y][x] = 1
		return 1
	}

	right := 0
	down := 0
	if x < len(obstacleGrid[y])-1 && obstacleGrid[y][x+1] == 0 {
		right = uniquePathsWithObstacles2(obstacleGrid, x+1, y)
	}
	if y < len(obstacleGrid)-1 && obstacleGrid[y+1][x] == 0 {
		down = uniquePathsWithObstacles2(obstacleGrid, x, y+1)
	}
	cache[y][x] = right + down
	return right + down
}

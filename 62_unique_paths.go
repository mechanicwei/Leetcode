package leetcode

func uniquePaths(m int, n int) int {
	cache := make(map[int]map[int]int)

	return uniquePathsWithCache(m, n, cache)
}

func uniquePathsWithCache(m, n int, cache map[int]map[int]int) int {
	if m == 1 || n == 1 {
		return 1
	}

	if cache[m] == nil {
		cache[m] = make(map[int]int)
	}
	if cache[n] == nil {
		cache[n] = make(map[int]int)
	}

	if cache[m][n] != 0 {
		return cache[m][n]
	}
	if cache[n][m] != 0 {
		return cache[n][m]
	}

	result := uniquePathsWithCache(m, n-1, cache) + uniquePathsWithCache(m-1, n, cache)
	cache[m][n] = result

	return result
}

func uniquePaths2(m int, n int) int {
	result := make(map[int]map[int]int)
	for i := 1; i <= m; i++ {
		if result[i] == nil {
			result[i] = make(map[int]int)
		}
		result[i][1] = 1

		for j := 2; j <= n; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}

	return result[m][n]
}

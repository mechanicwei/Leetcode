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

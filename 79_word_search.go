package leetcode

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if 1 == len(word) {
					return true
				}

				cache := make(map[int]map[int]bool)
				cache[i] = make(map[int]bool)
				cache[i][j] = true
				if existWithCache(board, word, 1, [2]int{i, j}, cache) {
					return true
				}
			}
		}
	}

	return false
}

func existWithCache(board [][]byte, word string, index int, s [2]int, cache map[int]map[int]bool) bool {
	i, j := s[0], s[1]
	if judge79(board, word, index, [2]int{i - 1, j}, copyMap79(cache)) {
		return true
	}
	if judge79(board, word, index, [2]int{i + 1, j}, copyMap79(cache)) {
		return true
	}
	if judge79(board, word, index, [2]int{i, j - 1}, copyMap79(cache)) {
		return true
	}
	if judge79(board, word, index, [2]int{i, j + 1}, copyMap79(cache)) {
		return true
	}

	return false
}

func judge79(board [][]byte, word string, index int, s [2]int, cache map[int]map[int]bool) bool {
	i, j := s[0], s[1]
	if cache[i][j] {
		return false
	}

	if i >= 0 && j >= 0 && i < len(board) && j < len(board[0]) && board[i][j] == word[index] {
		if cache[i] == nil {
			cache[i] = make(map[int]bool)
		}
		cache[i][j] = true

		if index+1 == len(word) {
			return true
		}
		if existWithCache(board, word, index+1, s, cache) {
			return true
		}
	}

	return false
}

func copyMap79(src map[int]map[int]bool) map[int]map[int]bool {
	dst := make(map[int]map[int]bool)
	for k, v := range src {
		if dst[k] == nil {
			dst[k] = make(map[int]bool)
		}

		for k1, v1 := range v {
			dst[k][k1] = v1
		}
	}
	return dst
}

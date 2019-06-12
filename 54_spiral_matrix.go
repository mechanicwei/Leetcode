package leetcode

func spiralOrder(matrix [][]int) []int {
	result := []int{}

	if len(matrix) == 0 {
		return result
	}

	sx := 0
	ex := len(matrix[0])
	sy := 0
	ey := len(matrix)

	for i, j := 0, 0; i < ey && i < ex; i, j = i+2, j+1 {
		appendX1(matrix, &result, j, sx+j, ex-1-j)
		appendY1(matrix, &result, ex-j-1, sy+j+1, ey-1-j)
		if ey-j-1 > j {
			appendX2(matrix, &result, ey-j-1, sx+j, ex-1-j)
		}
		if ex-j-1 > j {
			appendY2(matrix, &result, j, sy+j+1, ey-1-j)
		}
	}

	return result
}

func appendX1(matrix [][]int, result *[]int, y, s, e int) {
	row := matrix[y]
	for i := s; i <= e; i++ {
		*result = append(*result, row[i])
	}
}
func appendX2(matrix [][]int, result *[]int, y, s, e int) {
	row := matrix[y]
	for i := e; i >= s; i-- {
		*result = append(*result, row[i])
	}
}
func appendY1(matrix [][]int, result *[]int, x, s, e int) {
	if s >= e {
		return
	}

	for i := s; i < e; i++ {
		*result = append(*result, matrix[i][x])
	}
}
func appendY2(matrix [][]int, result *[]int, x, s, e int) {
	if s >= e {
		return
	}

	for i := e - 1; i >= s; i-- {
		*result = append(*result, matrix[i][x])
	}
}

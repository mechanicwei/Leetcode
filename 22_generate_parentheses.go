package leetcode

import "log"

func generateParenthesis(n int) []string {
	result := make(map[int]map[int][]string)

	for x := 0; x <= n; x++ {
		result[x] = make(map[int][]string)

		for y := x; y <= n; y++ {
			if y == 0 {
				continue
			}

			if x == y {
				if result[x-1][y] == nil {
					result[x][y] = []string{`(`}
				} else {
					newRes := make([]string, len(result[x-1][y]))
					for i, item := range result[x-1][y] {
						newRes[i] = `(` + item
					}
					result[x][y] = newRes
				}
			} else {
				newRes := make([]string, len(result[x-1][y])+len(result[x][y-1]))

				if len(newRes) == 0 {
					result[x][y] = []string{`)`}
				} else {
					for i, item := range result[x-1][y] {
						newRes[i] = `(` + item
					}
					for i, item := range result[x][y-1] {
						log.Println(i + len(result[x-1][y]))
						newRes[i+len(result[x-1][y])] = `)` + item
					}
					result[x][y] = newRes
				}
			}
		}
	}

	return result[n][n]
}

package leetcode

import "fmt"

// https://leetcode.com/problems/unique-binary-search-trees/
func numTrees(n int) int {
	used := make(map[int]bool, n)
	cache := make(map[string]int)

	r := numTrees2(n, used, cache)
	// fmt.Println(cache)
	return r
}

func numTrees2(n int, used map[int]bool, cache map[string]int) int {
	var result int
	var allUsed bool = true

	for i := 1; i <= n; i++ {
		if used[i] {
			continue
		}

		left := dupMap(used)
		right := dupMap(used)

		leftNum := 0
		rightNum := 0

		allUsed = false
		for j := 1; j <= n; j++ {
			if j < i {
				right[j] = true
				if !used[j] {
					leftNum++
				}
			} else if j > i {
				left[j] = true
				if !used[j] {
					rightNum++
				}
			} else {
				left[j] = true
				right[j] = true
			}
		}

		key := fmt.Sprintf("%d-%d-%d", i, leftNum, rightNum)
		if cache[key] != 0 {
			result += cache[key]
		} else {
			temp := numTrees2(n, left, cache) * numTrees2(n, right, cache)
			cache[key] = temp
			result += temp
		}
	}

	if allUsed {
		result = 1
	}

	return result
}

func dupMap(m map[int]bool) map[int]bool {
	m2 := make(map[int]bool, len(m))
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

// https://leetcode.com/problems/unique-binary-search-trees/discuss/31695/Golang-DP-solution

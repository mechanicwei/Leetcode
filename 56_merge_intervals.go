package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}

	if len(intervals) == 1 {
		return append(result, intervals...)
	}

	first, second := 0, 0
	for i, item := range intervals {
		if i == 0 {
			first, second = item[0], item[1]
			continue
		}

		if item[0] > second {
			result = append(result, []int{first, second})
			first, second = item[0], item[1]

		} else if item[1] < first {
			result = append(result, item)
		} else {
			first = min56(first, item[0])
			second = max56(second, item[1])
		}

		if i == len(intervals)-1 {
			result = append(result, []int{first, second})
		}
	}
	return result
}

func max56(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min56(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

package leetcode

func climbStairs(n int) int {
	result := map[int]int{}
	result[1] = 1
	result[2] = 2

	for i := 3; i <= n; i++ {
		result[i] = result[i-1] + result[i-2]
	}

	return result[n]
}

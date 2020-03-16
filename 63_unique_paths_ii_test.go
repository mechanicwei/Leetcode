package leetcode

import "fmt"

func Example_uniquePathsWithObstacles() {
	input := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	res := uniquePathsWithObstacles(input)
	fmt.Println(res)

	// Output:
	// 2
}

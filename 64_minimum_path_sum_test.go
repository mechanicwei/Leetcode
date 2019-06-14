package leetcode

import "fmt"

func Example_minPathSum() {
	input := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	fmt.Println(minPathSum(input))

	input2 := [][]int{
		[]int{9, 1, 4, 8},
	}
	fmt.Println(minPathSum(input2))

	// Output:
	// 7
	// 22
}

package leetcode

import "fmt"

func Example_spiralOrder() {
	m := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	fmt.Println(spiralOrder(m))

	// Output:
	// [1 2 3 6 9 8 7 4 5]
}

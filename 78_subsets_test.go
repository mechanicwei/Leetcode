package leetcode

import "fmt"

func Example_subsets() {
	res := subsets([]int{1, 2, 3})
	fmt.Println(res)

	// Output:
	// [[1] [] [1 2] [2] [1 3] [3] [1 2 3] [2 3]]
}

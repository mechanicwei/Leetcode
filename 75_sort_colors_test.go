package leetcode

import "fmt"

func Example_sortColors() {
	a := []int{2, 0, 2, 1, 1, 0}
	sortColors(a)
	fmt.Println(a)

	b := []int{0, 1, 2}
	sortColors(b)
	fmt.Println(b)

	// Output:
	// [0 0 1 1 2 2]
	// [0 1 2]
}

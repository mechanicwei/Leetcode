package leetcode

import "fmt"

func Example_removeDuplicates2() {
	data := []int{1, 1, 1, 2, 2, 3}
	// data := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	len := removeDuplicates2(data)
	fmt.Println(len)
	fmt.Println(data[:len])

	// Output:
	// 5
	// [1 1 2 2 3]
}

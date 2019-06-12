package leetcode

import "fmt"

func Example_merge() {
	a := [][]int{
		[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18},
	}
	b := [][]int{
		[]int{1, 4}, []int{0, 0},
	}
	c := [][]int{
		[]int{2, 3}, []int{4, 5}, []int{6, 7}, []int{8, 9}, []int{1, 10},
	}

	fmt.Println(merge(a))
	fmt.Println(merge(b))
	fmt.Println(merge(c))
	// Output:
	// [[1 6] [8 10] [15 18]]
	// [[0 0] [1 4]]
	// [[1 10]]
}

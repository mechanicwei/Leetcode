package leetcode

import "fmt"

func Example_exist() {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))

	board = [][]byte{
		[]byte{'a', 'a'},
	}
	fmt.Println(exist(board, "a"))

	board = [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'E', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCESEEEFS"))

	// Output:
	// true
	// true
	// false
	// true
	// true
}

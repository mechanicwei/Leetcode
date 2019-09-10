package leetcode

import "fmt"

func Example_generateParenthesis() {
	res := generateParenthesis(3)
	fmt.Println(res)

	// Output:
	// [((())) (()()) (())() ()(()) ()()()]
}

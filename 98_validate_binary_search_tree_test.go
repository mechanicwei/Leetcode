package leetcode

import "fmt"

func Example_isValidBST() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(isValidBST(root))

	// Output:
	// true
}

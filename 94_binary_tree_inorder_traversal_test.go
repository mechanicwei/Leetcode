package leetcode

import "fmt"

func Example_inorderTraversal() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	r := inorderTraversal(root)
	fmt.Println(r)

	// Output:
	// [1 3 2]

}

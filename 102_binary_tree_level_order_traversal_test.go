package leetcode

import "fmt"

func Example_levelOrder() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(levelOrder(root))
	// Output:
	// [[3] [9 20] [15 7]]
}

package leetcode

import "fmt"

func Example_maxPathSum() {
	// 	-10
	// 	/ \
	// 9  20
	// 	 /  \
	// 	15   7

	root := &TreeNode{
		Val: -10,
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

	val := maxPathSum(root)
	fmt.Println(val)

	// Output:
	// 42
}

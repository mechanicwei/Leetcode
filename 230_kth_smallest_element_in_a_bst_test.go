package leetcode

import "fmt"

func Example_kthSmallest() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	val := kthSmallest(root, 1)
	fmt.Println(val)

	// Output:
	// 1
}

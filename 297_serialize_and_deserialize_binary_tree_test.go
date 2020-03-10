package leetcode

import "fmt"

func Example_serializeTree() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	res := serializeTree(root)
	fmt.Println(res)

	// Output:
	// 3,5,1,6,2,0,8,null,null,7,4
}

func Example_deserializeTree() {
	src := `3,5,1,6,2,0,8,null,null,7,4`
	root := deserializeTree(src)
	fmt.Println(root.Left.Right.Left.Val)

	// Output:
	// 7
}

package leetcode

import "fmt"

func Example_zigzagLevelOrder() {
	// 	 3
	// 	/ \
	// 9  20
	// 	 /  \
	// 	15   7

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

	root2 := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: -1,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	res := zigzagLevelOrder(root)
	fmt.Println(res)
	res = zigzagLevelOrder(root2)
	fmt.Println(res)
	// Output:
	// [[3] [20 9] [15 7]]
	// [[0] [4 2] [1 3 -1] [8 6 1 5]]
}

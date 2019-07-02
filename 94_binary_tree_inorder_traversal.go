package leetcode

// https://leetcode.com/problems/binary-tree-inorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	temp := []*TreeNode{}
	current := root
	for {
		if current == nil {
			return result
		}

		if current.Left != nil {
			temp = append(temp, current)
			current = current.Left
		} else {
			for {
				result = append(result, current.Val)

				if current.Right == nil {
					if len(temp) == 0 {
						return result
					}

					current = temp[len(temp)-1]
					temp = temp[:len(temp)-1]
				} else {
					current = current.Right
					break
				}
			}
		}
	}

}

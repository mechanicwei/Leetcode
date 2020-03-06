package leetcode

import "math"

// https://leetcode.com/problems/binary-tree-maximum-path-sum/
// https://leetcode.com/problems/binary-tree-maximum-path-sum/discuss/467723/Go-golang-clean-solution
func maxPathSum(root *TreeNode) int {
	var max = int(math.NaN())
	maxPathSum2(root, &max)
	return max
}

func maxPathSum2(node *TreeNode, max *int) int {
	val := node.Val
	left := 0
	right := 0
	if node.Left != nil {
		left = maxPathSum2(node.Left, max)
	}
	if node.Right != nil {
		right = maxPathSum2(node.Right, max)
	}
	if left > 0 {
		val += left
	}
	if right > 0 {
		val += right
	}

	if val > *max {
		*max = val
	}

	if left > right {
		if left > 0 {
			return node.Val + left
		}
	} else {
		if right > 0 {
			return node.Val + right
		}
	}

	return node.Val
}

package leetcode

// https://leetcode.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	return isValidBST2(root, nil, nil)
}

func isValidBST2(root *TreeNode, min *int, max *int) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= *min {
		return false
	}
	if max != nil && root.Val >= *max {
		return false
	}

	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}

	return isValidBST2(root.Left, min, &root.Val) && isValidBST2(root.Right, &root.Val, max)
}

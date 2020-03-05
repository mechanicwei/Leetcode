package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	m := len(nums) / 2
	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[0:m]),
		Right: sortedArrayToBST(nums[m+1:]),
	}
}

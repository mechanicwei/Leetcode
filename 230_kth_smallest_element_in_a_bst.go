package leetcode

func kthSmallest(root *TreeNode, k int) int {
	res := []int{}
	kthSmallestWithRes(root, k, &res)
	return res[k-1]
}

func kthSmallestWithRes(node *TreeNode, k int, res *[]int) {
	if node == nil {
		return
	}
	if len(*res) == k {
		return
	}

	kthSmallestWithRes(node.Left, k, res)
	if len(*res) == k {
		return
	}
	*res = append(*res, node.Val)
	kthSmallestWithRes(node.Right, k, res)
}

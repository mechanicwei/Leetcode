package leetcode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generateTrees2(1, n)
}

func generateTrees2(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	res := []*TreeNode{}

	for i := start; i <= end; i++ {
		var lefts []*TreeNode
		var rights []*TreeNode

		lefts = generateTrees2(start, i-1)
		rights = generateTrees2(i+1, end)

		for _, node1 := range lefts {
			for _, node2 := range rights {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  node1,
					Right: node2,
				})
			}
		}
	}

	return res
}

package leetcode

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	currentLevel := []*TreeNode{root}
	nextLevel := []*TreeNode{}

	for len(currentLevel) != 0 {
		var row []int
		for i := 0; i < len(currentLevel); i++ {
			row = append(row, currentLevel[i].Val)
			if currentLevel[i].Left != nil {
				nextLevel = append(nextLevel, currentLevel[i].Left)
			}
			if currentLevel[i].Right != nil {
				nextLevel = append(nextLevel, currentLevel[i].Right)
			}

		}
		res = append(res, row)

		currentLevel, nextLevel = nextLevel, currentLevel[:0]
	}

	return res
}

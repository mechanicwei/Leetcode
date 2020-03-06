package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	level := []*TreeNode{root}
	toRight := false
	for len(level) > 0 {
		output := make([]int, 0)
		for _, node := range level {
			output = append(output, node.Val)
		}
		res = append(res, output)
		newLevel := []*TreeNode{}
		for i := 0; i < len(level); i++ {
			node := level[len(level)-1-i]
			if node == nil {
				continue
			}
			if toRight {
				if node.Left != nil {
					newLevel = append(newLevel, node.Left)
				}
				if node.Right != nil {
					newLevel = append(newLevel, node.Right)
				}
			} else {
				if node.Right != nil {
					newLevel = append(newLevel, node.Right)
				}
				if node.Left != nil {
					newLevel = append(newLevel, node.Left)
				}
			}
		}
		toRight = !toRight
		level = newLevel
	}
	return res
}

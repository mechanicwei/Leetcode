package leetcode

// https://leetcode.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	nodes := []*TreeNode{root}
	nodes2 := []*TreeNode{}

	for {
		if !isSymmetricForSlice(nodes) {
			return false
		}

		allNil := true
		for _, node := range nodes {
			if node == nil {
				nodes2 = append(nodes2, nil, nil)
			} else {
				allNil = false
				nodes2 = append(nodes2, node.Left, node.Right)
			}
		}

		if allNil {
			return true
		}

		nodes, nodes2 = nodes2, nodes[:0]
	}
}

func isSymmetricForSlice(slice []*TreeNode) bool {
	if len(slice) == 1 {
		return true
	}

	for i := 0; i < len(slice)/2; i++ {
		if slice[i] == nil && slice[len(slice)-i-1] != nil {
			return false
		}
		if slice[i] != nil && slice[len(slice)-i-1] == nil {
			return false
		}
		if slice[i] == nil && slice[len(slice)-i-1] == nil {
			continue
		}
		if slice[i].Val != slice[len(slice)-i-1].Val {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/symmetric-tree/discuss/33244/A-Golang-solution

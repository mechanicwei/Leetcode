package leetcode

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/discuss/243390/go-superfast
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathsP := []*TreeNode{}
	pathsQ := []*TreeNode{}
	lowestCommonAncestor2(root, p, q, nil, &pathsP, &pathsQ)

	var ancestor *TreeNode
	for i := 0; i < len(pathsP); i++ {
		if len(pathsQ) > i && pathsQ[i] == pathsP[i] {
			ancestor = pathsQ[i]
		} else {
			break
		}
	}

	return ancestor
}

func lowestCommonAncestor2(node, p, q *TreeNode, paths []*TreeNode, pathsP, pathsQ *[]*TreeNode) {
	if node == nil {
		return
	}
	if node == p {
		*pathsP = append(*pathsP, paths...)
		*pathsP = append(*pathsP, p)
	}
	if node == q {
		*pathsQ = append(*pathsQ, paths...)
		*pathsQ = append(*pathsQ, q)
	}

	if len(*pathsP) > 0 && len(*pathsQ) > 0 {
		return
	}

	paths = append(paths, node)

	lowestCommonAncestor2(node.Left, p, q, paths, pathsP, pathsQ)
	lowestCommonAncestor2(node.Right, p, q, paths, pathsP, pathsQ)
}

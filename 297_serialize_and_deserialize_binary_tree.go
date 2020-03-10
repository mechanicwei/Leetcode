package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/discuss/184386/Ruby-solution
func serializeTree(root *TreeNode) string {
	level := []*TreeNode{root}
	vals := []string{}
	for len(level) > 0 {
		newLevel := []*TreeNode{}
		stashNil := []*TreeNode{}
		addNode := func(node *TreeNode) {
			if node == nil {
				stashNil = append(stashNil, node)
			} else {
				newLevel = append(newLevel, stashNil...)
				stashNil = []*TreeNode{}
				newLevel = append(newLevel, node)
			}
		}

		for _, node := range level {
			if node == nil {
				vals = append(vals, `null`)
			} else {
				vals = append(vals, fmt.Sprintf(`%d`, node.Val))
			}

			if node == nil {
				addNode(nil)
				addNode(nil)
			} else {
				addNode(node.Left)
				addNode(node.Right)
			}
		}

		level = newLevel
	}

	return strings.Join(vals, `,`)
}

func deserializeTree(src string) *TreeNode {
	vals := strings.Split(src, `,`)
	if vals[0] == `null` {
		return nil
	}

	val, _ := strconv.Atoi(vals[0])
	root := &TreeNode{
		Val: val,
	}
	deserializeTree2(root, 0, vals)
	return root
}

func deserializeTree2(root *TreeNode, index int, vals []string) {
	if 2*index+1 > len(vals)-1 {
		return
	}

	leftVal := vals[2*index+1]
	if leftVal != `null` {
		val, _ := strconv.Atoi(leftVal)
		root.Left = &TreeNode{
			Val: val,
		}
	}

	rightVal := vals[2*index+2]
	if rightVal != `null` {
		val, _ := strconv.Atoi(rightVal)
		root.Right = &TreeNode{
			Val: val,
		}
	}

	if root.Left != nil {
		deserializeTree2(root.Left, 2*index+1, vals)
	}
	if root.Right != nil {
		deserializeTree2(root.Right, 2*index+2, vals)
	}
}

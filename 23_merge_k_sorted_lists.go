package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	for i := len(lists) / 2; i >= 0; i-- {
		downNode(lists, i)
	}

	head := &ListNode{}
	current := head
	for {
		current.Next, lists = popNode(lists)
		if current.Next == nil {
			break
		} else {
			current = current.Next
		}
	}
	return head.Next
}

func downNode(lists []*ListNode, i int) {
	if len(lists) == 0 {
		return
	}

	node := lists[i]
	var left, right *ListNode
	if 2*i+1 < len(lists) {
		left = lists[2*i+1]
	}
	if 2*i+2 < len(lists) {
		right = lists[2*i+2]
	}

	if left == nil && right == nil {
		return
	}

	if right == nil || (left != nil && left.Val < right.Val) {
		if node == nil || node.Val > left.Val {
			lists[i], lists[2*i+1] = lists[2*i+1], lists[i]
			downNode(lists, 2*i+1)
		}
	} else {
		if node == nil || node.Val > right.Val {
			lists[i], lists[2*i+2] = lists[2*i+2], lists[i]
			downNode(lists, 2*i+2)
		}
	}
}

func popNode(lists []*ListNode) (*ListNode, []*ListNode) {
	if len(lists) == 0 {
		return nil, lists
	}

	node := lists[0]
	if node == nil {
		return nil, lists
	}
	if node.Next == nil {
		lists[0] = lists[len(lists)-1]
		lists = lists[:len(lists)-1]
		downNode(lists, 0)
	} else {
		lists[0] = node.Next
		downNode(lists, 0)
	}

	return node, lists
}

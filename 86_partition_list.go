package leetcode

func partition(head *ListNode, x int) *ListNode {
	head1 := &ListNode{}
	head2 := &ListNode{}

	node1 := head1
	node2 := head2

	current := head
	for current != nil {
		if current.Val < x {
			node1.Next = current
			node1 = current
		} else {
			node2.Next = current
			node2 = current
		}

		current = current.Next
	}

	node1.Next = head2.Next
	node2.Next = nil

	return head1.Next
}

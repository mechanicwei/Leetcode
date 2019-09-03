package leetcode

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	lastN := head
	var prevOfLastN *ListNode

	currentN := 0
	for node != nil {
		if currentN > n-1 {
			prevOfLastN = lastN
			lastN = lastN.Next

		} else {
			currentN++
		}

		node = node.Next
	}

	if prevOfLastN == nil {
		return nil
	} else {
		prevOfLastN.Next = lastN.Next

		return head
	}
}

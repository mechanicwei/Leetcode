package leetcode

// https://leetcode.com/problems/sort-list/discuss/46858/28ms-Golang-solution-with-MergeSort
func sortList(head *ListNode) *ListNode {
	sortList2(head, nil)
	return head
}

func sortList2(head *ListNode, end *ListNode) {
	if head == nil || head.Next == nil || head == end {
		return
	}

	i := head.Next
	j := head.Next
	jParent := head
	for i != end {
		if i.Val >= head.Val {
			i = i.Next
		} else {
			i.Val, j.Val = j.Val, i.Val
			i = i.Next
			jParent = j
			j = j.Next
		}
	}

	head.Val, jParent.Val = jParent.Val, head.Val
	sortList2(head, jParent)
	sortList2(j, nil)
}

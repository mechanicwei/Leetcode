package leetcode

// https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1
	}

	ll1 := reverseList(l1)
	ll2 := reverseList(l2)

	head := &ListNode{}
	current := head
	var flag int // 是否进位
	for {
		var v1, v2 int
		if ll1 != nil {
			v1 = ll1.Val
		}
		if ll2 != nil {
			v2 = ll2.Val
		}
		current.Val = (v1 + v2 + flag) % 10
		if (v1 + v2 + flag) >= 10 {
			flag = 1
		} else {
			flag = 0
		}

		if ll1 != nil {
			ll1 = ll1.Next
		}
		if ll2 != nil {
			ll2 = ll2.Next
		}

		if ll1 == nil && ll2 == nil {
			if flag == 1 {
				current.Next = &ListNode{Val: 1}
			}

			break
		} else {
			current.Next = &ListNode{}
			current = current.Next
		}
	}

	return head
}

func reverseList(l *ListNode) *ListNode {
	var prev *ListNode
	for l != nil {
		next := l.Next
		l.Next = prev
		prev = l
		l = next
	}
	return prev
}

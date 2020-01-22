package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	res := head

	var in bool

	var m1, m2 *ListNode
	var n1, n2 *ListNode

	var previous *ListNode

	index := 1

	for {
		if m == 1 && index == 1 {
			m2 = head
			in = true
		} else if index == m-1 {
			m1 = head
			m2 = head.Next
			in = true
		} else if index == n {
			n1 = head
		} else if index > n {
			n2 = head
			break
		}

		if in {
			next := head.Next
			head.Next = previous
			previous = head
			head = next
		} else {
			head = head.Next
		}

		index++
	}

	if m1 != nil {
		m1.Next = n1
	}
	m2.Next = n2

	if m1 == nil {
		if n1 == nil {
			return m2
		} else {
			return n1
		}

	} else {
		return res
	}
}

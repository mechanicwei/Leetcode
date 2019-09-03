package leetcode

func Example_removeNthFromEnd() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	res := removeNthFromEnd(head, 2)
	printList(res)

	// Output:
	// 1->2->3->5
}

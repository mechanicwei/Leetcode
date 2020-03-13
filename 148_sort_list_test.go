package leetcode

func Example_sortList() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	res := sortList(head)
	printList(res)

	// Output:
	// 1->2->3->4
}

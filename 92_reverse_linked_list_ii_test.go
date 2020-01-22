package leetcode

func Example_reverseBetween() {
	head := &ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: &ListNode{
					Next: &ListNode{
						Val: 5,
					},
					Val: 4,
				},
				Val: 3,
			},
			Val: 2,
		},
		Val: 1,
	}

	h := reverseBetween(head, 2, 4)
	printList(h)

	// Output:
	// 1->4->3->2->5
}

func Example_reverseBetween2() {
	head := &ListNode{
		Val: 5,
	}

	h := reverseBetween(head, 1, 1)
	printList(h)

	// Output:
	// 5

}

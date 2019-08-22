package leetcode

import "fmt"

func Example_addTwoNumbers() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l := addTwoNumbers(l1, l2)
	testPrint(l)

	// Output:
	// 7
	// 0
	// 8
}

func testPrint(l *ListNode) {
	if l == nil {
		return
	}
	fmt.Println(l.Val)
	testPrint(l.Next)
}

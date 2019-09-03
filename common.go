package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	fmt.Printf("%d", head.Val)
	if head.Next != nil {
		fmt.Printf("->")
		printList(head.Next)
	}
}

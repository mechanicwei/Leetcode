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

func buildList(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	head := &ListNode{}
	current := head
	for _, val := range data {
		current.Next = &ListNode{
			Val: val,
		}
		current = current.Next
	}

	return head.Next
}

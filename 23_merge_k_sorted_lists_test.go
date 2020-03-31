package leetcode

func Example_mergeKLists() {
	lists := []*ListNode{
		buildList([]int{1, 4, 5}),
		buildList([]int{1, 3, 4}),
		buildList([]int{2, 6}),
	}
	// lists := []*ListNode{
	// 	buildList([]int{2}),
	// 	buildList([]int{}),
	// 	buildList([]int{-1}),
	// }

	head := mergeKLists(lists)
	printList(head)
	// Output:
	// 1->1->2->3->4->4->5->6
}

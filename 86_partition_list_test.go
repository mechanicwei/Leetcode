package leetcode

func Example_partition() {
	head := partition(buildList([]int{1, 4, 3, 2, 5, 2}), 3)
	printList(head)

	// Output:
	// 1->2->2->4->3->5
}

package leetcode

import "fmt"

func Example_numDecodings() {
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("20"))

	fmt.Println(numDecodings3("226"))
	fmt.Println(numDecodings3("20"))

	// Output:
	// 3
	// 1
	// 3
	// 1
}

package leetcode

import "fmt"

func Example_myPow() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.00000, -2))
	fmt.Println(myPow(2.00000, 0))
	fmt.Println(myPow(1.00000, -2147483648))
	fmt.Println(myPow(-1.00000, 2147483648))
	// Output:
	// 1024
	// 0.25
	// 1
	// 1
	// 1
}

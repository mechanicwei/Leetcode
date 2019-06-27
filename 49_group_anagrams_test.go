package leetcode

import "fmt"

func Example_groupAnagrams() {
	input1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input1))

	// Output:
	// [[eat tea ate] [tan nat] [bat]]
}

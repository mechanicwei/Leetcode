package leetcode

import "fmt"

func Example_wordBreak() {
	var s = "applepenapple"
	var wordDict = []string{"apple", "pen"}

	res := wordBreak(s, wordDict)
	fmt.Println(res)

	// Output:
	// true
}

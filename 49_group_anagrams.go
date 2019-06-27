package leetcode

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)

	for _, str := range strs {
		isMatched := false
		for i, _ := range result {
			group := &result[i]
			if matched49(str, (*group)[0]) {
				(*group) = append((*group), str)
				isMatched = true
				continue
			}
		}

		if !isMatched {
			result = append(result, []string{str})
		}
	}

	return result
}

func matched49(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, s := range a {
		map1[s] += 1
	}
	for _, s := range b {
		map2[s] += 1
	}

	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/group-anagrams/discuss/19267/Golang-concise-solution-beats-95-other-solutions
// https://leetcode.com/problems/group-anagrams/discuss/244569/golang-beat-100
// https://leetcode.com/problems/group-anagrams/discuss/212228/a-foolish-solution-by-golang-GO-276ms

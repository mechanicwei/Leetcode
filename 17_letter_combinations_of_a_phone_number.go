package leetcode

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	maps := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	result := []string{}

	for i := range digits {
		if len(result) == 0 {
			result = append(result, maps[digits[i]]...)
		} else {
			newResult := []string{}
			for _, item := range result {
				for _, str := range maps[digits[i]] {
					newResult = append(newResult, item+str)
				}
			}
			result = newResult
		}

	}

	return result
}

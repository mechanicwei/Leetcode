package leetcode

func isAnagram(s string, t string) bool {
	byteMap := make(map[byte]int)
	for i := range s {
		byteMap[s[i]] += 1
	}

	for i := range t {
		byteMap[t[i]] -= 1
	}

	for _, v := range byteMap {
		if v != 0 {
			return false
		}
	}

	return true
}

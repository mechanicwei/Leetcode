package leetcode

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	length := 0
	exist := map[byte]int{
		s[0]: 0,
	}
	start := 0
	for end := 1; end < len(s); end++ {
		if len(s)-start <= length {
			break
		}

		if index, ok := exist[s[end]]; ok {
			if index >= start {
				if end-start > length {
					length = end - start
				}
				start = index + 1
			}
		}
		exist[s[end]] = end
	}
	if len(s)-start > length {
		length = len(s) - start
	}

	return length
}

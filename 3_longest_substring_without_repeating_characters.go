package leetcode

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	length := 0
	for i := range s {
		exist := make(map[byte]int)
		exist[s[i]] = i
		flag := 0
		for j := i + 1; j < len(s); j++ {
			if _, ok := exist[s[j]]; ok {
				if (j - i) > length {
					length = (j - i)
				}
				flag = 1
				break
			}
			exist[s[j]] = j
		}
		if flag == 0 {
			if len(s)-i > length {
				length = len(s) - i
			}
			return length
		}
	}

	return length
}

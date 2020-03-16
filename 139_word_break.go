package leetcode

func wordBreak(s string, wordDict []string) bool {
	dictMap := make(map[string]bool)
	for _, word := range wordDict {
		dictMap[word] = true
	}

	cache := make(map[string]bool)
	return wordBreak2(s, dictMap, cache)
}

func wordBreak2(s string, wordDict map[string]bool, cache map[string]bool) bool {
	if s == "" {
		return true
	}
	if val, ok := cache[s]; ok {
		return val
	}
	for i := range s {
		if wordDict[s[:i+1]] {
			if wordBreak2(s[i+1:], wordDict, cache) {
				cache[s] = true
				return true
			} else {
				continue
			}
		} else {
			continue
		}
	}
	cache[s] = false
	return false
}

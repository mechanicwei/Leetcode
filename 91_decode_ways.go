package leetcode

// https://leetcode.com/problems/decode-ways/
func numDecodings(s string) int {
	return numDecodings2(s, 0)
}

func numDecodings2(s string, start int) int {
	if start > len(s)-1 {
		return 1
	}

	currentNum := byte(s[start]) - '0'
	if currentNum == 0 {
		return 0
	}
	if start == len(s)-1 {
		return 1
	}

	nextNum := byte(s[start+1]) - '0'
	if currentNum*10+nextNum <= 26 {
		return numDecodings2(s, start+2) + numDecodings2(s, start+1)
	} else {
		return numDecodings2(s, start+1)
	}
}

func numDecodings3(s string) int {
	cache := make(map[int]int)
	cache[0] = 1

	for i := 1; i <= len(s); i++ {
		currentNum := byte(s[len(s)-i]) - '0'
		if currentNum == 0 {
			cache[i] = 0
			continue
		}

		if i == 1 {
			cache[i] = 1
		} else {
			nextNum := byte(s[len(s)-i+1]) - '0'
			if currentNum*10+nextNum <= 26 {
				cache[i] = cache[i-1] + cache[i-2]
			} else {
				cache[i] = cache[i-1]
			}
		}
	}
	return cache[len(s)]
}

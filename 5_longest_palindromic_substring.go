package leetcode

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	end := s[len(s)-1]
	current := ""
	for i := 0; i < len(s)-1; i++ {
		if s[i] == end {
			if isPalindrome(s[i:]) {
				current = s[i:]
				break
			}
		}
	}

	if current == s {
		return current
	}

	prev := longestPalindrome(s[:len(s)-1])
	if len(current) > len(prev) {
		return current
	} else {
		return prev
	}
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

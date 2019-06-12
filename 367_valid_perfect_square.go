package leetcode

func isPerfectSquare(num int) bool {
	i := 1
	var previous_i int
	for {
		j := i * i
		if j == num {
			return true
		}
		if j < num {
			previous_i = i
			i *= 2
		} else {
			if i-previous_i == 1 {
				return false
			} else {
				i = (previous_i + i) / 2
			}
		}
	}
}

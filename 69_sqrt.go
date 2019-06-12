package leetcode

func mySqrt(x int) int {
	i := 1
	var previous_i int
	for {
		j := i * i
		if j == x {
			return i
		}
		if j < x {
			previous_i = i
			i *= 2
		} else {
			if i-previous_i == 1 {
				return previous_i
			} else {
				i = (previous_i + i) / 2
			}
		}
	}
}

package leetcode

func myPow(x float64, n int) float64 {
	if x == 1 {
		return x
	}
	if x == -1 {
		if n%2 == 0 {
			return -x
		} else {
			return x
		}
	}

	var result float64 = 1

	if n == 0 {
		return 1.00000
	} else if n > 0 {
		for i := 0; i < n; i++ {
			result *= x
		}
	} else {
		base := 1 / x
		for i := 0; i < -n; i++ {
			result *= base
		}
	}

	return result
}

// https://leetcode.com/problems/powx-n/discuss/259063/Go-Golang-0ms-100
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n < 0 {
		return 1 / myPow2(x, -n)
	}

	r := myPow2(x, n/2)
	if n%2 == 0 {
		return r * r
	} else {
		return r * r * x
	}
}

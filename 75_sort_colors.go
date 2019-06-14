package leetcode

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	var i, j int = 1, len(nums) - 1
	for ; i < len(nums); i++ {
		if nums[i] < nums[0] {
			continue
		}

		for ; j > 0; j-- {
			if nums[j] < nums[0] {
				break
			}
		}

		if i < j {
			swap75(nums, i, j)
		} else {
			break
		}
	}

	swap75(nums, 0, j)

	sortColors(nums[:j])
	if j+1 < len(nums) {
		sortColors(nums[j+1:])
	}
}

func swap75(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

// https://leetcode.com/problems/sort-colors/discuss/144289/Golang-one-pass-solution
func sortColors2(a []int) {
	l := len(a)
	if l < 2 {
		return
	}

	r := 0
	b := 0
	i := 0
	for i < l-b {
		if a[i] == 0 {
			if i != r {
				a[i], a[r] = a[r], a[i]
			} else {
				i++
			}
			r++
		} else if a[i] == 2 {
			if i != l-1-b {
				a[i], a[l-1-b] = a[l-1-b], a[i]
			} else {
				i++
			}
			b++
		} else {
			i++
		}
	}
}

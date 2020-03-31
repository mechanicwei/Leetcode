package leetcode

func removeDuplicates2(nums []int) int {
	validLength := 1
	num := 1
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] || num == 1 {
			if nums[i] != nums[i-1] {
				num = 1
			} else {
				num = 2
			}

			nums[j] = nums[i]
			j++
			validLength++
		}
	}

	return validLength
}

package leetcode

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	len := 0
	var prev int
	for i := range nums {
		if i == 0 {
			prev = nums[i]
			len++
			continue
		}

		if prev != nums[i] {
			nums[len] = nums[i]
			len++
			prev = nums[i]
		}
	}

	return len
}

package leetcode

func searchRange(nums []int, target int) []int {
	return searchRangeWithOffset(nums, target, 0)
}

func searchRangeWithOffset(nums []int, target, offset int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	mid := len(nums) / 2
	if nums[mid] > target {
		return searchRangeWithOffset(nums[:mid], target, offset)
	} else if nums[mid] < target {
		if len(nums) == 1 {
			return []int{-1, -1}
		}
		return searchRangeWithOffset(nums[mid+1:], target, offset+mid+1)
	} else {
		start := mid + offset
		end := mid + offset

		prev := searchRangeWithOffset(nums[:mid], target, offset)
		if prev[0] != -1 {
			start = prev[0]
		}

		next := searchRangeWithOffset(nums[mid+1:], target, offset+mid+1)
		if next[1] != -1 {
			end = next[1]
		}
		return []int{start, end}
	}
}

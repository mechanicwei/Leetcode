package leetcode

// https://leetcode.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	return searchWithOffset(nums, target, 0)
}

func searchWithOffset(nums []int, target, offset int) int {
	if len(nums) == 0 {
		return -1
	}

	mid := len(nums) / 2
	if nums[mid] == target {
		return mid + offset
	}
	if len(nums) == 1 {
		return -1
	}

	// 前半连续
	if nums[0] <= nums[mid-1] {
		if target >= nums[0] && target <= nums[mid-1] {
			return searchWithOffset(nums[:mid], target, offset)
		} else {
			return searchWithOffset(nums[mid+1:], target, mid+1+offset)
		}
	} else {
		if target >= nums[mid+1] && target <= nums[len(nums)-1] {
			return searchWithOffset(nums[mid+1:], target, mid+1+offset)
		} else {
			return searchWithOffset(nums[:mid], target, offset)
		}
	}
}

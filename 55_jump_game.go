package leetcode

func canJump(nums []int) bool {
	var cache = make(map[int]bool)
	return canJump2(nums, cache)
}

func canJump2(nums []int, cache map[int]bool) bool {
	if len(nums) != 1 && nums[0] == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}

	for i := 1; i < len(nums); i++ {
		if nums[len(nums)-i-1] >= i {
			if result, ok := cache[len(nums)]; ok {
				return result
			}
			if canJump2(nums[:len(nums)-1], cache) {
				cache[len(nums)-1] = true
				return true
			} else {
				cache[len(nums)-1] = false
			}
		}
	}

	return false
}

// https://leetcode.com/problems/jump-game/discuss/161398/simple-go-solution
func canJump3(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	} else if nums[0] == 0 && length > 1 {
		return false
	}

	maxJump := nums[0]
	for i := 1; i < length-1; i++ {
		next := nums[i] + i
		if next > maxJump {
			maxJump = next
		}

		if nums[i] == 0 && i >= maxJump {
			return false
		}
	}

	return true
}

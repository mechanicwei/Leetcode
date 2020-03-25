package leetcode

// func rob(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	option1 := nums[0] + rob(nums[2:])
// 	option2 := nums[1]
// 	if len(nums) > 3 {
// 		option2 += rob(nums[3:])
// 	}
// 	if option1 > option2 {
// 		return option1
// 	} else {
// 		return option2
// 	}
// }

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := map[int]int{
		len(nums) - 1: nums[len(nums)-1],
	}
	for i := len(nums) - 2; i >= 0; i-- {
		option1 := nums[i] + res[i+2]
		option2 := nums[i+1] + res[i+3]
		if option1 > option2 {
			res[i] = option1
		} else {
			res[i] = option2
		}
	}
	return res[0]
}

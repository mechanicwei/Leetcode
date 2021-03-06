package leetcode

func maxSubArray(nums []int) int {
	max := 0
	sum := 0
	for i, n := range nums {
		if i == 0 {
			max = n
			sum = n
			continue
		}

		if n > 0 {
			if sum > 0 {
				sum += n
				if sum > max {
					max = sum
				}
			} else {
				sum = n

				if sum > max {
					max = sum
				}
			}
		} else {
			if sum > 0 {
				sum += n
			} else {
				sum = n
				if sum > max {
					max = sum
				}
			}
		}
	}

	return max
}

// https://leetcode.com/problems/maximum-subarray/discuss/309414/Go-Simple-DP-solution
func maxSubArray2(nums []int) int {
	currSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	current := nums[len(nums)-1]
	maxSum := nums[len(nums)-1]

	for i := len(nums) - 1 - 1; i >= 0; i-- {
		current = max(nums[i], current+nums[i])
		maxSum = max(current, maxSum)
	}

	return max(maxSubArray3(nums[:len(nums)-1]), maxSum)
}

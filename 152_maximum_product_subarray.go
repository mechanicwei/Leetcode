package leetcode

func maxProduct(nums []int) int {
	allProduct := nums[0]
	maxProduct := nums[0]
	var positiveProduct int
	if nums[0] > 0 {
		positiveProduct = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			positiveProduct = 0
			allProduct = 0
		}

		if allProduct == 0 {
			allProduct = nums[i]
		} else {
			allProduct *= nums[i]
		}

		maxProduct = max152(maxProduct, allProduct)

		if nums[i] > 0 {
			if positiveProduct > 0 {
				positiveProduct *= nums[i]
			} else {
				positiveProduct = nums[i]
			}
		} else {
			positiveProduct = 0
		}

		maxProduct = max152(maxProduct, positiveProduct)
	}

	return maxProduct
}

func max152(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

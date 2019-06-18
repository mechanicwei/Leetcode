package leetcode

import (
	"math"
)

func subsets(nums []int) [][]int {
	len := int(math.Pow(2, float64(len(nums))))
	result := make([][]int, 0, len)
	subsets2(nums, &result)

	return result
}

func subsets2(nums []int, result *[][]int) {
	if len(nums) == 1 {
		*result = append(*result, []int{nums[0]}, []int{})
		return
	} else {
		subsets2(nums[:len(nums)-1], result)
	}

	lastNum := nums[len(nums)-1]
	length := len(*result)
	for i := 0; i < length; i++ {
		dst := make([]int, len((*result)[i]), len((*result)[i])+1)
		copy(dst, (*result)[i])
		*result = append(*result, append(dst, lastNum))
	}
}

// https://leetcode.com/problems/subsets/discuss/267896/golang-14-lines-0ms
func subsets3(nums []int) [][]int {
	ret := [][]int{{}}
	for i := range nums {
		for _, e := range ret {
			t := make([]int, len(e)+1)
			copy(t, e)
			t[len(t)-1] = nums[i]
			ret = append(ret, t)
		}
	}
	return ret
}

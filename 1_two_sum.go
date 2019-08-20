package leetcode

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int) // number -> index
	for i, num := range nums {
		if index, ok := cache[target-num]; ok {
			return []int{index, i}
		}

		cache[num] = i
	}
	return nil
}

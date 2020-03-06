package leetcode

func maxProfit(prices []int) int {
	i := 0 // 最小值index
	j := 0 // 最大值index
	profit := 0
	for index := 1; index < len(prices); index++ {
		if prices[index] < prices[i] {
			i = index
			j = index
		} else if prices[index] > prices[j] {
			j = index
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}

	return profit
}

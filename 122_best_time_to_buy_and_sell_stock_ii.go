package leetcode

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	i := 0
	profit := 0
	for index := 1; index < len(prices); index++ {
		if prices[index-1] > prices[index] {
			profit += prices[index-1] - prices[i]
			i = index
		}
	}
	if i != len(prices)-1 {
		profit += prices[len(prices)-1] - prices[i]
	}
	return profit
}

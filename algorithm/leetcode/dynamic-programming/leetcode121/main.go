package leetcode121



func maxProfit(prices []int) int {

	max := 0
	dp := make([]int, len(prices))
	dp[0] = prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i] = mini(dp[i - 1], prices[i])
		if prices[i] - dp[i] > max {
			max = prices[i] - dp[i]
		}
	}


	return max
}


func mini(x, y int) int {
	if x < y {
		return x
	}

	return y
}
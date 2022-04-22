package leetcode188

// 股票问题
// https://leetcode-cn.com/circle/article/qiAgHn/

func maxProfit(k int, prices []int) int {
	l := len(prices)

	if k <= 1 || l <= 1 {
		return 0
	}

	dp := make([][2]int, l)
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
		dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])
	}



	return 0
}


func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
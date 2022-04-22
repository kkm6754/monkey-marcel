package leetcode123


// 手里最多只能有1支股票
// 最后计算利润时，手里没有股票


//func maxProfit(prices []int) int {
//	if len(prices) < 2 {
//		return 0
//	}
//
//	// dp[i][0] 第i天交易后，手上没有股票，最大利润
//	// dp[i][1] 第i天交易后，手上有股票，最大利润
//	dp := make([][2]int, len(prices))
//	dp[0][0] = 0
//	dp[0][1] = -prices[0]
//	for i := 1; i < len(prices); i++ {
//		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i]) // 前一天没有股票， 前一天有股票+今天卖掉
//		dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i]) // 前一天有股票，前一天没有股票+今天买股票
//	}
//
//	return dp[len(prices) - 1][0]
//}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


// 一天最多能有2笔交易
// 到每一天结束，当前状态可以分为5种
//		没有操作
//		第一次买
//		第一次卖
//		第二次买
//		第二次卖
// 当天没有操作，就以前一天状态为准

// buy1 可以是沿用前一天第一次买；或者是前一天没有操作，当天第一次买
// sell1 可以是沿用前一天第一次卖；或者是前一天买入，当天第一次卖
// buy2 可以是沿用前一天第二次买；或者是前一天第一次卖，当天第二次买
// sell2 可以是沿用前一天第二次卖；或者是前一天第二次买，当天第二次卖


// 初始值
// buy1 第一天买入股票
// sell1 第一天买入，然后由卖出
// buy2 第一天买入、卖出，然后又买入
// sell2 第一天买入、卖出，然后又买入、卖出


func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1 + prices[i])
		buy2 = max(buy2, sell1 - prices[i])
		sell2 = max(sell2, buy2 + prices[i])
	}
	return sell2
}
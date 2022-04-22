package leetcode55

// 55. 跳跃游戏

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true // 每个位置是否能够到达

	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if dp[j] && nums[j] + j >= i { // 每个位置的数是最大步长
				dp[i] = true
				break
			}
		}
	}

	return dp[len(nums) - 1]
}

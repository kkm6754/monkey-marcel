package leetcode115

//115. 不同的子序列


func numDistinct(s string, t string) int {
	sl := len(s)
	tl := len(t)
	if sl == 0 || tl == 0 {
		return 0
	}

	dp := make([][]int, sl + 1) // dp[i][j]：从开头到s[i-1]的子串中，出现『从开头到t[j-1]的子串』的 次数。即：前 i 个字符的 s 子串中，出现前 j 个字符的 t 子串的次数。
	for i, _ := range dp {
		dp[i] = make([]int, tl + 1)
	}

	for i := 0; i < sl + 1; i++ {
		for j := 0; j < tl + 1; j++ {
			if j == 0 { // 目标字符串为""，出现次数是1
				dp[i][j] = 1
			} else if i == 0 { // 源字符串是""，出现次数是0
				dp[i][j] = 0
			}else {
				if s[i - 1] == t[j - 1] { // 第i个位置与第j个位置字符相同，最后一位字符相同，源字符串中的这个字符可以使用，也可以不实用
					dp[i][j] = dp[i - 1][j] + dp[i - 1][j - 1]
				}else {
					dp[i][j] = dp[i - 1][j]	// 最后一位字符不相同，源字符串只能去往前找
				}
			}
		}
	}


	return dp[sl][tl]
}
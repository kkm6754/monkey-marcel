package leetcode72

//72. 编辑距离


// state f[i][j] 前i个字符串转换成前j个字符串的编辑距离
// function: f[i][j] = 以下情况最小值
//		f[i-1][j] + 1  word1补一个字符变成word2
//		f[i][j-1] + 1  word1删除一个字符变成word2
//		f[i-1][j-1]  分成2种情况， i字符 == j字符  和 i字符 != j字符
// initialize: f[0][j]=j 	f[i][0]=i
// answer:	f[word1.length][word2.length]

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	// 其中一个是空字符串
	if l1 == 0 || l2 ==0 {
		return l1 + l2
	}

	mini := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	// 长、宽+1，方便计数，不做存储
	dp := make([][]int, l1 + 1)
	for i, _ := range dp {
		dp[i] = make([]int, l2 + 1)
	}

	for i, _ := range dp {
		dp[i][0] = i
	}
	for j,_ := range dp[0]{
		dp[0][j] = j
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1]
			if word1[i - 1] != word2[j - 1] {
				dp[i][j] = mini(left, mini(down, leftDown + 1))
			} else {
				dp[i][j] = mini(left, mini(down, leftDown))
			}
		}
	}

	return dp[l1][l2]
}
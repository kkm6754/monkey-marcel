package main

import "fmt"

// 双序列型动态规划

// state：
//		f[i][j]代表第一个sequence的前i个字符/数字，配上第二个sequence的前j个。。
// function：f[i][j] = 研究第i个和第j个的匹配关系
// initialize：f[i][0] 和 f[0][i]
// answer: f[s1.length][s2.length]


// 最长公共子序列
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

// 公共子序列 任意位置，只要相对位置相同就是
// 输入：
//A = "ABCD"
//B = "EACB"
//
//输出：
//2  LCS是AC



// X=(x1,x2,.....xn)
// Y={y1,y2,.....ym}

// state:
//		f[j][j] 前i个字符匹配前j个字符的lcs长度
// function：f[i][j] =
//		xi = yj		f[i][j] = f[n-1][m-1] + 1
//		xi != yj	f[i][j] = max(f[i-1][j], f[i][j-1])
// initialize: f[i][0]=0 f[0][j] = 0
// answer: f[a.length][b.length]



func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	dp := make([][]int, m + 1)
	for i, _ := range dp {
		dp[i] = make([]int, n + 1)
	}

	for i, c1 := range text1 { // i，j从0开始
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] +1
			}else {
				dp[i+1][j+1] = Max(dp[i][j+1], dp[i+1][j])
			}
		}
	}


	return dp[m][n]
}


func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
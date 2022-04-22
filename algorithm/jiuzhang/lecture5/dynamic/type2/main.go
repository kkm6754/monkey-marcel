package main

import "fmt"

// 单序列型动态规划


// 如果不是坐标型的动态规划
// 一般有 N 个数/字符，就开 N+1 个位置的数组处理
// 第0个位置单独作初始化


// 分割回文串II
// 给一个字符串s，将其分割成多个回文子串，问最少分割次数


// dp[i] 代表i+1长度（索引 0～i）的字符串，分割成回文串的最小次数
// dp[0] = 0 不用切割
// n长度的字符串，结果为 dp[n-1] 
// [0, i]字符串被j分割，[0, j] [j+1, i] ，当[j+1, i]是回文字符串时，dp[i] = dp[j] + 1

// [j+1, i] 回文字符串判断




func minCut(s string) int {
	if len(s) == 0 {
		return 0
	}

	n := len(s)
	isPali := make([][]bool, n, n) // 存放[i, j] 是否是回文字符串
	for i, _ := range isPali {
		isPali[i] = make([]bool, n)
	}
	fmt.Println(isPali)

	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if i == j {  // 单个字符是回文字符串
				isPali[i][j] = true
			} else if j - i == 1 && s[i] == s[j] { // 长度为2子串，2个字符相同
				isPali[i][j] = true
			} else if j - i > 1 && s[i] == s[j] && isPali[i+1][j-1] { // 首位字符相同，去头去尾，子串是回味字符串
				isPali[i][j] = true
			}
		}
	}
	fmt.Println(isPali)

	dp := make([]int, n) // dp[i] 表示 [0,i] 最小分割次数
	for i, _ := range dp { // 初始化，全部分割成一个字符的分割次数
		dp[i] = i
	}

	for i := 1; i < len(dp); i++ {
		if isPali[0][i] { // [0, i]就是回文字符串
			dp[i] = 0
			continue
		}

		for j := 0; j < i; j++ {
			if isPali[j + 1][i] {
				if dp[j] + 1 < dp[i]{
					dp[i] = dp[j] + 1
				}

			}
		}
	}

	return dp[n - 1]
}


func main() {
	s := "cbb"
	fmt.Println(minCut(s))
}
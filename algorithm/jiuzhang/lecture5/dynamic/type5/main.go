package main

import "fmt"

// 背包类动态规划

// 放n个物品到v容量到背包中，每个物品分别占用ci的容量，得到wi的价值
// 初始化：
//		恰好放满背包时的最大值，初始化时，F[0] = 0，其它值初始化 F[1...] = 负无穷
//		不用装满背包，只装到最大，初始化时，F[0...] = 0



// 0-1 背包问题
// 有 n 个物品和一个大小为 m 的背包. 给定数组 A 表示每个物品的大小和数组 V 表示每个物品的价值.
// 问最多能装入背包的总价值是多大?
// 每个物品只能放一次；可以放或者不放


// state：
//		f[i][v] 代表放前i个物品放入容量为v的背包中，总价值最大
// function：f[i][v] = max(f[i-1][v], f[i-1][v-ci] + vi) 分2种情况，第i个物品不放背包中 + 第i个物品放入背包中
// initialize：f[0][...] = f[...][0] = 0
// answer: f[n][m]



func backPackII (m int, A []int, V []int) int {
	dp := make([][]int, len(A) + 1)
	for i, _ := range dp {
		dp[i] = make([]int, m + 1)
	}

	for i := 0; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if A[i - 1] > j {
				dp[i][j] = dp[i - 1][j]
			} else {
				dp[i][j] = Max(dp[i - 1][j], dp[i - 1][j - A[i - 1]] + V[i-1])
			}
		}
	}

	return dp[len(A)][m]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


func main() {
	m := 10
	A := []int{2, 3, 5, 7}
	V := []int{1, 5, 2, 4}

	fmt.Println(backPackII(m, A, V))
}






// 完全背包问题
// 有 N 种物品和一个容量为 V 的背包，每种物品都有无限件可用。放入第 i 种物品 的费用是 Ci，价值是 Wi。
// 求解:将哪些物品装入背包，可使这些物品的耗费的费用总 和不超过背包容量，且价值总和最大。
// 每种物品有无限件可以取





// 多重背包问题
// 有 N 种物品和一个容量为 V 的背包。第 i 种物品最多有 Mi 件可用，每件耗费的 空间是 Ci，价值是 Wi。
// 求解将哪些物品装入背包可使这些物品的耗费的空间总和不超 过背包容量，且价值总和最大。
// 每种物品有指定数量可以取

// state：
//		f[i][v] 取前i种物品放到容量为v的背包中最大价值
// function: f[i][v] = max{f[i-1][v - k * Ci] + k * Wi} 对于第 i 种物品有 Mi +1 种策略:取 0 件，取 1 件......取 Mi 件
// initialize:
// answer:
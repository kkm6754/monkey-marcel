package leetcode174

import "fmt"

// 174. 地下城游戏


func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 1
	}

	m := len(dungeon)
	n := len(dungeon[0])

	mini := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}


	checkHealth := func(hel, val int) int {
		if val >= hel { // 注意相等情况，即血量刚好够
			return 1
		} else {
			return hel - val
		}
	}



	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	// 从目的地往回计算
	// 每个节点最小血量是1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m - 1 && j == n - 1 {
				dp[i][j] = checkHealth(1, dungeon[i][j])
			} else if i == m - 1 {
				dp[i][j] = checkHealth(dp[i][j + 1], dungeon[i][j])
			} else if j == n - 1 {
				dp[i][j] = checkHealth(dp[i + 1][j], dungeon[i][j])
			} else {
				dp[i][j] = checkHealth(mini(dp[i][j + 1], dp[i + 1][j]), dungeon[i][j])
			}
		}
	}
	fmt.Println(dp)

	res := dp[0][0]
	return res
}

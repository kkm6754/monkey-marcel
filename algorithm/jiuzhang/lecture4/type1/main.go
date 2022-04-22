package main

import "fmt"

// 坐标型动态规划

// 数组
// 爬楼
// 跳跃



// 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。

// state:
//		f[x][y] 从起点到(x,y)的最短路径
// function:
//		f[x][y] = min(f[x-1][y], f[x][y-1]) + A[x][y]
// intialize:
//		f[0][0] = (0, 0)
//		f[i][0] = sum(0, 0~i, 0) 第0列
// 		f[0][i] = sum(0, 0~0, i) 第0行
// answer:
//		f[m-1][n-1]



func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

    l := len(grid)
	dp := make([][]int, l)
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}


	return dp[l-1][len(grid[l-1])-1]
}


func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}



func main() {
	grid := make([][]int, 4)
	grid[0] = []int{1, 2, 4, 5, 9}
	grid[1] = []int{2, 2, 2, 4, 0}
	grid[2] = []int{4, 2, 8, 5, 4}
	grid[3] = []int{5, 4, 3, 4, 3}

	ret := minPathSum(grid)
	fmt.Println(ret)	
}
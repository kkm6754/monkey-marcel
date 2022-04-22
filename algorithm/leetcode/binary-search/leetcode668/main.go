package main

import "fmt"

//668. 乘法表中第k小的数


// 二维中找第k小的数

func findKthNumber(m int, n int, k int) int {
	if k == 1 {
		return 1
	}
	if k == m * n {
		return m * n
	}

	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}

	// 计算当前mid位于第几小的数
	midCount := func(m, n, mid int) int {
		res := 0
		for i := 1; i <= m; i++ {
			res += min(mid / i, n) // 第i行有多少个小于mid的数
		}
		return res
	}


	l := 1
	r := m * n
	for l < r {
		mid := l + (r - l) / 2
		if midCount(m, n, mid) < k {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func main() {
	fmt.Println(findKthNumber(3, 3, 6))
}
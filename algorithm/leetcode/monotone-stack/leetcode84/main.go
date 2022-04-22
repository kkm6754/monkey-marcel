package main

// 84. 柱状图中最大的矩形

//[2,1,5,6,2,3]
//[0,2,1,5,6,2,3,0]

func largestRectangleArea(heights []int) int {
	n :=  len(heights)
	if n == 0 {
		return 0
	}

	// 简单栈实现
	// st存储下标
	st, pos := make([]int, n + 2), 0
	push := func(v int) {
		st[pos] = v
		pos++
	}
	pop := func() int {
		pos--
		return st[pos]
	}
	top := func() int {
		return st[pos-1]
	}

	// 最大值函数
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	//首尾增加哨兵
	get := func(i int) int {
		if i == 0 || i == n+1 {
			return 0
		}
		return heights[i-1]
	}

	//处理程序
	res := 0
	for i := 0; i < n + 2; i++ {
		for ; pos > 0 && get(top()) > get(i); {
			res = max(get(pop()) * (i - top() - 1), res) // 注意 宽 的计算，已经取出了栈顶
		}
		push(i)
	}
	return res
}




package leetcode42


// 42. 接雨水
//  [0,1,0,2,1,0,1,3,2,1,2,1]


func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	

	st, pos := make([]int, n), 0
	push := func(v int) {
		st[pos] = v
		pos++
	}
	pop := func() int {
		pos--
		return st[pos]
	}
	top := func() int {
		return st[pos - 1]
	}


	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	res := 0
	for i := 0; i < n; i++ {
		for ;pos > 0 && height[i] > height[top()]; {
			head := pop()
			if pos == 0 { // 最左侧的没有凹洼处
				break
			}

			// 根据高的差，横向计算面积
			res = res + (i - top() -1) * (min(height[i], height[top()]) - height[head]) // 宽的计算，高的计算
		}

		push(i)
	}


	return res
}

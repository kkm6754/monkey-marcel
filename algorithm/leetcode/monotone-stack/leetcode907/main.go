package leetcode907

//907. 子数组的最小值之和

// 找到每个元素作为最小值，其子数组的个数
// m个元素 + nums[i]最小值 + n个元素
// nums[i] 作为最小值的子数组的个数 m + n + 1 + m*n = (m+1)*(n+1)

//[3,1,2,4]

//0 0 1*3 3
//1 2 6*1 6
//0 1 2*2 4
//0 0 1*4 4

//[11,81,94,43,3]

//0 1 2 	0 0 1*1 * 94
//0 1		0 1 1*2 * 81
//0
//0 3		2 0 3*1 * 43
//
// todo




func sumSubarrayMins(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	mod := 1000000007

	st, pos := make([]int, n), 0
	push := func(val int) {
		st[pos] = val
		pos++
	}
	pop := func() int {
		pos--
		return st[pos]
	}
	top := func() int {
		return st[pos - 1]
	}

	res := 0
	for i, v := range arr {
		for pos > 0 && v < arr[top()] { // 下一个元素比栈顶元素小
			head := top()
			val := arr[pop()]
			if pos > 0 {
				res = res + val * (head - top()) * (i - head)
			}else {
				res = res + val * 1 * (i - head)
			}
		}
		push(i)
	}
	for pos > 0 {
		head := pop()
		if pos == 0 {
			res = res + arr[head] * head * (n - head)
		}
		res = res + arr[head] * 1 * (n - head)
	}

	return res % mod
}
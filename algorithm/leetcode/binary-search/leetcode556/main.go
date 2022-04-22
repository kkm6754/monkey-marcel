package main

import (
	"fmt"
	"math"
	"sort"
)

//556. 下一个更大元素 III


func nextGreaterElement(n int) int {
	// 拆分十进制数，注意顺序已倒置
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n % 10)
		n = n / 10
	}
	fmt.Println(digits)

	if len(digits) <= 1 {
		return -1
	}

	// 找最小转折点
	breakIndex := 0
	for i := 1; i < len(digits); i++ {
		if digits[i - 1] > digits[i] {
			breakIndex = i
			break
		}
	}

	fmt.Println(breakIndex)

	if breakIndex == 0 {
		return -1
	}

	// 最小可替换点
	l, r := 0, breakIndex
	for l < r {
		m := l + (r - l) / 2
		if digits[m] <= digits[breakIndex] {
			l = m + 1
		}else {
			r = m
		}
	}
	// 交换
	digits[l], digits[breakIndex] = digits[breakIndex], digits[l]
	// 前面数字排序
	sort.Sort(sort.Reverse(sort.IntSlice(digits[: breakIndex])))

	res := digits[0]
	count := 10
	for i := 1; i < len(digits); i++ {
		res = res + digits[i] * count
		count = count * 10
	}

	// 最大值不能超过MaxInt
	if res > math.MaxInt32 || n < 0 {
		return -1
	}

	return res

}

func main() {
	n := 12443322
	fmt.Println(nextGreaterElement(n))
}
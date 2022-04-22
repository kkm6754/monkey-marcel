package leetcode28



func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext(next, needle)

	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}

		// 匹配，进位
		if haystack[i] == needle[j] {
			j++
		}

		// 结束
		if j == n {
			return i - n + 1
		}
	}
	return -1
}



// getNext 构造前缀表next
// params:
//		  next 前缀表数组
//		  s 模式串
func getNext(next []int, s string) {
	l := 0 // 当前长度下的 最长相同前后缀 的长度
	next[0] = l // 初始位置没有前后缀

	// r永远在l的右边，用作位置移动
	for r := 1; r < len(s); r++ {

		// l、r位置数据不同，l回退
		for l > 0 && s[r] != s[l] {
			l = next[l-1]
		}

		// l、r位置数据相同，l右移
		// 这里如果不是初始状态，l会记录着r-1长度下 最长相同前后缀 的长度
		if s[r] == s[l] {
			l++
		}
		next[r] = l
	}
}

package leetcode459

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	l := 0
	next := make([]int, n)
	for r := 1; r < n; r++ {
		for l > 0 && s[l] != s[r] {
			l = next[l - 1]
		}
		if s[l] == s[r] {
			l++
		}
		next[r] = l
	}


	//a b c a b c a b c a b c
	//0 0 0 1 2 3 4 5 6 7 8 9
	// next[n-1]  最长相同前后缀的长度
	if next[n - 1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}

	return false
}
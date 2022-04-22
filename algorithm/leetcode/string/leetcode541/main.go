package leetcode541


func reverseStr(s string, k int) string {
	ss := []byte(s)
	for i := 0; i < len(s); i += 2*k {
		if i + k < len(s) {
			res := ss[i : i + k]
			reverseString(res)
		} else {
			res := ss[i : len(ss)]
			reverseString(res)
		}
	}

	return string(ss)
}


func reverseString(s []byte)  {
	if len(s) < 2 {
		return
	}

	l, r := 0, len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
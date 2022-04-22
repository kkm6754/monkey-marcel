package leetcode344

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

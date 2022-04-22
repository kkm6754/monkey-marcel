package leetcode151


func reverseWords(s string) string {
	ss := []byte(s)
	res := make([]byte, 0)

	for i := len(ss) - 1; i >= 0; {
		l, r := i, i
		for i >= 0 && ss[i] == ' ' {
			i--
		}
		r = i
		for i >= 0 && ss[i] != ' ' {
			i--
		}
		l = i

		if l != r {
			res = append(res, ss[l+1 : r+1]...)
			res = append(res, ' ')
		}
	}

	return string(res[ : len(res) - 1])

}
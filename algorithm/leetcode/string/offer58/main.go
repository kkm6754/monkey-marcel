package offer58


func reverseLeftWords(s string, n int) string {
	bs := []byte(s)
	l := len(bs)
	for i := 0; i < n; i++ {
		tmp := bs[0]
		for j := 0; j < l - 1; j++ {
			bs[j] = bs[j + 1]
		}
		bs[l - 1] = tmp
	}

	return string(bs)
}
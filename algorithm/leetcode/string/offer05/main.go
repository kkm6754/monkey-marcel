package offer05


func replaceSpace(s string) string {
	res := ""
	ss := []byte(s)
	for _, b := range ss {
		if b == ' ' {
			res = res + "%20"
		} else {
			res = res + string(b)
		}
	}

	return res
}

func replaceSpace1(s string) string {
	b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)
}
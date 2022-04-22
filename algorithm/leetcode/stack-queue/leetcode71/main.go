package leetcode71

func simplifyPath(path string) string {
	arr := []byte(path)
	res := make([]byte, 0)
	po := make([]byte, 0)

	for len(arr) > 0 {
		if arr[0] == '.' {
			po = append(po, '.')
			arr = arr[1:]
			continue
		} else if l := len(po); l > 0 {
			for l >= 2 && len(res) > 1 {
				res = res[:len(res)-1]
				for res[len(res)-1] != '/' {
					res = res[:len(res)-1]
				}
				if len(po) == 2 {
					po = make([]byte, 0)
				} else {
					po = po[:len(po)-2]
				}
			}
			po = make([]byte, 0)
		}

		if arr[0] == '/' {
			if len(res) > 0 && res[len(res)-1] == '/' {
				arr = arr[1:]
			} else {
				res = append(res, '/')
				arr = arr[1:]
			}
		} else {
			res = append(res, arr[0])
			arr = arr[1:]
		}
	}

	// 最后是 .. ，需要返回上一层
	for len(po) >= 2 && len(res) > 1 {
		res = res[:len(res)-1]
		for res[len(res)-1] != '/' {
			res = res[:len(res)-1]
		}
		if len(po) == 2 {
			po = make([]byte, 0)
		} else {
			po = po[:len(po)-2]
		}
	}

	if len(res) > 1 && res[len(res)-1] == '/' {
		res = res[:len(res)-1]
	}

	return string(res)
}

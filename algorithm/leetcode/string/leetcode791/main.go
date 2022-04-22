package leetcode791

import "sort"

// 对自定义类型进行排序，重写下面三个函数即可
type Str struct {
	Data []byte
	Sort map[byte]int
}
func (s Str) Len() int {
	return len(s.Data)
}

func (s Str) Less(i, j int) bool {
	return s.Sort[s.Data[i]] < s.Sort[s.Data[j]]
}

func (s Str) Swap(i, j int) {
	s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
}

func customSortString(S string, T string) string {
	byteStr := []byte(S)
	input := []byte(T)

	var str = Str{
		Data: input,
		Sort: make(map[byte]int),
	}
	for i := 0; i < len(byteStr); i++ {
		str.Sort[byteStr[i]] = i
	}

	sort.Sort(str)
	return string(str.Data)
}



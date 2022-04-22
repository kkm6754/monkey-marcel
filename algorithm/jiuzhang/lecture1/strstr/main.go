package main

import "fmt"

// strstr 子串
// 找子串的位置


func strStr(haystack string, needle string) int {

	// 边界，字符串长度为0，一定要子串放在前
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}

	// 边界，子串长度过长
	if len(haystack) < len(needle) {
		return -1
	}

	// 比较范围 len(haystack)-len(needle)+1
	for i:=0; i < len(haystack)-len(needle)+1; i++ {
		for j:=0; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1
}

func main(){
	ret := strStr("22", "233")
	fmt.Print(ret)
}
package main

import "fmt"


func main() {
	

	sli := make([]int, 0, 3)
	tmp := make([]int, 0, len(sli))
	for val := range sli {
		tmp = append(tmp, val)
	}
	fmt.Println(tmp)

	

}
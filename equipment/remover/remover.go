package main

import (
	"fmt"
)

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	all := []int{0, 1, 2, 5, 4, 3, 6, 7, 8, 9}
	fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]
	all = RemoveIndex(all, 5)
	fmt.Println(all) //[0 1 2 3 4 6 7 8 9]
}

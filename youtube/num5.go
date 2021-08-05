package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int
	a = append(a, 5)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 7)
	a = append(a, 8)
	a = append(a, 9)
	sort.Ints(a)
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}

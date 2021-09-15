package main

import (
	"fmt"
)

func main() {

	var a, count int

	fmt.Scan(&a)

	for i := 0; i <= a; i++ {
		for j := 0; j < 60; j++ {
			for k := 0; k < 60; k++ {
				b := []int{i, j, k}
				result := IntInSlice(a, b)
				if result == true {
					count += 1
				}
			}
		}
	}
	fmt.Printf("%d", count)
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b%10 == a || b/10 == a {
			return true
		}
	}
	return false
}

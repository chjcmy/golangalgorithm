package main

import "fmt"

func revursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * revursive(n-1)
}

func main() {
	fmt.Printf("%d", revursive(5))
}

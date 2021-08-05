package main

import "fmt"

func recursive2(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return recursive2(b, a%b)
	}
}

func main() {
	fmt.Printf("%d", recursive2(192, 162))
}

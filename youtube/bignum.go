package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	card := make([][]int, a, b)

	for i, ints := range card[0] {
		card[0][i] = ints
	}

	fmt.Printf("len %d cap %d", len(card), cap(card))
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, result int

	fmt.Scanf("%d %d", &a, &b)

	card := make([][]int, a)

	for i := range card {
		card[i] = make([]int, b)
	}

	fmt.Printf("len %d cap %d \n", len(card), cap(card))

	for i := 0; i < len(card); i++ {
		for j := 0; j < cap(card); j++ {
			fmt.Scanf("%d", &card[i][j])
		}
	}

	for i := range card {
		sort.Ints(card[i])
	}

	result = 0

	for i := range card {
		if result < card[i][0] {
			result = card[i][0]
		}
	}

	fmt.Printf("%d", result)
}

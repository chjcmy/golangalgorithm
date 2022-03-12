package main

import (
	"fmt"
)

func main() {
	var a, b, result int

	fmt.Scanf("%d %d\n", &a, &b)

	arr := make([]int, a)

	for i := 0; i < a; i++ {
		fmt.Scanf("%d\n", &arr[i])
	}

	for i := a - 1; i >= 0; i-- {
		if b > arr[i] {
			result += b / arr[i]
			b %= arr[i]
			if b == 0 {
				break
			}
		}
	}
	fmt.Println(result)
}

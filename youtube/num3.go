package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int
	fmt.Scan(&a)

	arr := make([]int, a)

	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%d", &arr[i])
	}
	sort.Ints(arr)

	result := 0
	count := 0

	for i := 0; i < a; i++ {
		count += 1
		if count >= arr[i] {
			result += 1
			count = 0
		}
	}
	fmt.Printf("%d", result)

}

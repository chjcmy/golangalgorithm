package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Scan(&a)
	result, _ := strconv.Atoi(string(a[0]))

	for i := 0; i < len(a); i++ {
		num, _ := strconv.Atoi(string(a[i]))
		if num <= 1 || result <= 1 {
			result += num
		} else {
			result *= num
		}
	}
	print(result)
}

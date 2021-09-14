package main

import "fmt"

func main() {
	var a, b int
	result := 0
	fmt.Scan(&a, &b)

	for a > b {
		if a%b == 0 {
			a -= 1
			result += 1
		}
		a /= b
		result += 1
	}
	if a > 1 {
		a -= 1
		result += 1
	}
	fmt.Print(result)
}

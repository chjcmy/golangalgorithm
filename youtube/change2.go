package main

import "fmt"

func main() {
	var coin = [4]int{500, 100, 50, 10}
	var a int
	var results = 0

	fmt.Scan(&a)

	for i := range coin {
		if a > coin[i] {
			results = a / coin[i]
			a %= coin[i]
		}
	}
	fmt.Print(results)
}

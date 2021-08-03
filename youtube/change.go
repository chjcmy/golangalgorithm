package main

import "fmt"

func main() {
	var come int
	var apt [4]int
	fmt.Scan(&come)

	coun := [4]int{500, 100, 50, 10}

	for i := range coun {
		if come > coun[i] {
			apt[i] = come / coun[i]
			come %= coun[i]
		}
	}

	for i := 0; i < len(apt); i++ {
		fmt.Printf("%d ", apt[i])
	}
}

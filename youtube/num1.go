package main

import (
	. "fmt"
)

func main() {

	var n, k, target, result int
	Scan(&n, &k)

	for true {
		target = (n / k) * k
		result += n - target
		n = target
		if n < k {
			break
		}
		result++
		n /= k
	}
	result += n - 1
	print(result)
}

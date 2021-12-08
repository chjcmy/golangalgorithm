package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var fibo [42][2]int
	in.Scan()

	a, _ := strconv.Atoi(in.Text())

	fibo[0][0] = 1
	fibo[1][1] = 1

	for i := 2; i < 42; i++ {
		for j := 0; j < 2; j++ {
			fibo[i][j] = fibo[i-1][j] + fibo[i-2][j]
		}
	}

	for i := 0; i < a; i++ {
		in.Scan()
		b, _ := strconv.Atoi(in.Text())
		fmt.Printf("%d %d\n", fibo[b][0], fibo[b][1])
	}
}
